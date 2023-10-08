package block_termination_fixer

import (
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
)

// FixBlockTerminations either removes redundant instructions after first termination of a block.
func FixBlockTerminations(program quadruple.Program) quadruple.Program {
	return fixMoreThanOneTermination(program)
}

func fixMoreThanOneTermination(program quadruple.Program) quadruple.Program {
	c := newConverter()

	for _, fn := range program.Fns {
		c.fixFn(fn)
	}
	for _, obj := range program.Objects {
		c.fixObj(obj)
	}
	return program
}

func (c *converter) fixFn(fn *quadruple.Fn) {
	if len(fn.Blocks) == 0 {
		return
	}
	c.VisitBlock(fn.Blocks[0])
}

func (c *converter) fixObj(obj *quadruple.Object) {
	for _, method := range obj.Methods {
		c.fixFn(method)
	}
}

func newConverter() *converter {
	return &converter{}
}

type converter struct{}

func (c *converter) VisitUnreachable(u *quadruple.Unreachable) any {
	return true
}

func (c *converter) VisitBlock(b *quadruple.Block) any {
	for i, ex := range b.Instructions {
		if ex.Accept(c).(bool) {
			b.Instructions = b.Instructions[:i+1]
			break
		}
	}

	return b
}

func (c *converter) VisitGoto(g *quadruple.Goto) any {
	return true
}

func (c *converter) VisitPointerAssignment(pa *quadruple.PointerAssignment) any {
	return false
}

func (c *converter) VisitAssignment(a *quadruple.Assignment) any {
	return false
}

func (c *converter) VisitReturn(r *quadruple.Return) any {
	return true
}

func (c *converter) VisitVoidCall(fc *quadruple.VoidCall) any {
	return false
}

func (c *converter) VisitCondJump(cj *quadruple.CondJump) any {
	return true
}

var _ quadruple.InstructionVisitor = &converter{}
var _ quadruple.BlockVisitor = &converter{}
