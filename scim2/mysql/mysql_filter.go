package mysql

import (
	"github.com/spf13/cast"
	"github.com/suvera/goScim2/scim2"
	"log"
	"strings"
)

type FilterConverter struct {
	scim2.BaseDbFilterConverter
}

func NewMysqlFilterConverter() *FilterConverter {
	return &FilterConverter{}
}

func (f *FilterConverter) Convert(filter string, columnMappings map[string]string) {
	f.ColumnMappings = columnMappings
	f.Clause = scim2.DbFilterClause{
		WhereClause: "",
		Binds:       make(map[string]any),
	}

	scim2.ParseFilter(filter, f)
}

func (f *FilterConverter) OnAttributeExpression(expression *scim2.AttributeExpression) {
	column := f.GetMappedColumn(expression)
	if column == "" {
		log.Printf("Filter column '%s' is not supported in this context\n", expression.Attribute)
		return
	}

	bindVar := f.GetBindVariable()
	value := expression.Value
	valueType := expression.ValueType
	op := expression.Operator

	switch op {
	case scim2.FilterConditionEqual, scim2.FilterConditionNotEqual:
		compOp := " = :" + bindVar
		if op == scim2.FilterConditionNotEqual {
			compOp = " != :" + bindVar
		}
		if valueType == scim2.ValueTypeString && f.CaseInsensitive {
			column = "LOWER(" + column + ")"
			value = strings.ToLower(cast.ToString(value))
		}
		f.Clause.WhereClause += column + compOp
		f.Clause.Binds[bindVar] = value

	case scim2.FilterConditionContains, scim2.FilterConditionStartsWith, scim2.FilterConditionEndsWith:
		if valueType != scim2.ValueTypeString {
			log.Printf("Unsupported value type for operator %s", op)
			return
		}
		if f.CaseInsensitive {
			column = "LOWER(" + column + ")"
			value = strings.ToLower(cast.ToString(value))
		}
		f.Clause.WhereClause += column + " LIKE :" + bindVar
		strVal := cast.ToString(value)
		if op == scim2.FilterConditionStartsWith {
			strVal = strVal + "%"
		} else if op == scim2.FilterConditionEndsWith {
			strVal = "%" + strVal
		} else {
			strVal = "%" + strVal + "%"
		}
		f.Clause.Binds[bindVar] = strVal

	case scim2.FilterConditionLesserThan, scim2.FilterConditionGreaterThan, scim2.FilterConditionLesserThanEquals, scim2.FilterConditionGreaterThanEquals:
		if valueType != scim2.ValueTypeInteger && valueType != scim2.ValueTypeDecimal && valueType != scim2.ValueTypeDateTime {
			log.Printf("Unsupported value type for operator %s", op)
			return
		}
		dbOp := "<"
		if op == scim2.FilterConditionGreaterThan {
			dbOp = ">"
		} else if op == scim2.FilterConditionLesserThanEquals {
			dbOp = "<="
		} else if op == scim2.FilterConditionGreaterThanEquals {
			dbOp = ">="
		}

		f.Clause.WhereClause += column + " " + dbOp + " :" + bindVar
		f.Clause.Binds[bindVar] = value

	case scim2.FilterConditionPresent:
		f.Clause.WhereClause += column + " IS NOT NULL"
		if valueType == scim2.ValueTypeString {
			f.Clause.WhereClause += " AND " + column + " != ''"
		} else if valueType == scim2.ValueTypeInteger || valueType == scim2.ValueTypeDecimal {
			f.Clause.WhereClause += " AND " + column + " != 0"
		}

	default:
		log.Printf("Unsupported filter operator %s", op)
	}
}

func (f *FilterConverter) OnOpenParenthesis(negated bool) {
	if negated {
		f.Clause.WhereClause += " NOT ( "
	} else {
		f.Clause.WhereClause += " ( "
	}
}

func (f *FilterConverter) OnCloseParenthesis(negated bool) {
	f.Clause.WhereClause += " ) "
}

func (f *FilterConverter) OnLogicalOperator(operation scim2.FilterOperation) {
	f.Clause.WhereClause += " " + string(operation) + " "
}
