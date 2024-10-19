// Code generated from SCIMFilter.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filter // SCIMFilter
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SCIMFilterParser struct {
	*antlr.BaseParser
}

var SCIMFilterParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func scimfilterParserInit() {
	staticData := &SCIMFilterParserStaticData
	staticData.LiteralNames = []string{
		"", "'not'", "'('", "')'", "'['", "']'", "'pr'", "'and'", "'or'", "'eq'",
		"'ne'", "'co'", "'sw'", "'ew'", "'gt'", "'lt'", "'ge'", "'le'", "'\"'",
		"'\"\"'", "'-'", "'_'", "'.'", "'false'", "'null'", "'true'", "'\\r'",
		"'\\n'", "", "", "", "", "", "'\\'",
	}
	staticData.SymbolicNames = []string{
		"", "NOT", "OPEN_PARA", "CLOSE_PARA", "OPEN_BRACKET", "CLOSE_BRACKET",
		"PR", "AND", "OR", "EQ", "NE", "CO", "SW", "EW", "GT", "LT", "GE", "LE",
		"QUOTE", "DOUBLE_QUOTE", "HYPHEN", "UNDERSCORE", "DOT", "FALSE", "NULL",
		"TRUE", "NEW_RETURN", "NEW_LINE", "ALPHA", "DIGIT", "SP", "URI", "ATTR_NAME",
		"DOUBLE_SLASH", "QUOTE_SLASH", "ANY",
	}
	staticData.RuleNames = []string{
		"expression", "filter", "moreFilters", "filterExpression", "closedFilter",
		"notFilter", "valuePath", "attrExp", "filterOperator", "valuePresent",
		"compValue", "stringValue", "number", "compareOp", "attrPath", "subAttr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 35, 127, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 50, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 74, 8, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 85, 8, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 5, 11, 91, 8, 11, 10, 11, 12, 11, 94, 9, 11, 1, 11, 1, 11, 1,
		12, 4, 12, 99, 8, 12, 11, 12, 12, 12, 100, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 5, 14, 109, 8, 14, 10, 14, 12, 14, 112, 9, 14, 1, 14, 1,
		14, 5, 14, 116, 8, 14, 10, 14, 12, 14, 119, 9, 14, 1, 14, 3, 14, 122, 8,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 0, 0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 0, 3, 1, 0, 7, 8, 1, 0, 18, 18, 1, 0, 9, 17,
		126, 0, 32, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 49, 1,
		0, 0, 0, 8, 51, 1, 0, 0, 0, 10, 55, 1, 0, 0, 0, 12, 58, 1, 0, 0, 0, 14,
		73, 1, 0, 0, 0, 16, 75, 1, 0, 0, 0, 18, 77, 1, 0, 0, 0, 20, 84, 1, 0, 0,
		0, 22, 86, 1, 0, 0, 0, 24, 98, 1, 0, 0, 0, 26, 102, 1, 0, 0, 0, 28, 121,
		1, 0, 0, 0, 30, 123, 1, 0, 0, 0, 32, 33, 3, 2, 1, 0, 33, 34, 5, 0, 0, 1,
		34, 1, 1, 0, 0, 0, 35, 38, 3, 6, 3, 0, 36, 38, 3, 4, 2, 0, 37, 35, 1, 0,
		0, 0, 37, 36, 1, 0, 0, 0, 38, 3, 1, 0, 0, 0, 39, 40, 3, 6, 3, 0, 40, 41,
		5, 30, 0, 0, 41, 42, 3, 16, 8, 0, 42, 43, 5, 30, 0, 0, 43, 44, 3, 2, 1,
		0, 44, 5, 1, 0, 0, 0, 45, 50, 3, 14, 7, 0, 46, 50, 3, 8, 4, 0, 47, 50,
		3, 10, 5, 0, 48, 50, 3, 12, 6, 0, 49, 45, 1, 0, 0, 0, 49, 46, 1, 0, 0,
		0, 49, 47, 1, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 7, 1, 0, 0, 0, 51, 52, 5,
		2, 0, 0, 52, 53, 3, 2, 1, 0, 53, 54, 5, 3, 0, 0, 54, 9, 1, 0, 0, 0, 55,
		56, 5, 1, 0, 0, 56, 57, 3, 8, 4, 0, 57, 11, 1, 0, 0, 0, 58, 59, 3, 28,
		14, 0, 59, 60, 5, 4, 0, 0, 60, 61, 3, 2, 1, 0, 61, 62, 5, 5, 0, 0, 62,
		13, 1, 0, 0, 0, 63, 64, 3, 28, 14, 0, 64, 65, 5, 30, 0, 0, 65, 66, 3, 18,
		9, 0, 66, 74, 1, 0, 0, 0, 67, 68, 3, 28, 14, 0, 68, 69, 5, 30, 0, 0, 69,
		70, 3, 26, 13, 0, 70, 71, 5, 30, 0, 0, 71, 72, 3, 20, 10, 0, 72, 74, 1,
		0, 0, 0, 73, 63, 1, 0, 0, 0, 73, 67, 1, 0, 0, 0, 74, 15, 1, 0, 0, 0, 75,
		76, 7, 0, 0, 0, 76, 17, 1, 0, 0, 0, 77, 78, 5, 6, 0, 0, 78, 19, 1, 0, 0,
		0, 79, 85, 5, 23, 0, 0, 80, 85, 5, 24, 0, 0, 81, 85, 5, 25, 0, 0, 82, 85,
		3, 24, 12, 0, 83, 85, 3, 22, 11, 0, 84, 79, 1, 0, 0, 0, 84, 80, 1, 0, 0,
		0, 84, 81, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 83, 1, 0, 0, 0, 85, 21,
		1, 0, 0, 0, 86, 92, 5, 18, 0, 0, 87, 91, 8, 1, 0, 0, 88, 89, 5, 33, 0,
		0, 89, 91, 5, 18, 0, 0, 90, 87, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 94,
		1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 95, 96, 5, 18, 0, 0, 96, 23, 1, 0, 0, 0, 97, 99, 5,
		29, 0, 0, 98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0,
		100, 101, 1, 0, 0, 0, 101, 25, 1, 0, 0, 0, 102, 103, 7, 2, 0, 0, 103, 27,
		1, 0, 0, 0, 104, 105, 5, 31, 0, 0, 105, 106, 5, 20, 0, 0, 106, 110, 5,
		32, 0, 0, 107, 109, 3, 30, 15, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0,
		0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 122, 1, 0, 0, 0,
		112, 110, 1, 0, 0, 0, 113, 117, 5, 32, 0, 0, 114, 116, 3, 30, 15, 0, 115,
		114, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118,
		1, 0, 0, 0, 118, 122, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 122, 5, 32,
		0, 0, 121, 104, 1, 0, 0, 0, 121, 113, 1, 0, 0, 0, 121, 120, 1, 0, 0, 0,
		122, 29, 1, 0, 0, 0, 123, 124, 5, 22, 0, 0, 124, 125, 5, 32, 0, 0, 125,
		31, 1, 0, 0, 0, 10, 37, 49, 73, 84, 90, 92, 100, 110, 117, 121,
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

// SCIMFilterParserInit initializes any static state used to implement SCIMFilterParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSCIMFilterParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SCIMFilterParserInit() {
	staticData := &SCIMFilterParserStaticData
	staticData.once.Do(scimfilterParserInit)
}

// NewSCIMFilterParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSCIMFilterParser(input antlr.TokenStream) *SCIMFilterParser {
	SCIMFilterParserInit()
	this := new(SCIMFilterParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SCIMFilterParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SCIMFilter.g4"

	return this
}

// SCIMFilterParser tokens.
const (
	SCIMFilterParserEOF           = antlr.TokenEOF
	SCIMFilterParserNOT           = 1
	SCIMFilterParserOPEN_PARA     = 2
	SCIMFilterParserCLOSE_PARA    = 3
	SCIMFilterParserOPEN_BRACKET  = 4
	SCIMFilterParserCLOSE_BRACKET = 5
	SCIMFilterParserPR            = 6
	SCIMFilterParserAND           = 7
	SCIMFilterParserOR            = 8
	SCIMFilterParserEQ            = 9
	SCIMFilterParserNE            = 10
	SCIMFilterParserCO            = 11
	SCIMFilterParserSW            = 12
	SCIMFilterParserEW            = 13
	SCIMFilterParserGT            = 14
	SCIMFilterParserLT            = 15
	SCIMFilterParserGE            = 16
	SCIMFilterParserLE            = 17
	SCIMFilterParserQUOTE         = 18
	SCIMFilterParserDOUBLE_QUOTE  = 19
	SCIMFilterParserHYPHEN        = 20
	SCIMFilterParserUNDERSCORE    = 21
	SCIMFilterParserDOT           = 22
	SCIMFilterParserFALSE         = 23
	SCIMFilterParserNULL          = 24
	SCIMFilterParserTRUE          = 25
	SCIMFilterParserNEW_RETURN    = 26
	SCIMFilterParserNEW_LINE      = 27
	SCIMFilterParserALPHA         = 28
	SCIMFilterParserDIGIT         = 29
	SCIMFilterParserSP            = 30
	SCIMFilterParserURI           = 31
	SCIMFilterParserATTR_NAME     = 32
	SCIMFilterParserDOUBLE_SLASH  = 33
	SCIMFilterParserQUOTE_SLASH   = 34
	SCIMFilterParserANY           = 35
)

// SCIMFilterParser rules.
const (
	SCIMFilterParserRULE_expression       = 0
	SCIMFilterParserRULE_filter           = 1
	SCIMFilterParserRULE_moreFilters      = 2
	SCIMFilterParserRULE_filterExpression = 3
	SCIMFilterParserRULE_closedFilter     = 4
	SCIMFilterParserRULE_notFilter        = 5
	SCIMFilterParserRULE_valuePath        = 6
	SCIMFilterParserRULE_attrExp          = 7
	SCIMFilterParserRULE_filterOperator   = 8
	SCIMFilterParserRULE_valuePresent     = 9
	SCIMFilterParserRULE_compValue        = 10
	SCIMFilterParserRULE_stringValue      = 11
	SCIMFilterParserRULE_number           = 12
	SCIMFilterParserRULE_compareOp        = 13
	SCIMFilterParserRULE_attrPath         = 14
	SCIMFilterParserRULE_subAttr          = 15
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Filter() IFilterContext
	EOF() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Filter() IFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SCIMFilterParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SCIMFilterParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Filter()
	}
	{
		p.SetState(33)
		p.Match(SCIMFilterParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FilterExpression() IFilterExpressionContext
	MoreFilters() IMoreFiltersContext

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_filter
	return p
}

func InitEmptyFilterContext(p *FilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_filter
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) FilterExpression() IFilterExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterExpressionContext)
}

