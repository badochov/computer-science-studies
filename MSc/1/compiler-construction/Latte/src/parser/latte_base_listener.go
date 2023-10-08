// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Latte

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseLatteListener is a complete listener for a parse tree produced by LatteParser.
type BaseLatteListener struct{}

var _ LatteListener = &BaseLatteListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLatteListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLatteListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLatteListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLatteListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLatteListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLatteListener) ExitProgram(ctx *ProgramContext) {}

// EnterDefFunc is called when production DefFunc is entered.
func (s *BaseLatteListener) EnterDefFunc(ctx *DefFuncContext) {}

// ExitDefFunc is called when production DefFunc is exited.
func (s *BaseLatteListener) ExitDefFunc(ctx *DefFuncContext) {}

// EnterDefClass is called when production DefClass is entered.
func (s *BaseLatteListener) EnterDefClass(ctx *DefClassContext) {}

// ExitDefClass is called when production DefClass is exited.
func (s *BaseLatteListener) ExitDefClass(ctx *DefClassContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseLatteListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseLatteListener) ExitArg(ctx *ArgContext) {}

// EnterRef is called when production ref is entered.
func (s *BaseLatteListener) EnterRef(ctx *RefContext) {}

// ExitRef is called when production ref is exited.
func (s *BaseLatteListener) ExitRef(ctx *RefContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLatteListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLatteListener) ExitBlock(ctx *BlockContext) {}

// EnterEmpty is called when production Empty is entered.
func (s *BaseLatteListener) EnterEmpty(ctx *EmptyContext) {}

// ExitEmpty is called when production Empty is exited.
func (s *BaseLatteListener) ExitEmpty(ctx *EmptyContext) {}

// EnterIndexAss is called when production IndexAss is entered.
func (s *BaseLatteListener) EnterIndexAss(ctx *IndexAssContext) {}

// ExitIndexAss is called when production IndexAss is exited.
func (s *BaseLatteListener) ExitIndexAss(ctx *IndexAssContext) {}

// EnterAttrAss is called when production AttrAss is entered.
func (s *BaseLatteListener) EnterAttrAss(ctx *AttrAssContext) {}

// ExitAttrAss is called when production AttrAss is exited.
func (s *BaseLatteListener) ExitAttrAss(ctx *AttrAssContext) {}

// EnterBlockStmt is called when production BlockStmt is entered.
func (s *BaseLatteListener) EnterBlockStmt(ctx *BlockStmtContext) {}

// ExitBlockStmt is called when production BlockStmt is exited.
func (s *BaseLatteListener) ExitBlockStmt(ctx *BlockStmtContext) {}

