package compile

import (
	"fmt"
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func Compile(program quadruple.Program) *LatteCode {
	c := newCompiler(program)
	c.createDeclarations()
	c.fillFunctions()

	return &LatteCode{module: c.module}
}

var boolType = &types.IntType{BitSize: 1}
var charType = &types.IntType{BitSize: 8}
var intType = &types.IntType{BitSize: 64}

var vFnType = &types.PointerType{
	ElemType: &types.FuncType{
		RetType: &types.VoidType{},
	},
}

func newCompiler(program quadruple.Program) *compiler {
	return &compiler{
		module:  ir.NewModule(),
		strs:    make(map[string]*ir.Global),
		fns:     make(map[*quadruple.Fn]*ir.Func),
		vtables: make(map[string]*ir.Global),
		types:   make(map[string]*types.StructType),
		program: program,
	}
}

type compiler struct {
	module     *ir.Module
	strs       map[string]*ir.Global
	fns        map[*quadruple.Fn]*ir.Func
	vtables    map[string]*ir.Global
	blocksMap  map[*quadruple.Block]*ir.Block
	types      map[string]*types.StructType
	currBlock  *ir.Block
	varMap     map[quadruple.Var]value.Value
	program    quadruple.Program
	phisToFill []phiToFill
}

func (c *compiler) VisitFn(fn *quadruple.Fn) any {
	return c.fns[fn]
}

func (c *compiler) VisitMethod(method *quadruple.Method) any {
	obj := c.replaceVal(method.Object)
	objType := obj.Type().(*types.PointerType).ElemType
	vTableType := objType.(*types.StructType).Fields[0]

	vpointer := c.currBlock.NewGetElementPtr(
		objType,
		obj,
		constant.NewInt(&types.IntType{BitSize: 64}, 0),
		constant.NewInt(&types.IntType{BitSize: 32}, 0),
	)
	vtable := c.currBlock.NewLoad(vTableType, vpointer)
	vFn := c.currBlock.NewGetElementPtr(
		vtable.Type().(*types.PointerType).ElemType,
		vtable,
		constant.NewInt(&types.IntType{BitSize: 64}, 0),
		constant.NewInt(&types.IntType{BitSize: 64}, int64(method.Idx)),
	)
	wrongTypeFn := c.currBlock.NewLoad(vFn.Type().(*types.PointerType).ElemType, vFn)

	quadrupleObj := method.Object.Type().Object
	quadrupleFn := quadrupleObj.Methods[method.Idx]
	llvmFn := c.fns[quadrupleFn]
	fn := c.currBlock.NewBitCast(wrongTypeFn, llvmFn.Type())

	return fn
}

type phiToFill struct {
	ephi *quadruple.EPhi
	phi  *ir.InstPhi
}

func (c *compiler) createDeclarations() {
	for _, obj := range c.program.Objects {
		c.declareObjStubs(obj)
	}

	for _, fn := range c.program.Fns {
		c.declareFn(fn)
	}
	for _, obj := range c.program.Objects {
		c.declareObj(obj)
	}
	c.convertStrings(c.program.Strings)
}

func (c *compiler) declareFn(fn *quadruple.Fn) *ir.Func {
	if f, ok := c.fns[fn]; ok {
		return f
	}

	retType := c.convertType(fn.RetType())
	params := make([]*ir.Param, 0, len(fn.Args))
	for _, a := range fn.Args {
		params = append(params, &ir.Param{
			LocalIdent: ir.NewLocalIdent(a.Name),
			Typ:        c.convertType(a.Type),
		})
	}

	f := c.module.NewFunc(fn.Name, retType, params...)
	c.fns[fn] = f

	return f
}

func (c *compiler) convertStrings(strs map[string]struct{}) {
	for s := range strs {
		irS := constant.NewCharArrayFromString(s + "\000")
		name := fmt.Sprintf(".string_%s", s)
		glob := ir.NewGlobalDef(name, irS)

		c.module.Globals = append(c.module.Globals, glob)
		c.strs[s] = glob
	}
}

func (c *compiler) declareObj(obj *quadruple.Object) {
	vtable := c.createVTable(obj.Methods)
	vtableDef := c.module.NewGlobalDef(fmt.Sprintf(".%s__vtable", obj.Name), &vtable)
	c.vtables[obj.Name] = vtableDef

	convertedFields := make([]types.Type, 0, len(obj.Attributes))
	convertedFields = append(convertedFields, vtableDef.Type())
	for _, f := range obj.Attributes[1:] {
		convertedFields = append(convertedFields, c.convertType(f))
	}

	structType := c.types[obj.Name]
	structType.Fields = convertedFields

	c.module.NewTypeDef(obj.Name, structType)
}

func (c *compiler) declareObjStubs(obj *quadruple.Object) {
	c.types[obj.Name] = &types.StructType{}
}

func (c *compiler) convertType(f quadruple.Type) types.Type {
	if f.Array {
		f.Array = false
		t := c.convertType(f)
		return &types.PointerType{ElemType: t}
	}
	if f.Pointer {
		f.Pointer = false
		t := c.convertType(f)
		return &types.PointerType{ElemType: t}
	}
	if f.Object != nil {
		return &types.PointerType{ElemType: c.types[f.Object.Name]}
	}
	return convertTypeEnum(f.Type)
}

func convertTypeEnum(t quadruple.TypeEnum) types.Type {
	switch t {
	case quadruple.VoidType:
		return &types.VoidType{}
	case quadruple.BoolType:
		return boolType
	case quadruple.StringType:
		return &types.PointerType{
			ElemType: charType,
		}
	case quadruple.IntType:
		return intType
	default:
		panic("IMPOSSIBLE")
	}
}

func (c *compiler) createVTable(methods []*quadruple.Fn) constant.Array {
	ret := constant.Array{
		Typ: &types.ArrayType{
			Len:      uint64(len(methods)),
			ElemType: vFnType,
		},
		Elems: make([]constant.Constant, 0, len(methods)),
	}

	for _, m := range methods {
		fn := c.declareFn(m)

		ret.Elems = append(ret.Elems, constant.NewBitCast(fn, vFnType))
	}

	return ret
}

func (c *compiler) fillFunctions() {
	for fnP, fn := range c.fns {
		c.fillFunction(fnP, fn)
	}
}

func (c *compiler) fillFunction(p *quadruple.Fn, fn *ir.Func) {
	c.phisToFill = nil
	c.declareBlocks(p, fn)
	c.fillBlocks(p)
	c.fillPhis()
}

func (c *compiler) declareBlocks(p *quadruple.Fn, fn *ir.Func) {
	c.blocksMap = make(map[*quadruple.Block]*ir.Block, len(p.Blocks))
	for _, b := range p.Blocks {
		bl := fn.NewBlock(fmt.Sprintf(".block_%d", b.Id))
		c.blocksMap[b] = bl
	}
}

func (c *compiler) fillBlocks(p *quadruple.Fn) {
	c.varMap = make(map[quadruple.Var]value.Value)

	f := c.fns[p]
	for i, param := range f.Params {
		val := quadruple.Var{
			Type_: p.Args[i].Type,
			Id:    p.Args[i].Name,
		}
		c.varMap[val] = param
	}

	for _, b := range p.Blocks {
		c.VisitBlock(b)
	}
}

func (c *compiler) replaceVal(val quadruple.Val) value.Value {
	switch t := val.(type) {
	case quadruple.Var:
		return c.varMap[t]
	case quadruple.IntVal:
		return constant.NewInt(intType, t.Int)
	case quadruple.BoolVal:
		if t.Bool {
			return constant.NewInt(boolType, 1)
		}
		return constant.NewInt(boolType, 0)
	case quadruple.StrVal:
		return constant.NewBitCast(c.strs[t.String], &types.PointerType{ElemType: charType})
	case quadruple.NullVal:
		return constant.NewNull(c.convertType(t.Type_).(*types.PointerType))
	default:
		panic("IMPOSSIBLE")
	}
}

func (c *compiler) VisitEUnOp(e *quadruple.EUnOp) any {
	v := c.replaceVal(e.Val)
	switch e.Op {
	case quadruple.LogicNot:
		return c.currBlock.NewAdd(v, constant.NewInt(boolType, 1))
	default:
		panic("IMPOSSIBLE")
	}
}

func (c *compiler) VisitECast(e *quadruple.ECast) any {
	v := c.replaceVal(e.Val)

	return c.currBlock.NewBitCast(v, c.convertType(e.To))
}

func (c *compiler) VisitEBinOp(e *quadruple.EBinOp) any {
	lv := c.replaceVal(e.LVal)
	rv := c.replaceVal(e.RVal)

	switch e.Op {
	case quadruple.Gt:
		return c.currBlock.NewICmp(enum.IPredSGT, lv, rv)
	case quadruple.Ge:
		return c.currBlock.NewICmp(enum.IPredSGE, lv, rv)
	case quadruple.Ne:
		return c.currBlock.NewICmp(enum.IPredNE, lv, rv)
	case quadruple.Eq:
		return c.currBlock.NewICmp(enum.IPredEQ, lv, rv)
	case quadruple.Times:
		return c.currBlock.NewMul(lv, rv)
	case quadruple.Divide:
		return c.currBlock.NewSDiv(lv, rv)
	case quadruple.Modulo:
		return c.currBlock.NewSRem(lv, rv)
	case quadruple.Plus:
		return c.currBlock.NewAdd(lv, rv)
	case quadruple.Minus:
		return c.currBlock.NewSub(lv, rv)
	default:
		panic("IMPOSSIBLE")
	}
}

func (c *compiler) VisitEVal(e *quadruple.EVal) any {
	panic("IMPOSSIBLE")
}

func (c *compiler) VisitECall(e *quadruple.ECall) any {
	fn := e.Callable.Accept(c).(value.Value)

	args := make([]value.Value, 0, len(e.Params))
	for _, arg := range e.Params {
		args = append(args, c.replaceVal(arg))
	}

	return c.fnCall(fn, fn.Type().(*types.PointerType).ElemType.(*types.FuncType), args...)
}

func (c *compiler) VisitEGetAttrPtr(e *quadruple.EGetAttrPtr) any {
	v := c.replaceVal(e.Obj)
	llvmT := v.Type().(*types.PointerType).ElemType

	return c.currBlock.NewGetElementPtr(
		llvmT,
		v,
		constant.NewInt(&types.IntType{BitSize: 64}, 0),
		constant.NewInt(&types.IntType{BitSize: 32}, int64(e.AttrIdx)),
	)
}

func (c *compiler) VisitEGetArrLen(e *quadruple.EGetArrLen) any {
	bc := c.currBlock.NewBitCast(c.replaceVal(e.Arr), &types.PointerType{ElemType: intType})
	elPtr := c.currBlock.NewGetElementPtr(intType, bc, constant.NewInt(intType, -1))
	return c.currBlock.NewLoad(intType, elPtr)
}

func (c *compiler) VisitEGetArrElPtr(e *quadruple.EGetArrElPtr) any {
	arr := c.replaceVal(e.Arr)
	idx := c.replaceVal(e.Idx)

	elType := arr.Type().(*types.PointerType).ElemType
	return c.currBlock.NewGetElementPtr(elType, arr, idx)
}

func (c *compiler) VisitELoadPtr(e *quadruple.ELoadPtr) any {
	llvmT := c.convertType(e.Ptr.Type()).(*types.PointerType).ElemType

	return c.currBlock.NewLoad(llvmT, c.replaceVal(e.Ptr))
}

func (c *compiler) VisitENewArray(e *quadruple.ENewArray) any {
	typ := c.convertType(e.Type)

	elSize := c.sizeOf(typ)
	size := c.replaceVal(e.Size)
	calloc := c.fns[c.program.Fns[".new_arr"]]
	arr := c.currBlock.NewCall(calloc, size, elSize)

	return c.currBlock.NewBitCast(arr, &types.PointerType{ElemType: typ})
}

func (c *compiler) VisitENewObj(e *quadruple.ENewObj) any {
	typ := c.types[e.Obj.Name]

	size := c.sizeOf(typ)
	calloc := c.fns[c.program.Fns["calloc"]]
	obj := c.currBlock.NewCall(calloc, constant.NewInt(intType, 1), size)
	objCorrectType := c.currBlock.NewBitCast(obj, &types.PointerType{ElemType: typ})
	vtablePtr := c.currBlock.NewGetElementPtr(
		typ,
		objCorrectType,
		constant.NewInt(&types.IntType{BitSize: 64}, 0),
		constant.NewInt(&types.IntType{BitSize: 32}, 0),
	)
	c.currBlock.NewStore(c.vtables[e.Obj.Name], vtablePtr)

	return objCorrectType
}

func (c *compiler) sizeOf(typ types.Type) value.Value {
	null := constant.NewNull(&types.PointerType{ElemType: typ})
	sizePtr := c.currBlock.NewGetElementPtr(typ, null, constant.NewInt(&types.IntType{BitSize: 32}, 1))
	return c.currBlock.NewPtrToInt(sizePtr, intType)
}

func (c *compiler) VisitEPhi(e *quadruple.EPhi) any {
	inc := make([]*ir.Incoming, 0, len(e.Vals))
	toFill := false
	for _, v := range e.Vals {
		newV := c.replaceVal(v.Val)
		if newV == nil {
			toFill = true
		} else {
			inc = append(inc, &ir.Incoming{
				X:    newV,
				Pred: c.blocksMap[v.Block],
			})
		}
	}

	phi := c.currBlock.NewPhi(inc...)
	if toFill {
		c.phisToFill = append(c.phisToFill, phiToFill{
			ephi: e,
			phi:  phi,
		})
	}
	return phi
}

type blockCtx struct {
	prev            *blockCtx
	block           *quadruple.Block
	valReplacements map[quadruple.Val]quadruple.Val
}

func (c *compiler) VisitBlock(b *quadruple.Block) any {
	c.currBlock = c.blocksMap[b]

	for _, ins := range b.Instructions {
		ins.Accept(c)
	}

	return nil
}

func (c *compiler) VisitUnreachable(u *quadruple.Unreachable) any {
	return c.currBlock.NewUnreachable()
}

func (c *compiler) VisitGoto(g *quadruple.Goto) any {
	return c.currBlock.NewBr(c.blocksMap[g.Block])
}

func (c *compiler) VisitPointerAssignment(pa *quadruple.PointerAssignment) any {
	return c.currBlock.NewStore(c.replaceVal(pa.Val), c.replaceVal(pa.Pointer))
}

func (c *compiler) VisitAssignment(a *quadruple.Assignment) any {
	inst := a.Expr.Accept(c).(value.Value)
	c.varMap[a.Var] = inst

	return inst
}

func (c *compiler) VisitReturn(r *quadruple.Return) any {
	if r.Val == nil {
		return c.currBlock.NewRet(nil)
	}
	return c.currBlock.NewRet(c.replaceVal(*r.Val))
}

func (c *compiler) VisitVoidCall(fc *quadruple.VoidCall) any {
	fn := fc.Callable.Accept(c).(value.Value)

	args := make([]value.Value, 0, len(fc.Params))
	for _, arg := range fc.Params {
		args = append(args, c.replaceVal(arg))
	}

	return c.fnCall(fn, fn.Type().(*types.PointerType).ElemType.(*types.FuncType), args...)
}

func (c *compiler) VisitCondJump(cj *quadruple.CondJump) any {
	return c.currBlock.NewCondBr(c.replaceVal(cj.Cond), c.blocksMap[cj.TrueBlock], c.blocksMap[cj.FalseBlock])
}

func (c *compiler) fillPhis() {
	for _, phi := range c.phisToFill {
		inc := make([]*ir.Incoming, 0, len(phi.ephi.Vals))
		for _, v := range phi.ephi.Vals {
			newV := c.replaceVal(v.Val)
			inc = append(inc, &ir.Incoming{
				X:    newV,
				Pred: c.blocksMap[v.Block],
			})
		}
		phi.phi.Incs = inc
	}
}

func (c *compiler) fnCall(fn value.Value, fnT *types.FuncType, args ...value.Value) *ir.InstCall {
	newArgs := make([]value.Value, 0, len(args))
	for i, arg := range args {
		paramT := fnT.Params[i]
		if arg.Type().Equal(paramT) {
			newArgs = append(newArgs, arg)
		} else {
			newArgs = append(newArgs, c.currBlock.NewBitCast(arg, paramT))
		}
	}

	return c.currBlock.NewCall(fn, newArgs...)
}

var _ quadruple.InstructionVisitor = &compiler{}
var _ quadruple.ExprVisitor = &compiler{}
var _ quadruple.BlockVisitor = &compiler{}