func (s *FilterContext) MoreFilters() IMoreFiltersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMoreFiltersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMoreFiltersContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *SCIMFilterParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SCIMFilterParserRULE_filter)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.FilterExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.MoreFilters()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMoreFiltersContext is an interface to support dynamic dispatch.
type IMoreFiltersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FilterExpression() IFilterExpressionContext
	AllSP() []antlr.TerminalNode
	SP(i int) antlr.TerminalNode
	FilterOperator() IFilterOperatorContext
	Filter() IFilterContext

	// IsMoreFiltersContext differentiates from other interfaces.
	IsMoreFiltersContext()
}

type MoreFiltersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoreFiltersContext() *MoreFiltersContext {
	var p = new(MoreFiltersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_moreFilters
	return p
}

func InitEmptyMoreFiltersContext(p *MoreFiltersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_moreFilters
}

func (*MoreFiltersContext) IsMoreFiltersContext() {}

func NewMoreFiltersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoreFiltersContext {
	var p = new(MoreFiltersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_moreFilters

	return p
}

func (s *MoreFiltersContext) GetParser() antlr.Parser { return s.parser }

func (s *MoreFiltersContext) FilterExpression() IFilterExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterExpressionContext)
}

func (s *MoreFiltersContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(SCIMFilterParserSP)
}

