package error_listener

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/badochov/MRJP/Latte/src/status"
)

type ErrorListener struct{}

func (el ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	status.Error("line %d:%d %s\n", line, column, msg)
}

func (el ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (el ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (el ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

var _ antlr.ErrorListener = &ErrorListener{}

func New() *ErrorListener {
	return &ErrorListener{}
}
