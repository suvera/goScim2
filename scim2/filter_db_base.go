package scim2

import "github.com/spf13/cast"

const bindPrefix = "svr_"

type DbFilterConverter interface {
	Convert(filter string, columnMappings map[string]string)
	GetClause() DbFilterClause
	GetMappedColumn(expression *AttributeExpression) string
}
type BaseDbFilterConverter struct {
	Clause          DbFilterClause
	ColumnMappings  map[string]string
	BindCounter     int
	CaseInsensitive bool
}

func (b *BaseDbFilterConverter) GetMappedColumn(expression *AttributeExpression) string {
	if expression == nil || b.ColumnMappings == nil {
		return ""
	}

	column := b.ColumnMappings[expression.Attribute]
	if column == "" {
		column = b.ColumnMappings[expression.GetFqdn()]
	}

	return column
}

func (b *BaseDbFilterConverter) GetClause() DbFilterClause {
	return b.Clause
}

func (b *BaseDbFilterConverter) Convert(filter string, columnMappings map[string]string) {
	b.ColumnMappings = columnMappings
	panic("implement BaseDbFilterConverter.Convert method")
}

func (b *BaseDbFilterConverter) GetBindVariable() string {
	str := bindPrefix + cast.ToString(b.BindCounter)
	b.BindCounter++
	return str
}

func (b *BaseDbFilterConverter) IsCaseInsensitive() bool {
	return b.CaseInsensitive
}

func (b *BaseDbFilterConverter) SetCaseInsensitive(caseInsensitive bool) {
	b.CaseInsensitive = caseInsensitive
}

func (b *BaseDbFilterConverter) SetBindCounter(counter int) {
	b.BindCounter = counter
}
