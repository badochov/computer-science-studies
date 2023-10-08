// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Instant

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type InstantParser struct {
	*antlr.BaseParser
}

var instantParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func instantParserInit() {
	staticData := &instantParserStaticData
	staticData.literalNames = []string{
		"", "';'", "'='", "'-'", "'+'", "'('", "')'", "'*'", "'/'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "INT", "IDENT", "WS",
	}
	staticData.ruleNames = []string{
		"program", "stmt", "expr", "mulOp",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 54, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 0, 5, 0, 12, 8, 0, 10, 0, 12, 0, 15, 9, 0, 3, 0, 17, 8, 0, 1, 0,
		3, 0, 20, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 26, 8, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 35, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 47, 8, 2, 10, 2, 12, 2, 50, 9, 2, 1,
		3, 1, 3, 1, 3, 0, 1, 4, 4, 0, 2, 4, 6, 0, 1, 1, 0, 7, 8, 58, 0, 16, 1,
		0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 34, 1, 0, 0, 0, 6, 51, 1, 0, 0, 0, 8, 13,
		3, 2, 1, 0, 9, 10, 5, 1, 0, 0, 10, 12, 3, 2, 1, 0, 11, 9, 1, 0, 0, 0, 12,
		15, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 17, 1, 0, 0,
		0, 15, 13, 1, 0, 0, 0, 16, 8, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 19, 1,
		0, 0, 0, 18, 20, 5, 1, 0, 0, 19, 18, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20,
		1, 1, 0, 0, 0, 21, 22, 5, 10, 0, 0, 22, 23, 5, 2, 0, 0, 23, 26, 3, 4, 2,
		0, 24, 26, 3, 4, 2, 0, 25, 21, 1, 0, 0, 0, 25, 24, 1, 0, 0, 0, 26, 3, 1,
		0, 0, 0, 27, 28, 6, 2, -1, 0, 28, 35, 5, 10, 0, 0, 29, 35, 5, 9, 0, 0,
		30, 31, 5, 5, 0, 0, 31, 32, 3, 4, 2, 0, 32, 33, 5, 6, 0, 0, 33, 35, 1,
		0, 0, 0, 34, 27, 1, 0, 0, 0, 34, 29, 1, 0, 0, 0, 34, 30, 1, 0, 0, 0, 35,
		48, 1, 0, 0, 0, 36, 37, 10, 4, 0, 0, 37, 38, 3, 6, 3, 0, 38, 39, 3, 4,
		2, 5, 39, 47, 1, 0, 0, 0, 40, 41, 10, 3, 0, 0, 41, 42, 5, 3, 0, 0, 42,
		47, 3, 4, 2, 4, 43, 44, 10, 2, 0, 0, 44, 45, 5, 4, 0, 0, 45, 47, 3, 4,
		2, 2, 46, 36, 1, 0, 0, 0, 46, 40, 1, 0, 0, 0, 46, 43, 1, 0, 0, 0, 47, 50,
		1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 5, 1, 0, 0, 0,
		50, 48, 1, 0, 0, 0, 51, 52, 7, 0, 0, 0, 52, 7, 1, 0, 0, 0, 7, 13, 16, 19,
		25, 34, 46, 48,
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

// InstantParserInit initializes any static state used to implement InstantParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInstantParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InstantParserInit() {
	staticData := &instantParserStaticData
	staticData.once.Do(instantParserInit)
}

// NewInstantParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInstantParser(input antlr.TokenStream) *InstantParser {
	InstantParserInit()
	this := new(InstantParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &instantParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// InstantParser tokens.
const (
	InstantParserEOF   = antlr.TokenEOF
	InstantParserT__0  = 1
	InstantParserT__1  = 2
	InstantParserT__2  = 3
	InstantParserT__3  = 4
	InstantParserT__4  = 5
	InstantParserT__5  = 6
	InstantParserT__6  = 7
	InstantParserT__7  = 8
	InstantParserINT   = 9
	InstantParserIDENT = 10
	InstantParserWS    = 11
)

// InstantParser rules.
const (
	InstantParserRULE_program = 0
	InstantParserRULE_stmt    = 1
	InstantParserRULE_expr    = 2
	InstantParserRULE_mulOp   = 3
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InstantParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InstantParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *InstantParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InstantParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1568) != 0 {
		{
			p.SetState(8)
			p.Stmt()
		}
		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(9)
					p.Match(InstantParserT__0)
				}
				{
					p.SetState(10)
					p.Stmt()
				}

			}
			p.SetState(15)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
		}

	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == InstantParserT__0 {
		{
			p.SetState(18)
			p.Match(InstantParserT__0)
		}

	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InstantParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InstantParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyFrom(ctx *StmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SAssignContext struct {
	*StmtContext
}

func NewSAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SAssignContext {
	var p = new(SAssignContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SAssignContext) IDENT() antlr.TerminalNode {
	return s.GetToken(InstantParserIDENT, 0)
}

func (s *SAssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterSAssign(s)
	}
}

func (s *SAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitSAssign(s)
	}
}

type SExprContext struct {
	*StmtContext
}

func NewSExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SExprContext {
	var p = new(SExprContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterSExpr(s)
	}
}

func (s *SExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitSExpr(s)
	}
}

