package dto

import "time"

// User represents the users table
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName     string    `gorm:"type:varchar(100);not null" json:"full_name"`
	Email        string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	PasswordHash string    `gorm:"type:text;not null" json:"-"`
	PhoneNumber  string    `gorm:"type:varchar(15)" json:"phone_number"`
	ProfileImage string    `gorm:"type:text" json:"profile_image"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
