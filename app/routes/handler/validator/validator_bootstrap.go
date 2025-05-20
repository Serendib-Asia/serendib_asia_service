package validator

import (
	"fmt"

	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils"

	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GenericBaseValidator is a generic base function that validates a request.
func GenericBaseValidator[T any](requestID string, ctx *fiber.Ctx) (T, *custom.ErrorResult) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	log.Logger.Debug(log.TraceMsgFuncStart(GenericBaseValidatorMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(GenericBaseValidatorMethod), commonLogFields...)

	var (
		request T
		body    = string(ctx.Body())
		err     error
	)

	err = ctx.BodyParser(&request)
	if err != nil {
		log.Logger.Error(constant.InvalidInputAndPassErr, append(commonLogFields, []zap.Field{zap.String(constant.ErrorRequestBody, body), zap.Error(err)}...)...)
		errorResult := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.InvalidRequestErrorMessage, err.Error())
		return request, &errorResult
	}

	err = validate.Struct(&request)
	if err != nil {
		log.Logger.Error(constant.InvalidInputAndPassErr, append(commonLogFields, []zap.Field{zap.String(constant.ErrorRequestBody, body), zap.Error(err)}...)...)
		return request, BuildValidationErrorResponse(requestID, err)
	}

	return request, nil
}

// GenericValidator is a generic function that validates a request with custome validators.
func GenericValidator[T any](requestID string, ctx *fiber.Ctx, validators ...func([]zap.Field, any) *custom.ErrorResult) (T, *custom.ErrorResult) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	log.Logger.Debug(log.TraceMsgFuncStart(GenericValidatorMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(GenericValidatorMethod), commonLogFields...)

	var (
		errResult = &custom.ErrorResult{}
	)

	request, errRes := GenericBaseValidator[T](requestID, ctx)
	if errRes != nil {
		logFields := log.TraceCustomError(commonLogFields, *errRes)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(GenericBaseValidatorMethod), logFields...)
		errResult.ErrorList = append(errResult.ErrorList, errRes.ErrorList...)
	}

	if len(errResult.ErrorList) == 0 {
		for _, validator := range validators {
			errRes = validator([]zap.Field{}, request)
			if errRes != nil {
				logFields := log.TraceCustomError(commonLogFields, *errRes)
				log.Logger.Error(log.TraceMsgErrorOccurredFrom(customValidators), logFields...)
				errResult.ErrorList = append(errResult.ErrorList, errRes.ErrorList...)
			}
		}
	}

	if len(errResult.ErrorList) > 0 {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errResult))
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(GenericValidatorMethod), logFields...)
		errRes := custom.BuildBadReqErrResultWithList(errResult.ErrorList...)
		return request, &errRes
	}

	return request, nil
}

// ValidateDateRange is a function that validates the date range.
func ValidateDateRange(commonLogFields []zapcore.Field, startDate, endDate string) *custom.ErrorResult {
	log.Logger.Debug(log.TraceMsgFuncStart(ValidateDataRangeMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidateDataRangeMethod), commonLogFields...)

	if startDate == constant.Empty || endDate == constant.Empty {
		log.Logger.Error(constant.ErrEmptyDateMsg, commonLogFields...)
		errRes := custom.BuildBadReqErrResult(constant.ErrEmptyDateCode, constant.ErrEmptyDateMsg, constant.Empty)
		return &errRes
	}

	// Parse start date
	start, errRes := utils.DateParser(commonLogFields, DatePattern, startDate, StartDate)
	if errRes != nil {
		logFields := log.TraceCustomError(commonLogFields, *errRes)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(ValidateDataRangeMethod), logFields...)
		return errRes
	}

	// Parse end date
	end, errRes := utils.DateParser(commonLogFields, DatePattern, endDate, EndDate)
	if errRes != nil {
		logFields := log.TraceCustomError(commonLogFields, *errRes)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(ValidateDataRangeMethod), logFields...)
		return errRes
	}

	// Compare dates
	errRes = utils.DateCompare(commonLogFields, start, end, StartDate, EndDate)
	if errRes != nil {
		logFields := log.TraceCustomError(commonLogFields, *errRes)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(ValidateDataRangeMethod), logFields...)
		return errRes
	}

	return nil
}

