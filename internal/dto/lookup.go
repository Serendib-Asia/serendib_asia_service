package dto

// PurposeType represents the purpose_types table
type PurposeType struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(20);unique;not null" json:"name"`
}

// PropertyType represents the property_types table
type PropertyType struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

// FurnitureType represents the furniture_types table
type FurnitureType struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

// PropertyCondition represents the property_conditions table
type PropertyCondition struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

// Utility represents the utilities table
type Utility struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

// Amenity represents the amenities table
type Amenity struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}
