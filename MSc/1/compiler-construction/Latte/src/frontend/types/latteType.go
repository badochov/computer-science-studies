package types

import "fmt"

type Type struct {
	Name        TypeName // TypeName or Name
	Array       *Type
	Struct      bool
}

func (t Type) String() string {
	if t.Array != nil {
		return fmt.Sprintf("%s[]", *t.Array)
	}
	return fmt.Sprint(t.Name)
}

type TypeName = string

var String Type = Type{Name: "string"}
var Int Type = Type{Name: "int"}
var Bool Type = Type{Name: "bool"}
var Void Type = Type{Name: "void"}
var Null Type = Type{Name: "null"}
