package ast

type IProgram[Data any] interface {
	VisitProgram(visitor ProgramVisitor[Data]) any
}

type ProgramVisitor[Data any] interface {
	VisitProgram(*Program[Data]) any
}

type Program[Data any] struct {
	Data Data

	TopDefs []ITopDef[Data]
}

func (p *Program[Data]) VisitProgram(visitor ProgramVisitor[Data]) any {
	return visitor.VisitProgram(p)
}

var _ IProgram[any] = &Program[any]{}
