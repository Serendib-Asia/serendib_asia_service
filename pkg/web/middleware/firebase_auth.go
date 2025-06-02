package middleware

import (
	"context"
	"strings"

	"github.com/chazool/serendib_asia_service/pkg/config"
	"github.com/chazool/serendib_asia_service/pkg/config/firebase"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"
	"github.com/chazool/serendib_asia_service/pkg/web"

	firebaseApp "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

// FirebaseAuthMiddleware creates a middleware that validates Firebase JWT tokens
func FirebaseAuthMiddleware() fiber.Handler {
	// Initialize Firebase app
	config := config.GetConfig().FirebaseConfig
	firebaseConfig := firebaseApp.Config{
		ProjectID: config.ProjectID,
	}
	serviceAccountJSON := firebase.GetServiceAccountJSON()

	opt := option.WithCredentialsJSON(serviceAccountJSON)

	app, err := firebaseApp.NewApp(context.Background(), &firebaseConfig, opt)
	if err != nil {
		log.Logger.Fatal("Error initializing Firebase app", log.TraceError(nil, err)...)
	}

	// Get Firebase Auth client
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Logger.Fatal("Error getting Firebase Auth client", log.TraceError(nil, err)...)
	}

	return func(c *fiber.Ctx) error {
		requestID := web.GetRequestID(c)
		commonLogFields := log.CommonLogField(requestID)

		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Logger.Error(constant.ErrEmptyAuthHeaderMsg, commonLogFields...)
			errRes := custom.BuildBadReqErrResult(constant.ErrEmptyAuthHeaderCode, constant.ErrEmptyAuthHeaderMsg, constant.Empty)
			return c.Status(fiber.StatusUnauthorized).JSON(errRes)
		}

		// Extract token from Bearer header
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Logger.Error(constant.ErrInvalidAuthHeaderMsg, commonLogFields...)
			errRes := custom.BuildBadReqErrResult(constant.ErrInvalidAuthHeaderCode, constant.ErrInvalidAuthHeaderMsg, constant.Empty)
			return c.Status(fiber.StatusUnauthorized).JSON(errRes)
		}

		token := parts[1]

		// Verify the token
		decodedToken, err := client.VerifyIDToken(context.Background(), token)
		if err != nil {
			log.Logger.Error(constant.ErrInParsingTokenMsg, log.TraceError(commonLogFields, err)...)
			errRes := custom.BuildBadReqErrResult(constant.ErrInvalidTokenCode, constant.ErrInParsingTokenMsg, constant.Empty)
			return c.Status(fiber.StatusUnauthorized).JSON(errRes)
		}

		// Store user information in context
		c.Locals("user_id", decodedToken.UID)
		c.Locals("user_email", decodedToken.Claims["email"])
		c.Locals("user_name", decodedToken.Claims["name"])

		return c.Next()
	}
}
