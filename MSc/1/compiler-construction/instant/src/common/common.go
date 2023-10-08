package common

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser2 "github.com/badochov/MRJP/instant/src/parser"
	"os"
)

func ValidateArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("%s <file_name>", args)
	}
	return nil
}

func GetFileName(args []string) string {
	return args[1]
}

type CompileFunction = func(program string) error

func Main(compileFn CompileFunction) error {
	return main(os.Args, compileFn)
}

func main(args []string, compileFn CompileFunction) error {
	if err := ValidateArgs(args); err != nil {
		return err
	}

	fileName := GetFileName(args)

	return compileFn(fileName)
}

type ListenerCodeWriter interface {
	parser2.InstantListener
	Save(originalName string) error
}

func CreateCompileFunction(listener ListenerCodeWriter) CompileFunction {
	return func(fileName string) error {
		data, err := os.ReadFile(fileName)
		if err != nil {
			return err
		}

		is := antlr.NewInputStream(string(data))

		lexer := parser2.NewInstantLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser2.NewInstantParser(stream)
		antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

		return listener.Save(fileName)
	}
}
