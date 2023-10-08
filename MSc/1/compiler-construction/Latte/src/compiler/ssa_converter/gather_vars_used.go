package ssa_converter

import (
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
	"golang.org/x/exp/maps"
)

func gatherVarsUsed(b *quadruple.Block) []quadruple.Var {
	return newVarsGatherer().VisitBlock(b).([]quadruple.Var)
}

func newVarsGatherer() *varsGatherer {
	return &varsGatherer{
		varsUsed: make(map[quadruple.Var]struct{}),
	}
}

type varsGatherer struct {
	varsUsed map[quadruple.Var]struct{}
}

func (c *varsGatherer) VisitFn(fn *quadruple.Fn) any {
	return nil
}

func (c *varsGatherer) VisitMethod(method *quadruple.Method) any {
	c.gatherVal(method.Object)

	return nil
}

func (c *varsGatherer) VisitECast(e *quadruple.ECast) any {
	c.gatherVal(e.Val)

	return nil
}

func (c *varsGatherer) VisitEUnOp(e *quadruple.EUnOp) any {
	c.gatherVal(e.Val)

	return nil
}

func (c *varsGatherer) VisitEVal(e *quadruple.EVal) any {
	c.gatherVal(e.Val)

	return nil
}

func (c *varsGatherer) VisitEBinOp(e *quadruple.EBinOp) any {
	c.gatherVal(e.RVal)
	c.gatherVal(e.LVal)

	return nil
}

func (c *varsGatherer) VisitECall(e *quadruple.ECall) any {
	e.Callable.Accept(c)
	for _, param := range e.Params {
		c.gatherVal(param)
	}

	return nil
}

func (c *varsGatherer) VisitEGetAttrPtr(e *quadruple.EGetAttrPtr) any {
	c.gatherVal(e.Obj)

	return nil
}

func (c *varsGatherer) VisitEGetArrElPtr(e *quadruple.EGetArrElPtr) any {
	c.gatherVal(e.Idx)
	c.gatherVal(e.Arr)

	return nil
}

func (c *varsGatherer) VisitEGetArrLen(e *quadruple.EGetArrLen) any {
	c.gatherVal(e.Arr)

	return nil
}

func (c *varsGatherer) VisitELoadPtr(e *quadruple.ELoadPtr) any {
	c.gatherVal(e.Ptr)

	return nil
}

func (c *varsGatherer) VisitENewArray(e *quadruple.ENewArray) any {
	c.gatherVal(e.Size)

	return nil
}

func (c *varsGatherer) VisitENewObj(e *quadruple.ENewObj) any {
	return nil
}

func (c *varsGatherer) VisitEPhi(e *quadruple.EPhi) any {
	panic("VARS GATHERER NEVER ENCOUNTERS PHI")
}

func (c *varsGatherer) VisitBlock(b *quadruple.Block) any {
	for _, inst := range b.Instructions {
		inst.Accept(c)
	}

	return maps.Keys(c.varsUsed)
}

func (c *varsGatherer) VisitUnreachable(u *quadruple.Unreachable) any {
	return nil
}

func (c *varsGatherer) VisitGoto(g *quadruple.Goto) any {
	return nil
}

func (c *varsGatherer) VisitPointerAssignment(pa *quadruple.PointerAssignment) any {
	c.gatherVal(pa.Pointer)
	c.gatherVal(pa.Val)

	return nil
}

func (c *varsGatherer) VisitAssignment(a *quadruple.Assignment) any {
	a.Expr.Accept(c)
	c.gatherVar(a.Var)

	return nil
}

func (c *varsGatherer) VisitReturn(r *quadruple.Return) any {
	if r.Val != nil {
		c.gatherVal(*r.Val)
	}

	return nil
}

func (c *varsGatherer) VisitVoidCall(fc *quadruple.VoidCall) any {
	fc.Callable.Accept(c)
	for _, arg := range fc.Params {
		c.gatherVal(arg)
	}

	return nil
}

func (c *varsGatherer) VisitCondJump(cj *quadruple.CondJump) any {
	c.gatherVal(cj.Cond)

	return nil
}

func (c *varsGatherer) gatherVar(var_ quadruple.Var) {
	c.varsUsed[var_] = struct{}{}
}

func (c *varsGatherer) gatherVal(val quadruple.Val) {
	var_, isVar := val.(quadruple.Var)
	if isVar {
		c.gatherVar(var_)
	}
}

var _ quadruple.InstructionVisitor = &varsGatherer{}
var _ quadruple.ExprVisitor = &varsGatherer{}
var _ quadruple.BlockVisitor = &varsGatherer{}
