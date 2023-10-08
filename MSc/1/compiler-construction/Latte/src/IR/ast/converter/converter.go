package converter

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	ast2 "github.com/badochov/MRJP/Latte/src/IR/ast"
	"github.com/badochov/MRJP/Latte/src/parser"
	"github.com/badochov/MRJP/Latte/src/status"
	"strconv"
)

type converter struct {
	currType ast2.IType[antlr.ParserRuleContext]
	temps    int
}

var _ parser.LatteVisitor = &converter{}

func (c *converter) VisitIndexAss(ctx *parser.IndexAssContext) interface{} {
	return &ast2.StmtIndexAss[antlr.ParserRuleContext]{
		Data:      ctx.GetParser().GetParserRuleContext(),
		ArrExpr:   ctx.Expr(0).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		IndexExpr: ctx.Expr(1).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		Expr:      ctx.Expr(2).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitAttrAss(ctx *parser.AttrAssContext) interface{} {
	return &ast2.StmtAttrAss[antlr.ParserRuleContext]{
		Data:    ctx.GetParser().GetParserRuleContext(),
		ObjExpr: ctx.Expr(0).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		Attr:    ctx.ID().GetText(),
		Expr:    ctx.Expr(1).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitForEach(ctx *parser.ForEachContext) interface{} {
	iId := c.nextTempVarName()
	lengthId := c.nextTempVarName()
	arrayId := c.nextTempVarName()
	itemType := ctx.Type_().Accept(c).(ast2.IType[antlr.ParserRuleContext])

	itemI := &ast2.ItemAss[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   iId,
		Expr: &ast2.ExprInt[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			Int:  0,
		},
	}
	itemArray := &ast2.ItemAss[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   arrayId,
		Expr: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
	itemLength := &ast2.ItemAss[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   lengthId,
		Expr: &ast2.ExprAttr[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			ObjExpr: &ast2.ExprId[antlr.ParserRuleContext]{
				Data: ctx.GetParser().GetParserRuleContext(),
				Id:   arrayId,
			},
			Attr: "length",
		},
	}
	declArr := &ast2.StmtDecl[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Type: &ast2.TypeArray[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			Type: itemType.(ast2.ITypeSingle[antlr.ParserRuleContext]),
		},
		Items: []ast2.IItem[antlr.ParserRuleContext]{
			itemArray,
		},
	}
	decl := &ast2.StmtDecl[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Type: &ast2.TypeInt[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
		},
		Items: []ast2.IItem[antlr.ParserRuleContext]{
			itemI,
			itemLength,
		},
	}
	itemEl := &ast2.ItemAss[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   ctx.ID().GetText(),
		Expr: &ast2.ExprIndex[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			ArrExpr: &ast2.ExprId[antlr.ParserRuleContext]{
				Data: ctx.GetParser().GetParserRuleContext(),
				Id:   arrayId,
			},
			IndexExpr: &ast2.ExprId[antlr.ParserRuleContext]{
				Data: ctx.GetParser().GetParserRuleContext(),
				Id:   iId,
			},
		},
	}
	itemDecl := &ast2.StmtDecl[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Type: itemType,
		Items: []ast2.IItem[antlr.ParserRuleContext]{
			itemEl,
		},
	}

	cond := &ast2.ExprRelOp[antlr.ParserRuleContext]{
		Data:  ctx.GetParser().GetParserRuleContext(),
		RelOp: ast2.Lt,
		LExpr: &ast2.ExprId[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			Id:   iId,
		},
		RExpr: &ast2.ExprId[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			Id:   lengthId,
		},
	}

	while := &ast2.StmtWhile[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Cond: cond,
		Stmt: &ast2.StmtBlock[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			Block: &ast2.Block[antlr.ParserRuleContext]{
				Data: ctx.GetParser().GetParserRuleContext(),
				Stmts: []ast2.IStmt[antlr.ParserRuleContext]{
					itemDecl,
					ctx.Stmt().Accept(c).(ast2.IStmt[antlr.ParserRuleContext]),
					c.incr(iId, ctx.GetParser().GetParserRuleContext()),
				},
			},
		},
	}

	return &ast2.StmtBlock[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Block: &ast2.Block[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			Stmts: []ast2.IStmt[antlr.ParserRuleContext]{
				declArr,
				decl,
				while,
			},
		},
	}
}

func (c *converter) nextTempVarName() string {
	c.temps++
	return strconv.Itoa(c.temps) // Vars cannot begin with number so this name would be illegal hence no name clash.
}

func (c *converter) VisitTypeArray(ctx *parser.TypeArrayContext) interface{} {
	return &ast2.TypeArray[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Type: ctx.Type_single().Accept(c).(ast2.ITypeSingle[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitTypeSingle(ctx *parser.TypeSingleContext) interface{} {
	return ctx.Type_single().Accept(c)
}

func (c *converter) VisitClassName(ctx *parser.ClassNameContext) interface{} {
	return &ast2.TypeClassName[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   ctx.ID().GetText(),
	}
}

func (c *converter) VisitClassAttribute(ctx *parser.ClassAttributeContext) interface{} {
	var ret []ast2.IClassElement[antlr.ParserRuleContext]
	for _, id := range ctx.AllID() {
		ret = append(ret, &ast2.Attribute[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			Type: ctx.Type_().Accept(c).(ast2.IType[antlr.ParserRuleContext]),
			Id:   id.GetText(),
		})
	}

	return ret
}

func (c *converter) VisitClassMethod(ctx *parser.ClassMethodContext) interface{} {
	method := &ast2.Method[antlr.ParserRuleContext]{
		Data:  ctx.GetParser().GetParserRuleContext(),
		Id:    ctx.ID().GetText(),
		Block: ctx.Block().Accept(c).(ast2.IBlock[antlr.ParserRuleContext]),
		Type:  ctx.Type_().Accept(c).(ast2.IType[antlr.ParserRuleContext]),
	}
	if arg := ctx.Arg(); arg != nil {
		method.Args = arg.Accept(c).([]ast2.IArg[antlr.ParserRuleContext])
	}
	return []ast2.IClassElement[antlr.ParserRuleContext]{method}
}

func (c *converter) VisitENewArray(ctx *parser.ENewArrayContext) interface{} {
	return &ast2.ExprNewArray[antlr.ParserRuleContext]{
		Data:     ctx.GetParser().GetParserRuleContext(),
		Type:     ctx.Type_single().Accept(c).(ast2.ITypeSingle[antlr.ParserRuleContext]),
		SizeExpr: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitENewClass(ctx *parser.ENewClassContext) interface{} {
	return &ast2.ExprNewClass[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Type: ctx.Type_single().Accept(c).(*ast2.TypeClassName[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitEMethodCall(ctx *parser.EMethodCallContext) interface{} {
	fn := &ast2.ExprMethodCall[antlr.ParserRuleContext]{
		Data:    ctx.GetParser().GetParserRuleContext(),
		Id:      ctx.ID().GetText(),
		ObjExpr: ctx.Expr(0).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}

	for _, expr := range ctx.AllExpr()[1:] {
		fn.Params = append(fn.Params, expr.Accept(c).(ast2.IExpr[antlr.ParserRuleContext]))
	}

	return fn
}

func (c *converter) VisitEIndex(ctx *parser.EIndexContext) interface{} {
	return &ast2.ExprIndex[antlr.ParserRuleContext]{
		Data:      ctx.GetParser().GetParserRuleContext(),
		ArrExpr:   ctx.Expr(0).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		IndexExpr: ctx.Expr(1).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitENull(ctx *parser.ENullContext) interface{} {
	typEx := ctx.Type_()
	var t ast2.IType[antlr.ParserRuleContext]
	if typEx != nil {
		t = typEx.Accept(c).(ast2.IType[antlr.ParserRuleContext])
	}

	return &ast2.ExprNull[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Type: t,
	}
}

func (c *converter) VisitEAttr(ctx *parser.EAttrContext) interface{} {
	return &ast2.ExprAttr[antlr.ParserRuleContext]{
		Data:    ctx.GetParser().GetParserRuleContext(),
		ObjExpr: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		Attr:    ctx.ID().GetText(),
	}
}

func ConvertAst(tree antlr.ParseTree) ast2.IProgram[antlr.ParserRuleContext] {
	return new(converter).Visit(tree).(ast2.IProgram[antlr.ParserRuleContext])
}

func (c *converter) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(c)
}

func (c *converter) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic(node.GetText()) // Impossible
}

func (c *converter) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic(node.GetText()) // Impossible
}

func (c *converter) VisitChildren(ctx antlr.RuleNode) interface{} {
	panic(ctx.GetText()) // Impossible
}

func (c *converter) VisitProgram(ctx *parser.ProgramContext) interface{} {
	p := &ast2.Program[antlr.ParserRuleContext]{Data: ctx.GetParser().GetParserRuleContext()}
	for _, def := range ctx.AllTopDef() {
		d := def.Accept(c)
		p.TopDefs = append(p.TopDefs, d.(ast2.ITopDef[antlr.ParserRuleContext]))
	}

	return p
}

func (c *converter) VisitDefFunc(ctx *parser.DefFuncContext) interface{} {
	fnName := ctx.ID().GetText()
	topDef := &ast2.Func[antlr.ParserRuleContext]{
		Data:  ctx.GetParser().GetParserRuleContext(),
		Id:    fnName,
		Block: ctx.Block().Accept(c).(ast2.IBlock[antlr.ParserRuleContext]),
		Type:  ctx.Type_().Accept(c).(ast2.IType[antlr.ParserRuleContext]),
	}
	if arg := ctx.Arg(); arg != nil {
		topDef.Args = arg.Accept(c).([]ast2.IArg[antlr.ParserRuleContext])
	}
	return topDef
}

func (c *converter) VisitDefClass(ctx *parser.DefClassContext) interface{} {
	topDef := &ast2.Class[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   ctx.ID(0).GetText(),
	}
	if extendsNode := ctx.ID(1); extendsNode != nil {
		extends := extendsNode.GetText()
		topDef.Extends = &extends
	}
	for _, el := range ctx.AllClassElement() {
		topDef.ClassElements = append(topDef.ClassElements, el.Accept(c).([]ast2.IClassElement[antlr.ParserRuleContext])...)
	}
	return topDef
}

func (c *converter) VisitArg(ctx *parser.ArgContext) interface{} {
	var args []ast2.IArg[antlr.ParserRuleContext]
	for i, t := range ctx.AllType_() {
		args = append(args, &ast2.Arg[antlr.ParserRuleContext]{
			Id:   ctx.ID(i).GetText(),
			Type: t.Accept(c).(ast2.IType[antlr.ParserRuleContext]),
			Ref:  ctx.Ref(i).Accept(c).(bool),
			Data: ctx.GetParser().GetParserRuleContext(),
		})
	}
	return args
}

func (c *converter) VisitRef(ctx *parser.RefContext) interface{} {
	return !ctx.IsEmpty()
}

func (c *converter) VisitBlock(ctx *parser.BlockContext) interface{} {
	block := &ast2.Block[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
	}

	for _, stmt := range ctx.AllStmt() {
		block.Stmts = append(block.Stmts, stmt.Accept(c).(ast2.IStmt[antlr.ParserRuleContext]))
	}

	return block
}

func (c *converter) VisitEmpty(ctx *parser.EmptyContext) interface{} {
	return &ast2.StmtEmpty[antlr.ParserRuleContext]{Data: ctx.GetParser().GetParserRuleContext()}
}

func (c *converter) VisitBlockStmt(ctx *parser.BlockStmtContext) interface{} {
	return &ast2.StmtBlock[antlr.ParserRuleContext]{
		Data:  ctx.GetParser().GetParserRuleContext(),
		Block: ctx.Block().Accept(c).(ast2.IBlock[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitDecl(ctx *parser.DeclContext) interface{} {
	c.currType = ctx.Type_().Accept(c).(ast2.IType[antlr.ParserRuleContext])
	decl := &ast2.StmtDecl[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Type: c.currType,
	}

	for _, item := range ctx.AllItem() {
		decl.Items = append(decl.Items, item.Accept(c).(ast2.IItem[antlr.ParserRuleContext]))
	}

	return decl
}

func (c *converter) VisitAss(ctx *parser.AssContext) interface{} {
	return &ast2.StmtAss[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   ctx.ID().GetText(),
		Expr: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitIncr(ctx *parser.IncrContext) interface{} {
	id := ctx.ID().GetText()
	return c.incr(id, ctx.GetParser().GetParserRuleContext())
}

func (c *converter) incr(id string, data antlr.ParserRuleContext) *ast2.StmtAss[antlr.ParserRuleContext] {
	return &ast2.StmtAss[antlr.ParserRuleContext]{
		Data: data,
		Id:   id,
		Expr: &ast2.ExprAddOp[antlr.ParserRuleContext]{
			Data: data,
			LExpr: &ast2.ExprId[antlr.ParserRuleContext]{
				Data: data,
				Id:   id,
			},
			RExpr: &ast2.ExprInt[antlr.ParserRuleContext]{
				Data: data,
				Int:  1,
			},
			AddOp: ast2.Plus,
		},
	}
}

func (c *converter) VisitDecr(ctx *parser.DecrContext) interface{} {
	id := ctx.ID().GetText()
	return &ast2.StmtAss[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   id,
		Expr: &ast2.ExprAddOp[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			LExpr: &ast2.ExprId[antlr.ParserRuleContext]{
				Data: ctx.GetParser().GetParserRuleContext(),
				Id:   id,
			},
			RExpr: &ast2.ExprInt[antlr.ParserRuleContext]{
				Data: ctx.GetParser().GetParserRuleContext(),
				Int:  1,
			},
			AddOp: ast2.Minus,
		},
	}
}

func (c *converter) VisitRet(ctx *parser.RetContext) interface{} {
	return &ast2.StmtRet[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Expr: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitVRet(ctx *parser.VRetContext) interface{} {
	return &ast2.StmtVRet[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
	}
}

func (c *converter) VisitCond(ctx *parser.CondContext) interface{} {
	return &ast2.StmtCond[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Cond: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		Stmt: ctx.Stmt().Accept(c).(ast2.IStmt[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitCondElse(ctx *parser.CondElseContext) interface{} {
	return &ast2.StmtCondElse[antlr.ParserRuleContext]{
		Data:     ctx.GetParser().GetParserRuleContext(),
		Cond:     ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		Stmt:     ctx.Stmt(0).Accept(c).(ast2.IStmt[antlr.ParserRuleContext]),
		ElseStmt: ctx.Stmt(1).Accept(c).(ast2.IStmt[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitWhile(ctx *parser.WhileContext) interface{} {
	return &ast2.StmtWhile[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Cond: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		Stmt: ctx.Stmt().Accept(c).(ast2.IStmt[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitSExp(ctx *parser.SExpContext) interface{} {
	return &ast2.StmtExp[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Expr: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitInt(ctx *parser.IntContext) interface{} {
	return &ast2.TypeInt[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
	}
}

func (c *converter) VisitStr(ctx *parser.StrContext) interface{} {
	return &ast2.TypeString[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
	}
}

func (c *converter) VisitBool(ctx *parser.BoolContext) interface{} {
	return &ast2.TypeBool[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
	}
}

func (c *converter) VisitVoid(ctx *parser.VoidContext) interface{} {
	return &ast2.TypeVoid[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
	}
}

func (c *converter) VisitItem(ctx *parser.ItemContext) interface{} {
	id := ctx.ID().GetText()
	if expr := ctx.Expr(); expr != nil {
		return &ast2.ItemAss[antlr.ParserRuleContext]{
			Data: ctx.GetParser().GetParserRuleContext(),
			Id:   id,
			Expr: expr.Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		}
	}
	expr := c.getExprForCurrType(ctx.GetParser().GetParserRuleContext())

	return &ast2.ItemAss[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   id,
		Expr: expr,
	}
}

func (c *converter) VisitEId(ctx *parser.EIdContext) interface{} {
	return &ast2.ExprId[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   ctx.ID().GetText(),
	}
}

func (c *converter) VisitEFunCall(ctx *parser.EFunCallContext) interface{} {
	fn := &ast2.ExprFunCall[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Id:   ctx.ID().GetText(),
	}

	for _, expr := range ctx.AllExpr() {
		fn.Params = append(fn.Params, expr.Accept(c).(ast2.IExpr[antlr.ParserRuleContext]))
	}

	return fn
}

func (c *converter) VisitERelOp(ctx *parser.ERelOpContext) interface{} {
	return &ast2.ExprRelOp[antlr.ParserRuleContext]{
		Data:  ctx.GetParser().GetParserRuleContext(),
		LExpr: ctx.Expr(0).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		RExpr: ctx.Expr(1).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		RelOp: ctx.RelOp().Accept(c).(ast2.RelOp),
	}
}

func (c *converter) VisitETrue(ctx *parser.ETrueContext) interface{} {
	return &ast2.ExprBool[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Bool: true,
	}
}

func (c *converter) VisitEOr(ctx *parser.EOrContext) interface{} {
	return &ast2.ExprOr[antlr.ParserRuleContext]{
		Data:  ctx.GetParser().GetParserRuleContext(),
		LExpr: ctx.Expr(0).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		RExpr: ctx.Expr(1).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitEInt(ctx *parser.EIntContext) interface{} {
	i, _ := strconv.ParseInt(ctx.GetText(), 0, 64)
	return &ast2.ExprInt[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Int:  i,
	}
}

func (c *converter) VisitEUnOp(ctx *parser.EUnOpContext) interface{} {
	return &ast2.ExprUnOp[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Expr: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		UnOp: ctx.UnOp().Accept(c).(ast2.UnOp),
	}
}

func (c *converter) VisitEStr(ctx *parser.EStrContext) interface{} {
	raw := ctx.GetText()
	unquoted, err := strconv.Unquote(raw)
	if err != nil {
		status.HandleError(err, ctx.GetParser().GetParserRuleContext())
	}

	return &ast2.ExprString[antlr.ParserRuleContext]{
		Data:   ctx.GetParser().GetParserRuleContext(),
		String: unquoted,
	}
}

func (c *converter) VisitEMulOp(ctx *parser.EMulOpContext) interface{} {
	return &ast2.ExprMulOp[antlr.ParserRuleContext]{
		Data:  ctx.GetParser().GetParserRuleContext(),
		LExpr: ctx.Expr(0).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		RExpr: ctx.Expr(1).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		MulOp: ctx.MulOp().Accept(c).(ast2.MulOp),
	}
}

func (c *converter) VisitEAnd(ctx *parser.EAndContext) interface{} {
	return &ast2.ExprAnd[antlr.ParserRuleContext]{
		Data:  ctx.GetParser().GetParserRuleContext(),
		LExpr: ctx.Expr(0).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		RExpr: ctx.Expr(1).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitEParen(ctx *parser.EParenContext) interface{} {
	return &ast2.ExprParen[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Expr: ctx.Expr().Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
	}
}

func (c *converter) VisitEFalse(ctx *parser.EFalseContext) interface{} {
	return &ast2.ExprBool[antlr.ParserRuleContext]{
		Data: ctx.GetParser().GetParserRuleContext(),
		Bool: false,
	}
}

func (c *converter) VisitEAddOp(ctx *parser.EAddOpContext) interface{} {
	return &ast2.ExprAddOp[antlr.ParserRuleContext]{
		Data:  ctx.GetParser().GetParserRuleContext(),
		LExpr: ctx.Expr(0).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		RExpr: ctx.Expr(1).Accept(c).(ast2.IExpr[antlr.ParserRuleContext]),
		AddOp: ctx.AddOp().Accept(c).(ast2.AddOp),
	}
}

func (c *converter) VisitAddOp(ctx *parser.AddOpContext) interface{} {
	switch ctx.GetText() {
	case "+":
		return ast2.Plus
	case "-":
		return ast2.Minus
	default:
		panic("IMPOSSIBLE")
	}
}

func (c *converter) VisitUnOp(ctx *parser.UnOpContext) interface{} {
	switch ctx.GetText() {
	case "!":
		return ast2.LogicNot
	case "-":
		return ast2.UnaryMinus
	default:
		panic("IMPOSSIBLE")
	}
}

func (c *converter) VisitMulOp(ctx *parser.MulOpContext) interface{} {
	switch ctx.GetText() {
	case "*":
		return ast2.Times
	case "/":
		return ast2.Divide
	case "%":
		return ast2.Modulo
	default:
		panic("IMPOSSIBLE")
	}
}

func (c *converter) VisitRelOp(ctx *parser.RelOpContext) interface{} {
	switch ctx.GetText() {
	case ">":
		return ast2.Gt
	case ">=":
		return ast2.Ge
	case "<":
		return ast2.Lt
	case "<=":
		return ast2.Le
	case "!=":
		return ast2.Ne
	case "==":
		return ast2.Eq
	default:
		panic("IMPOSSIBLE")
	}
}

func (c *converter) getExprForCurrType(data antlr.ParserRuleContext) ast2.IExpr[antlr.ParserRuleContext] {
	switch t := c.currType.(type) {
	case *ast2.TypeInt[antlr.ParserRuleContext]:
		return &ast2.ExprInt[antlr.ParserRuleContext]{
			Data: data,
			Int:  0,
		}
	case *ast2.TypeBool[antlr.ParserRuleContext]:
		return &ast2.ExprBool[antlr.ParserRuleContext]{
			Data: data,
			Bool: false,
		}
	case *ast2.TypeString[antlr.ParserRuleContext]:
		return &ast2.ExprString[antlr.ParserRuleContext]{
			Data:   data,
			String: "",
		}
	default:
		return &ast2.ExprNull[antlr.ParserRuleContext]{
			Data: data,
			Type: t,
		}
	}
}
