package status

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
)

func Ok() {
	fmt.Fprintln(os.Stderr, "OK")
	os.Exit(0)
}

func Error(format string, args ...any) {
	fmt.Fprintln(os.Stderr, "ERROR")
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func HandleError(err error, ruleCtx antlr.ParserRuleContext) {
	if err == nil {
		return
	}
	Error("line %d column %d: %v", ruleCtx.GetStart().GetLine(), ruleCtx.GetStart().GetColumn(), err)
}

func Fatal(err error) {
	Error("%s", err)
}
