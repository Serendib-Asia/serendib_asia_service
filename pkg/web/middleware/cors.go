package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CorsMiddleware provide Fiber's built-in middlewares
// see: https://docs.gofiber.io/v1.x/api/middleware/
func CorsMiddleware(app *fiber.App) {
	app.Use(cors.New())
}
