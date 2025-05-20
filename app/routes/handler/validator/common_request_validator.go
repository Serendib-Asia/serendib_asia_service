package validator

import (
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils"

	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/custom"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// ValidateCommonRequest used to validate icx common req
func ValidateCommonRequest(requestID string, ctx *fiber.Ctx) (dto.CommonFilterRequest, *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ValidateCommonRequestMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidateCommonRequestMethod), commonLogFields...)

	request, errRes := GenericValidator[dto.CommonFilterRequest](requestID, ctx, ValidateStartAndEndDate)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(GenericValidatorMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return request, errRes
	}

	return request, nil
}

// ValidateStartAndEndDate used to validate start and end date
func ValidateStartAndEndDate(commonLogFields []zap.Field, request any) *custom.ErrorResult {
	r, errRes := utils.StructCastor[dto.CommonFilterRequest](commonLogFields, request)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(StructCasterMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return errRes
	}

	errRes = ValidateDateRange(commonLogFields, r.StartDate, r.EndDate)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(validateStartAndEndDateMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return errRes
	}

	return nil
}

// ValidateStartAndEndWeek used to validate start and end date
func ValidateStartAndEndWeek(commonLogFields []zap.Field, request any) *custom.ErrorResult {
	r, errRes := utils.StructCastor[dto.AgentInfoRequest](commonLogFields, request)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(StructCasterMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return errRes
	}

	errRes = ValidateWeekRange(commonLogFields, r.StartWeek, r.EndWeek, r.StartYear == r.EndYear)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(ValidateWeekRangeMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return errRes
	}

	return nil
}

// ValidateStartAndEndYear used to validate start and end date
func ValidateStartAndEndYear(commonLogFields []zap.Field, request any) *custom.ErrorResult {
	r, errRes := utils.StructCastor[dto.AgentInfoRequest](commonLogFields, request)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(StructCasterMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return errRes
	}

	errRes = ValidateYearRange(commonLogFields, r.StartYear, r.EndYear, StartYearField, EndYearField)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(ValidateYearRangeMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return errRes
	}

	return nil
}

// ValidatePaginatedCommonFilterRequest used to validate paginated common filter request
func ValidatePaginatedCommonFilterRequest(requestID string, ctx *fiber.Ctx) (dto.PaginatedCommonFilterRequest, *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ValidatePaginatedCommonFilterRequestMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidatePaginatedCommonFilterRequestMethod), commonLogFields...)

	request, errRes := GenericValidator[dto.PaginatedCommonFilterRequest](requestID, ctx, ValidateStartAndEndDate)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(GenericValidatorMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return request, errRes
	}

	return request, nil
}

func ValidateDateRangeRequest(requestID string, ctx *fiber.Ctx) (dto.DateRangeRequest, *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ValidateDateRangeRequestMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidateDateRangeRequestMethod), commonLogFields...)

	request, errRes := GenericValidator[dto.DateRangeRequest](requestID, ctx, ValidateStartAndEndDate)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(GenericValidatorMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return request, errRes
	}

	return request, nil
}

func ValidatePaginatedDateRangeRequest(requestID string, ctx *fiber.Ctx) (dto.PaginatedDateRangeRequest, *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ValidatePaginatedDateRangeRequestMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidatePaginatedDateRangeRequestMethod), commonLogFields...)

	request, errRes := GenericValidator[dto.PaginatedDateRangeRequest](requestID, ctx, ValidateStartAndEndDate)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(GenericValidatorMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return request, errRes
	}

	return request, nil
}

func ValidateDateTrendRequest(requestID string, ctx *fiber.Ctx) (dto.DateTrendRequest, *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ValidateDateTrendRequestMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidateDateTrendRequestMethod), commonLogFields...)

	request, errRes := GenericValidator[dto.DateTrendRequest](requestID, ctx, ValidateStartAndEndDate)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(GenericValidatorMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return request, errRes
	}

	return request, nil
}
