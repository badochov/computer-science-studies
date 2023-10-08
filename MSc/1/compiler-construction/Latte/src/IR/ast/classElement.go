package ast

type IClassElement[Data any] interface {
	VisitClassElement(visitor ClassElementVisitor[Data]) any
}

type ClassElementVisitor[Data any] interface {
	VisitClassAttribute(attribute *Attribute[Data]) any
	VisitClassMethod(*Method[Data]) any
}

type Method[Data any] struct {
	Data Data

	Type  IType[Data]
	Id    string
	Args  []IArg[Data]
	Block IBlock[Data]
}

type Attribute[Data any] struct {
	Data Data

	Type IType[Data]
	Id   string
}

func (f *Method[Data]) VisitClassElement(visitor ClassElementVisitor[Data]) any {
	return visitor.VisitClassMethod(f)
}

func (f *Attribute[Data]) VisitClassElement(visitor ClassElementVisitor[Data]) any {
	return visitor.VisitClassAttribute(f)
}

var _ IClassElement[any] = &Attribute[any]{}
var _ IClassElement[any] = &Method[any]{}