func (s *MoreFiltersContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserSP, i)
}

func (s *MoreFiltersContext) FilterOperator() IFilterOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterOperatorContext)
}

func (s *MoreFiltersContext) Filter() IFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *MoreFiltersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoreFiltersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoreFiltersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterMoreFilters(s)
	}
}

func (s *MoreFiltersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitMoreFilters(s)
	}
}

func (p *SCIMFilterParser) MoreFilters() (localctx IMoreFiltersContext) {
	localctx = NewMoreFiltersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SCIMFilterParserRULE_moreFilters)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.FilterExpression()
	}
	{
		p.SetState(40)
		p.Match(SCIMFilterParserSP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.FilterOperator()
	}
	{
		p.SetState(42)
		p.Match(SCIMFilterParserSP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Filter()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilterExpressionContext is an interface to support dynamic dispatch.
type IFilterExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttrExp() IAttrExpContext
	ClosedFilter() IClosedFilterContext
	NotFilter() INotFilterContext
	ValuePath() IValuePathContext

	// IsFilterExpressionContext differentiates from other interfaces.
	IsFilterExpressionContext()
}

type FilterExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterExpressionContext() *FilterExpressionContext {
	var p = new(FilterExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_filterExpression
	return p
}

func InitEmptyFilterExpressionContext(p *FilterExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_filterExpression
}

func (*FilterExpressionContext) IsFilterExpressionContext() {}

func NewFilterExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterExpressionContext {
	var p = new(FilterExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_filterExpression

	return p
}

func (s *FilterExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterExpressionContext) AttrExp() IAttrExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrExpContext)
}

func (s *FilterExpressionContext) ClosedFilter() IClosedFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosedFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosedFilterContext)
}

