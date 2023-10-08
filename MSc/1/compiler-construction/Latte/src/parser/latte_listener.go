// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Latte

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// LatteListener is a complete listener for a parse tree produced by LatteParser.
type LatteListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDefFunc is called when entering the DefFunc production.
	EnterDefFunc(c *DefFuncContext)

	// EnterDefClass is called when entering the DefClass production.
	EnterDefClass(c *DefClassContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterRef is called when entering the ref production.
	EnterRef(c *RefContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterEmpty is called when entering the Empty production.
	EnterEmpty(c *EmptyContext)

	// EnterIndexAss is called when entering the IndexAss production.
	EnterIndexAss(c *IndexAssContext)

	// EnterAttrAss is called when entering the AttrAss production.
	EnterAttrAss(c *AttrAssContext)

	// EnterBlockStmt is called when entering the BlockStmt production.
	EnterBlockStmt(c *BlockStmtContext)

	// EnterDecl is called when entering the Decl production.
	EnterDecl(c *DeclContext)

	// EnterAss is called when entering the Ass production.
	EnterAss(c *AssContext)

	// EnterIncr is called when entering the Incr production.
	EnterIncr(c *IncrContext)

	// EnterDecr is called when entering the Decr production.
	EnterDecr(c *DecrContext)

	// EnterRet is called when entering the Ret production.
	EnterRet(c *RetContext)

	// EnterVRet is called when entering the VRet production.
	EnterVRet(c *VRetContext)

	// EnterCond is called when entering the Cond production.
	EnterCond(c *CondContext)

	// EnterCondElse is called when entering the CondElse production.
	EnterCondElse(c *CondElseContext)

	// EnterWhile is called when entering the While production.
	EnterWhile(c *WhileContext)

	// EnterForEach is called when entering the ForEach production.
	EnterForEach(c *ForEachContext)

	// EnterSExp is called when entering the SExp production.
	EnterSExp(c *SExpContext)

	// EnterTypeArray is called when entering the TypeArray production.
	EnterTypeArray(c *TypeArrayContext)

	// EnterTypeSingle is called when entering the TypeSingle production.
	EnterTypeSingle(c *TypeSingleContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterStr is called when entering the Str production.
	EnterStr(c *StrContext)

	// EnterBool is called when entering the Bool production.
	EnterBool(c *BoolContext)

	// EnterVoid is called when entering the Void production.
	EnterVoid(c *VoidContext)

	// EnterClassName is called when entering the ClassName production.
	EnterClassName(c *ClassNameContext)

	// EnterClassAttribute is called when entering the ClassAttribute production.
	EnterClassAttribute(c *ClassAttributeContext)

	// EnterClassMethod is called when entering the ClassMethod production.
	EnterClassMethod(c *ClassMethodContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterEId is called when entering the EId production.
	EnterEId(c *EIdContext)

	// EnterEFunCall is called when entering the EFunCall production.
	EnterEFunCall(c *EFunCallContext)

	// EnterENewArray is called when entering the ENewArray production.
	EnterENewArray(c *ENewArrayContext)

	// EnterERelOp is called when entering the ERelOp production.
	EnterERelOp(c *ERelOpContext)

	// EnterETrue is called when entering the ETrue production.
	EnterETrue(c *ETrueContext)

	// EnterEOr is called when entering the EOr production.
	EnterEOr(c *EOrContext)

	// EnterEInt is called when entering the EInt production.
	EnterEInt(c *EIntContext)

	// EnterENewClass is called when entering the ENewClass production.
	EnterENewClass(c *ENewClassContext)

	// EnterEUnOp is called when entering the EUnOp production.
	EnterEUnOp(c *EUnOpContext)

	// EnterEStr is called when entering the EStr production.
	EnterEStr(c *EStrContext)

	// EnterEMulOp is called when entering the EMulOp production.
	EnterEMulOp(c *EMulOpContext)

	// EnterEAnd is called when entering the EAnd production.
	EnterEAnd(c *EAndContext)

	// EnterEMethodCall is called when entering the EMethodCall production.
	EnterEMethodCall(c *EMethodCallContext)

	// EnterEParen is called when entering the EParen production.
	EnterEParen(c *EParenContext)

	// EnterEFalse is called when entering the EFalse production.
	EnterEFalse(c *EFalseContext)

	// EnterEIndex is called when entering the EIndex production.
	EnterEIndex(c *EIndexContext)

	// EnterEAddOp is called when entering the EAddOp production.
	EnterEAddOp(c *EAddOpContext)

	// EnterENull is called when entering the ENull production.
	EnterENull(c *ENullContext)

	// EnterEAttr is called when entering the EAttr production.
	EnterEAttr(c *EAttrContext)

	// EnterUnOp is called when entering the unOp production.
	EnterUnOp(c *UnOpContext)

	// EnterAddOp is called when entering the addOp production.
	EnterAddOp(c *AddOpContext)

	// EnterMulOp is called when entering the mulOp production.
	EnterMulOp(c *MulOpContext)

	// EnterRelOp is called when entering the relOp production.
	EnterRelOp(c *RelOpContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDefFunc is called when exiting the DefFunc production.
	ExitDefFunc(c *DefFuncContext)

	// ExitDefClass is called when exiting the DefClass production.
	ExitDefClass(c *DefClassContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitRef is called when exiting the ref production.
	ExitRef(c *RefContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitEmpty is called when exiting the Empty production.
	ExitEmpty(c *EmptyContext)

	// ExitIndexAss is called when exiting the IndexAss production.
	ExitIndexAss(c *IndexAssContext)

	// ExitAttrAss is called when exiting the AttrAss production.
	ExitAttrAss(c *AttrAssContext)

	// ExitBlockStmt is called when exiting the BlockStmt production.
	ExitBlockStmt(c *BlockStmtContext)

	// ExitDecl is called when exiting the Decl production.
	ExitDecl(c *DeclContext)

	// ExitAss is called when exiting the Ass production.
	ExitAss(c *AssContext)

	// ExitIncr is called when exiting the Incr production.
	ExitIncr(c *IncrContext)

	// ExitDecr is called when exiting the Decr production.
	ExitDecr(c *DecrContext)

	// ExitRet is called when exiting the Ret production.
	ExitRet(c *RetContext)

	// ExitVRet is called when exiting the VRet production.
	ExitVRet(c *VRetContext)

	// ExitCond is called when exiting the Cond production.
	ExitCond(c *CondContext)

	// ExitCondElse is called when exiting the CondElse production.
	ExitCondElse(c *CondElseContext)

	// ExitWhile is called when exiting the While production.
	ExitWhile(c *WhileContext)

	// ExitForEach is called when exiting the ForEach production.
	ExitForEach(c *ForEachContext)

	// ExitSExp is called when exiting the SExp production.
	ExitSExp(c *SExpContext)

	// ExitTypeArray is called when exiting the TypeArray production.
	ExitTypeArray(c *TypeArrayContext)

	// ExitTypeSingle is called when exiting the TypeSingle production.
	ExitTypeSingle(c *TypeSingleContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitStr is called when exiting the Str production.
	ExitStr(c *StrContext)

	// ExitBool is called when exiting the Bool production.
	ExitBool(c *BoolContext)

	// ExitVoid is called when exiting the Void production.
	ExitVoid(c *VoidContext)

	// ExitClassName is called when exiting the ClassName production.
	ExitClassName(c *ClassNameContext)

	// ExitClassAttribute is called when exiting the ClassAttribute production.
	ExitClassAttribute(c *ClassAttributeContext)

	// ExitClassMethod is called when exiting the ClassMethod production.
	ExitClassMethod(c *ClassMethodContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitEId is called when exiting the EId production.
	ExitEId(c *EIdContext)

	// ExitEFunCall is called when exiting the EFunCall production.
	ExitEFunCall(c *EFunCallContext)

	// ExitENewArray is called when exiting the ENewArray production.
	ExitENewArray(c *ENewArrayContext)

	// ExitERelOp is called when exiting the ERelOp production.
	ExitERelOp(c *ERelOpContext)

	// ExitETrue is called when exiting the ETrue production.
	ExitETrue(c *ETrueContext)

	// ExitEOr is called when exiting the EOr production.
	ExitEOr(c *EOrContext)

	// ExitEInt is called when exiting the EInt production.
	ExitEInt(c *EIntContext)

	// ExitENewClass is called when exiting the ENewClass production.
	ExitENewClass(c *ENewClassContext)

	// ExitEUnOp is called when exiting the EUnOp production.
	ExitEUnOp(c *EUnOpContext)

	// ExitEStr is called when exiting the EStr production.
	ExitEStr(c *EStrContext)

	// ExitEMulOp is called when exiting the EMulOp production.
	ExitEMulOp(c *EMulOpContext)

	// ExitEAnd is called when exiting the EAnd production.
	ExitEAnd(c *EAndContext)

	// ExitEMethodCall is called when exiting the EMethodCall production.
	ExitEMethodCall(c *EMethodCallContext)

	// ExitEParen is called when exiting the EParen production.
	ExitEParen(c *EParenContext)

	// ExitEFalse is called when exiting the EFalse production.
	ExitEFalse(c *EFalseContext)

	// ExitEIndex is called when exiting the EIndex production.
	ExitEIndex(c *EIndexContext)

	// ExitEAddOp is called when exiting the EAddOp production.
	ExitEAddOp(c *EAddOpContext)

	// ExitENull is called when exiting the ENull production.
	ExitENull(c *ENullContext)

	// ExitEAttr is called when exiting the EAttr production.
	ExitEAttr(c *EAttrContext)

	// ExitUnOp is called when exiting the unOp production.
	ExitUnOp(c *UnOpContext)

	// ExitAddOp is called when exiting the addOp production.
	ExitAddOp(c *AddOpContext)

	// ExitMulOp is called when exiting the mulOp production.
	ExitMulOp(c *MulOpContext)

	// ExitRelOp is called when exiting the relOp production.
	ExitRelOp(c *RelOpContext)
}
