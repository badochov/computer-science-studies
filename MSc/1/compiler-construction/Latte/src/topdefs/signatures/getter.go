package signatures

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/badochov/MRJP/Latte/src/frontend/types"
	"github.com/badochov/MRJP/Latte/src/parser"
	"github.com/badochov/MRJP/Latte/src/status"
	"golang.org/x/exp/slices"
	"strings"
)

type sigsGetter struct {
	*parser.BaseLatteVisitor
	fns       map[string]Function
	classes   map[string]classSignature
	currClass *classSignature
}

type TopDefs struct {
	Fns     map[string]Function
	Classes map[string]Class
}

func Get(tree antlr.ParseTree) TopDefs {
	return tree.Accept(newFnSigsGetter()).(TopDefs)
}

func newFnSigsGetter() *sigsGetter {
	return &sigsGetter{&parser.BaseLatteVisitor{}, defaultFuncSigs(), make(map[string]classSignature), nil}
}

func (v *sigsGetter) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, def := range ctx.AllTopDef() {
		def.Accept(v)
	}
	return v.topDefs()
}

func (v *sigsGetter) topDefs() TopDefs {
	topDefs := TopDefs{
		Fns:     v.fns,
		Classes: make(map[string]Class),
	}

	for name := range v.classes {
		getClass(name, v.classes, topDefs.Classes)
	}
	validateTypes(topDefs)

	return topDefs
}

func isTypeValid(p types.Type, classes map[string]Class) bool {
	if slices.Contains([]types.Type{types.Int, types.Bool, types.String, types.Void}, p) {
		return true
	}

	if p.Struct {
		_, ok := classes[p.Name]
		return ok
	}

	if p.Array != nil {
		return isTypeValid(*p.Array, classes)
	}
	return false
}

func validateTypes(defs TopDefs) {
	validateFunctionTypes(defs.Fns, defs.Classes)
	validateClasses(defs.Classes)
}

func validateClasses(classes map[string]Class) {
	for className, class := range classes {
		validateAttributes(className, class, classes)
		validateMethods(className, class, classes)
	}

	for _, class := range classes {
		validateVirtualMethods(class)
	}
}

func validateVirtualMethods(class Class) {
	parents := class.GetParents()

	for i := len(parents) - 1; i >= 0; i-- {
		currClass := parents[i]
		for methodName, methodSig := range class.Methods {
			sig, ok := currClass.Methods[methodName]
			if !ok {
				continue
			}
			if err := compareSig(sig, methodSig); err != nil {
				status.Fatal(fmt.Errorf("class '%s' tries to overwrites virtual method '%s' inherited from class '%s'\n%v", class.Name, methodName, currClass.Name, err))
			}
		}
	}
}

func compareSig(sig Function, sig2 Function) error {
	if sig.RetType == sig2.RetType {
		return nil
	}

	for i, argT := range sig.ArgTypes {
		argT2 := sig2.ArgTypes[i]
		if argT.Ref != argT2.Ref || argT.Type != argT2.Type {
			return fmt.Errorf("expected: %s\ngot     : %s\n", sig, sig2)
		}
	}

	return nil
}

func validateMethods(className string, class Class, classes map[string]Class) {
	for methodName, methodSig := range class.Methods {
		if !isTypeValid(methodSig.RetType, classes) {
			status.Fatal(fmt.Errorf("return type '%s' of method '%s' on class '%s' is not a valid type", methodName, methodSig.RetType, className))
		}
		for i, t := range methodSig.ArgTypes {
			if !isTypeValid(t.Type, classes) {
				status.Fatal(fmt.Errorf("type of argument no %d '%s' of method '%s' on class '%s' is not a valid type", i, methodName, methodSig.RetType, className))
			}
		}
	}
}

func validateAttributes(className string, class Class, classes map[string]Class) {
	for attrName, attrType := range class.Attributes {
		if !isTypeValid(attrType, classes) {
			status.Fatal(fmt.Errorf("type of '%s' attribute '%s' of class '%s' is not a valid type", attrName, attrType, className))
		}
	}
}

func validateFunctionTypes(fns map[string]Function, classes map[string]Class) {
	for name, fn := range fns {
		if !isTypeValid(fn.RetType, classes) {
			status.Fatal(fmt.Errorf("return type '%s' of function '%s' is not a valid type", name, fn.RetType))
		}
		for i, t := range fn.ArgTypes {
			if !isTypeValid(t.Type, classes) {
				status.Fatal(fmt.Errorf("type of argument no %d '%s' of function '%s' is not a valid type", i, name, fn.RetType))
			}
		}
	}
}

func getClass(className string, classes map[string]classSignature, converted map[string]Class) Class {
	if c, ok := converted[className]; ok {
		return c
	}
	if c, ok := classes[className]; ok {
		cl := Class{
			Attributes: c.Attributes,
			Methods:    c.Methods,
			Name:       c.Name,
		}
		if c.Extends != nil {
			nc := getClass(*c.Extends, classes, converted)
			cl.Extends = &nc
		}
		converted[className] = cl
		return cl
	}

	status.Fatal(fmt.Errorf("class %s is not defined", className))
	panic("FATAL EXITS")
}

func (v *sigsGetter) VisitDefFunc(ctx *parser.DefFuncContext) interface{} {
	fnName := ctx.ID().GetText()
	if _, ok := v.fns[fnName]; ok {
		status.HandleError(fmt.Errorf("redeclaration of function '%s'", fnName), ctx)
	}
	v.fns[fnName] = v.getFnType(ctx)
	return nil
}

