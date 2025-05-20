package dto

import "time"

// Property represents the properties table
type Property struct {
	ID              uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          uint              `gorm:"not null" json:"user_id"`
	User            User              `gorm:"foreignKey:UserID" json:"user"`
	Title           string            `gorm:"type:varchar(150)" json:"title"`
	Description     string            `gorm:"type:text" json:"description"`
	PurposeID       uint              `gorm:"not null" json:"purpose_id"`
	Purpose         PurposeType       `gorm:"foreignKey:PurposeID" json:"purpose"`
	PropertyTypeID  uint              `gorm:"not null" json:"property_type_id"`
	PropertyType    PropertyType      `gorm:"foreignKey:PropertyTypeID" json:"property_type"`
	FurnitureTypeID uint              `gorm:"not null" json:"furniture_type_id"`
	FurnitureType   FurnitureType     `gorm:"foreignKey:FurnitureTypeID" json:"furniture_type"`
	ConditionID     uint              `gorm:"not null" json:"condition_id"`
	Condition       PropertyCondition `gorm:"foreignKey:ConditionID" json:"condition"`
	Bedrooms        int               `json:"bedrooms"`
	Bathrooms       int               `json:"bathrooms"`
	Size            float64           `json:"size"`
	SizeUnit        string            `gorm:"type:varchar(20)" json:"size_unit"`
	City            string            `gorm:"type:varchar(50)" json:"city"`
	Address         string            `gorm:"type:text" json:"address"`
	PostalCode      string            `gorm:"type:varchar(10)" json:"postal_code"`
	Latitude        float64           `json:"latitude"`
	Longitude       float64           `json:"longitude"`
	Price           float64           `json:"price"`
	PriceUnit       string            `gorm:"type:varchar(20)" json:"price_unit"`
	IsNegotiable    bool              `gorm:"default:false" json:"is_negotiable"`
	RentalPeriod    string            `gorm:"type:varchar(20)" json:"rental_period"`
	IsRefundable    bool              `gorm:"default:false" json:"is_refundable"`
	PricingType     string            `gorm:"type:varchar(10);check:pricing_type IN ('sell', 'rent', 'stay')" json:"pricing_type"`
	CreatedAt       time.Time         `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	// Many-to-many relationships
	Amenities []Amenity       `gorm:"many2many:property_amenities;" json:"amenities"`
	Utilities []Utility       `gorm:"many2many:property_utilities;" json:"utilities"`
	Images    []PropertyImage `json:"images"`
}

// PropertyImage represents the property_images table
type PropertyImage struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	PropertyID uint   `gorm:"not null" json:"property_id"`
	URL        string `gorm:"type:text;not null" json:"url"`
	IsPrimary  bool   `gorm:"default:false" json:"is_primary"`
}

// Favourite represents the favourites table
type Favourite struct {
	ID         uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint     `gorm:"not null" json:"user_id"`
	PropertyID uint     `gorm:"not null" json:"property_id"`
	User       User     `gorm:"foreignKey:UserID" json:"user"`
	Property   Property `gorm:"foreignKey:PropertyID" json:"property"`
}
