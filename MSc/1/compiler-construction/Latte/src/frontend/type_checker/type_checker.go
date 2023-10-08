package type_checker

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/badochov/MRJP/Latte/src/frontend/types"
	"github.com/badochov/MRJP/Latte/src/parser"
	"github.com/badochov/MRJP/Latte/src/status"
	"github.com/badochov/MRJP/Latte/src/topdefs/signatures"
	"strconv"
)

type TypeChecker struct {
	classes       map[string]signatures.Class
	blockCtx      *blockCtx
	retType       *types.Type
	currDeclType  *types.Type
	currClassType *types.Type

	onDeclarationError error
}

var _ parser.LatteVisitor = &TypeChecker{}

func New(topDefs signatures.TopDefs) *TypeChecker {
	return &TypeChecker{
		classes: topDefs.Classes,
		blockCtx: &blockCtx{
			vars: make(map[string]types.Type),
			fns:  topDefs.Fns,
		},
		retType:      nil,
		currDeclType: nil,
	}
}

func (v *TypeChecker) VisitRef(ctx *parser.RefContext) interface{} {
	panic("REF SHOULD NEVER BE VISITED")
}

func (v *TypeChecker) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *TypeChecker) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic(node.GetText()) // Impossible
}

func (v *TypeChecker) VisitErrorNode(node antlr.ErrorNode) interface{} {
	status.Error("%s", node.GetText())
	return nil
}

func (v *TypeChecker) VisitChildren(ctx antlr.RuleNode) interface{} {
	panic(ctx.GetText()) // Impossible
}

func (v *TypeChecker) VisitDefFunc(ctx *parser.DefFuncContext) interface{} {
	rT := ctx.Type_().Accept(v).(types.Type)
	v.retType = &rT

	v.enterBlock()
	defer v.exitBlock()
	if arg := ctx.Arg(); arg != nil {
		arg.Accept(v)
	}

	return ctx.Block().Accept(v)
}

func (v *TypeChecker) VisitDefClass(ctx *parser.DefClassContext) interface{} {
	className := ctx.ID(0).GetText()
	v.enterClass(className)
	defer v.exitClass(className)

	for _, el := range ctx.AllClassElement() {
		el.Accept(v)
	}
	return nil
}

func (v *TypeChecker) VisitIndexAss(ctx *parser.IndexAssContext) interface{} {
	arr := ctx.Expr(0).Accept(v).(types.Type)
	if arr.Array == nil {
		status.HandleError(fmt.Errorf("can't index types '%s'", arr), ctx.GetParser().GetParserRuleContext())
	}
	index := ctx.Expr(1).Accept(v).(types.Type)
	if index != types.Int {
		status.HandleError(fmt.Errorf("can't index with type '%s'", index), ctx.GetParser().GetParserRuleContext())
	}

	val := ctx.Expr(2).Accept(v).(types.Type)
	status.HandleError(v.assignsToType(*arr.Array, val), ctx.GetParser().GetParserRuleContext())

	return nil
}

func (v *TypeChecker) VisitAttrAss(ctx *parser.AttrAssContext) interface{} {
	obj := ctx.Expr(0).Accept(v).(types.Type)
	attr := ctx.ID().GetText()
	if obj.Struct {
		expr := ctx.Expr(1).Accept(v).(types.Type)
		class := v.classes[obj.Name]
		for classP := &class; classP != nil; classP = classP.Extends {
			if attrType, ok := classP.Attributes[attr]; ok {
				status.HandleError(v.assignsToType(attrType, expr), ctx.GetParser().GetParserRuleContext())
				return nil
			}
		}
		status.HandleError(fmt.Errorf("class '%s' doesn't have attribute '%s'", obj, attr), ctx.GetParser().GetParserRuleContext())
	}

	status.HandleError(fmt.Errorf("can't assign to '%s' of '%s'", attr, obj), ctx.GetParser().GetParserRuleContext())
	panic("HANDLE ERROR EXITS")
}

