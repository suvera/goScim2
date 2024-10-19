package scim2

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/cast"
	"github.com/suvera/goScim2/scim2/filter"
	"strings"
	"time"
)

type FilterListenerInternal struct {
	filter.BaseSCIMFilterListener
	listener          ScimFilterListener
	attributes        []string
	values            []any
	operators         []FilterCondition
	notFilters        []bool
	parentAttrNames   []string
	hasParentAttrName []bool
	eof               bool
	debug             bool
}

type AttributeExpression struct {
	Schema    string          `json:"schema"`
	Attribute string          `json:"attribute"`
	Operator  FilterCondition `json:"operator"`
	Value     any             `json:"value"`
	ValueType ValueType       `json:"valueType"`
}

func NewAttributeExpression() *AttributeExpression {
	return &AttributeExpression{}
}

func (s *AttributeExpression) GetFqdn() string {
	if s.Schema == "" {
		return s.Attribute
	}
	return s.Schema + ":" + s.Attribute
}

func (s *AttributeExpression) SetAttribute(attribute string) {
	s.Attribute = attribute
	if strings.Contains(attribute, ":") {
		pos := strings.LastIndex(attribute, ":")
		s.Schema = attribute[:pos]
		s.Attribute = attribute[pos+1:]
	}
}

func (s *AttributeExpression) String() string {
	return fmt.Sprintf("AttributeExpression(schema=%s, name=%s, operator=%s, value=%v, valueType=%s)",
		s.Schema, s.Attribute, s.Operator, s.Value, s.ValueType)
}

// ScimFilterListener
//   - Listener for expressions
type ScimFilterListener interface {
	OnAttributeExpression(expression *AttributeExpression)
	OnOpenParenthesis(negated bool)
	OnCloseParenthesis(negated bool)
	OnLogicalOperator(operation FilterOperation)
}

// NewFilterListenerInternal
//   - Internal listener for ScimFilter parser
func NewFilterListenerInternal(listener ScimFilterListener) *FilterListenerInternal {
	return &FilterListenerInternal{
		listener: listener,
	}
}

