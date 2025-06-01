package handler

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/app/services"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"
	"github.com/chazool/serendib_asia_service/pkg/web"
	"github.com/chazool/serendib_asia_service/pkg/web/responsebuilder"

	"github.com/gofiber/fiber/v2"
)

const (
	// Property handler methods
	HandleCreatePropertyMethod       = "HandleCreateProperty"
	HandleGetPropertyMethod          = "HandleGetProperty"
	HandleUpdatePropertyMethod       = "HandleUpdateProperty"
	HandleDeletePropertyMethod       = "HandleDeleteProperty"
	HandleListPropertiesMethod       = "HandleListProperties"
	HandleListPropertiesByUserMethod = "HandleListPropertiesByUser"
)

// HandleCreateProperty handles the creation of a new property
// @Summary Create a new property
// @Description Creates a new property with the provided details
// @Tags properties
// @Accept json
// @Produce json
// @Param property body dto.PropertyRequest true "Property details"
// @Success 200 {object} dto.Property
// @Failure 400 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties [post]
func HandleCreateProperty(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleCreatePropertyMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleCreatePropertyMethod), commonLogFields...)

	var (
		statusCode      int
		errorResult     *custom.ErrorResult
		errRes          custom.ErrorResult
		request         dto.PropertyRequest
		response        *dto.Property
		propertyService = services.CreatePropertyService(requestID, nil)
	)

	if err := ctx.BodyParser(&request); err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleCreatePropertyMethod), logFields...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.InvalidRequestErrorMessage, err.Error())
		errorResult = &errRes
		statusCode, errRes = HandleError(errorResult)
	} else {
		response, errorResult = propertyService.Create(request)
		if errorResult != nil {
			logFields := log.TraceCustomError(commonLogFields, *errorResult)
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.PropertyServiceCreateMethod), logFields...)
			statusCode, errRes = HandleError(errorResult)
		}
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

// HandleGetProperty handles retrieving a property by ID
// @Summary Get a property by ID
// @Description Retrieves a property's details by its ID
// @Tags properties
// @Accept json
// @Produce json
// @Param id path int true "Property ID"
// @Success 200 {object} dto.PropertyResponse
// @Failure 400 {object} custom.ErrorResult
// @Failure 404 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties/{id} [get]
func HandleGetProperty(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleGetPropertyMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleGetPropertyMethod), commonLogFields...)

	var (
		statusCode      int = fiber.StatusOK
		errorResult     *custom.ErrorResult
		errRes          custom.ErrorResult
		response        dto.Property
		propertyService = services.CreatePropertyService(requestID, nil)
	)

	propertyID, err := GetIDFromParams(ctx)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleGetPropertyMethod), commonLogFields...)
		errorResult = err
		statusCode, errRes = HandleError(errorResult)
	} else {
		response, errorResult = propertyService.GetByID(propertyID)
		if errorResult != nil {
			logFields := log.TraceCustomError(commonLogFields, *errorResult)
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.PropertyServiceGetByIDMethod), logFields...)
			statusCode, errRes = HandleError(errorResult)
		}
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

// HandleUpdateProperty handles updating a property
// @Summary Update a property
// @Description Updates a property's details
// @Tags properties
// @Accept json
// @Produce json
// @Param id query int true "Property ID"
// @Param property body dto.PropertyRequest true "Property details"
// @Success 200 {object} dto.PropertyResponse
// @Failure 400 {object} custom.ErrorResult
// @Failure 404 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties [put]
func HandleUpdateProperty(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleUpdatePropertyMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleUpdatePropertyMethod), commonLogFields...)

	var (
		statusCode      int
		errorResult     *custom.ErrorResult
		errRes          custom.ErrorResult
		request         dto.PropertyRequest
		response        dto.Property
		propertyService = services.CreatePropertyService(requestID, nil)
	)

	propertyID, err := GetIDFromParams(ctx)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleGetPropertyMethod), commonLogFields...)
		errorResult = err
		statusCode, errRes = HandleError(errorResult)
	} else {
		if err := ctx.BodyParser(&request); err != nil {
			logFields := log.TraceError(commonLogFields, err)
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleUpdatePropertyMethod), logFields...)
			errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.InvalidRequestErrorMessage, err.Error())
			errorResult = &errRes
			statusCode, errRes = HandleError(errorResult)
		} else {
			response, errorResult = propertyService.Update(uint(propertyID), request)
			if errorResult != nil {
				logFields := log.TraceCustomError(commonLogFields, *errorResult)
				log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.PropertyServiceUpdateMethod), logFields...)
				statusCode, errRes = HandleError(errorResult)
			}
		}
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