func (v *TypeChecker) VisitClassAttribute(ctx *parser.ClassAttributeContext) interface{} {
	// Already checked by signature getter.
	return nil
}

func (v *TypeChecker) VisitClassMethod(ctx *parser.ClassMethodContext) interface{} {
	rT := ctx.Type_().Accept(v).(types.Type)
	v.retType = &rT

	v.enterBlock()
	defer v.exitBlock()
	if arg := ctx.Arg(); arg != nil {
		arg.Accept(v)
	}
	status.HandleError(v.blockCtx.DeclareType("self", *v.currClassType), ctx)

	return ctx.Block().Accept(v)
}

func (v *TypeChecker) assignsToType(varType types.Type, exprType types.Type) error {
	if varType != exprType {
		if varType.Array != nil && exprType.Array != nil && *varType.Array == *exprType.Array {
			return nil
		}

		// Check for null
		if exprType == types.Null && (varType.Array != nil || varType.Struct) {
			return nil
		}

		// Check inheritance
		if varType.Struct && exprType.Struct {
			varName := varType.Name
			exprT := v.classes[varType.Name]
			exprTP := &exprT

			for ; exprTP != nil; exprTP = exprTP.Extends {
				if exprTP.Name == varName {
					return nil
				}
			}
		}
		return fmt.Errorf("cannot assign type '%s' to '%s'", exprType, varType)
	}
	return nil
}

func (v *TypeChecker) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, def := range ctx.AllTopDef() {
		def.Accept(v)
	}

	return nil
}

func (v *TypeChecker) VisitArg(ctx *parser.ArgContext) interface{} {
	for i, t := range ctx.AllType_() {
		argT := t.Accept(v).(types.Type)
		argName := ctx.ID(i).GetText()
		if argT == types.Void {
			status.HandleError(fmt.Errorf("argument '%s' is of type void", argName), ctx)
		}
		status.HandleError(v.blockCtx.DeclareType(argName, argT), ctx)
	}
	return nil
}

func (v *TypeChecker) enterBlock() {
	v.blockCtx = &blockCtx{
		prev: v.blockCtx,
		vars: make(map[string]types.Type),
		fns:  make(map[string]signatures.Function),
	}
}

func (v *TypeChecker) exitBlock() {
	v.blockCtx = v.blockCtx.prev
}

func (v *TypeChecker) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.enterBlock()
	defer v.exitBlock()

	defer v.SetOnDeclarationError(nil)()

	for _, stmt := range ctx.AllStmt() {
		stmt.Accept(v)
	}

	return nil
}

func (v *TypeChecker) SetOnDeclarationError(err error) func() {
	prev := v.onDeclarationError
	v.onDeclarationError = err
	return func() {
		v.onDeclarationError = prev
	}
}

func (v *TypeChecker) VisitEmpty(ctx *parser.EmptyContext) interface{} {
	return nil
}

func (v *TypeChecker) VisitBlockStmt(ctx *parser.BlockStmtContext) interface{} {
	return ctx.Block().Accept(v)
}

func (v *TypeChecker) VisitForEach(ctx *parser.ForEachContext) interface{} {
	arrT := ctx.Expr().Accept(v).(types.Type)
	if arrT.Array == nil {
		status.HandleError(fmt.Errorf("for each can only be performed on array"), ctx.GetParser().GetParserRuleContext())
	}
	t := ctx.Type_().Accept(v).(types.Type)
	status.HandleError(v.assignsToType(t, *arrT.Array), ctx.GetParser().GetParserRuleContext())
	v.enterBlock()
	defer v.exitBlock()
	status.HandleError(v.blockCtx.DeclareType(ctx.ID().GetText(), t), ctx.GetParser().GetParserRuleContext())

	defer v.SetOnDeclarationError(declarationInLoopError)()
	ctx.Stmt().Accept(v)

	return nil
}

