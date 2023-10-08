package ast

type IArg[Data any] interface {
	VisitArg(visitor ArgVisitor[Data]) any
}

type ArgVisitor[Data any] interface {
	VisitArg(*Arg[Data]) any
}

type Arg[Data any] struct {
	Data Data

	Type IType[Data]
	Ref  bool
	Id   string
}

func (a *Arg[Data]) VisitArg(visitor ArgVisitor[Data]) any {
	return visitor.VisitArg(a)
}

var _ IArg[any] = &Arg[any]{}
