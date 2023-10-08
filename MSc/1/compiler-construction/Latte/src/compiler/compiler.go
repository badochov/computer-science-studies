package compiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/badochov/MRJP/Latte/src/IR/ast"
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
	"github.com/badochov/MRJP/Latte/src/IR/quadruple/converter"
	"github.com/badochov/MRJP/Latte/src/compiler/block_termination_fixer"
	"github.com/badochov/MRJP/Latte/src/compiler/compile"
	"github.com/badochov/MRJP/Latte/src/compiler/optimizer"
	"github.com/badochov/MRJP/Latte/src/compiler/ssa_converter"
	"github.com/badochov/MRJP/Latte/src/status"
	"github.com/badochov/MRJP/Latte/src/topdefs/signatures"
)

func Compile[Data antlr.ParserRuleContext](program ast.IProgram[Data], topDefs signatures.TopDefs, name string) {
	quadrupleCode := converter.ConvertAst(program, topDefs)

	optimized := optimize(quadrupleCode)

	latteCode := compile.Compile(optimized)
	if err := latteCode.Save(name); err != nil {
		status.Fatal(err)
	}
}

func optimize(quadrupleCode quadruple.Program) quadruple.Program {
	fixedTermination := block_termination_fixer.FixBlockTerminations(quadrupleCode)
	ssa := ssa_converter.ToSSA(fixedTermination)
	constPropagated := optimizer.ConstPropagation(ssa)
	lcse := optimizer.LCSE(constPropagated)

	return lcse
}
