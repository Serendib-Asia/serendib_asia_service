package routes

import (
	"github.com/chazool/serendib_asia_service/app/routes/handler"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// APIRoutes sets up the API routes for the application.
func APIRoutes(app *fiber.App) {
	// root doc path
	app.Get("/docs", handler.HandleDoc)
	// static route
	app.Get("/docs/staticDoc", handler.HandleStaticDoc)
	// swagger route
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	route := app.Group("/api/v1")

	// property related endpoints
	property := route.Group("/properties")
	property.Post("/", handler.HandleCreateProperty)
	property.Get("/:id", handler.HandleGetProperty)
	property.Put("/:id", handler.HandleUpdateProperty)
	property.Delete("/:id", handler.HandleDeleteProperty)
	property.Get("/", handler.HandleListProperties)
	property.Get("/user/:id", handler.HandleListPropertiesByUser)

	// property image routes
	property.Post("/:propertyId/images", handler.HandleUploadImage)
	property.Delete("/images/:imageId", handler.HandleDeleteImage)
	property.Put("/images/:imageId/primary", handler.HandleSetPrimaryImage)
	property.Get("/:propertyId/images", handler.HandleListImages)

	// lookup tables endpoints
	lookup := route.Group("/lookups")
	lookup.Get("/purpose-types", handler.HandleGetPurposeTypes)
	lookup.Get("/property-types", handler.HandleGetPropertyTypes)
	lookup.Get("/furniture-types", handler.HandleGetFurnitureTypes)
	lookup.Get("/conditions", handler.HandleGetConditions)
	lookup.Get("/utilities", handler.HandleGetUtilities)
	lookup.Get("/amenities", handler.HandleGetAmenities)

	// user management endpoints
	user := route.Group("/users")
	userHandler := handler.CreateUserHandler("")
	// register user
	user.Post("/register", userHandler.Register)
	// login user
	user.Post("/login", userHandler.Login)
	// get user profile
	user.Get("/profile", userHandler.GetProfile)
	// update user profile
	user.Put("/profile", userHandler.UpdateProfile)
	// update user password
	user.Put("/password", userHandler.UpdatePassword)

	// property favorites endpoints
	favorites := route.Group("/favorites")
	favoriteHandler := handler.CreateFavoriteHandler("")
	// add to favorites
	favorites.Post("", favoriteHandler.AddFavorite)
	// remove from favorites
	favorites.Delete("/:id", favoriteHandler.RemoveFavorite)
	// list user favorites
	favorites.Get("", favoriteHandler.ListFavorites)
}
