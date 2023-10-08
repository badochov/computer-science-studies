package optimizer

import (
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
)

func removeDeadVariables(usedVariables map[string]struct{}, fn *quadruple.Fn) bool {
	r := newDeadVariableRemover(usedVariables)

	removed := false
	for _, b := range fn.Blocks {
		removed = removed || r.VisitBlock(b).(bool)
	}

	return removed
}

func newDeadVariableRemover(used map[string]struct{}) *deadVariableRemover {
	return &deadVariableRemover{
		usedVariables: used,
	}
}

type deadVariableRemover struct {
	usedVariables map[string]struct{}
}

func (c *deadVariableRemover) shouldRemove(id string) bool {
	_, ok := c.usedVariables[id]
	return !ok
}

func (c *deadVariableRemover) VisitUnreachable(u *quadruple.Unreachable) any {
	return u
}

func (c *deadVariableRemover) VisitBlock(b *quadruple.Block) any {
	instructions := make([]quadruple.Instruction, 0)
	removedVariable := false

	for _, ex := range b.Instructions {
		inst := ex.Accept(c)
		if inst == nil {
			removedVariable = true
		} else {
			instructions = append(instructions, inst.(quadruple.Instruction))
		}
	}
	b.Instructions = instructions

	return removedVariable
}

func (c *deadVariableRemover) VisitGoto(g *quadruple.Goto) any {
	return g
}

func (c *deadVariableRemover) VisitPointerAssignment(pa *quadruple.PointerAssignment) any {
	return pa
}

func (c *deadVariableRemover) VisitAssignment(a *quadruple.Assignment) any {
	if c.shouldRemove(a.Var.Id) {
		if call, ok := a.Expr.(*quadruple.ECall); ok {
			return &quadruple.VoidCall{
				Callable: call.Callable,
				Params:   call.Params,
			}
		}
		return nil
	}

	return a
}

func (c *deadVariableRemover) VisitReturn(r *quadruple.Return) any {
	return r
}

func (c *deadVariableRemover) VisitVoidCall(fc *quadruple.VoidCall) any {
	return fc
}

func (c *deadVariableRemover) VisitCondJump(cj *quadruple.CondJump) any {
	return cj
}

var _ quadruple.InstructionVisitor = &deadVariableRemover{}
var _ quadruple.BlockVisitor = &deadVariableRemover{}
