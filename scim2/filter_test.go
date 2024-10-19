package scim2

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

type FakeScimFilterListener struct {
	Results []any
}

func NewFakeScimFilterListener() *FakeScimFilterListener {
	return &FakeScimFilterListener{
		Results: make([]any, 0),
	}
}

func (f *FakeScimFilterListener) OnAttributeExpression(expression *AttributeExpression) {
	fmt.Printf("AttributeExpression: %v\n", expression)
	f.Results = append(f.Results, expression)
}

func (f *FakeScimFilterListener) OnOpenParenthesis(negated bool) {
	fmt.Printf("OpenParenthesis: %v\n", negated)
	f.Results = append(f.Results, "("+cast.ToString(negated))
}

func (f *FakeScimFilterListener) OnCloseParenthesis(negated bool) {
	fmt.Printf("OnCloseParenthesis: %v\n", negated)
	f.Results = append(f.Results, negated)
}

func (f *FakeScimFilterListener) OnLogicalOperator(operation FilterOperation) {
	fmt.Printf("OnLogicalOperator: %v\n", operation)
	f.Results = append(f.Results, operation)
}

func TestParseFilter(t *testing.T) {
	input := "emails[type eq \"work\" and value co \"@example.com\"] or  ims[type eq \"xmpp\" and value co \"@foo.com\"]"

	listener := NewFakeScimFilterListener()
	listener.Results = make([]any, 0)
	ParseFilter(input, listener)

	if listener.Results == nil {
		t.Errorf("listener.Results is nil")
	}
}