// HandleDeleteProperty handles deleting a property
// @Summary Delete a property
// @Description Deletes a property by its ID
// @Tags properties
// @Accept json
// @Produce json
// @Param id query int true "Property ID"
// @Success 200 {object} dto.PropertyResponse
// @Failure 400 {object} custom.ErrorResult
// @Failure 404 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties [delete]
func HandleDeleteProperty(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleDeletePropertyMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleDeletePropertyMethod), commonLogFields...)

	var (
		statusCode      int
		errorResult     *custom.ErrorResult
		errRes          custom.ErrorResult
		response        dto.Property
		propertyService = services.CreatePropertyService(requestID, nil)
	)

	propertyID, err := GetIDFromParams(ctx)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleGetPropertyMethod), commonLogFields...)
		errorResult = err
		statusCode, errRes = HandleError(errorResult)
	} else {
		response, errorResult = propertyService.Delete(uint(propertyID))
		if errorResult != nil {
			logFields := log.TraceCustomError(commonLogFields, *errorResult)
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.PropertyServiceDeleteMethod), logFields...)
			statusCode, errRes = HandleError(errorResult)
		}
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

// HandleListProperties handles listing properties
// @Summary List properties
// @Description Lists properties with pagination
// @Tags properties
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} []dto.PropertyResponse
// @Failure 400 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties [get]
func HandleListProperties(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleListPropertiesMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleListPropertiesMethod), commonLogFields...)

	var (
		statusCode      int
		errorResult     *custom.ErrorResult
		errRes          custom.ErrorResult
		response        []dto.Property
		propertyService = services.CreatePropertyService(requestID, nil)
	)

	page := ctx.QueryInt("page", 1)
	pageSize := ctx.QueryInt("page_size", 10)

	response, errorResult = propertyService.List(page, pageSize)
	if errorResult != nil {
		logFields := log.TraceCustomError(commonLogFields, *errorResult)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.PropertyServiceListMethod), logFields...)
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

// HandleListPropertiesByUser handles listing properties for a specific user
// @Summary List properties by user
// @Description Lists properties for a specific user with pagination
// @Tags properties
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} []dto.PropertyResponse
// @Failure 400 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties/user [get]
func HandleListPropertiesByUser(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleListPropertiesByUserMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleListPropertiesByUserMethod), commonLogFields...)

	var (
		statusCode      int
		errorResult     *custom.ErrorResult
		errRes          custom.ErrorResult
		response        []dto.Property
		propertyService = services.CreatePropertyService(requestID, nil)
	)

	// Get user ID from context
	userID, err := GetIDFromParams(ctx)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleListPropertiesByUserMethod), commonLogFields...)
		errorResult = err
		statusCode, errRes = HandleError(errorResult)
	} else {
		page := ctx.QueryInt("page", 1)
		pageSize := ctx.QueryInt("limit", 10)
		offset := (page - 1) * pageSize

		response, errorResult = propertyService.ListByUserID(userID, offset, pageSize)
		if errorResult != nil {
			logFields := log.TraceCustomError(commonLogFields, *errorResult)
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.PropertyServiceListByUserIDMethod), logFields...)
			statusCode, errRes = HandleError(errorResult)
		}
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
