package quadruple

import (
	"encoding/json"
	"golang.org/x/exp/maps"
)

type Program struct {
	Objects map[string]*Object
	Fns     map[string]*Fn
	Strings map[string]struct{}
}

type Object struct {
	Name           string
	Extends        *Object
	Methods        []*Fn
	Attributes     []Type
	MethodsMap     map[string]int
	AttributeToIdx map[string]int
}

func (o Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{ Name string }{Name: o.Name})
}

func (o Object) Copy() Object {
	this := Object{
		Methods:        make([]*Fn, len(o.Methods)),
		Attributes:     make([]Type, len(o.Attributes)),
		AttributeToIdx: make(map[string]int, len(o.AttributeToIdx)),
		MethodsMap:     make(map[string]int, len(o.MethodsMap)),
	}
	copy(this.Methods, o.Methods)
	copy(this.Attributes, o.Attributes)
	maps.Copy(this.AttributeToIdx, o.AttributeToIdx)
	maps.Copy(this.MethodsMap, o.MethodsMap)

	return this
}

func (o Object) GetMethod(id string) *Fn {
	return o.Methods[o.MethodsMap[id]]
}

type Block struct {
	Id int

	Instructions []Instruction
}

func (b Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{ Id int }{Id: b.Id})
}

type InstructionVisitor interface {
	VisitGoto(g *Goto) any
	VisitUnreachable(u *Unreachable) any
	VisitPointerAssignment(pa *PointerAssignment) any
	VisitAssignment(a *Assignment) any
	VisitReturn(r *Return) any
	VisitVoidCall(fc *VoidCall) any
	VisitCondJump(cj *CondJump) any
}

type BlockVisitor interface {
	VisitBlock(b *Block) any
}

type Instruction interface {
	Accept(visitor InstructionVisitor) any
}

type TypeEnum int8

const (
	IntType TypeEnum = iota
	BoolType
	StringType
	VoidType
	ObjectType
	UntypedNullType
)

type Type struct {
	Type    TypeEnum
	Pointer bool
	Array   bool
	Object  *Object
}

type Val interface {
	Type() Type
}

type NullVal struct {
	Type_ Type
}

func (v NullVal) Type() Type {
	return v.Type_
}

type IntVal struct {
	Int int64
}

func (v IntVal) Type() Type {
	return Type{Type: IntType}
}

var _ Val = Var{}
var _ Val = IntVal{}
var _ Val = BoolVal{}
var _ Val = NullVal{}
var _ Val = StrVal{}

type BoolVal struct {
	Bool bool
}

func (v BoolVal) Type() Type {
	return Type{Type: BoolType}
}

type StrVal struct {
	String string
}

func (v StrVal) Type() Type {
	return Type{Type: StringType}
}

type Var struct {
	Type_ Type
	Id    string
}

func (v Var) Type() Type {
	return v.Type_
}

type Return struct {
	Val *Val
}

func (r *Return) Accept(visitor InstructionVisitor) any {
	return visitor.VisitReturn(r)
}

type Unreachable struct {
}

func (u *Unreachable) Accept(visitor InstructionVisitor) any {
	return visitor.VisitUnreachable(u)
}

type Arg struct {
	Name string
	Type Type
	Ref  bool
}

type Fn struct {
	Name string

	Ret  Type
	Args []Arg

	Blocks []*Block
}

func (f *Fn) Accept(visitor CallableVisitor) any {
	return visitor.VisitFn(f)
}

func (f *Fn) RetType() Type {
	return f.Ret
}

func (f *Fn) Arguments() []Arg {
	return f.Args
}

type VoidCall struct {
	Callable Callable
	Params   []Val
}

func (v *VoidCall) Accept(visitor InstructionVisitor) any {
	return visitor.VisitVoidCall(v)
}

type CondJump struct {
	Cond       Val
	TrueBlock  *Block
	FalseBlock *Block
}

func (c *CondJump) Accept(visitor InstructionVisitor) any {
	return visitor.VisitCondJump(c)
}

type Goto struct {
	Block *Block
}

func (g *Goto) Accept(visitor InstructionVisitor) any {
	return visitor.VisitGoto(g)
}

type PointerAssignment struct {
	Pointer Val
	Val     Val
}

func (p *PointerAssignment) Accept(visitor InstructionVisitor) any {
	return visitor.VisitPointerAssignment(p)
}

type Assignment struct {
	Var  Var
	Expr Expr
}

func (a *Assignment) Accept(visitor InstructionVisitor) any {
	return visitor.VisitAssignment(a)
}

var _ Instruction = &Goto{}
var _ Instruction = &PointerAssignment{}
var _ Instruction = &Assignment{}
var _ Instruction = &Return{}
var _ Instruction = &Unreachable{}
var _ Instruction = &VoidCall{}
var _ Instruction = &CondJump{}

