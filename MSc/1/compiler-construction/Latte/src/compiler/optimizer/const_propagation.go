package optimizer

import (
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
)

func ConstPropagation(program quadruple.Program) quadruple.Program {
	c := newConstPropagator()

	for _, fn := range program.Fns {
		c.convertFn(fn)
	}
	for _, obj := range program.Objects {
		c.convertObj(obj)
	}
	return program
}

func (c *constPropagator) convertFn(fn *quadruple.Fn) {
	c.hasChanged = true
	c.replacements = make(map[quadruple.Val]quadruple.Val)
	c.removed = make(map[quadruple.Val]struct{})

	deadVariablesRemoved := true
	for c.hasChanged || deadVariablesRemoved {
		c.hasChanged = false
		c.usedVariables = make(map[string]struct{})

		for _, block := range fn.Blocks {
			c.VisitBlock(block)
		}

		deadVariablesRemoved = removeDeadVariables(c.usedVariables, fn)
	}
}

func (c *constPropagator) convertObj(obj *quadruple.Object) {
	for _, method := range obj.Methods {
		c.convertFn(method)
	}
}

func newConstPropagator() *constPropagator {
	return &constPropagator{}
}

type constPropagator struct {
	replacements map[quadruple.Val]quadruple.Val
	removed      map[quadruple.Val]struct{}
	hasChanged   bool

	currVar *quadruple.Var

	usedVariables map[string]struct{}
}

func (c *constPropagator) VisitFn(fn *quadruple.Fn) any {
	return fn
}

func (c *constPropagator) VisitMethod(method *quadruple.Method) any {
	return &quadruple.Method{
		Object: c.replaceVal(method.Object),
		Idx:    method.Idx,
	}
}

func (c *constPropagator) VisitECast(e *quadruple.ECast) any {
	val := c.replaceVal(e.Val)
	if val.Type().Type == quadruple.UntypedNullType {
		c.hasChanged = true
		return &quadruple.EVal{
			Val: quadruple.NullVal{Type_: e.To},
		}
	}
	if val.Type().Object == e.To.Object {
		c.hasChanged = true
		return &quadruple.EVal{Val: val}
	}
	return &quadruple.ECast{Val: val, To: e.To}
}

func (c *constPropagator) VisitEUnOp(e *quadruple.EUnOp) any {
	switch e.Op {
	case quadruple.LogicNot:
		if b, ok := e.Val.(quadruple.BoolVal); ok {
			c.hasChanged = true
			return &quadruple.EVal{Val: quadruple.BoolVal{Bool: !b.Bool}}
		}
	}

	return &quadruple.EUnOp{
		Val: c.replaceVal(e.Val),
		Op:  e.Op,
	}
}

func (c *constPropagator) VisitEVal(e *quadruple.EVal) any {
	return &quadruple.EVal{Val: c.replaceVal(e.Val)}
}