func (v *TypeChecker) VisitDecl(ctx *parser.DeclContext) interface{} {
	if v.onDeclarationError != nil {
		status.HandleError(v.onDeclarationError, ctx.GetParser().GetParserRuleContext())
	}

	t := ctx.Type_().Accept(v).(types.Type)
	if t == types.Void {
		status.HandleError(fmt.Errorf("variable type can't 'void'"), ctx.GetParser().GetParserRuleContext())
	}
	v.currDeclType = &t
	for _, item := range ctx.AllItem() {
		item.Accept(v)
	}
	v.currDeclType = nil
	return nil
}

func (v *TypeChecker) VisitAss(ctx *parser.AssContext) interface{} {
	t, err := v.blockCtx.GetType(ctx.ID().GetText())
	status.HandleError(err, ctx) // Improve status.HandleError handling.

	exT := ctx.Expr().Accept(v).(types.Type)

	status.HandleError(v.assignsToType(t, exT), ctx)
	return nil
}

func (v *TypeChecker) VisitIncr(ctx *parser.IncrContext) interface{} {
	t, err := v.blockCtx.GetType(ctx.ID().GetText())
	status.HandleError(err, ctx) // Improve status.HandleError handling.
	if t != types.Int {
		status.HandleError(fmt.Errorf("can't increment '%s'", t), ctx)
	}
	return types.Int
}

func (v *TypeChecker) VisitDecr(ctx *parser.DecrContext) interface{} {
	t, err := v.blockCtx.GetType(ctx.ID().GetText())
	status.HandleError(err, ctx) // Improve status.HandleError handling.
	if t != types.Int {
		status.HandleError(fmt.Errorf("can't decrement '%s'", t), ctx)
	}
	return types.Int
}

func (v *TypeChecker) VisitRet(ctx *parser.RetContext) interface{} {
	if *v.retType == types.Void {
		status.HandleError(fmt.Errorf("can't return value from function that returns void '%s'", *v.retType), ctx)
	}
	exprT := ctx.Expr().Accept(v).(types.Type)
	status.HandleError(v.assignsToType(*v.retType, exprT), ctx)

	return nil
}

func (v *TypeChecker) VisitVRet(ctx *parser.VRetContext) interface{} {
	if *v.retType != types.Void {
		status.HandleError(fmt.Errorf("can't return void in a function with return type '%s'", *v.retType), ctx)
	}
	return nil
}

func (v *TypeChecker) VisitCond(ctx *parser.CondContext) interface{} {
	exprT := ctx.Expr().Accept(v).(types.Type)
	if types.Bool != exprT {
		status.HandleError(fmt.Errorf("expected bool expression in if condition, got '%s'", exprT), ctx)
	}

	defer v.SetOnDeclarationError(declarationInIfError)()
	ctx.Stmt().Accept(v)

	return nil
}

func (v *TypeChecker) VisitCondElse(ctx *parser.CondElseContext) interface{} {
	exprT := ctx.Expr().Accept(v).(types.Type)

	if types.Bool != exprT {
		status.HandleError(fmt.Errorf("expected bool expression in if condition, got '%s'", exprT), ctx)
	}

	defer v.SetOnDeclarationError(declarationInIfError)()
	for _, stmt := range ctx.AllStmt() {
		stmt.Accept(v)
	}

	return nil
}

var declarationInLoopError = fmt.Errorf("cannot declare variable in loop without enclosing block")
var declarationInIfError = fmt.Errorf("cannot declare variable in if without enclosing block")

func (v *TypeChecker) VisitWhile(ctx *parser.WhileContext) interface{} {
	exprT := ctx.Expr().Accept(v).(types.Type)
	if types.Bool != exprT {
		status.HandleError(fmt.Errorf("expected bool expression in while condition, got '%s'", exprT), ctx)
	}

	defer v.SetOnDeclarationError(declarationInLoopError)()
	ctx.Stmt().Accept(v)

	return nil
}

