// Code generated from SCIMFilter.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filter

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SCIMFilterLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SCIMFilterLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func scimfilterlexerLexerInit() {
	staticData := &SCIMFilterLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"NOT", "OPEN_PARA", "CLOSE_PARA", "OPEN_BRACKET", "CLOSE_BRACKET", "PR",
		"AND", "OR", "EQ", "NE", "CO", "SW", "EW", "GT", "LT", "GE", "LE", "QUOTE",
		"DOUBLE_QUOTE", "HYPHEN", "UNDERSCORE", "DOT", "FALSE", "NULL", "TRUE",
		"NEW_RETURN", "NEW_LINE", "ALPHA", "DIGIT", "SP", "URI", "ATTR_NAME",
		"DOUBLE_SLASH", "QUOTE_SLASH", "ANY",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 192, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 4, 29, 157, 8, 29, 11, 29, 12, 29,
		158, 1, 30, 4, 30, 162, 8, 30, 11, 30, 12, 30, 163, 1, 30, 1, 30, 4, 30,
		168, 8, 30, 11, 30, 12, 30, 169, 4, 30, 172, 8, 30, 11, 30, 12, 30, 173,
		1, 31, 4, 31, 177, 8, 31, 11, 31, 12, 31, 178, 1, 31, 5, 31, 182, 8, 31,
		10, 31, 12, 31, 185, 9, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 0,
		0, 35, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 1, 0, 6, 2, 0,
		65, 90, 97, 122, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32, 1, 0, 58, 58, 5, 0,
		45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 197, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0,
		55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0,
		0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0,
		0, 1, 71, 1, 0, 0, 0, 3, 75, 1, 0, 0, 0, 5, 77, 1, 0, 0, 0, 7, 79, 1, 0,
		0, 0, 9, 81, 1, 0, 0, 0, 11, 83, 1, 0, 0, 0, 13, 86, 1, 0, 0, 0, 15, 90,
		1, 0, 0, 0, 17, 93, 1, 0, 0, 0, 19, 96, 1, 0, 0, 0, 21, 99, 1, 0, 0, 0,
		23, 102, 1, 0, 0, 0, 25, 105, 1, 0, 0, 0, 27, 108, 1, 0, 0, 0, 29, 111,
		1, 0, 0, 0, 31, 114, 1, 0, 0, 0, 33, 117, 1, 0, 0, 0, 35, 120, 1, 0, 0,
		0, 37, 122, 1, 0, 0, 0, 39, 125, 1, 0, 0, 0, 41, 127, 1, 0, 0, 0, 43, 129,
		1, 0, 0, 0, 45, 131, 1, 0, 0, 0, 47, 137, 1, 0, 0, 0, 49, 142, 1, 0, 0,
		0, 51, 147, 1, 0, 0, 0, 53, 149, 1, 0, 0, 0, 55, 151, 1, 0, 0, 0, 57, 153,
		1, 0, 0, 0, 59, 156, 1, 0, 0, 0, 61, 161, 1, 0, 0, 0, 63, 176, 1, 0, 0,
		0, 65, 186, 1, 0, 0, 0, 67, 188, 1, 0, 0, 0, 69, 190, 1, 0, 0, 0, 71, 72,
		5, 110, 0, 0, 72, 73, 5, 111, 0, 0, 73, 74, 5, 116, 0, 0, 74, 2, 1, 0,
		0, 0, 75, 76, 5, 40, 0, 0, 76, 4, 1, 0, 0, 0, 77, 78, 5, 41, 0, 0, 78,
		6, 1, 0, 0, 0, 79, 80, 5, 91, 0, 0, 80, 8, 1, 0, 0, 0, 81, 82, 5, 93, 0,
		0, 82, 10, 1, 0, 0, 0, 83, 84, 5, 112, 0, 0, 84, 85, 5, 114, 0, 0, 85,
		12, 1, 0, 0, 0, 86, 87, 5, 97, 0, 0, 87, 88, 5, 110, 0, 0, 88, 89, 5, 100,
		0, 0, 89, 14, 1, 0, 0, 0, 90, 91, 5, 111, 0, 0, 91, 92, 5, 114, 0, 0, 92,
		16, 1, 0, 0, 0, 93, 94, 5, 101, 0, 0, 94, 95, 5, 113, 0, 0, 95, 18, 1,
		0, 0, 0, 96, 97, 5, 110, 0, 0, 97, 98, 5, 101, 0, 0, 98, 20, 1, 0, 0, 0,
		99, 100, 5, 99, 0, 0, 100, 101, 5, 111, 0, 0, 101, 22, 1, 0, 0, 0, 102,
		103, 5, 115, 0, 0, 103, 104, 5, 119, 0, 0, 104, 24, 1, 0, 0, 0, 105, 106,
		5, 101, 0, 0, 106, 107, 5, 119, 0, 0, 107, 26, 1, 0, 0, 0, 108, 109, 5,
		103, 0, 0, 109, 110, 5, 116, 0, 0, 110, 28, 1, 0, 0, 0, 111, 112, 5, 108,
		0, 0, 112, 113, 5, 116, 0, 0, 113, 30, 1, 0, 0, 0, 114, 115, 5, 103, 0,
		0, 115, 116, 5, 101, 0, 0, 116, 32, 1, 0, 0, 0, 117, 118, 5, 108, 0, 0,
		118, 119, 5, 101, 0, 0, 119, 34, 1, 0, 0, 0, 120, 121, 5, 34, 0, 0, 121,
		36, 1, 0, 0, 0, 122, 123, 5, 34, 0, 0, 123, 124, 5, 34, 0, 0, 124, 38,
		1, 0, 0, 0, 125, 126, 5, 45, 0, 0, 126, 40, 1, 0, 0, 0, 127, 128, 5, 95,
		0, 0, 128, 42, 1, 0, 0, 0, 129, 130, 5, 46, 0, 0, 130, 44, 1, 0, 0, 0,
		131, 132, 5, 102, 0, 0, 132, 133, 5, 97, 0, 0, 133, 134, 5, 108, 0, 0,
		134, 135, 5, 115, 0, 0, 135, 136, 5, 101, 0, 0, 136, 46, 1, 0, 0, 0, 137,
		138, 5, 110, 0, 0, 138, 139, 5, 117, 0, 0, 139, 140, 5, 108, 0, 0, 140,
		141, 5, 108, 0, 0, 141, 48, 1, 0, 0, 0, 142, 143, 5, 116, 0, 0, 143, 144,
		5, 114, 0, 0, 144, 145, 5, 117, 0, 0, 145, 146, 5, 101, 0, 0, 146, 50,
		1, 0, 0, 0, 147, 148, 5, 13, 0, 0, 148, 52, 1, 0, 0, 0, 149, 150, 5, 10,
		0, 0, 150, 54, 1, 0, 0, 0, 151, 152, 7, 0, 0, 0, 152, 56, 1, 0, 0, 0, 153,
		154, 7, 1, 0, 0, 154, 58, 1, 0, 0, 0, 155, 157, 7, 2, 0, 0, 156, 155, 1,
		0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0,
		0, 159, 60, 1, 0, 0, 0, 160, 162, 7, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162,
		163, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 171,
		1, 0, 0, 0, 165, 167, 7, 3, 0, 0, 166, 168, 7, 4, 0, 0, 167, 166, 1, 0,
		0, 0, 168, 169, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0,
		170, 172, 1, 0, 0, 0, 171, 165, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173,
		171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 62, 1, 0, 0, 0, 175, 177, 7,
		0, 0, 0, 176, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 176, 1, 0, 0,
		0, 178, 179, 1, 0, 0, 0, 179, 183, 1, 0, 0, 0, 180, 182, 7, 4, 0, 0, 181,
		180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 64, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 187, 5, 92,
		0, 0, 187, 66, 1, 0, 0, 0, 188, 189, 7, 5, 0, 0, 189, 68, 1, 0, 0, 0, 190,
		191, 9, 0, 0, 0, 191, 70, 1, 0, 0, 0, 7, 0, 158, 163, 169, 173, 178, 183,
		0,
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

// SCIMFilterLexerInit initializes any static state used to implement SCIMFilterLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSCIMFilterLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SCIMFilterLexerInit() {
	staticData := &SCIMFilterLexerLexerStaticData
	staticData.once.Do(scimfilterlexerLexerInit)
}

// NewSCIMFilterLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSCIMFilterLexer(input antlr.CharStream) *SCIMFilterLexer {
	SCIMFilterLexerInit()
	l := new(SCIMFilterLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SCIMFilterLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SCIMFilter.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SCIMFilterLexer tokens.
const (
	SCIMFilterLexerNOT           = 1
	SCIMFilterLexerOPEN_PARA     = 2
	SCIMFilterLexerCLOSE_PARA    = 3
	SCIMFilterLexerOPEN_BRACKET  = 4
	SCIMFilterLexerCLOSE_BRACKET = 5
	SCIMFilterLexerPR            = 6
	SCIMFilterLexerAND           = 7
	SCIMFilterLexerOR            = 8
	SCIMFilterLexerEQ            = 9
	SCIMFilterLexerNE            = 10
	SCIMFilterLexerCO            = 11
	SCIMFilterLexerSW            = 12
	SCIMFilterLexerEW            = 13
	SCIMFilterLexerGT            = 14
	SCIMFilterLexerLT            = 15
	SCIMFilterLexerGE            = 16
	SCIMFilterLexerLE            = 17
	SCIMFilterLexerQUOTE         = 18
	SCIMFilterLexerDOUBLE_QUOTE  = 19
	SCIMFilterLexerHYPHEN        = 20
	SCIMFilterLexerUNDERSCORE    = 21
	SCIMFilterLexerDOT           = 22
	SCIMFilterLexerFALSE         = 23
	SCIMFilterLexerNULL          = 24
	SCIMFilterLexerTRUE          = 25
	SCIMFilterLexerNEW_RETURN    = 26
	SCIMFilterLexerNEW_LINE      = 27
	SCIMFilterLexerALPHA         = 28
	SCIMFilterLexerDIGIT         = 29
	SCIMFilterLexerSP            = 30
	SCIMFilterLexerURI           = 31
	SCIMFilterLexerATTR_NAME     = 32
	SCIMFilterLexerDOUBLE_SLASH  = 33
	SCIMFilterLexerQUOTE_SLASH   = 34
	SCIMFilterLexerANY           = 35
)
