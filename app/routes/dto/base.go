package dto

import (
	"time"

	"gorm.io/gorm"
)

// Base is a struct that represents the base fields for all other structs.
type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
