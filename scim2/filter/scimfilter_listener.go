// Code generated from SCIMFilter.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filter // SCIMFilter
import "github.com/antlr4-go/antlr/v4"

// SCIMFilterListener is a complete listener for a parse tree produced by SCIMFilterParser.
type SCIMFilterListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterMoreFilters is called when entering the moreFilters production.
	EnterMoreFilters(c *MoreFiltersContext)

	// EnterFilterExpression is called when entering the filterExpression production.
	EnterFilterExpression(c *FilterExpressionContext)

	// EnterClosedFilter is called when entering the closedFilter production.
	EnterClosedFilter(c *ClosedFilterContext)

	// EnterNotFilter is called when entering the notFilter production.
	EnterNotFilter(c *NotFilterContext)

	// EnterValuePath is called when entering the valuePath production.
	EnterValuePath(c *ValuePathContext)

	// EnterAttrExp is called when entering the attrExp production.
	EnterAttrExp(c *AttrExpContext)

	// EnterFilterOperator is called when entering the filterOperator production.
	EnterFilterOperator(c *FilterOperatorContext)

	// EnterValuePresent is called when entering the valuePresent production.
	EnterValuePresent(c *ValuePresentContext)

	// EnterCompValue is called when entering the compValue production.
	EnterCompValue(c *CompValueContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterCompareOp is called when entering the compareOp production.
	EnterCompareOp(c *CompareOpContext)

	// EnterAttrPath is called when entering the attrPath production.
	EnterAttrPath(c *AttrPathContext)

	// EnterSubAttr is called when entering the subAttr production.
	EnterSubAttr(c *SubAttrContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitMoreFilters is called when exiting the moreFilters production.
	ExitMoreFilters(c *MoreFiltersContext)

	// ExitFilterExpression is called when exiting the filterExpression production.
	ExitFilterExpression(c *FilterExpressionContext)

	// ExitClosedFilter is called when exiting the closedFilter production.
	ExitClosedFilter(c *ClosedFilterContext)

	// ExitNotFilter is called when exiting the notFilter production.
	ExitNotFilter(c *NotFilterContext)

	// ExitValuePath is called when exiting the valuePath production.
	ExitValuePath(c *ValuePathContext)

	// ExitAttrExp is called when exiting the attrExp production.
	ExitAttrExp(c *AttrExpContext)

	// ExitFilterOperator is called when exiting the filterOperator production.
	ExitFilterOperator(c *FilterOperatorContext)

	// ExitValuePresent is called when exiting the valuePresent production.
	ExitValuePresent(c *ValuePresentContext)

	// ExitCompValue is called when exiting the compValue production.
	ExitCompValue(c *CompValueContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitCompareOp is called when exiting the compareOp production.
	ExitCompareOp(c *CompareOpContext)

	// ExitAttrPath is called when exiting the attrPath production.
	ExitAttrPath(c *AttrPathContext)

	// ExitSubAttr is called when exiting the subAttr production.
	ExitSubAttr(c *SubAttrContext)
}
