package ast

import (
	"github.com/badochov/MRJP/Latte/src/frontend/types"
)

type IType[Data any] interface {
	VisitType(visitor TypeVisitor[Data]) any
}

type ITypeSingle[Data any] interface {
	VisitTypeSingle(visitor TypeSingleVisitor[Data]) any
}

type ITypeArray[Data any] interface {
	VisitTypeArray(visitor TypeArrayVisitor[Data]) any
}

type TypeSingleVisitor[Data any] interface {
	VisitTypeInt(i *TypeInt[Data]) any
	VisitTypeString(i *TypeString[Data]) any
	VisitTypeBool(i *TypeBool[Data]) any
	VisitTypeVoid(i *TypeVoid[Data]) any
	VisitClassName(i *TypeClassName[Data]) any
}

type TypeArrayVisitor[Data any] interface {
	VisitTypeArray(i *TypeArray[Data]) any
}

type TypeVisitor[Data any] interface {
	TypeSingleVisitor[Data]
	TypeArrayVisitor[Data]
}

type TypeInt[Data any] struct {
	Data Data
}

func (i *TypeInt[Data]) VisitType(visitor TypeVisitor[Data]) any {
	return i.VisitTypeSingle(visitor)
}

func (i *TypeInt[Data]) VisitTypeSingle(visitor TypeSingleVisitor[Data]) any {
	return visitor.VisitTypeInt(i)
}

type TypeString[Data any] struct {
	Data Data
}

func (i *TypeString[Data]) VisitType(visitor TypeVisitor[Data]) any {
	return i.VisitTypeSingle(visitor)
}

func (i *TypeString[Data]) VisitTypeSingle(visitor TypeSingleVisitor[Data]) any {
	return visitor.VisitTypeString(i)
}

type TypeBool[Data any] struct {
	Data Data
}

func (i *TypeBool[Data]) VisitType(visitor TypeVisitor[Data]) any {
	return i.VisitTypeSingle(visitor)
}

func (i *TypeBool[Data]) VisitTypeSingle(visitor TypeSingleVisitor[Data]) any {
	return visitor.VisitTypeBool(i)
}

type TypeVoid[Data any] struct {
	Data Data
}

func (i *TypeVoid[Data]) VisitType(visitor TypeVisitor[Data]) any {
	return i.VisitTypeSingle(visitor)
}

func (i *TypeVoid[Data]) VisitTypeSingle(visitor TypeSingleVisitor[Data]) any {
	return visitor.VisitTypeVoid(i)
}

type TypeClassName[Data any] struct {
	Data Data

	Id string
}

func (i *TypeClassName[Data]) VisitType(visitor TypeVisitor[Data]) any {
	return i.VisitTypeSingle(visitor)
}

func (i *TypeClassName[Data]) VisitTypeSingle(visitor TypeSingleVisitor[Data]) any {
	return visitor.VisitClassName(i)
}

type TypeArray[Data any] struct {
	Data Data

	Type ITypeSingle[Data]
}

func (i *TypeArray[Data]) VisitType(visitor TypeVisitor[Data]) any {
	return i.VisitTypeArray(visitor)
}

func (i *TypeArray[Data]) VisitTypeArray(visitor TypeArrayVisitor[Data]) any {
	return visitor.VisitTypeArray(i)
}

var _ ITypeSingle[any] = &TypeInt[any]{}
var _ ITypeSingle[any] = &TypeString[any]{}
var _ ITypeSingle[any] = &TypeBool[any]{}
var _ ITypeSingle[any] = &TypeVoid[any]{}
var _ ITypeSingle[any] = &TypeClassName[any]{}

var _ ITypeArray[any] = &TypeArray[any]{}

var _ IType[any] = &TypeInt[any]{}
var _ IType[any] = &TypeString[any]{}
var _ IType[any] = &TypeBool[any]{}
var _ IType[any] = &TypeVoid[any]{}
var _ IType[any] = &TypeClassName[any]{}
var _ IType[any] = &TypeArray[any]{}

func GetType[Data any](i IType[Data]) types.Type {
	if i == nil {
		return types.Null
	}
	return i.VisitType(typeGetter[Data]{}).(types.Type)
}

func GetTypeSingle[Data any](i ITypeSingle[Data]) types.Type {
	return i.VisitTypeSingle(typeGetter[Data]{}).(types.Type)
}

func GetTypeArray[Data any](i ITypeArray[Data]) types.Type {
	return i.VisitTypeArray(typeGetter[Data]{}).(types.Type)
}

type typeGetter[Data any] struct{}

func (t typeGetter[Data]) VisitTypeArray(i *TypeArray[Data]) any {
	typ := i.Type.VisitTypeSingle(t).(types.Type)
	return types.Type{Array: &typ}
}

func (t typeGetter[Data]) VisitClassName(i *TypeClassName[Data]) any {
	return types.Type{
		Name:   i.Id,
		Struct: true,
	}
}

func (t typeGetter[Data]) VisitTypeInt(i *TypeInt[Data]) any {
	return types.Int
}

func (t typeGetter[Data]) VisitTypeString(i *TypeString[Data]) any {
	return types.String
}

func (t typeGetter[Data]) VisitTypeBool(i *TypeBool[Data]) any {
	return types.Bool
}

func (t typeGetter[Data]) VisitTypeVoid(i *TypeVoid[Data]) any {
	return types.Void
}

var _ TypeVisitor[any] = &typeGetter[any]{}