func (s *FilterListenerInternal) EnterExpression(ctx *filter.ExpressionContext) {
	if s.debug {
		fmt.Println("enterExpression: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitExpression(ctx *filter.ExpressionContext) {
	if s.debug {
		fmt.Println("exitExpression: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) EnterFilter(ctx *filter.FilterContext) {
	if s.debug {
		fmt.Println("enterFilter: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitFilter(ctx *filter.FilterContext) {
	if s.debug {
		fmt.Println("exitFilter: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) EnterFilterExpression(ctx *filter.FilterExpressionContext) {
	if s.debug {
		fmt.Println("enterFilterExpression: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitFilterExpression(ctx *filter.FilterExpressionContext) {
	if s.debug {
		fmt.Println("exitFilterExpression: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) EnterFilterOperator(ctx *filter.FilterOperatorContext) {
	if s.debug {
		fmt.Println("enterFilterOperator: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitFilterOperator(ctx *filter.FilterOperatorContext) {
	if s.debug {
		fmt.Println("exitFilterOperator: " + ctx.GetText())
	}
	s.listener.OnLogicalOperator(FilterOperationFromString(ctx.GetText()))
}

func (s *FilterListenerInternal) EnterValuePresent(ctx *filter.ValuePresentContext) {
	if s.debug {
		fmt.Println("enterValuePresent: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitValuePresent(ctx *filter.ValuePresentContext) {
	if s.debug {
		fmt.Println("exitValuePresent: " + ctx.GetText())
	}
	s.operators = append(s.operators, FilterConditionPresent)
	s.values = append(s.values, NullObjectInstance)
}

func (s *FilterListenerInternal) EnterNotFilter(ctx *filter.NotFilterContext) {
	if s.debug {
		fmt.Println("enterNotFilter: " + ctx.GetText())
	}
	s.notFilters = append(s.notFilters, true)
	s.listener.OnOpenParenthesis(true)
}

func (s *FilterListenerInternal) ExitNotFilter(ctx *filter.NotFilterContext) {
	if s.debug {
		fmt.Println("exitNotFilter: " + ctx.GetText())
	}
	negated := false
	if len(s.notFilters) > 0 {
		negated = s.notFilters[len(s.notFilters)-1]
		s.notFilters = s.notFilters[:len(s.notFilters)-1]
	}
	s.listener.OnCloseParenthesis(negated)
}

func (s *FilterListenerInternal) EnterAttrExp(ctx *filter.AttrExpContext) {
	if s.debug {
		fmt.Println("enterAttrExp: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitAttrExp(ctx *filter.AttrExpContext) {
	if s.debug {
		fmt.Println("exitAttrExp: " + ctx.GetText())
	}

	a := NewAttributeExpression()
	if len(s.operators) == 0 {
		fmt.Println("No operator found when ending the attribute expression")
		return
	}
	attrName := s.attributes[len(s.attributes)-1]
	s.attributes = s.attributes[:len(s.attributes)-1]
	a.SetAttribute(attrName)

	if len(s.operators) > 0 {
		a.Operator = s.operators[len(s.operators)-1]
		s.operators = s.operators[:len(s.operators)-1]
	}

	if len(s.values) == 0 {
		s.listener.OnAttributeExpression(a)
		return
	}
	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	a.ValueType = ValueTypeString
	if _, ok := val.(NullObject); ok {
		a.Value = nil
		a.ValueType = ValueTypeNull
	} else {
		a.Value = val
		switch val.(type) {
		case bool:
			a.ValueType = ValueTypeBoolean
		case int:
			a.ValueType = ValueTypeInteger
		case float64, float32:
			a.ValueType = ValueTypeDecimal
		case time.Time, *time.Time:
			a.ValueType = ValueTypeDateTime
		}
	}
	s.listener.OnAttributeExpression(a)
}

func (s *FilterListenerInternal) EnterAttrPath(ctx *filter.AttrPathContext) {
	if s.debug {
		fmt.Println("enterAttrPath: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitAttrPath(ctx *filter.AttrPathContext) {
	if s.debug {
		fmt.Println("exitAttrPath: " + ctx.GetText())
	}
	if len(s.hasParentAttrName) > 0 {
		s.hasParentAttrName = nil
		s.parentAttrNames = append(s.parentAttrNames, ctx.GetText())
		return
	}

	name := ctx.GetText()
	if len(s.parentAttrNames) > 0 {
		for i := len(s.parentAttrNames) - 1; i >= 0; i-- {
			name = s.parentAttrNames[i] + "." + name
		}
	}
	s.attributes = append(s.attributes, name)
}

func (s *FilterListenerInternal) EnterCompareOp(ctx *filter.CompareOpContext) {
	if s.debug {
		fmt.Println("enterCompareOp: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitCompareOp(ctx *filter.CompareOpContext) {
	if s.debug {
		fmt.Println("exitCompareOp: " + ctx.GetText())
	}
	s.operators = append(s.operators, FilterConditionFromString(ctx.GetText()))
}

func (s *FilterListenerInternal) EnterCompValue(ctx *filter.CompValueContext) {
	if s.debug {
		fmt.Println("enterCompValue: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitCompValue(ctx *filter.CompValueContext) {
	if s.debug {
		fmt.Println("exitCompValue: " + ctx.GetText())
	}
	val := ctx.GetText()
	if len(val) > 1 && val[0] == '"' && val[len(val)-1] == '"' {
		val = val[1 : len(val)-1]
		if IsDate(val) {
			date, err := ParseDate(val)
			if err != nil {
				if s.debug {
					fmt.Println("Error parsing date: " + val)
				}
				s.values = append(s.values, val)
			} else {
				s.values = append(s.values, date)
			}
		} else {
			s.values = append(s.values, val)
		}
	} else if val == "true" || val == "false" {
		s.values = append(s.values, val == "true")
	} else if val == "null" {
		s.values = append(s.values, NullObjectInstance)
	} else if IsInteger(val) {
		s.values = append(s.values, cast.ToInt64(val))
	} else if IsFloat(val) {
		s.values = append(s.values, cast.ToFloat64(val))
	} else {
		s.values = append(s.values, val)
	}
}

func (s *FilterListenerInternal) EnterValuePath(ctx *filter.ValuePathContext) {
	if s.debug {
		fmt.Println("enterValuePath: " + ctx.GetText())
	}
	s.hasParentAttrName = append(s.hasParentAttrName, true)
	s.listener.OnOpenParenthesis(false)
}

func (s *FilterListenerInternal) ExitValuePath(ctx *filter.ValuePathContext) {
	if s.debug {
		fmt.Println("exitValuePath: " + ctx.GetText())
	}
	s.parentAttrNames = s.parentAttrNames[:len(s.parentAttrNames)-1]
	s.listener.OnCloseParenthesis(false)
}

func (s *FilterListenerInternal) EnterClosedFilter(ctx *filter.ClosedFilterContext) {
	if s.debug {
		fmt.Println("enterClosedFilter: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitClosedFilter(ctx *filter.ClosedFilterContext) {
	if s.debug {
		fmt.Println("exitClosedFilter: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) EnterStringValue(ctx *filter.StringValueContext) {
	if s.debug {
		fmt.Println("enterStringValue: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitStringValue(ctx *filter.StringValueContext) {
	if s.debug {
		fmt.Println("exitStringValue: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) EnterSubAttr(ctx *filter.SubAttrContext) {
	if s.debug {
		fmt.Println("enterSubAttr: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitSubAttr(ctx *filter.SubAttrContext) {
	if s.debug {
		fmt.Println("exitSubAttr: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) EnterNumber(ctx *filter.NumberContext) {
	if s.debug {
		fmt.Println("enterNumber: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) ExitNumber(ctx *filter.NumberContext) {
	if s.debug {
		fmt.Println("exitNumber: " + ctx.GetText())
	}
}

func (s *FilterListenerInternal) VisitErrorNode(node antlr.ErrorNode) {
	if s.debug {
		fmt.Println("visitErrorNode: " + node.GetText())
	}
}

func (s *FilterListenerInternal) VisitTerminal(node antlr.TerminalNode) {
	if s.debug {
		fmt.Println("visitTerminal: " + node.GetText())
	}
	if node.GetText() == "<EOF>" {
		// stop parsing
		s.eof = true
	}
}
