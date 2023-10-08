package ast

type ITopDef[Data any] interface {
	VisitTopDef(visitor TopDefVisitor[Data]) any
}

type TopDefVisitor[Data any] interface {
	VisitFunc(*Func[Data]) any
	VisitClass(*Class[Data]) any
}

type Func[Data any] struct {
	Data Data

	Type  IType[Data]
	Id    string
	Args  []IArg[Data]
	Block IBlock[Data]
}

type Class[Data any] struct {
	Data Data

	Id            string
	Extends       *string
	ClassElements []IClassElement[Data]
}

func (f *Func[Data]) VisitTopDef(visitor TopDefVisitor[Data]) any {
	return visitor.VisitFunc(f)
}

func (f *Class[Data]) VisitTopDef(visitor TopDefVisitor[Data]) any {
	return visitor.VisitClass(f)
}

var _ ITopDef[any] = &Func[any]{}
var _ ITopDef[any] = &Class[any]{}
