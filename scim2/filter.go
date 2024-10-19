package scim2

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/suvera/goScim2/scim2/filter"
)

func ParseFilter(input string, listener ScimFilterListener) *ScimError {
	var err *ScimError

	defer func() {
		if r := recover(); r != nil {
			log.Error("Recovered from ", r)
			rErr := NewScimError(400, fmt.Sprintf("Filter syntax error %v", r), "invalidFilter")
			err = &rErr
		}
	}()

	is := antlr.NewInputStream(input)
	errListener := NewErrorListener()
	listenerInternal := NewFilterListenerInternal(listener)
	//listenerInternal.debug = true

	lexer := filter.NewSCIMFilterLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)

	// Create the parser
	parser := filter.NewSCIMFilterParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errListener)

	antlr.ParseTreeWalkerDefault.Walk(listenerInternal, parser.Expression())

	if errListener.HasErrors() {
		err = &errListener.Errors[0]
	}

	return err
}
