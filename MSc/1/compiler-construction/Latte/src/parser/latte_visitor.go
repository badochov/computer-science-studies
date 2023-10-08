// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Latte

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by LatteParser.
type LatteVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LatteParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by LatteParser#DefFunc.
	VisitDefFunc(ctx *DefFuncContext) interface{}

	// Visit a parse tree produced by LatteParser#DefClass.
	VisitDefClass(ctx *DefClassContext) interface{}

	// Visit a parse tree produced by LatteParser#arg.
	VisitArg(ctx *ArgContext) interface{}

	// Visit a parse tree produced by LatteParser#ref.
	VisitRef(ctx *RefContext) interface{}

	// Visit a parse tree produced by LatteParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LatteParser#Empty.
	VisitEmpty(ctx *EmptyContext) interface{}

	// Visit a parse tree produced by LatteParser#IndexAss.
	VisitIndexAss(ctx *IndexAssContext) interface{}

	// Visit a parse tree produced by LatteParser#AttrAss.
	VisitAttrAss(ctx *AttrAssContext) interface{}

	// Visit a parse tree produced by LatteParser#BlockStmt.
	VisitBlockStmt(ctx *BlockStmtContext) interface{}

	// Visit a parse tree produced by LatteParser#Decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by LatteParser#Ass.
	VisitAss(ctx *AssContext) interface{}

	// Visit a parse tree produced by LatteParser#Incr.
	VisitIncr(ctx *IncrContext) interface{}

	// Visit a parse tree produced by LatteParser#Decr.
	VisitDecr(ctx *DecrContext) interface{}

	// Visit a parse tree produced by LatteParser#Ret.
	VisitRet(ctx *RetContext) interface{}

	// Visit a parse tree produced by LatteParser#VRet.
	VisitVRet(ctx *VRetContext) interface{}

	// Visit a parse tree produced by LatteParser#Cond.
	VisitCond(ctx *CondContext) interface{}

	// Visit a parse tree produced by LatteParser#CondElse.
	VisitCondElse(ctx *CondElseContext) interface{}

	// Visit a parse tree produced by LatteParser#While.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by LatteParser#ForEach.
	VisitForEach(ctx *ForEachContext) interface{}

	// Visit a parse tree produced by LatteParser#SExp.
	VisitSExp(ctx *SExpContext) interface{}

	// Visit a parse tree produced by LatteParser#TypeArray.
	VisitTypeArray(ctx *TypeArrayContext) interface{}

	// Visit a parse tree produced by LatteParser#TypeSingle.
	VisitTypeSingle(ctx *TypeSingleContext) interface{}

	// Visit a parse tree produced by LatteParser#Int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by LatteParser#Str.
	VisitStr(ctx *StrContext) interface{}

	// Visit a parse tree produced by LatteParser#Bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by LatteParser#Void.
	VisitVoid(ctx *VoidContext) interface{}

	// Visit a parse tree produced by LatteParser#ClassName.
	VisitClassName(ctx *ClassNameContext) interface{}

	// Visit a parse tree produced by LatteParser#ClassAttribute.
	VisitClassAttribute(ctx *ClassAttributeContext) interface{}

	// Visit a parse tree produced by LatteParser#ClassMethod.
	VisitClassMethod(ctx *ClassMethodContext) interface{}

	// Visit a parse tree produced by LatteParser#item.
	VisitItem(ctx *ItemContext) interface{}

	// Visit a parse tree produced by LatteParser#EId.
	VisitEId(ctx *EIdContext) interface{}

	// Visit a parse tree produced by LatteParser#EFunCall.
	VisitEFunCall(ctx *EFunCallContext) interface{}

	// Visit a parse tree produced by LatteParser#ENewArray.
	VisitENewArray(ctx *ENewArrayContext) interface{}

	// Visit a parse tree produced by LatteParser#ERelOp.
	VisitERelOp(ctx *ERelOpContext) interface{}

	// Visit a parse tree produced by LatteParser#ETrue.
	VisitETrue(ctx *ETrueContext) interface{}

	// Visit a parse tree produced by LatteParser#EOr.
	VisitEOr(ctx *EOrContext) interface{}

	// Visit a parse tree produced by LatteParser#EInt.
	VisitEInt(ctx *EIntContext) interface{}

	// Visit a parse tree produced by LatteParser#ENewClass.
	VisitENewClass(ctx *ENewClassContext) interface{}

	// Visit a parse tree produced by LatteParser#EUnOp.
	VisitEUnOp(ctx *EUnOpContext) interface{}

	// Visit a parse tree produced by LatteParser#EStr.
	VisitEStr(ctx *EStrContext) interface{}

	// Visit a parse tree produced by LatteParser#EMulOp.
	VisitEMulOp(ctx *EMulOpContext) interface{}

	// Visit a parse tree produced by LatteParser#EAnd.
	VisitEAnd(ctx *EAndContext) interface{}

	// Visit a parse tree produced by LatteParser#EMethodCall.
	VisitEMethodCall(ctx *EMethodCallContext) interface{}

	// Visit a parse tree produced by LatteParser#EParen.
	VisitEParen(ctx *EParenContext) interface{}

	// Visit a parse tree produced by LatteParser#EFalse.
	VisitEFalse(ctx *EFalseContext) interface{}

	// Visit a parse tree produced by LatteParser#EIndex.
	VisitEIndex(ctx *EIndexContext) interface{}

	// Visit a parse tree produced by LatteParser#EAddOp.
	VisitEAddOp(ctx *EAddOpContext) interface{}

	// Visit a parse tree produced by LatteParser#ENull.
	VisitENull(ctx *ENullContext) interface{}

	// Visit a parse tree produced by LatteParser#EAttr.
	VisitEAttr(ctx *EAttrContext) interface{}

	// Visit a parse tree produced by LatteParser#unOp.
	VisitUnOp(ctx *UnOpContext) interface{}

	// Visit a parse tree produced by LatteParser#addOp.
	VisitAddOp(ctx *AddOpContext) interface{}

	// Visit a parse tree produced by LatteParser#mulOp.
	VisitMulOp(ctx *MulOpContext) interface{}

	// Visit a parse tree produced by LatteParser#relOp.
	VisitRelOp(ctx *RelOpContext) interface{}
}
