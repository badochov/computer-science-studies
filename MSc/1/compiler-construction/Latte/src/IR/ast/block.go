package ast

type IBlock[Data any] interface {
	VisitBlock(visitor BlockVisitor[Data]) any
}

type BlockVisitor[Data any] interface {
	VisitBlock(*Block[Data]) any
}

type Block[Data any] struct {
	Data Data

	Stmts []IStmt[Data]
}

func (b *Block[Data]) VisitBlock(visitor BlockVisitor[Data]) any {
	return visitor.VisitBlock(b)
}

var _ IBlock[any] = &Block[any]{}
