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
	// User handler methods
	UserHandlerRegisterMethod       = "UserHandlerRegister"
	UserHandlerLoginMethod          = "UserHandlerLogin"
	UserHandlerGetProfileMethod     = "UserHandlerGetProfile"
	UserHandlerUpdateProfileMethod  = "UserHandlerUpdateProfile"
	UserHandlerUpdatePasswordMethod = "UserHandlerUpdatePassword"
)

type UserHandler struct {
	_              struct{}
	handlerContext Context
	userSvc        *services.UserService
}

// CreateUserHandler creates a new instance of UserHandler
func CreateUserHandler(requestID string) *UserHandler {
	return &UserHandler{
		handlerContext: CreateHandlerContext(requestID),
		userSvc:        services.CreateUserService(requestID, nil),
	}
}

// Register handles user registration
func (h *UserHandler) Register(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserHandlerRegisterMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserHandlerRegisterMethod), commonLogFields...)

	// Parse request
	var request dto.UserRegisterRequest
	if err := c.BodyParser(&request); err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerRegisterMethod), log.TraceError(commonLogFields, err)...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.BindingErrorMessage, "Request")
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}

	// Register user
	response, err := h.userSvc.Register(&request)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerRegisterMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// Login handles user login
func (h *UserHandler) Login(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserHandlerLoginMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserHandlerLoginMethod), commonLogFields...)

	// Parse request
	var request dto.UserLoginRequest
	if err := c.BodyParser(&request); err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerLoginMethod), log.TraceError(commonLogFields, err)...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.BindingErrorMessage, "Request")
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}

	// Login user
	response, err := h.userSvc.Login(&request)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerLoginMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// GetProfile retrieves user profile
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserHandlerGetProfileMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserHandlerGetProfileMethod), commonLogFields...)

	// Get user ID from context
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerGetProfileMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	// Get profile
	response, err := h.userSvc.GetProfile(userID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerGetProfileMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateProfile updates user profile
func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserHandlerUpdateProfileMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserHandlerUpdateProfileMethod), commonLogFields...)

	// Get user ID from context
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerUpdateProfileMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	// Parse request
	var request dto.UserUpdateProfileRequest
	if err := c.BodyParser(&request); err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerUpdateProfileMethod), log.TraceError(commonLogFields, err)...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.BindingErrorMessage, "Request")
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}

	// Update profile
	response, err := h.userSvc.UpdateProfile(userID, &request)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerUpdateProfileMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdatePassword updates user password
func (h *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	commonLogFields := log.CommonLogField(h.handlerContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserHandlerUpdatePasswordMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserHandlerUpdatePasswordMethod), commonLogFields...)

	// Get user ID from context
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerUpdatePasswordMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	// Parse request
	var request dto.UserUpdatePasswordRequest
	if err := c.BodyParser(&request); err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerUpdatePasswordMethod), log.TraceError(commonLogFields, err)...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.BindingErrorMessage, "Request")
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}

	// Update password
	err = h.userSvc.UpdatePassword(userID, &request)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserHandlerUpdatePasswordMethod), log.TraceError(commonLogFields, err)...)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusOK)
}