func (c *constPropagator) VisitEBinOp(e *quadruple.EBinOp) any {
	lVal := c.replaceVal(e.LVal)
	rVal := c.replaceVal(e.RVal)

	if lValInt, islValInt := lVal.(quadruple.IntVal); islValInt {
		if rValInt, isrValInt := rVal.(quadruple.IntVal); isrValInt {
			var val quadruple.Val
			switch e.Op {
			case quadruple.Gt:
				val = quadruple.BoolVal{Bool: lValInt.Int < rValInt.Int}
			case quadruple.Ge:
				val = quadruple.BoolVal{Bool: lValInt.Int <= rValInt.Int}
			case quadruple.Ne:
				val = quadruple.BoolVal{Bool: lValInt.Int != rValInt.Int}
			case quadruple.Eq:
				val = quadruple.BoolVal{Bool: lValInt.Int == rValInt.Int}
			case quadruple.Times:
				val = quadruple.IntVal{Int: lValInt.Int * rValInt.Int}
			case quadruple.Divide:
				val = quadruple.IntVal{Int: lValInt.Int / rValInt.Int}
			case quadruple.Modulo:
				val = quadruple.IntVal{Int: lValInt.Int % rValInt.Int}
			case quadruple.Plus:
				val = quadruple.IntVal{Int: lValInt.Int + rValInt.Int}
			case quadruple.Minus:
				val = quadruple.IntVal{Int: lValInt.Int - rValInt.Int}
			default:
				panic("IMPOSSIBLE")
			}
			c.hasChanged = true
			return &quadruple.EVal{Val: val}
		}
	}

	if lValBool, islValBool := lVal.(quadruple.BoolVal); islValBool {
		if rValBool, isrValBool := rVal.(quadruple.BoolVal); isrValBool {
			var val quadruple.Val
			switch e.Op {
			case quadruple.Ne:
				val = quadruple.BoolVal{Bool: lValBool.Bool != rValBool.Bool}
			case quadruple.Eq:
				val = quadruple.BoolVal{Bool: lValBool.Bool == rValBool.Bool}
			default:
				panic("IMPOSSIBLE")
			}
			c.hasChanged = true
			return &quadruple.EVal{Val: val}
		}
	}

	if lValStr, islValStr := lVal.(quadruple.StrVal); islValStr {
		if rValStr, isrValStr := rVal.(quadruple.StrVal); isrValStr {
			switch e.Op {
			case quadruple.Ne:
				c.hasChanged = true
				return &quadruple.EVal{Val: quadruple.BoolVal{Bool: lValStr.String != rValStr.String}}
			case quadruple.Eq:
				c.hasChanged = true
				return &quadruple.EVal{Val: quadruple.BoolVal{Bool: lValStr.String == rValStr.String}}
			case quadruple.Plus:
				// Nothing to do here.
			default:
				panic("IMPOSSIBLE")
			}
		}
	}

	return &quadruple.EBinOp{
		LVal: lVal,
		RVal: rVal,
		Op:   e.Op,
	}
}

func (c *constPropagator) VisitECall(e *quadruple.ECall) any {
	args := make([]quadruple.Val, 0, len(e.Params))
	for _, arg := range e.Params {
		args = append(args, c.replaceVal(arg))
	}

	return &quadruple.ECall{
		Callable: e.Callable.Accept(c).(quadruple.Callable),
		Params:   args,
	}
}

func (c *constPropagator) VisitEGetAttrPtr(e *quadruple.EGetAttrPtr) any {
	return &quadruple.EGetAttrPtr{
		Obj:     c.replaceVal(e.Obj),
		AttrIdx: e.AttrIdx,
	}
}

func (c *constPropagator) VisitEGetArrElPtr(e *quadruple.EGetArrElPtr) any {
	return &quadruple.EGetArrElPtr{
		Arr: c.replaceVal(e.Arr),
		Idx: c.replaceVal(e.Idx),
	}
}

func (c *constPropagator) VisitEGetArrLen(e *quadruple.EGetArrLen) any {
	return &quadruple.EGetArrLen{
		Arr: c.replaceVal(e.Arr),
	}
}

func (c *constPropagator) VisitELoadPtr(e *quadruple.ELoadPtr) any {
	return &quadruple.ELoadPtr{
		Ptr: c.replaceVal(e.Ptr),
	}
}

func (c *constPropagator) VisitENewArray(e *quadruple.ENewArray) any {
	return &quadruple.ENewArray{
		Size: c.replaceVal(e.Size),
		Type: e.Type,
	}
}

func (c *constPropagator) VisitENewObj(e *quadruple.ENewObj) any {
	return &quadruple.ENewObj{
		Obj: e.Obj,
	}
}