func (s *FilterExpressionContext) NotFilter() INotFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotFilterContext)
}

func (s *FilterExpressionContext) ValuePath() IValuePathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuePathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuePathContext)
}

func (s *FilterExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterFilterExpression(s)
	}
}

func (s *FilterExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitFilterExpression(s)
	}
}

func (p *SCIMFilterParser) FilterExpression() (localctx IFilterExpressionContext) {
	localctx = NewFilterExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SCIMFilterParserRULE_filterExpression)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.AttrExp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.ClosedFilter()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(47)
			p.NotFilter()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(48)
			p.ValuePath()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClosedFilterContext is an interface to support dynamic dispatch.
type IClosedFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_PARA() antlr.TerminalNode
	Filter() IFilterContext
	CLOSE_PARA() antlr.TerminalNode

	// IsClosedFilterContext differentiates from other interfaces.
	IsClosedFilterContext()
}

type ClosedFilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClosedFilterContext() *ClosedFilterContext {
	var p = new(ClosedFilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_closedFilter
	return p
}

func InitEmptyClosedFilterContext(p *ClosedFilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_closedFilter
}

func (*ClosedFilterContext) IsClosedFilterContext() {}

func NewClosedFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClosedFilterContext {
	var p = new(ClosedFilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_closedFilter

	return p
}

func (s *ClosedFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *ClosedFilterContext) OPEN_PARA() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserOPEN_PARA, 0)
}

func (s *ClosedFilterContext) Filter() IFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *ClosedFilterContext) CLOSE_PARA() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserCLOSE_PARA, 0)
}

func (s *ClosedFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosedFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClosedFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterClosedFilter(s)
	}
}

func (s *ClosedFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitClosedFilter(s)
	}
}

func (p *SCIMFilterParser) ClosedFilter() (localctx IClosedFilterContext) {
	localctx = NewClosedFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SCIMFilterParserRULE_closedFilter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(SCIMFilterParserOPEN_PARA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Filter()
	}
	{
		p.SetState(53)
		p.Match(SCIMFilterParserCLOSE_PARA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INotFilterContext is an interface to support dynamic dispatch.
type INotFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode
	ClosedFilter() IClosedFilterContext

	// IsNotFilterContext differentiates from other interfaces.
	IsNotFilterContext()
}

type NotFilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotFilterContext() *NotFilterContext {
	var p = new(NotFilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_notFilter
	return p
}

func InitEmptyNotFilterContext(p *NotFilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_notFilter
}

func (*NotFilterContext) IsNotFilterContext() {}

func NewNotFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotFilterContext {
	var p = new(NotFilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_notFilter

	return p
}

func (s *NotFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *NotFilterContext) NOT() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserNOT, 0)
}

func (s *NotFilterContext) ClosedFilter() IClosedFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosedFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosedFilterContext)
}

func (s *NotFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterNotFilter(s)
	}
}