func (v *sigsGetter) VisitClassAttribute(ctx *parser.ClassAttributeContext) interface{} {
	for _, id := range ctx.AllID() {
		name := id.GetText()
		if _, ok := v.currClass.Attributes[name]; ok {
			status.HandleError(fmt.Errorf("redeclaration of attribute '%s'", name), ctx)
		}
		v.currClass.Attributes[name] = ctx.Type_().Accept(v).(types.Type)
	}
	return nil
}

func (v *sigsGetter) VisitClassMethod(ctx *parser.ClassMethodContext) interface{} {
	retT := ctx.Type_().Accept(v).(types.Type)
	var argT []Arg
	arg := ctx.Arg()
	if arg != nil {
		argT = arg.Accept(v).([]Arg)
	}
	fn := Function{retT, argT}

	v.currClass.Methods[ctx.ID().GetText()] = fn
	return nil
}

func (v *sigsGetter) VisitDefClass(ctx *parser.DefClassContext) interface{} {
	className := ctx.ID(0).GetText()
	if _, ok := v.classes[className]; ok {
		status.HandleError(fmt.Errorf("redeclaration of class '%s'", className), ctx)
	}

	v.currClass = &classSignature{
		Attributes: make(map[string]types.Type),
		Methods:    make(map[string]Function),
		Name:       className,
	}

	if extendsNode := ctx.ID(1); extendsNode != nil {
		extends := extendsNode.GetText()
		v.currClass.Extends = &extends
	}
	v.classes[ctx.ID(0).GetText()] = *v.currClass

	for _, el := range ctx.AllClassElement() {
		el.Accept(v)
	}
	return *v.currClass
}

func (v *sigsGetter) getFnType(ctx *parser.DefFuncContext) Function {
	retT := ctx.Type_().Accept(v).(types.Type)
	var argT []Arg
	arg := ctx.Arg()
	if arg != nil {
		argT = arg.Accept(v).([]Arg)
	}
	return Function{retT, argT}
}

func (v *sigsGetter) VisitArg(ctx *parser.ArgContext) interface{} {
	ret := make([]Arg, 0, len(ctx.AllType_()))
	for i, t := range ctx.AllType_() {
		ty := t.Accept(v).(types.Type)
		ref := ctx.Ref(i).Accept(v).(bool)
		ret = append(ret, Arg{
			Ref:  ref,
			Name: ctx.ID(i).GetText(),
			Type: ty,
		})
	}
	return ret
}

func (v *sigsGetter) VisitRef(ctx *parser.RefContext) interface{} {
	return !ctx.IsEmpty()
}

func (v *sigsGetter) VisitClassName(ctx *parser.ClassNameContext) interface{} {
	return types.Type{
		Name:   ctx.ID().GetText(),
		Struct: true,
	}
}

func (v *sigsGetter) VisitTypeArray(ctx *parser.TypeArrayContext) interface{} {
	t := ctx.Type_single().Accept(v).(types.Type)
	return types.Type{
		Array: &t,
	}
}

func (v *sigsGetter) VisitTypeSingle(ctx *parser.TypeSingleContext) interface{} {
	return ctx.Type_single().Accept(v)
}

func (v *sigsGetter) VisitInt(ctx *parser.IntContext) interface{} {
	return types.Int
}

func (v *sigsGetter) VisitStr(ctx *parser.StrContext) interface{} {
	return types.String
}

func (v *sigsGetter) VisitBool(ctx *parser.BoolContext) interface{} {
	return types.Bool
}

func (v *sigsGetter) VisitVoid(ctx *parser.VoidContext) interface{} {
	return types.Void
}

type Arg struct {
	Ref  bool
	Name string
	Type types.Type
}

func (a Arg) String() string {
	if a.Ref {
		return fmt.Sprintf("&%s", a.Type)
	}
	return a.Type.String()
}

type Function struct {
	RetType  types.Type
	ArgTypes []Arg
}

func (f Function) String() string {
	args := make([]string, len(f.ArgTypes))
	for i, arg := range f.ArgTypes {
		args[i] = arg.String()
	}
	return fmt.Sprintf("%s(%s)", f.RetType, strings.Join(args, ","))
}

type classSignature struct {
	Extends    *string
	Name       string
	Attributes map[string]types.Type
	Methods    map[string]Function
}

type Class struct {
	Extends    *Class
	Name       string
	Attributes map[string]types.Type
	Methods    map[string]Function
}

func (c Class) GetParents() []Class {
	var parents []Class
	currClass := c
	for currClass.Extends != nil {
		currClass = *currClass.Extends
		parents = append(parents, currClass)
	}
	return parents
}

func defaultFuncSigs() map[string]Function {
	return map[string]Function{
		"printInt":    {types.Void, []Arg{{false, "num", types.Int}}},
		"printString": {types.Void, []Arg{{false, "str", types.String}}},
		"error":       {types.Void, nil},
		"readInt":     {types.Int, nil},
		"readString":  {types.String, nil},

		".str_add": {types.String, []Arg{{false, "str", types.String}, {false, "str2", types.String}}},
		".new_arr": {types.String, []Arg{{false, "count", types.Int}, {false, "size", types.Int}}},
		"calloc":   {types.String, []Arg{{false, "count", types.Int}, {false, "size", types.Int}}},
	}
}
