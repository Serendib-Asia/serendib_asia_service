package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/snabb/isoweek"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type customString string
type customInt int32

func NewString(s string) *customString {
	str := customString(s)
	return &str
}

func NewInt(i int32) *customInt {
	val := customInt(i)
	return &val
}

func (s *customString) ToUINT(commonLogFields []zap.Field) (*uint, *custom.ErrorResult) {
	str := string(*s)

	ID, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(fmt.Sprintf(constant.ErrorOccurredWhen, constant.ErrStringToUintParseMsg), logFields...)
		errResult := custom.BuildBadReqErrResult(constant.ErrStringToUintParseCode, constant.ErrStringToUintParseMsg, err.Error())
		return nil, &errResult
	}

	id := uint(ID)

	return &id, nil
}

// DateParser is a function that parses a date string to a time.Time object.
func DateParser(commonLogFields []zapcore.Field, datePattern, dateString, fieldName string) (time.Time, *custom.ErrorResult) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.DateParserMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.DateParserMethod), commonLogFields...)

	// Parse date
	parseDate, err := time.Parse(datePattern, dateString)
	if err != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, err))
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(constant.DateParserMethod), logFields...)

		errRes := custom.BuildBadReqErrResult(constant.ErrDateParseCode, fmt.Sprintf(constant.ErrOccouredWhenParseDate, fieldName), constant.Empty)
		return parseDate, &errRes
	}

	return parseDate, nil
}

// DateCompare is a function that compares two dates.
func DateCompare(commonLogFields []zapcore.Field, startDate, endDate time.Time, startDateField, endDateField string) *custom.ErrorResult {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.DateCompareMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.DateCompareMethod), commonLogFields...)

	if endDate.Before(startDate) {
		errMsg := fmt.Sprintf(constant.ErrorOccurredWhenCompareGreaterThan, endDateField, startDateField)
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errMsg))
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(constant.DateCompareMethod), logFields...)

		errRes := custom.BuildBadReqErrResult(constant.ErrDateCompareCode, errMsg, constant.Empty)
		return &errRes
	}

	return nil
}

// JSONUnmarshal is a function that unmarshal a byte array to a struct.
func JSONUnmarshal[T any](commonLogFields []zapcore.Field, data []byte) (T, *custom.ErrorResult) {
	log.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncStart, constant.DataUnmarshalMethod), commonLogFields...)
	defer log.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncEnd, constant.DataUnmarshalMethod), commonLogFields...)

	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Logger.Error(constant.UnexpectedWhenUnmarshalError, append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		errResult := custom.BuildBadReqErrResult(constant.ErrDataUnmarshalCode, constant.UnexpectedWhenUnmarshalError, err.Error())
		errRes := &errResult
		return v, errRes
	}

	return v, nil
}

func StructCastor[T any](commonLogFields []zapcore.Field, d any) (*T, *custom.ErrorResult) {
	data, err := json.Marshal(d)
	if err != nil {
		log.Logger.Error(constant.UnexpectedWhenMarshalError, append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		errResult := custom.BuildBadReqErrResult(constant.ErrDataMarshalCode, constant.UnexpectedWhenMarshalError, err.Error())
		return nil, &errResult
	}

	r, errRes := JSONUnmarshal[T](commonLogFields, data)
	if errRes != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errRes))
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(constant.DataUnmarshalMethod), logFields...)
		return nil, errRes
	}

	return &r, nil
}

func AppendToArrayIfNotEmpty[T any](commonLogFields []zapcore.Field, array *[]T, value T) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.AppendToArrayIfNotEmpty), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.AppendToArrayIfNotEmpty), commonLogFields...)

	rawValue := reflect.ValueOf(value).Interface()
	if rawValue != "" && rawValue != nil {
		*array = append(*array, value)
	}
}

func ValidateWeekAndGetWeekStartDate(commonLogFields []zapcore.Field, year, week int) (string, *custom.ErrorResult) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.ValidateAndGetWeekStartDateMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.ValidateAndGetWeekStartDateMethod), commonLogFields...)

	if !isoweek.Validate(year, week) {
		log.Logger.Error(log.TraceMsgErrorOccurredWhen(constant.ErrOccurredWhenValidateWeek), commonLogFields...)
		errRes := custom.BuildBadReqErrResult(constant.ErrWeekValidateCode, constant.ErrOccurredWhenValidateWeek, constant.Empty)
		return constant.Empty, &errRes
	}

	returnedYear, month, date := isoweek.StartDate(year, week)
	weekStartDate := fmt.Sprintf(constant.ThreePlaceholderWithHyphen, returnedYear, month, date)
	return weekStartDate, nil
}

func (s *customInt) ToUINT(commonLogFields []zap.Field) (*uint, *custom.ErrorResult) {
	intValue := int(*s)

	if intValue < 0 {
		err := fmt.Errorf("cannot convert negative value %d to uint", intValue)
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(fmt.Sprintf(constant.ErrorOccurredWhen, constant.ErrIntToUintParseMsg), logFields...)
		errResult := custom.BuildBadReqErrResult(constant.ErrIntToUintParseCode, constant.ErrIntToUintParseMsg, err.Error())
		return nil, &errResult
	}

	uintValue := uint(intValue)

	return &uintValue, nil
}

func (s *customInt) ToString(commonLogFields []zap.Field) (*string, *custom.ErrorResult) {
	intValue := int(*s)

	if intValue < 0 {
		err := fmt.Errorf("cannot convert negative value %d to string", intValue)
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(fmt.Sprintf(constant.ErrorOccurredWhen, constant.ErrIntToStringParseMsg), logFields...)
		errResult := custom.BuildBadReqErrResult(constant.ErrIntToStringParseCode, constant.ErrIntToStringParseMsg, err.Error())
		return nil, &errResult
	}

	// Convert to string
	strValue := fmt.Sprintf("%d", intValue)
	return &strValue, nil
}

func StringToInt(commonLogFields []zap.Field, value string) (*int, *custom.ErrorResult) {
	// Attempt to parse the string to an integer
	intValue, err := strconv.Atoi(value)

	if err != nil {
		err := fmt.Errorf("cannot convert negative value %d to string", intValue)
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(fmt.Sprintf(constant.ErrorOccurredWhen, constant.ErrIntToStringParseMsg), logFields...)
		errResult := custom.BuildBadReqErrResult(constant.ErrIntToStringParseCode, constant.ErrIntToStringParseMsg, err.Error())
		return nil, &errResult
	}

	return &intValue, nil
}

// IsStrValueTrue checks if a value is equivalent to "true"
func IsStrValueTrue(commonLogFields []zap.Field, value string) bool {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.IsStrValueTrueMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.IsStrValueTrueMethod), commonLogFields...)

	return strings.EqualFold(value, constant.LowerTrueStr)
}