func (v *TypeChecker) VisitSExp(ctx *parser.SExpContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *TypeChecker) VisitClassName(ctx *parser.ClassNameContext) interface{} {
	name := ctx.ID().GetText()
	return types.Type{
		Name:   name,
		Struct: true,
	}
}

func (v *TypeChecker) VisitInt(ctx *parser.IntContext) interface{} {
	return types.Int
}

func (v *TypeChecker) VisitTypeArray(ctx *parser.TypeArrayContext) interface{} {
	t := ctx.Type_single().Accept(v).(types.Type)
	return types.Type{
		Array: &t,
	}
}

func (v *TypeChecker) VisitTypeSingle(ctx *parser.TypeSingleContext) interface{} {
	return ctx.Type_single().Accept(v)
}

func (v *TypeChecker) VisitStr(ctx *parser.StrContext) interface{} {
	return types.String
}

func (v *TypeChecker) VisitBool(ctx *parser.BoolContext) interface{} {
	return types.Bool
}

func (v *TypeChecker) VisitVoid(ctx *parser.VoidContext) interface{} {
	return types.Void
}

func (v *TypeChecker) VisitItem(ctx *parser.ItemContext) interface{} {
	if expr := ctx.Expr(); expr != nil {
		exprT := expr.Accept(v).(types.Type)
		status.HandleError(v.assignsToType(*v.currDeclType, exprT), ctx)
	}
	status.HandleError(v.blockCtx.DeclareType(ctx.ID().GetText(), *v.currDeclType), ctx)

	return nil
}

func (v *TypeChecker) VisitEId(ctx *parser.EIdContext) interface{} {
	t, err := v.blockCtx.GetType(ctx.ID().GetText())
	status.HandleError(err, ctx)

	return t
}

func (v *TypeChecker) VisitEFunCall(ctx *parser.EFunCallContext) interface{} {
	fnName := ctx.ID().GetText()
	fnSig, err := v.blockCtx.GetFn(fnName)
	status.HandleError(err, ctx)
	if len(fnSig.ArgTypes) != len(ctx.AllExpr()) {
		status.HandleError(fmt.Errorf("function '%s' takes %d arguments but was provided with %d", fnName, len(fnSig.ArgTypes), len(ctx.AllExpr())), ctx)
	}
	for i, expr := range ctx.AllExpr() {
		argT := expr.Accept(v).(types.Type)
		status.HandleError(v.assignsToType(fnSig.ArgTypes[i].Type, argT), ctx)
	}

	return fnSig.RetType
}

func (v *TypeChecker) VisitERelOp(ctx *parser.ERelOpContext) interface{} {
	expr1T := ctx.Expr(0).Accept(v).(types.Type)
	expr2T := ctx.Expr(1).Accept(v).(types.Type)
	op := ctx.RelOp().Accept(v).(BinOp)

	t, err := op(expr1T, expr2T)
	status.HandleError(err, ctx)

	return t
}

func (v *TypeChecker) VisitETrue(ctx *parser.ETrueContext) interface{} {
	return types.Bool
}

func (v *TypeChecker) VisitEOr(ctx *parser.EOrContext) interface{} {
	expr1T := ctx.Expr(0).Accept(v).(types.Type)
	expr2T := ctx.Expr(1).Accept(v).(types.Type)

	if expr1T != types.Bool || expr2T != types.Bool {
		status.HandleError(fmt.Errorf("can't perform '%s' || '%s'", expr1T, expr2T), ctx)
	}
	return types.Bool
}

func (v *TypeChecker) VisitEInt(ctx *parser.EIntContext) interface{} {
	if _, err := strconv.ParseInt(ctx.GetText(), 0, 64); err != nil {
		status.Fatal(err)
	}
	return types.Int
}

