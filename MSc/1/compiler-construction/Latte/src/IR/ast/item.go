package ast

type IItem[Data any] interface {
	VisitItem(visitor ItemVisitor[Data]) any
}

type ItemVisitor[Data any] interface {
	VisitItemAss(i *ItemAss[Data]) any
}

type ItemAss[Data any] struct {
	Data Data

	Id   string
	Expr IExpr[Data]
}

func (i *ItemAss[Data]) VisitItem(visitor ItemVisitor[Data]) any {
	return visitor.VisitItemAss(i)
}

var _ IItem[any] = &ItemAss[any]{}