func (s *NotFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitNotFilter(s)
	}
}

func (p *SCIMFilterParser) NotFilter() (localctx INotFilterContext) {
	localctx = NewNotFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SCIMFilterParserRULE_notFilter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(SCIMFilterParserNOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.ClosedFilter()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValuePathContext is an interface to support dynamic dispatch.
type IValuePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttrPath() IAttrPathContext
	OPEN_BRACKET() antlr.TerminalNode
	Filter() IFilterContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsValuePathContext differentiates from other interfaces.
	IsValuePathContext()
}

type ValuePathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuePathContext() *ValuePathContext {
	var p = new(ValuePathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_valuePath
	return p
}

func InitEmptyValuePathContext(p *ValuePathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_valuePath
}

func (*ValuePathContext) IsValuePathContext() {}

func NewValuePathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuePathContext {
	var p = new(ValuePathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_valuePath

	return p
}

func (s *ValuePathContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuePathContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *ValuePathContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserOPEN_BRACKET, 0)
}

func (s *ValuePathContext) Filter() IFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *ValuePathContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserCLOSE_BRACKET, 0)
}

func (s *ValuePathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuePathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterValuePath(s)
	}
}

func (s *ValuePathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitValuePath(s)
	}
}

func (p *SCIMFilterParser) ValuePath() (localctx IValuePathContext) {
	localctx = NewValuePathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SCIMFilterParserRULE_valuePath)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.AttrPath()
	}
	{
		p.SetState(59)
		p.Match(SCIMFilterParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Filter()
	}
	{
		p.SetState(61)
		p.Match(SCIMFilterParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttrExpContext is an interface to support dynamic dispatch.
type IAttrExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttrPath() IAttrPathContext
	AllSP() []antlr.TerminalNode
	SP(i int) antlr.TerminalNode
	ValuePresent() IValuePresentContext
	CompareOp() ICompareOpContext
	CompValue() ICompValueContext

	// IsAttrExpContext differentiates from other interfaces.
	IsAttrExpContext()
}

type AttrExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrExpContext() *AttrExpContext {
	var p = new(AttrExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_attrExp
	return p
}

func InitEmptyAttrExpContext(p *AttrExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_attrExp
}

func (*AttrExpContext) IsAttrExpContext() {}

func NewAttrExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrExpContext {
	var p = new(AttrExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_attrExp

	return p
}

func (s *AttrExpContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *AttrExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(SCIMFilterParserSP)
}

func (s *AttrExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserSP, i)
}

func (s *AttrExpContext) ValuePresent() IValuePresentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuePresentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuePresentContext)
}

func (s *AttrExpContext) CompareOp() ICompareOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareOpContext)
}

func (s *AttrExpContext) CompValue() ICompValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompValueContext)
}

func (s *AttrExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterAttrExp(s)
	}
}

func (s *AttrExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitAttrExp(s)
	}
}

func (p *SCIMFilterParser) AttrExp() (localctx IAttrExpContext) {
	localctx = NewAttrExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SCIMFilterParserRULE_attrExp)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.AttrPath()
		}
		{
			p.SetState(64)
			p.Match(SCIMFilterParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.ValuePresent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.AttrPath()
		}
		{
			p.SetState(68)
			p.Match(SCIMFilterParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.CompareOp()
		}
		{
			p.SetState(70)
			p.Match(SCIMFilterParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.CompValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilterOperatorContext is an interface to support dynamic dispatch.
type IFilterOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsFilterOperatorContext differentiates from other interfaces.
	IsFilterOperatorContext()
}

type FilterOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterOperatorContext() *FilterOperatorContext {
	var p = new(FilterOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_filterOperator
	return p
}

func InitEmptyFilterOperatorContext(p *FilterOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_filterOperator
}

func (*FilterOperatorContext) IsFilterOperatorContext() {}

func NewFilterOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterOperatorContext {
	var p = new(FilterOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_filterOperator

	return p
}

func (s *FilterOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserAND, 0)
}

func (s *FilterOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserOR, 0)
}

func (s *FilterOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterFilterOperator(s)
	}
}

func (s *FilterOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitFilterOperator(s)
	}
}

func (p *SCIMFilterParser) FilterOperator() (localctx IFilterOperatorContext) {
	localctx = NewFilterOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SCIMFilterParserRULE_filterOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SCIMFilterParserAND || _la == SCIMFilterParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValuePresentContext is an interface to support dynamic dispatch.
type IValuePresentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PR() antlr.TerminalNode

	// IsValuePresentContext differentiates from other interfaces.
	IsValuePresentContext()
}

type ValuePresentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuePresentContext() *ValuePresentContext {
	var p = new(ValuePresentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_valuePresent
	return p
}

func InitEmptyValuePresentContext(p *ValuePresentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_valuePresent
}

func (*ValuePresentContext) IsValuePresentContext() {}

func NewValuePresentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuePresentContext {
	var p = new(ValuePresentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_valuePresent

	return p
}

func (s *ValuePresentContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuePresentContext) PR() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserPR, 0)
}

func (s *ValuePresentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuePresentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuePresentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterValuePresent(s)
	}
}

