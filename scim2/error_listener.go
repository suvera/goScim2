package scim2

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/cast"
)

type ErrorListener struct {
	Errors []ScimError `json:"errors"`
}

func NewErrorListener() *ErrorListener {
	l := new(ErrorListener)
	l.Errors = make([]ScimError, 0)
	return l
}

func (this *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	this.Errors = append(this.Errors, NewScimError(400,
		"Filter syntax error at line "+cast.ToString(line)+" char "+cast.ToString(column)+": "+msg+", caused by "+e.GetMessage(),
		"invalidFilter"))
}

func (this *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	this.Errors = append(this.Errors, NewScimError(400,
		"Ambiguity error at line "+cast.ToString(startIndex)+" char "+cast.ToString(stopIndex),
		"invalidFilter"))
}

func (this *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	this.Errors = append(this.Errors, NewScimError(400,
		"Attempting full context error at line "+cast.ToString(startIndex)+" char "+cast.ToString(stopIndex),
		"invalidFilter"))
}

func (this *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	this.Errors = append(this.Errors, NewScimError(400,
		"Context sensitivity error at line "+cast.ToString(startIndex)+" char "+cast.ToString(stopIndex),
		"invalidFilter"))
}

func (this *ErrorListener) HasErrors() bool {
	return len(this.Errors) > 0
}

func (this *ErrorListener) GetError() *ScimError {
	if this.HasErrors() {
		return &this.Errors[0]
	}
	return nil
}
