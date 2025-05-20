package web

import (
	"fmt"
	"time"

	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

// SetupFiber is a function that sets up the Fiber app.
func SetupFiber(idleTimeout time.Duration) *fiber.App {
	// log-----
	return fiber.New(fiber.Config{
		//Prefork:               false,
		IdleTimeout:           idleTimeout,
		DisableStartupMessage: false,
	})
}

// Shutdown is a function that shuts down the Fiber app.
func Shutdown(app *fiber.App) error {
	// log.Logger.Info("Shutting down Fiber...")
	err := app.Shutdown()
	// log.Logger.Info("Shutdown complete")
	return err
}

// GetRequestID is a function that gets the request ID from the Fiber context.
func GetRequestID(ctx *fiber.Ctx) string {
	return ctx.Locals(requestid.ConfigDefault.ContextKey).(string)
}

// GetHeaderFromRequest is a function that gets the header value from the Fiber context.
func GetHeaderFromRequest(commonLogFields []zap.Field, ctx *fiber.Ctx, headerKey string) string {
	headerValue := ctx.Get(headerKey)
	if headerValue == constant.Empty {
		log.Logger.Debug(fmt.Sprintf(constant.EmptyHeaderDetails, headerKey), commonLogFields...)
	}

	return headerValue
}
