// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Instant

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseInstantListener is a complete listener for a parse tree produced by InstantParser.
type BaseInstantListener struct{}

var _ InstantListener = &BaseInstantListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInstantListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInstantListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInstantListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInstantListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseInstantListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseInstantListener) ExitProgram(ctx *ProgramContext) {}

// EnterSAssign is called when production SAssign is entered.
func (s *BaseInstantListener) EnterSAssign(ctx *SAssignContext) {}

// ExitSAssign is called when production SAssign is exited.
func (s *BaseInstantListener) ExitSAssign(ctx *SAssignContext) {}

// EnterSExpr is called when production SExpr is entered.
func (s *BaseInstantListener) EnterSExpr(ctx *SExprContext) {}

// ExitSExpr is called when production SExpr is exited.
func (s *BaseInstantListener) ExitSExpr(ctx *SExprContext) {}

// EnterEId is called when production EId is entered.
func (s *BaseInstantListener) EnterEId(ctx *EIdContext) {}

// ExitEId is called when production EId is exited.
func (s *BaseInstantListener) ExitEId(ctx *EIdContext) {}

// EnterEParen is called when production EParen is entered.
func (s *BaseInstantListener) EnterEParen(ctx *EParenContext) {}

// ExitEParen is called when production EParen is exited.
func (s *BaseInstantListener) ExitEParen(ctx *EParenContext) {}

// EnterEInt is called when production EInt is entered.
func (s *BaseInstantListener) EnterEInt(ctx *EIntContext) {}

// ExitEInt is called when production EInt is exited.
func (s *BaseInstantListener) ExitEInt(ctx *EIntContext) {}

// EnterESub is called when production ESub is entered.
func (s *BaseInstantListener) EnterESub(ctx *ESubContext) {}

// ExitESub is called when production ESub is exited.
func (s *BaseInstantListener) ExitESub(ctx *ESubContext) {}

// EnterEAdd is called when production EAdd is entered.
func (s *BaseInstantListener) EnterEAdd(ctx *EAddContext) {}

// ExitEAdd is called when production EAdd is exited.
func (s *BaseInstantListener) ExitEAdd(ctx *EAddContext) {}

// EnterEMul is called when production EMul is entered.
func (s *BaseInstantListener) EnterEMul(ctx *EMulContext) {}

// ExitEMul is called when production EMul is exited.
func (s *BaseInstantListener) ExitEMul(ctx *EMulContext) {}

// EnterMulOp is called when production mulOp is entered.
func (s *BaseInstantListener) EnterMulOp(ctx *MulOpContext) {}

// ExitMulOp is called when production mulOp is exited.
func (s *BaseInstantListener) ExitMulOp(ctx *MulOpContext) {}
