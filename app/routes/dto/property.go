package dto

import (
	"time"
)

// Property represents the property entity
type Property struct {
	*Base
	ID                uint              `gorm:"not null; column:id; primaryKey; autoIncrement"`
	UserID            uint              `gorm:"not null; column:user_id; index:idx_properties_on_user_id, type:btree"`
	Title             string            `gorm:"not null; column:title; type:varchar(150)"`
	Description       string            `gorm:"column:description; type:text"`
	PurposeID         uint              `gorm:"not null; column:purpose_id; index:idx_properties_on_purpose_id, type:btree"`
	PropertyTypeID    uint              `gorm:"not null; column:property_type_id; index:idx_properties_on_property_type_id, type:btree"`
	FurnitureTypeID   uint              `gorm:"column:furniture_type_id; index:idx_properties_on_furniture_type_id, type:btree"`
	ConditionID       uint              `gorm:"column:condition_id; index:idx_properties_on_condition_id, type:btree"`
	Bedrooms          int               `gorm:"column:bedrooms"`
	Bathrooms         int               `gorm:"column:bathrooms"`
	Size              float64           `gorm:"column:size"`
	SizeUnit          string            `gorm:"column:size_unit; type:varchar(20)"`
	City              string            `gorm:"not null; column:city; type:varchar(50)"`
	Address           string            `gorm:"not null; column:address; type:text"`
	PostalCode        string            `gorm:"column:postal_code; type:varchar(10)"`
	Latitude          float64           `gorm:"column:latitude"`
	Longitude         float64           `gorm:"column:longitude"`
	Price             float64           `gorm:"not null; column:price"`
	PriceUnit         string            `gorm:"not null; column:price_unit; type:varchar(20)"`
	IsNegotiable      bool              `gorm:"column:is_negotiable; default:false"`
	RentalPeriod      string            `gorm:"column:rental_period; type:varchar(20)"`
	IsRefundable      bool              `gorm:"column:is_refundable; default:false"`
	PricingType       string            `gorm:"not null; column:pricing_type; type:varchar(10)"`
	CreatedAt         time.Time         `gorm:"not null; column:created_at; default:CURRENT_TIMESTAMP"`
	PropertyAmenities []PropertyAmenity `gorm:"foreignKey:PropertyID"`
	PropertyUtilities []PropertyUtility `gorm:"foreignKey:PropertyID"`
	PropertyImages    []PropertyImage   `gorm:"foreignKey:PropertyID"`
}

// PropertyAmenity represents the many-to-many relationship between properties and amenities
type PropertyAmenity struct {
	*Base
	PropertyID uint `gorm:"not null; column:property_id; primaryKey"`
	AmenityID  uint `gorm:"not null; column:amenity_id; primaryKey"`
}

// PropertyUtility represents the many-to-many relationship between properties and utilities
type PropertyUtility struct {
	*Base
	PropertyID uint `gorm:"not null; column:property_id; primaryKey"`
	UtilityID  uint `gorm:"not null; column:utility_id; primaryKey"`
}

// PropertyImage represents the property images
type PropertyImage struct {
	*Base
	ID         uint   `gorm:"not null; column:id; primaryKey; autoIncrement"`
	PropertyID uint   `gorm:"not null; column:property_id; index:idx_property_images_on_property_id, type:btree"`
	URL        string `gorm:"not null; column:url; type:text"`
	IsPrimary  bool   `gorm:"column:is_primary; default:false"`
}

// PropertyRequest represents the request for creating/updating a property
type PropertyRequest struct {
	Title           string  `json:"title" validate:"required,max=150"`
	Description     string  `json:"description"`
	PurposeID       int     `json:"purpose_id" validate:"required"`
	PropertyTypeID  int     `json:"property_type_id" validate:"required"`
	FurnitureTypeID int     `json:"furniture_type_id"`
	ConditionID     int     `json:"condition_id"`
	Bedrooms        int     `json:"bedrooms"`
	Bathrooms       int     `json:"bathrooms"`
	Size            float64 `json:"size"`
	SizeUnit        string  `json:"size_unit" validate:"max=20"`
	City            string  `json:"city" validate:"required,max=50"`
	Address         string  `json:"address" validate:"required"`
	PostalCode      string  `json:"postal_code" validate:"max=10"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Price           float64 `json:"price" validate:"required"`
	PriceUnit       string  `json:"price_unit" validate:"required,max=20"`
	IsNegotiable    bool    `json:"is_negotiable"`
	RentalPeriod    string  `json:"rental_period" validate:"max=20"`
	IsRefundable    bool    `json:"is_refundable"`
	PricingType     string  `json:"pricing_type" validate:"required,oneof=sell rent stay"`
	AmenityIDs      []int   `json:"amenity_ids"`
	UtilityIDs      []int   `json:"utility_ids"`
}

// PropertyResponse represents the response for a property
type PropertyResponse struct {
	ID              uint      `json:"id"`
	UserID          uint      `json:"user_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	PurposeID       int       `json:"purpose_id"`
	PropertyTypeID  int       `json:"property_type_id"`
	FurnitureTypeID int       `json:"furniture_type_id"`
	ConditionID     int       `json:"condition_id"`
	Bedrooms        int       `json:"bedrooms"`
	Bathrooms       int       `json:"bathrooms"`
	Size            float64   `json:"size"`
	SizeUnit        string    `json:"size_unit"`
	City            string    `json:"city"`
	Address         string    `json:"address"`
	PostalCode      string    `json:"postal_code"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	Price           float64   `json:"price"`
	PriceUnit       string    `json:"price_unit"`
	IsNegotiable    bool      `json:"is_negotiable"`
	RentalPeriod    string    `json:"rental_period"`
	IsRefundable    bool      `json:"is_refundable"`
	PricingType     string    `json:"pricing_type"`
	CreatedAt       time.Time `json:"created_at"`
	Amenities       []int     `json:"amenities"`
	Utilities       []int     `json:"utilities"`
	Images          []string  `json:"images"`
}

// PropertyListResponse represents the list of property responses
type PropertyListResponse []PropertyResponse