func (v *TypeChecker) VisitEUnOp(ctx *parser.EUnOpContext) interface{} {
	exprT := ctx.Expr().Accept(v).(types.Type)
	op := ctx.UnOp().Accept(v).(UnOp)

	t, err := op(exprT)
	status.HandleError(err, ctx)

	return t
}

func (v *TypeChecker) VisitEStr(ctx *parser.EStrContext) interface{} {
	return types.String
}

func (v *TypeChecker) VisitEMulOp(ctx *parser.EMulOpContext) interface{} {
	expr1T := ctx.Expr(0).Accept(v).(types.Type)
	expr2T := ctx.Expr(1).Accept(v).(types.Type)
	op := ctx.MulOp().Accept(v).(BinOp)

	t, err := op(expr1T, expr2T)
	status.HandleError(err, ctx)

	return t
}

func (v *TypeChecker) VisitEAnd(ctx *parser.EAndContext) interface{} {
	expr1T := ctx.Expr(0).Accept(v).(types.Type)
	expr2T := ctx.Expr(1).Accept(v).(types.Type)

	if expr1T != types.Bool || expr2T != types.Bool {
		status.HandleError(fmt.Errorf("can't perform '%s' && '%s'", expr1T, expr2T), ctx)
	}

	return types.Bool
}

func (v *TypeChecker) VisitEParen(ctx *parser.EParenContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *TypeChecker) VisitEFalse(ctx *parser.EFalseContext) interface{} {
	return types.Bool
}

func (v *TypeChecker) VisitEAddOp(ctx *parser.EAddOpContext) interface{} {
	expr1T := ctx.Expr(0).Accept(v).(types.Type)
	expr2T := ctx.Expr(1).Accept(v).(types.Type)
	op := ctx.AddOp().Accept(v).(BinOp)

	t, err := op(expr1T, expr2T)
	status.HandleError(err, ctx)

	return t
}

func (v *TypeChecker) VisitENewArray(ctx *parser.ENewArrayContext) interface{} {
	t := ctx.Type_single().Accept(v).(types.Type)
	expr := ctx.Expr().Accept(v).(types.Type)
	if expr != types.Int {
		status.HandleError(fmt.Errorf("size of an array must be an integer"), ctx.GetParser().GetParserRuleContext())
	}

	return types.Type{
		Array: &t,
	}
}

func (v *TypeChecker) VisitENewClass(ctx *parser.ENewClassContext) interface{} {
	t := ctx.Type_single().Accept(v).(types.Type)
	if !t.Struct {
		status.HandleError(fmt.Errorf("new can be only perform on struct"), ctx.GetParser().GetParserRuleContext())
	}

	return t
}

func (v *TypeChecker) VisitEIndex(ctx *parser.EIndexContext) interface{} {
	arr := ctx.Expr(0).Accept(v).(types.Type)
	if arr.Array == nil {
		status.HandleError(fmt.Errorf("can't index types '%s'", arr), ctx.GetParser().GetParserRuleContext())
	}
	index := ctx.Expr(1).Accept(v).(types.Type)
	if index != types.Int {
		status.HandleError(fmt.Errorf("can't index with type '%s'", index), ctx.GetParser().GetParserRuleContext())
	}

	return *arr.Array
}

func (v *TypeChecker) VisitENull(ctx *parser.ENullContext) interface{} {
	typ := ctx.Type_()
	if typ == nil {
		return types.Null
	}
	t := typ.Accept(v).(types.Type)
	if !t.Struct && t.Array == nil {
		status.HandleError(fmt.Errorf("only classes and arrays can be null"), ctx.GetParser().GetParserRuleContext())
	}
	return t
}

