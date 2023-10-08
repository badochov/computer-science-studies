package frontend

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/badochov/MRJP/Latte/src/IR/ast"
	"github.com/badochov/MRJP/Latte/src/IR/ast/converter"
	"github.com/badochov/MRJP/Latte/src/frontend/error_listener"
	"github.com/badochov/MRJP/Latte/src/frontend/return_checker"
	"github.com/badochov/MRJP/Latte/src/frontend/tree_simplifier"
	"github.com/badochov/MRJP/Latte/src/frontend/type_checker"
	"github.com/badochov/MRJP/Latte/src/parser"
	"github.com/badochov/MRJP/Latte/src/topdefs/signatures"
)

func lexicalAnalysis(stream antlr.TokenStream) {
	defer stream.Seek(0)
	p := parser.NewLatteParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(error_listener.New())
	antlr.ParseTreeWalkerDefault.Walk(&parser.BaseLatteListener{}, p.Program())
}

func semanticAnalysis(tree antlr.ParseTree, topDefs signatures.TopDefs) {
	tc := type_checker.New(topDefs)
	tc.Visit(tree)

	tc.Validate()
}

func getTopDefs(tree antlr.ParseTree) signatures.TopDefs {
	return signatures.Get(tree)
}

func getParseTree(stream antlr.TokenStream) antlr.ParseTree {
	defer stream.Seek(0)
	p := parser.NewLatteParser(stream)
	return p.Program()
}

func CheckAndSimplifyConst(stream antlr.TokenStream) (ast.IProgram[antlr.ParserRuleContext], signatures.TopDefs) {
	lexicalAnalysis(stream)

	t := getParseTree(stream)
	topDefs := getTopDefs(t)
	semanticAnalysis(t, topDefs)

	customAst := converter.ConvertAst(t)
	simplified := tree_simplifier.Simplify(customAst, topDefs)
	return_checker.Check(simplified)

	return simplified, topDefs
}
