package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/badochov/MRJP/Latte/src/compiler"
	"github.com/badochov/MRJP/Latte/src/frontend"
	"github.com/badochov/MRJP/Latte/src/parser"
	"github.com/badochov/MRJP/Latte/src/status"
	"os"
)

func ValidateArgs(args []string) {
	if len(args) != 2 {
		status.Error("Usage: %s <file_name>", args)
	}
}

func GetFileName(args []string) string {
	return args[1]
}

func getFileName(args []string) string {
	ValidateArgs(args)
	return GetFileName(args)
}

func getProgram(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		status.Error("Error while reading file: %s", err)
	}
	return string(data)
}

func toTokenStream(program string) antlr.TokenStream {
	inputStream := antlr.NewInputStream(program)
	lexer := parser.NewLatteLexer(inputStream)
	return antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
}

func main() {
	fileName := getFileName(os.Args)
	program := getProgram(fileName)
	stream := toTokenStream(program)

	converted, topDefs := frontend.CheckAndSimplifyConst(stream)

	compiler.Compile(converted, topDefs, fileName)

	status.Ok()
}
