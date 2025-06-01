package handler

import (
	"strconv"

	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/gofiber/fiber/v2"
)

// Context represents the handler context
type Context struct {
	RequestID string
}

// CreateHandlerContext creates a new handler context
func CreateHandlerContext(requestID string) Context {
	return Context{
		RequestID: requestID,
	}
}

// GetUserIDFromContext extracts the user ID from the request context
func GetUserIDFromContext(c *fiber.Ctx) (uint, *custom.ErrorResult) {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		errRes := custom.BuildBadReqErrResult(constant.ErrAccessTokenCode, "User ID not found in context", "Authorization")
		return 0, &errRes
	}
	return userID, nil
}

// GetIDFromParams extracts the property ID from the request parameters
func GetIDFromParams(c *fiber.Ctx) (uint, *custom.ErrorResult) {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, "Invalid property ID", "ID")
		return 0, &errRes
	}
	return uint(id), nil
}
