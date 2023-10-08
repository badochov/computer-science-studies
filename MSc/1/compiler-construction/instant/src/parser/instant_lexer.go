// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type InstantLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var instantlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func instantlexerLexerInit() {
	staticData := &instantlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "';'", "'='", "'-'", "'+'", "'('", "')'", "'*'", "'/'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "INT", "IDENT", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "Letter",
		"Capital", "Small", "Digit", "INT", "IDENT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 78, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 50, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 4, 12, 59, 8, 12, 11, 12, 12, 12, 60, 1, 13, 1, 13,
		1, 13, 1, 13, 5, 13, 67, 8, 13, 10, 13, 12, 13, 70, 9, 13, 1, 14, 4, 14,
		73, 8, 14, 11, 14, 12, 14, 74, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 0, 19, 0, 21, 0, 23, 0, 25, 9, 27,
		10, 29, 11, 1, 0, 5, 1, 0, 65, 90, 1, 0, 97, 122, 1, 0, 48, 57, 2, 0, 39,
		39, 95, 95, 3, 0, 9, 10, 13, 13, 32, 32, 79, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 33, 1, 0, 0, 0, 5,
		35, 1, 0, 0, 0, 7, 37, 1, 0, 0, 0, 9, 39, 1, 0, 0, 0, 11, 41, 1, 0, 0,
		0, 13, 43, 1, 0, 0, 0, 15, 45, 1, 0, 0, 0, 17, 49, 1, 0, 0, 0, 19, 51,
		1, 0, 0, 0, 21, 53, 1, 0, 0, 0, 23, 55, 1, 0, 0, 0, 25, 58, 1, 0, 0, 0,
		27, 62, 1, 0, 0, 0, 29, 72, 1, 0, 0, 0, 31, 32, 5, 59, 0, 0, 32, 2, 1,
		0, 0, 0, 33, 34, 5, 61, 0, 0, 34, 4, 1, 0, 0, 0, 35, 36, 5, 45, 0, 0, 36,
		6, 1, 0, 0, 0, 37, 38, 5, 43, 0, 0, 38, 8, 1, 0, 0, 0, 39, 40, 5, 40, 0,
		0, 40, 10, 1, 0, 0, 0, 41, 42, 5, 41, 0, 0, 42, 12, 1, 0, 0, 0, 43, 44,
		5, 42, 0, 0, 44, 14, 1, 0, 0, 0, 45, 46, 5, 47, 0, 0, 46, 16, 1, 0, 0,
		0, 47, 50, 3, 19, 9, 0, 48, 50, 3, 21, 10, 0, 49, 47, 1, 0, 0, 0, 49, 48,
		1, 0, 0, 0, 50, 18, 1, 0, 0, 0, 51, 52, 7, 0, 0, 0, 52, 20, 1, 0, 0, 0,
		53, 54, 7, 1, 0, 0, 54, 22, 1, 0, 0, 0, 55, 56, 7, 2, 0, 0, 56, 24, 1,
		0, 0, 0, 57, 59, 3, 23, 11, 0, 58, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0,
		60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 26, 1, 0, 0, 0, 62, 68, 3,
		17, 8, 0, 63, 67, 3, 17, 8, 0, 64, 67, 7, 3, 0, 0, 65, 67, 3, 23, 11, 0,
		66, 63, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 65, 1, 0, 0, 0, 67, 70, 1,
		0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 28, 1, 0, 0, 0, 70,
		68, 1, 0, 0, 0, 71, 73, 7, 4, 0, 0, 72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0,
		0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77,
		6, 14, 0, 0, 77, 30, 1, 0, 0, 0, 6, 0, 49, 60, 66, 68, 74, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// InstantLexerInit initializes any static state used to implement InstantLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewInstantLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func InstantLexerInit() {
	staticData := &instantlexerLexerStaticData
	staticData.once.Do(instantlexerLexerInit)
}

// NewInstantLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewInstantLexer(input antlr.CharStream) *InstantLexer {
	InstantLexerInit()
	l := new(InstantLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &instantlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Instant.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// InstantLexer tokens.
const (
	InstantLexerT__0  = 1
	InstantLexerT__1  = 2
	InstantLexerT__2  = 3
	InstantLexerT__3  = 4
	InstantLexerT__4  = 5
	InstantLexerT__5  = 6
	InstantLexerT__6  = 7
	InstantLexerT__7  = 8
	InstantLexerINT   = 9
	InstantLexerIDENT = 10
	InstantLexerWS    = 11
)
