package repository

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/chazool/serendib_asia_service/pkg/log"

	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sqlBuilder struct {
	SQLTemplate      string
	Request          dto.CommonFilterRequest
	RepeatWhereCount int
	Limit            int
	TableName        string
}

// GenerateBetweenQuery generates a SQL query for selecting records between two values.
// It takes a slice of zapcore.Field for logging, table name, column name, start and end values as inputs.
// The function returns a string which is the SQL query.
func GenerateBetweenQuery(logFields []zapcore.Field, table, column, start, end string) string {
	log.Logger.Debug(log.TraceMsgFuncStart(GenerateBetweenQueryMethod), logFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(GenerateBetweenQueryMethod), logFields...)

	return fmt.Sprintf(queryBetween, table, column, start, end)
}

// GenerateInQuery generates a SQL query for selecting records where a column's value is in a list of values.
// It takes a slice of zapcore.Field for logging, table name, column name, and a slice of values as inputs.
// The function returns a string which is the SQL query.
func GenerateInQuery[T any](logFields []zapcore.Field, table, column string, values []T) string {
	log.Logger.Debug(log.TraceMsgFuncStart(GenerateInQueryMethod), logFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(GenerateInQueryMethod), logFields...)

	// Skip IN clause generation for empty or nil values
	if len(values) == 0 {
		log.Logger.Debug(fmt.Sprintf("%s.%s filter empty - using 1=1", table, column))
		return "1=1"
	}

	var query strings.Builder
	query.WriteString(fmt.Sprintf(inQueryPrefix, table, column))
	for i, value := range values {
		if i != 0 {
			query.WriteString(CommaWithSpace)
		}

		// Directly handle integers and other types appropriately
		switch v := reflect.ValueOf(value); v.Kind() {
		case reflect.String:
			query.WriteString(fmt.Sprintf(StringPlaceholder, v.String()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			query.WriteString(fmt.Sprintf(ValuePlaceholder, v.Int()))
		}
	}

	query.WriteString(CloseBracket)
	log.Logger.Debug(GeneratedQuery, append(logFields, zap.String(GeneratedQuery, query.String()))...)

	return query.String()
}

// GenerateEqualsQuery generates a SQL query for selecting records where a column equals to a given value.
// It takes a slice of zapcore.Field for logging, table name, column name, and the value as inputs.
// The function returns a string which is the SQL query.
func GenerateEqualsQuery[T any](logFields []zapcore.Field, table, column string, value T) string {
	log.Logger.Debug(log.TraceMsgFuncStart(GenerateInQueryMethod), logFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(GenerateInQueryMethod), logFields...)

	if reflect.ValueOf(value).Interface() == constant.Empty {
		return constant.Empty
	}

	var query strings.Builder

	query.WriteString(fmt.Sprintf(equalsQueryPrefix, table, column))
	query.WriteString(strings.ToUpper(fmt.Sprintf(ValuePlaceholderWithSpaces, reflect.ValueOf(value))))

	log.Logger.Debug(GeneratedQuery, append(logFields, zap.String(GeneratedQuery, query.String()))...)

	return query.String()
}

// CombineQueries combines multiple queries using the provided operator.
// It takes a slice of zapcore.Field for logging, operator, and a list of queries as inputs.
// The function returns a string which is the combined query.
func CombineQueries(logFields []zapcore.Field, op string, queries ...string) string {
	log.Logger.Debug(log.TraceMsgFuncStart(CombineQueriesMethod), logFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(CombineQueriesMethod), logFields...)

	return strings.Join(queries, fmt.Sprintf(StringPlaceholderWithSpaces, op))
}

func buildSQLWhereClause(commonLogFields []zap.Field, request dto.CommonFilterRequest, tableName string) string {
	log.Logger.Debug(log.TraceMsgFuncStart(buildSQLWhereClauseMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(buildSQLWhereClauseMethod), commonLogFields...)
	var whereList []string

	if len(request.StartDate) > constant.Zero && len(request.EndDate) > constant.Zero {
		whereList = append(whereList, GenerateBetweenQuery(commonLogFields, tableName, startDateColumn, request.StartDate, request.EndDate))
	}
	if len(request.Client) > constant.Zero {
		whereList = append(whereList, GenerateInQuery(commonLogFields, tableName, clientIDColumn, request.Client))
	}
	if len(request.Priority) > constant.Zero {
		whereList = append(whereList, GenerateInQuery(commonLogFields, tableName, priorityIDColumn, request.Priority))
	}
	if len(request.Category) > constant.Zero {
		whereList = append(whereList, GenerateInQuery(commonLogFields, tableName, categoryIDColumn, request.Category))
	}
	if len(request.Service) > constant.Zero {
		whereList = append(whereList, GenerateInQuery(commonLogFields, tableName, serviceIDColumn, request.Service))
	}

	whereClause := CombineQueries(commonLogFields, andOperator, whereList...)
	log.Logger.Debug(GeneratedWhereClause, zap.String(GeneratedWhereClause, whereClause))

	return whereClause
}

func buildSQLQuery(commonLogFields []zap.Field, template string, replaceValues ...any) string {
	log.Logger.Debug(log.TraceMsgFuncStart(BuildSQLQueryMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(BuildSQLQueryMethod), commonLogFields...)

	finalQuery := fmt.Sprintf(template, replaceValues...)
	log.Logger.Debug(constant.BuildedQuery, append(commonLogFields, zap.String(constant.SQLQuery, finalQuery))...)

	return finalQuery
}

func getSQLQuery(commonLogFields []zap.Field, builder sqlBuilder) string {
	log.Logger.Debug(log.TraceMsgFuncStart(getSQLQueryMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(getSQLQueryMethod), commonLogFields...)

	whereClause := buildSQLWhereClause(commonLogFields, builder.Request, builder.TableName)

	var arr []any
	for i := constant.Zero; i < builder.RepeatWhereCount; i++ {
		arr = append(arr, whereClause)
	}

	if builder.Limit > constant.Zero {
		arr = append(arr, builder.Limit)
	}

	finalQuery := buildSQLQuery(commonLogFields, builder.SQLTemplate, arr...)
	return finalQuery
}

func generateDateIntervalQuery(logFields []zapcore.Field, table, column string, days int) string {
	log.Logger.Debug(log.TraceMsgFuncStart(generateDateIntervalQueryMethod), logFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(generateDateIntervalQueryMethod), logFields...)

	var query strings.Builder
	query.WriteString(fmt.Sprintf(lessThanQueryPrefix+" CURRENT_DATE - INTERVAL '%d days'", table, column, days))
	log.Logger.Debug(GeneratedQuery, append(logFields, zap.String(GeneratedQuery, query.String()))...)

	return query.String()
}

func generateOrderByClause(logFields []zapcore.Field, column, order string) string {
	log.Logger.Debug(log.TraceMsgFuncStart(generateOrderByClauseMethod), logFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(generateOrderByClauseMethod), logFields...)

	var query strings.Builder
	query.WriteString(fmt.Sprintf(orderByQueryPrefix, column, order))
	log.Logger.Debug(GeneratedQuery, append(logFields, zap.String(GeneratedQuery, query.String()))...)

	return query.String()
}