func (c *constPropagator) VisitEPhi(e *quadruple.EPhi) any {
	if len(e.Vals) == 0 {
		return nil
	}

	vals := make([]quadruple.PhiValData, 0, len(e.Vals))
	for _, val := range e.Vals {
		newV := c.replaceVal(val.Val)
		if newV == nil {
			return nil
		}
		vals = append(vals, quadruple.PhiValData{
			Block: val.Block,
			Val:   newV,
		})
	}

	phi := &quadruple.EPhi{
		Vals: vals,
	}

	allTheSame := true
	for _, v := range phi.Vals {
		if v.Val != phi.Vals[0].Val {
			allTheSame = false
			break
		}
	}
	if allTheSame {
		c.hasChanged = true
		return &quadruple.EVal{Val: phi.Vals[0].Val}
	}

	// Change all but one same as current variable
	var diff quadruple.Val
	for _, v := range phi.Vals {
		if v.Val != *c.currVar {
			if diff == nil {
				diff = v.Val
			} else {
				diff = nil
				break
			}
		}
	}
	if diff != nil {
		return &quadruple.EVal{Val: diff}
	}

	return phi
}

func (c *constPropagator) VisitUnreachable(u *quadruple.Unreachable) any {
	return &quadruple.Unreachable{}
}

func (c *constPropagator) VisitBlock(b *quadruple.Block) any {
	instructions := make([]quadruple.Instruction, 0)

	for _, ex := range b.Instructions {
		inst := ex.Accept(c)
		if inst != nil {
			instructions = append(instructions, inst.(quadruple.Instruction))
		}
	}
	b.Instructions = instructions

	return nil
}

func (c *constPropagator) VisitGoto(g *quadruple.Goto) any {
	return &quadruple.Goto{
		Block: g.Block,
	}
}

func (c *constPropagator) VisitPointerAssignment(pa *quadruple.PointerAssignment) any {
	return &quadruple.PointerAssignment{
		Pointer: c.replaceVal(pa.Pointer),
		Val:     c.replaceVal(pa.Val),
	}
}

func (c *constPropagator) VisitAssignment(a *quadruple.Assignment) any {
	_, wasUsed := c.usedVariables[a.Var.Id]

	c.currVar = &a.Var
	defer func() {
		c.currVar = nil
	}()

	ex := a.Expr.Accept(c)
	if ex == nil {
		c.hasChanged = true
		c.removed[a.Var] = struct{}{}
		return nil
	}
	expr := ex.(quadruple.Expr)
	if val, isVal := expr.(*quadruple.EVal); isVal {
		c.replacements[a.Var] = c.replaceVal(val.Val)
		c.hasChanged = true

		return nil
	}
	if !wasUsed {
		if _, ok := expr.(*quadruple.EPhi); ok {
			delete(c.usedVariables, a.Var.Id)
		}
	}

	return &quadruple.Assignment{
		Var:  a.Var,
		Expr: expr,
	}
}

func (c *constPropagator) VisitReturn(r *quadruple.Return) any {
	val := r.Val
	if val != nil {
		updated := c.replaceVal(*val)
		val = &updated
	}

	return &quadruple.Return{
		Val: val,
	}
}

func (c *constPropagator) VisitVoidCall(fc *quadruple.VoidCall) any {
	args := make([]quadruple.Val, 0, len(fc.Params))
	for _, arg := range fc.Params {
		args = append(args, c.replaceVal(arg))
	}

	return &quadruple.VoidCall{
		Callable: fc.Callable.Accept(c).(quadruple.Callable),
		Params:   args,
	}
}

func (c *constPropagator) VisitCondJump(cj *quadruple.CondJump) any {
	return &quadruple.CondJump{
		Cond:       c.replaceVal(cj.Cond),
		TrueBlock:  cj.TrueBlock,
		FalseBlock: cj.FalseBlock,
	}
}

func (c *constPropagator) replaceVal(val quadruple.Val) quadruple.Val {
	if _, removed := c.removed[val]; removed {
		return nil
	}
	if replacement, ok := c.replacements[val]; ok {
		if v, ok := replacement.(quadruple.Var); ok {
			c.usedVariables[v.Id] = struct{}{}
		}
		return replacement
	}
	if v, ok := val.(quadruple.Var); ok {
		c.usedVariables[v.Id] = struct{}{}
	}
	return val
}

var _ quadruple.InstructionVisitor = &constPropagator{}
var _ quadruple.ExprVisitor = &constPropagator{}
var _ quadruple.BlockVisitor = &constPropagator{}