// EnterDecl is called when production Decl is entered.
func (s *BaseLatteListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production Decl is exited.
func (s *BaseLatteListener) ExitDecl(ctx *DeclContext) {}

// EnterAss is called when production Ass is entered.
func (s *BaseLatteListener) EnterAss(ctx *AssContext) {}

// ExitAss is called when production Ass is exited.
func (s *BaseLatteListener) ExitAss(ctx *AssContext) {}

// EnterIncr is called when production Incr is entered.
func (s *BaseLatteListener) EnterIncr(ctx *IncrContext) {}

// ExitIncr is called when production Incr is exited.
func (s *BaseLatteListener) ExitIncr(ctx *IncrContext) {}

// EnterDecr is called when production Decr is entered.
func (s *BaseLatteListener) EnterDecr(ctx *DecrContext) {}

// ExitDecr is called when production Decr is exited.
func (s *BaseLatteListener) ExitDecr(ctx *DecrContext) {}

// EnterRet is called when production Ret is entered.
func (s *BaseLatteListener) EnterRet(ctx *RetContext) {}

// ExitRet is called when production Ret is exited.
func (s *BaseLatteListener) ExitRet(ctx *RetContext) {}

// EnterVRet is called when production VRet is entered.
func (s *BaseLatteListener) EnterVRet(ctx *VRetContext) {}

// ExitVRet is called when production VRet is exited.
func (s *BaseLatteListener) ExitVRet(ctx *VRetContext) {}

// EnterCond is called when production Cond is entered.
func (s *BaseLatteListener) EnterCond(ctx *CondContext) {}

// ExitCond is called when production Cond is exited.
func (s *BaseLatteListener) ExitCond(ctx *CondContext) {}

// EnterCondElse is called when production CondElse is entered.
func (s *BaseLatteListener) EnterCondElse(ctx *CondElseContext) {}

// ExitCondElse is called when production CondElse is exited.
func (s *BaseLatteListener) ExitCondElse(ctx *CondElseContext) {}

// EnterWhile is called when production While is entered.
func (s *BaseLatteListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production While is exited.
func (s *BaseLatteListener) ExitWhile(ctx *WhileContext) {}

// EnterForEach is called when production ForEach is entered.
func (s *BaseLatteListener) EnterForEach(ctx *ForEachContext) {}

// ExitForEach is called when production ForEach is exited.
func (s *BaseLatteListener) ExitForEach(ctx *ForEachContext) {}

// EnterSExp is called when production SExp is entered.
func (s *BaseLatteListener) EnterSExp(ctx *SExpContext) {}

// ExitSExp is called when production SExp is exited.
func (s *BaseLatteListener) ExitSExp(ctx *SExpContext) {}

// EnterTypeArray is called when production TypeArray is entered.
func (s *BaseLatteListener) EnterTypeArray(ctx *TypeArrayContext) {}

// ExitTypeArray is called when production TypeArray is exited.
func (s *BaseLatteListener) ExitTypeArray(ctx *TypeArrayContext) {}

// EnterTypeSingle is called when production TypeSingle is entered.
func (s *BaseLatteListener) EnterTypeSingle(ctx *TypeSingleContext) {}

// ExitTypeSingle is called when production TypeSingle is exited.
func (s *BaseLatteListener) ExitTypeSingle(ctx *TypeSingleContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseLatteListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseLatteListener) ExitInt(ctx *IntContext) {}

// EnterStr is called when production Str is entered.
func (s *BaseLatteListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production Str is exited.
func (s *BaseLatteListener) ExitStr(ctx *StrContext) {}

// EnterBool is called when production Bool is entered.
func (s *BaseLatteListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production Bool is exited.
func (s *BaseLatteListener) ExitBool(ctx *BoolContext) {}

// EnterVoid is called when production Void is entered.
func (s *BaseLatteListener) EnterVoid(ctx *VoidContext) {}

// ExitVoid is called when production Void is exited.
func (s *BaseLatteListener) ExitVoid(ctx *VoidContext) {}

// EnterClassName is called when production ClassName is entered.
func (s *BaseLatteListener) EnterClassName(ctx *ClassNameContext) {}

// ExitClassName is called when production ClassName is exited.
func (s *BaseLatteListener) ExitClassName(ctx *ClassNameContext) {}

// EnterClassAttribute is called when production ClassAttribute is entered.
func (s *BaseLatteListener) EnterClassAttribute(ctx *ClassAttributeContext) {}

// ExitClassAttribute is called when production ClassAttribute is exited.
func (s *BaseLatteListener) ExitClassAttribute(ctx *ClassAttributeContext) {}

// EnterClassMethod is called when production ClassMethod is entered.
func (s *BaseLatteListener) EnterClassMethod(ctx *ClassMethodContext) {}

// ExitClassMethod is called when production ClassMethod is exited.
func (s *BaseLatteListener) ExitClassMethod(ctx *ClassMethodContext) {}

// EnterItem is called when production item is entered.
func (s *BaseLatteListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BaseLatteListener) ExitItem(ctx *ItemContext) {}

// EnterEId is called when production EId is entered.
func (s *BaseLatteListener) EnterEId(ctx *EIdContext) {}

// ExitEId is called when production EId is exited.
func (s *BaseLatteListener) ExitEId(ctx *EIdContext) {}

// EnterEFunCall is called when production EFunCall is entered.
func (s *BaseLatteListener) EnterEFunCall(ctx *EFunCallContext) {}

// ExitEFunCall is called when production EFunCall is exited.
func (s *BaseLatteListener) ExitEFunCall(ctx *EFunCallContext) {}

// EnterENewArray is called when production ENewArray is entered.
func (s *BaseLatteListener) EnterENewArray(ctx *ENewArrayContext) {}

// ExitENewArray is called when production ENewArray is exited.
func (s *BaseLatteListener) ExitENewArray(ctx *ENewArrayContext) {}

// EnterERelOp is called when production ERelOp is entered.
func (s *BaseLatteListener) EnterERelOp(ctx *ERelOpContext) {}

// ExitERelOp is called when production ERelOp is exited.
func (s *BaseLatteListener) ExitERelOp(ctx *ERelOpContext) {}

// EnterETrue is called when production ETrue is entered.
func (s *BaseLatteListener) EnterETrue(ctx *ETrueContext) {}

// ExitETrue is called when production ETrue is exited.
func (s *BaseLatteListener) ExitETrue(ctx *ETrueContext) {}

// EnterEOr is called when production EOr is entered.
func (s *BaseLatteListener) EnterEOr(ctx *EOrContext) {}

// ExitEOr is called when production EOr is exited.
func (s *BaseLatteListener) ExitEOr(ctx *EOrContext) {}

// EnterEInt is called when production EInt is entered.
func (s *BaseLatteListener) EnterEInt(ctx *EIntContext) {}

// ExitEInt is called when production EInt is exited.
func (s *BaseLatteListener) ExitEInt(ctx *EIntContext) {}

// EnterENewClass is called when production ENewClass is entered.
func (s *BaseLatteListener) EnterENewClass(ctx *ENewClassContext) {}

// ExitENewClass is called when production ENewClass is exited.
func (s *BaseLatteListener) ExitENewClass(ctx *ENewClassContext) {}

// EnterEUnOp is called when production EUnOp is entered.
func (s *BaseLatteListener) EnterEUnOp(ctx *EUnOpContext) {}

// ExitEUnOp is called when production EUnOp is exited.
func (s *BaseLatteListener) ExitEUnOp(ctx *EUnOpContext) {}

// EnterEStr is called when production EStr is entered.
func (s *BaseLatteListener) EnterEStr(ctx *EStrContext) {}

// ExitEStr is called when production EStr is exited.
func (s *BaseLatteListener) ExitEStr(ctx *EStrContext) {}

// EnterEMulOp is called when production EMulOp is entered.
func (s *BaseLatteListener) EnterEMulOp(ctx *EMulOpContext) {}

// ExitEMulOp is called when production EMulOp is exited.
func (s *BaseLatteListener) ExitEMulOp(ctx *EMulOpContext) {}

// EnterEAnd is called when production EAnd is entered.
func (s *BaseLatteListener) EnterEAnd(ctx *EAndContext) {}

// ExitEAnd is called when production EAnd is exited.
func (s *BaseLatteListener) ExitEAnd(ctx *EAndContext) {}

// EnterEMethodCall is called when production EMethodCall is entered.
func (s *BaseLatteListener) EnterEMethodCall(ctx *EMethodCallContext) {}

// ExitEMethodCall is called when production EMethodCall is exited.
func (s *BaseLatteListener) ExitEMethodCall(ctx *EMethodCallContext) {}

// EnterEParen is called when production EParen is entered.
func (s *BaseLatteListener) EnterEParen(ctx *EParenContext) {}

// ExitEParen is called when production EParen is exited.
func (s *BaseLatteListener) ExitEParen(ctx *EParenContext) {}

// EnterEFalse is called when production EFalse is entered.
func (s *BaseLatteListener) EnterEFalse(ctx *EFalseContext) {}

// ExitEFalse is called when production EFalse is exited.
func (s *BaseLatteListener) ExitEFalse(ctx *EFalseContext) {}

// EnterEIndex is called when production EIndex is entered.
func (s *BaseLatteListener) EnterEIndex(ctx *EIndexContext) {}

// ExitEIndex is called when production EIndex is exited.
func (s *BaseLatteListener) ExitEIndex(ctx *EIndexContext) {}

// EnterEAddOp is called when production EAddOp is entered.
func (s *BaseLatteListener) EnterEAddOp(ctx *EAddOpContext) {}

// ExitEAddOp is called when production EAddOp is exited.
func (s *BaseLatteListener) ExitEAddOp(ctx *EAddOpContext) {}

// EnterENull is called when production ENull is entered.
func (s *BaseLatteListener) EnterENull(ctx *ENullContext) {}

// ExitENull is called when production ENull is exited.
func (s *BaseLatteListener) ExitENull(ctx *ENullContext) {}

// EnterEAttr is called when production EAttr is entered.
func (s *BaseLatteListener) EnterEAttr(ctx *EAttrContext) {}

// ExitEAttr is called when production EAttr is exited.
func (s *BaseLatteListener) ExitEAttr(ctx *EAttrContext) {}

// EnterUnOp is called when production unOp is entered.
func (s *BaseLatteListener) EnterUnOp(ctx *UnOpContext) {}

// ExitUnOp is called when production unOp is exited.
func (s *BaseLatteListener) ExitUnOp(ctx *UnOpContext) {}

// EnterAddOp is called when production addOp is entered.
func (s *BaseLatteListener) EnterAddOp(ctx *AddOpContext) {}

// ExitAddOp is called when production addOp is exited.
func (s *BaseLatteListener) ExitAddOp(ctx *AddOpContext) {}

// EnterMulOp is called when production mulOp is entered.
func (s *BaseLatteListener) EnterMulOp(ctx *MulOpContext) {}

// ExitMulOp is called when production mulOp is exited.
func (s *BaseLatteListener) ExitMulOp(ctx *MulOpContext) {}

// EnterRelOp is called when production relOp is entered.
func (s *BaseLatteListener) EnterRelOp(ctx *RelOpContext) {}

// ExitRelOp is called when production relOp is exited.
func (s *BaseLatteListener) ExitRelOp(ctx *RelOpContext) {}
