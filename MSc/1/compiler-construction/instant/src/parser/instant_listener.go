// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Instant

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// InstantListener is a complete listener for a parse tree produced by InstantParser.
type InstantListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSAssign is called when entering the SAssign production.
	EnterSAssign(c *SAssignContext)

	// EnterSExpr is called when entering the SExpr production.
	EnterSExpr(c *SExprContext)

	// EnterEId is called when entering the EId production.
	EnterEId(c *EIdContext)

	// EnterEParen is called when entering the EParen production.
	EnterEParen(c *EParenContext)

	// EnterEInt is called when entering the EInt production.
	EnterEInt(c *EIntContext)

	// EnterESub is called when entering the ESub production.
	EnterESub(c *ESubContext)

	// EnterEAdd is called when entering the EAdd production.
	EnterEAdd(c *EAddContext)

	// EnterEMul is called when entering the EMul production.
	EnterEMul(c *EMulContext)

	// EnterMulOp is called when entering the mulOp production.
	EnterMulOp(c *MulOpContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSAssign is called when exiting the SAssign production.
	ExitSAssign(c *SAssignContext)

	// ExitSExpr is called when exiting the SExpr production.
	ExitSExpr(c *SExprContext)

	// ExitEId is called when exiting the EId production.
	ExitEId(c *EIdContext)

	// ExitEParen is called when exiting the EParen production.
	ExitEParen(c *EParenContext)

	// ExitEInt is called when exiting the EInt production.
	ExitEInt(c *EIntContext)

	// ExitESub is called when exiting the ESub production.
	ExitESub(c *ESubContext)

	// ExitEAdd is called when exiting the EAdd production.
	ExitEAdd(c *EAddContext)

	// ExitEMul is called when exiting the EMul production.
	ExitEMul(c *EMulContext)

	// ExitMulOp is called when exiting the mulOp production.
	ExitMulOp(c *MulOpContext)
}