func (p *InstantParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InstantParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.Match(InstantParserIDENT)
		}
		{
			p.SetState(22)
			p.Match(InstantParserT__1)
		}
		{
			p.SetState(23)
			p.expr(0)
		}

	case 2:
		localctx = NewSExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.expr(0)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InstantParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InstantParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EIdContext struct {
	*ExprContext
}

func NewEIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EIdContext {
	var p = new(EIdContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIdContext) IDENT() antlr.TerminalNode {
	return s.GetToken(InstantParserIDENT, 0)
}

func (s *EIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterEId(s)
	}
}

func (s *EIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitEId(s)
	}
}

type EParenContext struct {
	*ExprContext
}

func NewEParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EParenContext {
	var p = new(EParenContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EParenContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterEParen(s)
	}
}

func (s *EParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitEParen(s)
	}
}

type EIntContext struct {
	*ExprContext
}

func NewEIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EIntContext {
	var p = new(EIntContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIntContext) INT() antlr.TerminalNode {
	return s.GetToken(InstantParserINT, 0)
}

func (s *EIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterEInt(s)
	}
}

func (s *EIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitEInt(s)
	}
}

type ESubContext struct {
	*ExprContext
}

func NewESubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ESubContext {
	var p = new(ESubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ESubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ESubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ESubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ESubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterESub(s)
	}
}

func (s *ESubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitESub(s)
	}
}

type EAddContext struct {
	*ExprContext
}

func NewEAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAddContext {
	var p = new(EAddContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAddContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EAddContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EAddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterEAdd(s)
	}
}

func (s *EAddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitEAdd(s)
	}
}

type EMulContext struct {
	*ExprContext
}

func NewEMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EMulContext {
	var p = new(EMulContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EMulContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EMulContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EMulContext) MulOp() IMulOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulOpContext)
}

func (s *EMulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterEMul(s)
	}
}

func (s *EMulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitEMul(s)
	}
}

func (p *InstantParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *InstantParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, InstantParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case InstantParserIDENT:
		localctx = NewEIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(28)
			p.Match(InstantParserIDENT)
		}

	case InstantParserINT:
		localctx = NewEIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(29)
			p.Match(InstantParserINT)
		}

	case InstantParserT__4:
		localctx = NewEParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.Match(InstantParserT__4)
		}
		{
			p.SetState(31)
			p.expr(0)
		}
		{
			p.SetState(32)
			p.Match(InstantParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEMulContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InstantParserRULE_expr)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(37)
					p.MulOp()
				}
				{
					p.SetState(38)
					p.expr(5)
				}

			case 2:
				localctx = NewESubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InstantParserRULE_expr)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(41)
					p.Match(InstantParserT__2)
				}
				{
					p.SetState(42)
					p.expr(4)
				}

			case 3:
				localctx = NewEAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InstantParserRULE_expr)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(44)
					p.Match(InstantParserT__3)
				}
				{
					p.SetState(45)
					p.expr(2)
				}

			}

		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IMulOpContext is an interface to support dynamic dispatch.
type IMulOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulOpContext differentiates from other interfaces.
	IsMulOpContext()
}

type MulOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulOpContext() *MulOpContext {
	var p = new(MulOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InstantParserRULE_mulOp
	return p
}

func (*MulOpContext) IsMulOpContext() {}

func NewMulOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulOpContext {
	var p = new(MulOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InstantParserRULE_mulOp

	return p
}

func (s *MulOpContext) GetParser() antlr.Parser { return s.parser }
func (s *MulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.EnterMulOp(s)
	}
}

func (s *MulOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstantListener); ok {
		listenerT.ExitMulOp(s)
	}
}

func (p *InstantParser) MulOp() (localctx IMulOpContext) {
	this := p
	_ = this

	localctx = NewMulOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InstantParserRULE_mulOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		_la = p.GetTokenStream().LA(1)

		if !(_la == InstantParserT__6 || _la == InstantParserT__7) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *InstantParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *InstantParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
