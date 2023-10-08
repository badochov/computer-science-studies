// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Latte

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseLatteVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLatteVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitDefFunc(ctx *DefFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitDefClass(ctx *DefClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitArg(ctx *ArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitRef(ctx *RefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEmpty(ctx *EmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitIndexAss(ctx *IndexAssContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitAttrAss(ctx *AttrAssContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitBlockStmt(ctx *BlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitDecl(ctx *DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitAss(ctx *AssContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitIncr(ctx *IncrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitDecr(ctx *DecrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitRet(ctx *RetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitVRet(ctx *VRetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitCond(ctx *CondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitCondElse(ctx *CondElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitWhile(ctx *WhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitForEach(ctx *ForEachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSExp(ctx *SExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTypeArray(ctx *TypeArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTypeSingle(ctx *TypeSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitStr(ctx *StrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitBool(ctx *BoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitVoid(ctx *VoidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitClassName(ctx *ClassNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitClassAttribute(ctx *ClassAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitClassMethod(ctx *ClassMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitItem(ctx *ItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEId(ctx *EIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEFunCall(ctx *EFunCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitENewArray(ctx *ENewArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitERelOp(ctx *ERelOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitETrue(ctx *ETrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEOr(ctx *EOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEInt(ctx *EIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitENewClass(ctx *ENewClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEUnOp(ctx *EUnOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEStr(ctx *EStrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEMulOp(ctx *EMulOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEAnd(ctx *EAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEMethodCall(ctx *EMethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEParen(ctx *EParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEFalse(ctx *EFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEIndex(ctx *EIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEAddOp(ctx *EAddOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitENull(ctx *ENullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEAttr(ctx *EAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitUnOp(ctx *UnOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitAddOp(ctx *AddOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitMulOp(ctx *MulOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitRelOp(ctx *RelOpContext) interface{} {
	return v.VisitChildren(ctx)
}