func (s *ValuePresentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitValuePresent(s)
	}
}

func (p *SCIMFilterParser) ValuePresent() (localctx IValuePresentContext) {
	localctx = NewValuePresentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SCIMFilterParserRULE_valuePresent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(SCIMFilterParserPR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompValueContext is an interface to support dynamic dispatch.
type ICompValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FALSE() antlr.TerminalNode
	NULL() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	Number() INumberContext
	StringValue() IStringValueContext

	// IsCompValueContext differentiates from other interfaces.
	IsCompValueContext()
}

type CompValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompValueContext() *CompValueContext {
	var p = new(CompValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_compValue
	return p
}

func InitEmptyCompValueContext(p *CompValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_compValue
}

func (*CompValueContext) IsCompValueContext() {}

func NewCompValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompValueContext {
	var p = new(CompValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_compValue

	return p
}

func (s *CompValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CompValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserFALSE, 0)
}

func (s *CompValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserNULL, 0)
}

func (s *CompValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserTRUE, 0)
}

func (s *CompValueContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *CompValueContext) StringValue() IStringValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *CompValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterCompValue(s)
	}
}

func (s *CompValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitCompValue(s)
	}
}

func (p *SCIMFilterParser) CompValue() (localctx ICompValueContext) {
	localctx = NewCompValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SCIMFilterParserRULE_compValue)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SCIMFilterParserFALSE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Match(SCIMFilterParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SCIMFilterParserNULL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(SCIMFilterParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SCIMFilterParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)
			p.Match(SCIMFilterParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SCIMFilterParserDIGIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(82)
			p.Number()
		}

	case SCIMFilterParserQUOTE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(83)
			p.StringValue()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode
	AllDOUBLE_SLASH() []antlr.TerminalNode
	DOUBLE_SLASH(i int) antlr.TerminalNode

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_stringValue
	return p
}

func InitEmptyStringValueContext(p *StringValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_stringValue
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(SCIMFilterParserQUOTE)
}

func (s *StringValueContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserQUOTE, i)
}

func (s *StringValueContext) AllDOUBLE_SLASH() []antlr.TerminalNode {
	return s.GetTokens(SCIMFilterParserDOUBLE_SLASH)
}

func (s *StringValueContext) DOUBLE_SLASH(i int) antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserDOUBLE_SLASH, i)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *SCIMFilterParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SCIMFilterParserRULE_stringValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(SCIMFilterParserQUOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719214590) != 0 {
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(87)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == SCIMFilterParserQUOTE {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		case 2:
			{
				p.SetState(88)
				p.Match(SCIMFilterParserDOUBLE_SLASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(89)
				p.Match(SCIMFilterParserQUOTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(SCIMFilterParserQUOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(SCIMFilterParserDIGIT)
}

func (s *NumberContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserDIGIT, i)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *SCIMFilterParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SCIMFilterParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SCIMFilterParserDIGIT {
		{
			p.SetState(97)
			p.Match(SCIMFilterParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompareOpContext is an interface to support dynamic dispatch.
type ICompareOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	CO() antlr.TerminalNode
	SW() antlr.TerminalNode
	EW() antlr.TerminalNode
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode
	GE() antlr.TerminalNode
	LE() antlr.TerminalNode

	// IsCompareOpContext differentiates from other interfaces.
	IsCompareOpContext()
}

type CompareOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareOpContext() *CompareOpContext {
	var p = new(CompareOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_compareOp
	return p
}

func InitEmptyCompareOpContext(p *CompareOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_compareOp
}

func (*CompareOpContext) IsCompareOpContext() {}

func NewCompareOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareOpContext {
	var p = new(CompareOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_compareOp

	return p
}

func (s *CompareOpContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserEQ, 0)
}

func (s *CompareOpContext) NE() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserNE, 0)
}

func (s *CompareOpContext) CO() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserCO, 0)
}

func (s *CompareOpContext) SW() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserSW, 0)
}

func (s *CompareOpContext) EW() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserEW, 0)
}