// ValidateWeekRange is a function that validates the week range.
func ValidateWeekRange(commonLogFields []zapcore.Field, startedWeek, endWeek int, isSameYear bool) *custom.ErrorResult {
	log.Logger.Debug(log.TraceMsgFuncStart(ValidateWeekRangeMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidateWeekRangeMethod), commonLogFields...)

	weekComp := func(start, max int, from string) *custom.ErrorResult {
		if start > max {
			errMsg := fmt.Sprintf(constant.ErrorOccurredWhenWeekValidation, from)
			log.Logger.Error(fmt.Sprintf(invalidWeekRangeMessage, start, max), append(commonLogFields, zap.Any(constant.ErrorNote, errMsg))...)
			errRes := custom.BuildBadReqErrResult(constant.ErrWeekValidateCode, errMsg, constant.Empty)
			return &errRes
		}
		return nil
	}

	if isSameYear && startedWeek > endWeek {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(ValidateWeekRangeMethod), append(commonLogFields, zap.Any(constant.ErrorNote, sameYearInvalidWeekRange))...)
		errRes := custom.BuildBadReqErrResult(constant.ErrWeekValidateCode, sameYearInvalidWeekRange, constant.Empty)
		return &errRes
	}

	if err := weekComp(startedWeek, maximumWeek, StartWeekField); err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(WeekComparison), log.TraceCustomError(commonLogFields, *err)...)
		return err
	}

	if err := weekComp(endWeek, maximumWeek, EndWeekField); err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(WeekComparison), log.TraceCustomError(commonLogFields, *err)...)
		return err
	}

	return nil
}

// ValidateYearRange is a function that validates the year range.
func ValidateYearRange(commonLogFields []zapcore.Field, startedYear, endYear int, startYearField, endYearField string) *custom.ErrorResult {
	log.Logger.Debug(log.TraceMsgFuncStart(ValidateYearRangeMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidateYearRangeMethod), commonLogFields...)

	if startedYear > endYear {
		errMsg := fmt.Sprintf(constant.ErrOccurredWhenCompGreaterThanOrEql, endYearField, startYearField)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(ValidateYearRangeMethod), append(commonLogFields, zap.Any(constant.ErrorNote, errMsg))...)
		errRes := custom.BuildBadReqErrResult(constant.ErrYearValidationCode, errMsg, constant.Empty)
		return &errRes
	}
	return nil
}

func ValidateSentimentCategories(commonLogFields []zapcore.Field, categories dto.SentimentScoreCategories) *custom.ErrorResult {
	log.Logger.Debug(log.TraceMsgFuncStart(ValidateSentimentCategoriesMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidateSentimentCategoriesMethod), commonLogFields...)

	if !(categories.Overall || categories.SLA || categories.CustomerExperience || categories.Efficiency || categories.ChurnRisk) {
		errRes := custom.BuildBadReqErrResult(constant.ErrEmptyCategoriesCode, constant.ErrEmptyCategoriesMsg, constant.Empty)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(ValidateSentimentCategoriesMethod), log.TraceCustomError(commonLogFields, errRes)...)
		return &errRes
	}

	return nil
}

func ValidateBacklogCategories(commonLogFields []zapcore.Field, categories dto.BacklogCategories) *custom.ErrorResult {
	log.Logger.Debug(log.TraceMsgFuncStart(ValidateBacklogCategoriesMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidateBacklogCategoriesMethod), commonLogFields...)

	if !(categories.MostViewed || categories.LastOutbound || categories.UniqueResponders) {
		errRes := custom.BuildBadReqErrResult(constant.ErrEmptyCategoriesCode, constant.ErrEmptyCategoriesMsg, constant.Empty)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(ValidateBacklogCategoriesMethod), log.TraceCustomError(commonLogFields, errRes)...)
		return &errRes
	}

	return nil
}