func (v *TypeChecker) VisitEAttr(ctx *parser.EAttrContext) interface{} {
	obj := ctx.Expr().Accept(v).(types.Type)
	attr := ctx.ID().GetText()
	if obj.Array != nil {
		if attr == "length" {
			return types.Int
		}
		status.HandleError(fmt.Errorf("arrays doesn't have attribute '%s'", attr), ctx.GetParser().GetParserRuleContext())
	}
	if obj.Struct {
		class := v.classes[obj.Name]
		for classP := &class; classP != nil; classP = classP.Extends {
			if attrType, ok := classP.Attributes[attr]; ok {
				return attrType
			}
		}
		status.HandleError(fmt.Errorf("class '%s' doesn't have attribute '%s'", obj, attr), ctx.GetParser().GetParserRuleContext())
	}

	status.HandleError(fmt.Errorf("'%s' doesn't have attribute '%s'", obj, attr), ctx.GetParser().GetParserRuleContext())
	panic("HANDLE ERROR EXITS")
}

func (v *TypeChecker) VisitEMethodCall(ctx *parser.EMethodCallContext) interface{} {
	obj := ctx.Expr(0).Accept(v).(types.Type)
	fnName := ctx.ID().GetText()

	var fnSig *signatures.Function
	if obj.Struct {
		class := v.classes[obj.Name]
		for classP := &class; classP != nil; classP = classP.Extends {
			if fn, ok := classP.Methods[fnName]; ok {
				fnSig = &fn
				break
			}
		}
	}
	if fnSig == nil {
		status.HandleError(fmt.Errorf("can't call '%s' on '%s'", fnName, obj), ctx.GetParser().GetParserRuleContext())
	}
	numProvidedArgs := len(ctx.AllExpr()) - 1
	if len(fnSig.ArgTypes) != numProvidedArgs {
		status.HandleError(fmt.Errorf("function '%s' takes %d arguments but was provided with %d", fnName, len(fnSig.ArgTypes), numProvidedArgs), ctx)
	}
	for i, expr := range ctx.AllExpr()[1:] {
		argT := expr.Accept(v).(types.Type)
		status.HandleError(v.assignsToType(fnSig.ArgTypes[i].Type, argT), ctx)
	}

	return fnSig.RetType
}

type BinOp = func(types.Type, types.Type) (types.Type, error)
type UnOp = func(types.Type) (types.Type, error)

func (v *TypeChecker) VisitAddOp(ctx *parser.AddOpContext) interface{} {
	switch ctx.GetText() {
	case "+":
		return func(op1, op2 types.Type) (types.Type, error) {
			if op1 != op2 || (op1 != types.String && op2 != types.Int) {
				return types.Type{}, fmt.Errorf("can't add '%s' to '%s'", op1, op2)
			}
			return op1, nil
		}
	case "-":
		return func(op1, op2 types.Type) (types.Type, error) {
			if op1 != op2 || op1 != types.Int {
				return types.Type{}, fmt.Errorf("can't subtract '%s' from '%s'", op1, op2)
			}
			return op1, nil
		}
	default:
		panic(ctx.GetText()) // Not possible to enter
	}
}

func (v *TypeChecker) VisitUnOp(ctx *parser.UnOpContext) interface{} {
	switch ctx.GetText() {
	case "-":
		return func(op types.Type) (types.Type, error) {
			if op != types.Int {
				return types.Type{}, fmt.Errorf("can't perform -'%s'", op)
			}
			return op, nil
		}
	case "!":
		return func(op types.Type) (types.Type, error) {
			if op != types.Bool {
				return types.Type{}, fmt.Errorf("can't perform !'%s'", op)
			}
			return op, nil
		}
	default:
		panic(ctx.GetText()) // Not possible to enter
	}
}

