package handler

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/app/services"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/gofiber/fiber/v2"
)

const (
	// Property handler methods
	PropertyHandlerCreateMethod  = "PropertyHandlerCreate"
	PropertyHandlerGetByIDMethod = "PropertyHandlerGetByID"
	PropertyHandlerUpdateMethod  = "PropertyHandlerUpdate"
	PropertyHandlerDeleteMethod  = "PropertyHandlerDelete"
	PropertyHandlerListMethod    = "PropertyHandlerList"
)

type PropertyHandler struct {
	_              struct{}
	handlerContext Context
	propertySvc    *services.PropertyService
}

// CreatePropertyHandler creates a new instance of PropertyHandler
func CreatePropertyHandler(requestID string) *PropertyHandler {
	return &PropertyHandler{
		handlerContext: CreateHandlerContext(requestID),
		propertySvc:    services.CreatePropertyService(requestID, nil),
	}
}

// Create handles property creation
func (h *PropertyHandler) Create(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyHandlerCreateMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyHandlerCreateMethod), commonLogFields...)

	// Parse request
	var request dto.PropertyRequest
	if err := c.BodyParser(&request); err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerCreateMethod), log.TraceError(commonLogFields, err)...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.BindingErrorMessage, "Request")
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}

	// Create property
	response, err := h.propertySvc.Create(request)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerCreateMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// GetByID handles getting a property by ID
func (h *PropertyHandler) GetByID(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyHandlerGetByIDMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyHandlerGetByIDMethod), commonLogFields...)

	// Get property ID from params
	propertyID, err := GetPropertyIDFromParams(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerGetByIDMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Get property
	response, err := h.propertySvc.GetByID(propertyID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerGetByIDMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Update handles property update
func (h *PropertyHandler) Update(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyHandlerUpdateMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyHandlerUpdateMethod), commonLogFields...)

	// Get property ID from params
	propertyID, err := GetPropertyIDFromParams(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerUpdateMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Parse request
	var request dto.PropertyRequest
	if err := c.BodyParser(&request); err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerUpdateMethod), log.TraceError(commonLogFields, err)...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.BindingErrorMessage, "Request")
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}

	// Update property
	response, err := h.propertySvc.Update(propertyID, request)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerUpdateMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Delete handles property deletion
func (h *PropertyHandler) Delete(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyHandlerDeleteMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyHandlerDeleteMethod), commonLogFields...)

	// Get property ID from params
	propertyID, err := GetPropertyIDFromParams(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerDeleteMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Delete property
	response, err := h.propertySvc.Delete(propertyID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerDeleteMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// List handles listing properties
func (h *PropertyHandler) List(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyHandlerListMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyHandlerListMethod), commonLogFields...)

	// Get pagination params
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)

	// List properties
	response, err := h.propertySvc.List(page, pageSize)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyHandlerListMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
