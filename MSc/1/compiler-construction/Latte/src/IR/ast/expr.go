package ast

type IExpr[Data any] interface {
	VisitExpr(visitor ExprVisitor[Data]) any
}

type (
	ExprVisitor[Data any] interface {
		VisitExprParen(e *ExprParen[Data]) any
		VisitExprUnOp(e *ExprUnOp[Data]) any
		VisitExprMulOp(e *ExprMulOp[Data]) any
		VisitExprAddOp(e *ExprAddOp[Data]) any
		VisitExprRelOp(e *ExprRelOp[Data]) any
		VisitExprAnd(e *ExprAnd[Data]) any
		VisitExprOr(e *ExprOr[Data]) any
		VisitExprId(e *ExprId[Data]) any
		VisitExprInt(e *ExprInt[Data]) any
		VisitExprBool(e *ExprBool[Data]) any
		VisitExprString(e *ExprString[Data]) any
		VisitExprFunCall(e *ExprFunCall[Data]) any
		VisitExprIndex(e *ExprIndex[Data]) any
		VisitExprNull(e *ExprNull[Data]) any
		VisitExprNewClass(e *ExprNewClass[Data]) any
		VisitExprMethodCall(e *ExprMethodCall[Data]) any
		VisitExprAttr(e *ExprAttr[Data]) any
		VisitExprNewArray(e *ExprNewArray[Data]) any
	}
)

type UnOp int

const (
	LogicNot UnOp = iota
	UnaryMinus
)

type ExprUnOp[Data any] struct {
	Data Data

	UnOp UnOp
	Expr IExpr[Data]
}

func (e *ExprUnOp[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprUnOp(e)
}

type MulOp int

const (
	Times MulOp = iota
	Divide
	Modulo
)

type ExprMulOp[Data any] struct {
	Data Data

	MulOp MulOp
	LExpr IExpr[Data]
	RExpr IExpr[Data]
}

func (e *ExprMulOp[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprMulOp(e)
}

type AddOp int

const (
	Plus AddOp = iota
	Minus
)

type ExprAddOp[Data any] struct {
	Data Data

	AddOp AddOp
	LExpr IExpr[Data]
	RExpr IExpr[Data]
}

func (e *ExprAddOp[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprAddOp(e)
}

type RelOp int

const (
	Lt RelOp = iota
	Le
	Ge
	Gt
	Eq
	Ne
)

type ExprRelOp[Data any] struct {
	Data Data

	RelOp RelOp
	LExpr IExpr[Data]
	RExpr IExpr[Data]
}

func (e *ExprRelOp[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprRelOp(e)
}

type ExprAnd[Data any] struct {
	Data Data

	LExpr IExpr[Data]
	RExpr IExpr[Data]
}

func (e *ExprAnd[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprAnd(e)
}

type ExprOr[Data any] struct {
	Data Data

	LExpr IExpr[Data]
	RExpr IExpr[Data]
}

func (e *ExprOr[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprOr(e)
}

type ExprId[Data any] struct {
	Data Data

	Id string
}

func (e *ExprId[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprId(e)
}

type ExprInt[Data any] struct {
	Data Data

	Int int64
}

func (e *ExprInt[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprInt(e)
}

type ExprBool[Data any] struct {
	Data Data

	Bool bool
}

func (e *ExprBool[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprBool(e)
}

type ExprFunCall[Data any] struct {
	Data Data

	Id     string
	Params []IExpr[Data]
}

func (e *ExprFunCall[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprFunCall(e)
}

type ExprString[Data any] struct {
	Data Data

	String string
}

func (e *ExprString[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprString(e)
}

type ExprParen[Data any] struct {
	Data Data

	Expr IExpr[Data]
}

func (e *ExprParen[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprParen(e)
}

type ExprIndex[Data any] struct {
	Data Data

	ArrExpr   IExpr[Data]
	IndexExpr IExpr[Data]
}

func (e *ExprIndex[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprIndex(e)
}

type ExprMethodCall[Data any] struct {
	Data Data

	ObjExpr IExpr[Data]
	Id      string
	Params  []IExpr[Data]
}

func (e *ExprMethodCall[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprMethodCall(e)
}

type ExprAttr[Data any] struct {
	Data Data

	ObjExpr IExpr[Data]
	Attr    string
}

func (e *ExprAttr[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprAttr(e)
}

type ExprNewArray[Data any] struct {
	Data Data

	Type     ITypeSingle[Data]
	SizeExpr IExpr[Data]
}

func (e *ExprNewArray[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprNewArray(e)
}

type ExprNewClass[Data any] struct {
	Data Data

	Type *TypeClassName[Data]
}

func (e *ExprNewClass[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprNewClass(e)
}

type ExprNull[Data any] struct {
	Data Data

	Type IType[Data]
}

func (e *ExprNull[Data]) VisitExpr(visitor ExprVisitor[Data]) any {
	return visitor.VisitExprNull(e)
}

var _ IExpr[any] = &ExprUnOp[any]{}
var _ IExpr[any] = &ExprMulOp[any]{}
var _ IExpr[any] = &ExprAddOp[any]{}
var _ IExpr[any] = &ExprRelOp[any]{}
var _ IExpr[any] = &ExprAnd[any]{}
var _ IExpr[any] = &ExprOr[any]{}
var _ IExpr[any] = &ExprId[any]{}
var _ IExpr[any] = &ExprInt[any]{}
var _ IExpr[any] = &ExprBool[any]{}
var _ IExpr[any] = &ExprFunCall[any]{}
var _ IExpr[any] = &ExprString[any]{}
var _ IExpr[any] = &ExprParen[any]{}
var _ IExpr[any] = &ExprIndex[any]{}
var _ IExpr[any] = &ExprMethodCall[any]{}
var _ IExpr[any] = &ExprAttr[any]{}
var _ IExpr[any] = &ExprNewArray[any]{}
var _ IExpr[any] = &ExprNewClass[any]{}
var _ IExpr[any] = &ExprNull[any]{}