func (s *CompareOpContext) GT() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserGT, 0)
}

func (s *CompareOpContext) LT() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserLT, 0)
}

func (s *CompareOpContext) GE() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserGE, 0)
}

func (s *CompareOpContext) LE() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserLE, 0)
}

func (s *CompareOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterCompareOp(s)
	}
}

func (s *CompareOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitCompareOp(s)
	}
}

func (p *SCIMFilterParser) CompareOp() (localctx ICompareOpContext) {
	localctx = NewCompareOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SCIMFilterParserRULE_compareOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&261632) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	URI() antlr.TerminalNode
	HYPHEN() antlr.TerminalNode
	ATTR_NAME() antlr.TerminalNode
	AllSubAttr() []ISubAttrContext
	SubAttr(i int) ISubAttrContext

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_attrPath
	return p
}

func InitEmptyAttrPathContext(p *AttrPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_attrPath
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) URI() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserURI, 0)
}

func (s *AttrPathContext) HYPHEN() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserHYPHEN, 0)
}

func (s *AttrPathContext) ATTR_NAME() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserATTR_NAME, 0)
}

func (s *AttrPathContext) AllSubAttr() []ISubAttrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubAttrContext); ok {
			len++
		}
	}

	tst := make([]ISubAttrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubAttrContext); ok {
			tst[i] = t.(ISubAttrContext)
			i++
		}
	}

	return tst
}

func (s *AttrPathContext) SubAttr(i int) ISubAttrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubAttrContext); ok {
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

	return t.(ISubAttrContext)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterAttrPath(s)
	}
}

func (s *AttrPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitAttrPath(s)
	}
}

func (p *SCIMFilterParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SCIMFilterParserRULE_attrPath)
	var _la int

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(SCIMFilterParserURI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(SCIMFilterParserHYPHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(SCIMFilterParserATTR_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SCIMFilterParserDOT {
			{
				p.SetState(107)
				p.SubAttr()
			}

			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.Match(SCIMFilterParserATTR_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SCIMFilterParserDOT {
			{
				p.SetState(114)
				p.SubAttr()
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.Match(SCIMFilterParserATTR_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubAttrContext is an interface to support dynamic dispatch.
type ISubAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	ATTR_NAME() antlr.TerminalNode

	// IsSubAttrContext differentiates from other interfaces.
	IsSubAttrContext()
}

type SubAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrContext() *SubAttrContext {
	var p = new(SubAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_subAttr
	return p
}

func InitEmptySubAttrContext(p *SubAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SCIMFilterParserRULE_subAttr
}

func (*SubAttrContext) IsSubAttrContext() {}

func NewSubAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrContext {
	var p = new(SubAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SCIMFilterParserRULE_subAttr

	return p
}

func (s *SubAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrContext) DOT() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserDOT, 0)
}

func (s *SubAttrContext) ATTR_NAME() antlr.TerminalNode {
	return s.GetToken(SCIMFilterParserATTR_NAME, 0)
}

func (s *SubAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.EnterSubAttr(s)
	}
}

func (s *SubAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SCIMFilterListener); ok {
		listenerT.ExitSubAttr(s)
	}
}

func (p *SCIMFilterParser) SubAttr() (localctx ISubAttrContext) {
	localctx = NewSubAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SCIMFilterParserRULE_subAttr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(SCIMFilterParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(SCIMFilterParserATTR_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
