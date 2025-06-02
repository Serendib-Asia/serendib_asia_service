package handler

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/app/services"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/web"
	"github.com/chazool/serendib_asia_service/pkg/web/responsebuilder"

	"github.com/gofiber/fiber/v2"
)

const (
	HandleGetPurposeTypesMethod   = "HandleGetPurposeTypes"
	HandleGetPropertyTypesMethod  = "HandleGetPropertyTypes"
	HandleGetFurnitureTypesMethod = "HandleGetFurnitureTypes"
	HandleGetConditionsMethod     = "HandleGetConditions"
	HandleGetUtilitiesMethod      = "HandleGetUtilities"
	HandleGetAmenitiesMethod      = "HandleGetAmenities"
)

// HandleGetPurposeTypes handles retrieving all purpose types
// @Summary Get all purpose types
// @Description Retrieves all purpose types
// @Tags lookups
// @Accept json
// @Produce json
// @Success 200 {object} []dto.LookupResponse
// @Failure 500 {object} custom.ErrorResult
// @Router /api/lookups/purpose-types [get]
func HandleGetPurposeTypes(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleGetPurposeTypesMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleGetPurposeTypesMethod), commonLogFields...)

	var (
		statusCode    int
		errRes        custom.ErrorResult
		response      []dto.LookupResponse
		lookupService = services.CreateLookupService(requestID, nil)
	)

	response, errorResult := lookupService.GetPurposeTypes()
	if errorResult != nil {
		logFields := log.TraceCustomError(commonLogFields, *errorResult)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleGetPurposeTypesMethod), logFields...)
		statusCode, errRes = HandleError(errorResult)
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		Response:      response,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

// HandleGetPropertyTypes handles retrieving all property types
// @Summary Get all property types
// @Description Retrieves all property types
// @Tags lookups
// @Accept json
// @Produce json
// @Success 200 {object} []dto.LookupResponse
// @Failure 500 {object} custom.ErrorResult
// @Router /api/lookups/property-types [get]
func HandleGetPropertyTypes(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleGetPropertyTypesMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleGetPropertyTypesMethod), commonLogFields...)

	var (
		statusCode    int
		errRes        custom.ErrorResult
		response      []dto.LookupResponse
		lookupService = services.CreateLookupService(requestID, nil)
	)

	response, errorResult := lookupService.GetPropertyTypes()
	if errorResult != nil {
		logFields := log.TraceCustomError(commonLogFields, *errorResult)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleGetPropertyTypesMethod), logFields...)
		statusCode, errRes = HandleError(errorResult)
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		Response:      response,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

// HandleGetFurnitureTypes handles retrieving all furniture types
// @Summary Get all furniture types
// @Description Retrieves all furniture types
// @Tags lookups
// @Accept json
// @Produce json
// @Success 200 {object} []dto.LookupResponse
// @Failure 500 {object} custom.ErrorResult
// @Router /api/lookups/furniture-types [get]
func HandleGetFurnitureTypes(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleGetFurnitureTypesMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleGetFurnitureTypesMethod), commonLogFields...)

	var (
		statusCode    int
		errRes        custom.ErrorResult
		response      []dto.LookupResponse
		lookupService = services.CreateLookupService(requestID, nil)
	)

	response, errorResult := lookupService.GetFurnitureTypes()
	if errorResult != nil {
		logFields := log.TraceCustomError(commonLogFields, *errorResult)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleGetFurnitureTypesMethod), logFields...)
		statusCode, errRes = HandleError(errorResult)
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		Response:      response,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

// HandleGetConditions handles retrieving all property conditions
// @Summary Get all property conditions
// @Description Retrieves all property conditions
// @Tags lookups
// @Accept json
// @Produce json
// @Success 200 {object} []dto.LookupResponse
// @Failure 500 {object} custom.ErrorResult
// @Router /api/lookups/conditions [get]
func HandleGetConditions(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleGetConditionsMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleGetConditionsMethod), commonLogFields...)

	var (
		statusCode    int
		errRes        custom.ErrorResult
		response      []dto.LookupResponse
		lookupService = services.CreateLookupService(requestID, nil)
	)

	response, errorResult := lookupService.GetConditions()
	if errorResult != nil {
		logFields := log.TraceCustomError(commonLogFields, *errorResult)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleGetConditionsMethod), logFields...)
		statusCode, errRes = HandleError(errorResult)
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		Response:      response,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

// HandleGetUtilities handles retrieving all utilities
// @Summary Get all utilities
// @Description Retrieves all utilities
// @Tags lookups
// @Accept json
// @Produce json
// @Success 200 {object} []dto.LookupResponse
// @Failure 500 {object} custom.ErrorResult
// @Router /api/lookups/utilities [get]
func HandleGetUtilities(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleGetUtilitiesMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleGetUtilitiesMethod), commonLogFields...)

	var (
		statusCode    int
		errRes        custom.ErrorResult
		response      []dto.LookupResponse
		lookupService = services.CreateLookupService(requestID, nil)
	)

	response, errorResult := lookupService.GetUtilities()
	if errorResult != nil {
		logFields := log.TraceCustomError(commonLogFields, *errorResult)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleGetUtilitiesMethod), logFields...)
		statusCode, errRes = HandleError(errorResult)
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		Response:      response,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

// HandleGetAmenities handles retrieving all amenities
// @Summary Get all amenities
// @Description Retrieves all amenities
// @Tags lookups
// @Accept json
// @Produce json
// @Success 200 {object} []dto.LookupResponse
// @Failure 500 {object} custom.ErrorResult
// @Router /api/lookups/amenities [get]
func HandleGetAmenities(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleGetAmenitiesMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleGetAmenitiesMethod), commonLogFields...)

	var (
		statusCode    int
		errRes        custom.ErrorResult
		response      []dto.LookupResponse
		lookupService = services.CreateLookupService(requestID, nil)
	)

	response, errorResult := lookupService.GetAmenities()
	if errorResult != nil {
		logFields := log.TraceCustomError(commonLogFields, *errorResult)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleGetAmenitiesMethod), logFields...)
		statusCode, errRes = HandleError(errorResult)
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		Response:      response,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}
