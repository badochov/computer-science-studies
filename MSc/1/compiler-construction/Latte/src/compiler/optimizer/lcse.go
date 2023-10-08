package optimizer

import (
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
)

func LCSE(program quadruple.Program) quadruple.Program {
	c := newLCSE()

	for _, fn := range program.Fns {
		c.convertFn(fn)
	}
	for _, obj := range program.Objects {
		c.convertObj(obj)
	}
	return ConstPropagation(program)
}

type lcse struct {
	expressions map[quadruple.SerializedExpression]expr

	currBlock      *quadruple.Block
	pointerCounter int
}

func (l *lcse) VisitGoto(g *quadruple.Goto) any {
	return g
}

func (l *lcse) VisitUnreachable(u *quadruple.Unreachable) any {
	return u
}

func (l *lcse) VisitPointerAssignment(pa *quadruple.PointerAssignment) any {
	l.invalidatePointersIfNeeded(pa.Pointer)

	return pa
}

func (l *lcse) VisitAssignment(a *quadruple.Assignment) any {
	serialized := a.Expr.Serialize()
	if prev, ok := l.expressions[serialized]; ok {
		ret := &quadruple.Assignment{Var: a.Var, Expr: &quadruple.EVal{Val: prev.ass.Var}} // To be removed by const propagator.
		switch a.Expr.(type) {
		case *quadruple.ELoadPtr:
			if l.pointerExprValid(prev) {
				return ret
			}
		default:
			return ret
		}
	}

	switch t := a.Expr.(type) {
	case *quadruple.ECall:
		l.invalidatePointersIfNeeded(t.Params...)
	}

	switch a.Expr.(type) {
	case *quadruple.ENewObj, *quadruple.ENewArray, *quadruple.ECall:
	// Do nothing
	default:
		l.expressions[a.Expr.Serialize()] = expr{
			ass:            a,
			block:          l.currBlock,
			pointerCounter: l.pointerCounter,
		}
	}

	return a
}

func (l *lcse) VisitReturn(r *quadruple.Return) any {
	return r
}

func (l *lcse) VisitVoidCall(fc *quadruple.VoidCall) any {
	l.invalidatePointersIfNeeded(fc.Params...)

	return fc
}

func (l *lcse) VisitCondJump(cj *quadruple.CondJump) any {
	return cj
}

type expr struct {
	ass            *quadruple.Assignment
	block          *quadruple.Block
	pointerCounter int
}

func newLCSE() *lcse {
	return &lcse{}
}

func (l *lcse) convertFn(fn *quadruple.Fn) {
	for _, block := range fn.Blocks {
		l.VisitBlock(block)
	}
}

func (l *lcse) convertObj(obj *quadruple.Object) {
	for _, method := range obj.Methods {
		l.convertFn(method)
	}
}

func (l *lcse) VisitBlock(block *quadruple.Block) {
	l.expressions = make(map[quadruple.SerializedExpression]expr)
	l.currBlock = block
	l.pointerCounter = 0

	instructions := make([]quadruple.Instruction, 0, len(block.Instructions))
	for _, inst := range block.Instructions {
		newInst := inst.Accept(l)
		if newInst != nil {
			instructions = append(instructions, newInst.(quadruple.Instruction))
		}
	}

	block.Instructions = instructions
}

func (l *lcse) invalidatePointersIfNeeded(params ...quadruple.Val) {
	for _, p := range params {
		typ := p.Type()
		if typ.Array || typ.Object != nil || typ.Pointer {
			l.pointerCounter++
			return
		}
	}
}

func (l *lcse) pointerExprValid(prev expr) bool {
	return prev.pointerCounter == l.pointerCounter
}
