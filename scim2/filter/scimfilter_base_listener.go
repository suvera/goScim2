// Code generated from SCIMFilter.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filter // SCIMFilter
import "github.com/antlr4-go/antlr/v4"

// BaseSCIMFilterListener is a complete listener for a parse tree produced by SCIMFilterParser.
type BaseSCIMFilterListener struct{}

var _ SCIMFilterListener = &BaseSCIMFilterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSCIMFilterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSCIMFilterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSCIMFilterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSCIMFilterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSCIMFilterListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSCIMFilterListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseSCIMFilterListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseSCIMFilterListener) ExitFilter(ctx *FilterContext) {}

// EnterMoreFilters is called when production moreFilters is entered.
func (s *BaseSCIMFilterListener) EnterMoreFilters(ctx *MoreFiltersContext) {}

// ExitMoreFilters is called when production moreFilters is exited.
func (s *BaseSCIMFilterListener) ExitMoreFilters(ctx *MoreFiltersContext) {}

// EnterFilterExpression is called when production filterExpression is entered.
func (s *BaseSCIMFilterListener) EnterFilterExpression(ctx *FilterExpressionContext) {}

// ExitFilterExpression is called when production filterExpression is exited.
func (s *BaseSCIMFilterListener) ExitFilterExpression(ctx *FilterExpressionContext) {}

// EnterClosedFilter is called when production closedFilter is entered.
func (s *BaseSCIMFilterListener) EnterClosedFilter(ctx *ClosedFilterContext) {}

// ExitClosedFilter is called when production closedFilter is exited.
func (s *BaseSCIMFilterListener) ExitClosedFilter(ctx *ClosedFilterContext) {}

// EnterNotFilter is called when production notFilter is entered.
func (s *BaseSCIMFilterListener) EnterNotFilter(ctx *NotFilterContext) {}

// ExitNotFilter is called when production notFilter is exited.
func (s *BaseSCIMFilterListener) ExitNotFilter(ctx *NotFilterContext) {}

// EnterValuePath is called when production valuePath is entered.
func (s *BaseSCIMFilterListener) EnterValuePath(ctx *ValuePathContext) {}

// ExitValuePath is called when production valuePath is exited.
func (s *BaseSCIMFilterListener) ExitValuePath(ctx *ValuePathContext) {}

// EnterAttrExp is called when production attrExp is entered.
func (s *BaseSCIMFilterListener) EnterAttrExp(ctx *AttrExpContext) {}

// ExitAttrExp is called when production attrExp is exited.
func (s *BaseSCIMFilterListener) ExitAttrExp(ctx *AttrExpContext) {}

// EnterFilterOperator is called when production filterOperator is entered.
func (s *BaseSCIMFilterListener) EnterFilterOperator(ctx *FilterOperatorContext) {}

// ExitFilterOperator is called when production filterOperator is exited.
func (s *BaseSCIMFilterListener) ExitFilterOperator(ctx *FilterOperatorContext) {}

// EnterValuePresent is called when production valuePresent is entered.
func (s *BaseSCIMFilterListener) EnterValuePresent(ctx *ValuePresentContext) {}

// ExitValuePresent is called when production valuePresent is exited.
func (s *BaseSCIMFilterListener) ExitValuePresent(ctx *ValuePresentContext) {}

// EnterCompValue is called when production compValue is entered.
func (s *BaseSCIMFilterListener) EnterCompValue(ctx *CompValueContext) {}

// ExitCompValue is called when production compValue is exited.
func (s *BaseSCIMFilterListener) ExitCompValue(ctx *CompValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseSCIMFilterListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseSCIMFilterListener) ExitStringValue(ctx *StringValueContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseSCIMFilterListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSCIMFilterListener) ExitNumber(ctx *NumberContext) {}

// EnterCompareOp is called when production compareOp is entered.
func (s *BaseSCIMFilterListener) EnterCompareOp(ctx *CompareOpContext) {}

// ExitCompareOp is called when production compareOp is exited.
func (s *BaseSCIMFilterListener) ExitCompareOp(ctx *CompareOpContext) {}

// EnterAttrPath is called when production attrPath is entered.
func (s *BaseSCIMFilterListener) EnterAttrPath(ctx *AttrPathContext) {}

// ExitAttrPath is called when production attrPath is exited.
func (s *BaseSCIMFilterListener) ExitAttrPath(ctx *AttrPathContext) {}

// EnterSubAttr is called when production subAttr is entered.
func (s *BaseSCIMFilterListener) EnterSubAttr(ctx *SubAttrContext) {}

// ExitSubAttr is called when production subAttr is exited.
func (s *BaseSCIMFilterListener) ExitSubAttr(ctx *SubAttrContext) {}
