package ast

type Visitor[Data any] interface {
	TypeVisitor[Data]
	ClassElementVisitor[Data]
	BlockVisitor[Data]
	StmtVisitor[Data]
	ExprVisitor[Data]
	ItemVisitor[Data]
	ArgVisitor[Data]
	TopDefVisitor[Data]
	ProgramVisitor[Data]
}