func (v *TypeChecker) VisitMulOp(ctx *parser.MulOpContext) interface{} {
	switch ctx.GetText() {
	case "*":
		return func(op1, op2 types.Type) (types.Type, error) {
			if op1 != op2 || op2 != types.Int {
				return types.Type{}, fmt.Errorf("can't multiply '%s' by '%s'", op1, op2)
			}
			return op1, nil
		}
	case "/":
		return func(op1, op2 types.Type) (types.Type, error) {
			if op1 != op2 || op1 != types.Int {
				return types.Type{}, fmt.Errorf("can't divide '%s' by '%s'", op1, op2)
			}
			return op1, nil
		}
	case "%":
		return func(op1, op2 types.Type) (types.Type, error) {
			if op1 != op2 || op1 != types.Int {
				return types.Type{}, fmt.Errorf("can't take modolu of '%s' by '%s'", op1, op2)
			}
			return op1, nil
		}
	default:
		panic(ctx.GetText()) // Not possible to enter
	}
}

func (v *TypeChecker) VisitRelOp(ctx *parser.RelOpContext) interface{} {
	switch ctx.GetText() {
	case "<", ">", ">=", "<=":
		return func(op1, op2 types.Type) (types.Type, error) {
			if op1 != op2 || op1 != types.Int {
				return types.Type{}, fmt.Errorf("can't compare '%s' with '%s'", op1, op2)
			}
			return types.Bool, nil
		}
	case "==", "!=":
		return func(op1, op2 types.Type) (types.Type, error) {
			if v.assignsToType(op1, op2) != nil && v.assignsToType(op2, op1) != nil {
				return types.Type{}, fmt.Errorf("can't compare '%s' with '%s'", op1, op2)
			}
			return types.Bool, nil
		}
	default:
		panic(ctx.GetText()) // Not possible to enter
	}
}

func (v *TypeChecker) Validate() {
	v.validateMainExists()
}

func (v *TypeChecker) validateMainExists() {
	t, err := v.blockCtx.GetFn("main")
	if err != nil {
		status.Fatal(err)
	}
	if t.RetType != types.Int || len(t.ArgTypes) != 0 {
		status.Fatal(fmt.Errorf("main function function signature should be 'int main()'"))
	}
}

func (v *TypeChecker) enterClass(name string) {
	class := v.classes[name]
	parents := class.GetParents()
	for i := len(parents) - 1; i >= 0; i-- {
		v.blockCtx = &blockCtx{
			prev: v.blockCtx,
			vars: parents[i].Attributes,
			fns:  parents[i].Methods,
		}
	}
	v.blockCtx = &blockCtx{
		prev: v.blockCtx,
		vars: class.Attributes,
		fns:  class.Methods,
	}

	v.currClassType = &types.Type{
		Name:   name,
		Struct: true,
	}
}

func (v *TypeChecker) exitClass(name string) {
	for range v.classes[name].GetParents() {
		v.blockCtx = v.blockCtx.prev
	}
	v.blockCtx = v.blockCtx.prev

	v.currClassType = nil
}

type blockCtx struct {
	prev *blockCtx
	vars map[string]types.Type
	fns  map[string]signatures.Function
}

func (b *blockCtx) GetType(varName string) (types.Type, error) {
	t, ok := b.vars[varName]
	if !ok {
		if b.prev == nil {
			return types.Type{}, fmt.Errorf("undeclared variable '%s'", varName)
		}
		return b.prev.GetType(varName)
	}
	return t, nil
}

func (b *blockCtx) DeclareType(varName string, t types.Type) error {
	if _, ok := b.vars[varName]; ok {
		return fmt.Errorf("redeclaration of variable '%s'", varName)
	}
	b.vars[varName] = t

	return nil
}

func (b *blockCtx) GetFn(fnName string) (signatures.Function, error) {
	t, ok := b.fns[fnName]
	if !ok {
		if b.prev == nil {
			return signatures.Function{}, fmt.Errorf("undeclared function '%s'", fnName)
		}
		return b.prev.GetFn(fnName)
	}
	return t, nil
}

func (b *blockCtx) DeclareFn(fnName string, fn signatures.Function) error {
	if _, ok := b.fns[fnName]; ok {
		return fmt.Errorf("redeclaration of function '%s'", fnName)
	}
	b.fns[fnName] = fn

	return nil
}
