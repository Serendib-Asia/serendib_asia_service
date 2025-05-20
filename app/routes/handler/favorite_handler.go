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
	// Favorite handler methods
	FavoriteHandlerAddMethod    = "FavoriteHandlerAdd"
	FavoriteHandlerRemoveMethod = "FavoriteHandlerRemove"
	FavoriteHandlerListMethod   = "FavoriteHandlerList"
)

type FavoriteHandler struct {
	_              struct{}
	handlerContext Context
	favoriteSvc    services.FavoriteService
}

// CreateFavoriteHandler creates a new instance of FavoriteHandler
func CreateFavoriteHandler(requestID string) *FavoriteHandler {
	return &FavoriteHandler{
		handlerContext: CreateHandlerContext(requestID),
		favoriteSvc:    services.CreateFavoriteService(requestID),
	}
}

// AddFavorite adds a property to user's favorites
func (h *FavoriteHandler) AddFavorite(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(FavoriteHandlerAddMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(FavoriteHandlerAddMethod), commonLogFields...)

	// Get user ID from context
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteHandlerAddMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	// Parse request
	var request dto.FavoriteRequest
	if err := c.BodyParser(&request); err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteHandlerAddMethod), log.TraceError(commonLogFields, err)...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.BindingErrorMessage, "Request")
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}

	// Add to favorites
	err = h.favoriteSvc.AddFavorite(userID, request.PropertyID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteHandlerAddMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

// RemoveFavorite removes a property from user's favorites
func (h *FavoriteHandler) RemoveFavorite(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(FavoriteHandlerRemoveMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(FavoriteHandlerRemoveMethod), commonLogFields...)

	// Get user ID from context
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteHandlerRemoveMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	// Get property ID from params
	propertyID, err := GetPropertyIDFromParams(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteHandlerRemoveMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Remove from favorites
	err = h.favoriteSvc.RemoveFavorite(userID, propertyID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteHandlerRemoveMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

// ListFavorites lists user's favorite properties
func (h *FavoriteHandler) ListFavorites(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(FavoriteHandlerListMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(FavoriteHandlerListMethod), commonLogFields...)

	// Get user ID from context
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteHandlerListMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	// Get favorites
	favorites, err := h.favoriteSvc.ListFavorites(userID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteHandlerListMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(favorites)
}
