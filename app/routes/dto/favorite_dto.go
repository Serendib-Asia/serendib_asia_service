package dto

// FavoriteRequest represents the request to add a property to favorites
type FavoriteRequest struct {
	PropertyID uint `json:"property_id" validate:"required"` // INTEGER REFERENCES properties(id)
}

// RemoveFromFavoritesRequest represents the request to remove a property from favorites
type RemoveFromFavoritesRequest struct {
	PropertyID uint `json:"property_id" validate:"required"` // INTEGER REFERENCES properties(id)
}

// FavoriteResponse represents a favorite property
type FavoriteResponse struct {
	ID         uint           `json:"id"`          // SERIAL PRIMARY KEY
	UserID     uint           `json:"user_id"`     // INTEGER REFERENCES users(id)
	PropertyID uint           `json:"property_id"` // INTEGER REFERENCES properties(id)
	Property   PropertyDetail `json:"property"`    // Nested property details
}

// PropertyDetail represents the property details in a favorite
type PropertyDetail struct {
	ID          uint    `json:"id"`          // INTEGER REFERENCES properties(id)
	Title       string  `json:"title"`       // VARCHAR(150)
	Description string  `json:"description"` // TEXT
	Price       float64 `json:"price"`       // FLOAT
	PriceUnit   string  `json:"price_unit"`  // VARCHAR(20)
	City        string  `json:"city"`        // VARCHAR(50)
	Address     string  `json:"address"`     // TEXT
	URL         string  `json:"url"`         // TEXT (from property_images)
}

// FavoriteListResponse represents a list of favorite properties
type FavoriteListResponse struct {
	Items []FavoriteResponse `json:"items"`
	Total int64              `json:"total"`
}