type ExprVisitor interface {
	VisitEUnOp(e *EUnOp) any
	VisitEBinOp(e *EBinOp) any
	VisitEVal(e *EVal) any
	VisitECall(e *ECall) any
	VisitEGetAttrPtr(e *EGetAttrPtr) any
	VisitEGetArrLen(e *EGetArrLen) any
	VisitEGetArrElPtr(e *EGetArrElPtr) any
	VisitELoadPtr(e *ELoadPtr) any
	VisitENewArray(e *ENewArray) any
	VisitENewObj(e *ENewObj) any
	VisitEPhi(e *EPhi) any
	VisitECast(e *ECast) any
}

type Expr interface {
	Accept(visitor ExprVisitor) any
	Serialize() SerializedExpression
}

type SerializedExpression string

type BinOp int

const (
	Modulo BinOp = iota
	Times
	Divide
	Plus
	Minus
	Ge
	Gt
	Eq
	Ne
)

type EBinOp struct {
	LVal Val
	RVal Val
	Op   BinOp
}

func (e *EBinOp) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *EBinOp) Accept(visitor ExprVisitor) any {
	return visitor.VisitEBinOp(e)
}

type UnOp int

const (
	LogicNot UnOp = iota
)

type EUnOp struct {
	Val Val
	Op  UnOp
}

func (e *EUnOp) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *EUnOp) Accept(visitor ExprVisitor) any {
	return visitor.VisitEUnOp(e)
}

type EVal struct {
	Val Val
}

func (e *EVal) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *EVal) Accept(visitor ExprVisitor) any {
	return visitor.VisitEVal(e)
}

type CallableVisitor interface {
	VisitFn(fn *Fn) any
	VisitMethod(method *Method) any
}

type Callable interface {
	RetType() Type
	Arguments() []Arg
	Accept(visitor CallableVisitor) any
}

type ECall struct {
	Callable Callable
	Params   []Val
}

type Method struct {
	Object Val
	Idx    int
}

func (m *Method) Accept(visitor CallableVisitor) any {
	return visitor.VisitMethod(m)
}

func (m *Method) RetType() Type {
	return m.getFn().RetType()
}

func (m *Method) Arguments() []Arg {
	return m.getFn().Arguments()
}

func (m *Method) getFn() *Fn {
	return m.Object.Type().Object.Methods[m.Idx]
}

var _ Callable = &Fn{}
var _ Callable = &Method{}

func (e *ECall) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *ECall) Accept(visitor ExprVisitor) any {
	return visitor.VisitECall(e)
}

type EGetArrElPtr struct {
	Arr Val
	Idx Val
}

func (e *EGetArrElPtr) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *EGetArrElPtr) Accept(visitor ExprVisitor) any {
	return visitor.VisitEGetArrElPtr(e)
}

type EGetArrLen struct {
	Arr Val
}

func (e *EGetArrLen) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *EGetArrLen) Accept(visitor ExprVisitor) any {
	return visitor.VisitEGetArrLen(e)
}

type EGetAttrPtr struct {
	Obj     Val
	AttrIdx int
}

func (e *EGetAttrPtr) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *EGetAttrPtr) Accept(visitor ExprVisitor) any {
	return visitor.VisitEGetAttrPtr(e)
}

type ELoadPtr struct {
	Ptr Val
}

func (e *ELoadPtr) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *ELoadPtr) Accept(visitor ExprVisitor) any {
	return visitor.VisitELoadPtr(e)
}

type ENewArray struct {
	Type Type
	Size Val
}

func (e *ENewArray) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *ENewArray) Accept(visitor ExprVisitor) any {
	return visitor.VisitENewArray(e)
}

type ENewObj struct {
	Obj *Object
}

func (e *ENewObj) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *ENewObj) Accept(visitor ExprVisitor) any {
	return visitor.VisitENewObj(e)
}

type ECast struct {
	Val Val
	To  Type
}

func (e *ECast) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *ECast) Accept(visitor ExprVisitor) any {
	return visitor.VisitECast(e)
}

type PhiValData struct {
	Block *Block
	Val   Val
}

type EPhi struct {
	Vals []PhiValData
}

func (e *EPhi) Serialize() SerializedExpression {
	v, _ := json.Marshal(e)
	return SerializedExpression(v)
}

func (e *EPhi) Accept(visitor ExprVisitor) any {
	return visitor.VisitEPhi(e)
}

var _ Expr = &EUnOp{}
var _ Expr = &ECast{}
var _ Expr = &EBinOp{}
var _ Expr = &EVal{}
var _ Expr = &ECall{}
var _ Expr = &EGetAttrPtr{}
var _ Expr = &EGetArrElPtr{}
var _ Expr = &EGetArrLen{}
var _ Expr = &ELoadPtr{}
var _ Expr = &ENewArray{}
var _ Expr = &ENewObj{}
var _ Expr = &EPhi{}
