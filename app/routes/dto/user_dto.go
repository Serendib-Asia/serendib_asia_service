package dto

import (
	"time"
)

// UserRegisterRequest represents the request to register a new user
type UserRegisterRequest struct {
	FullName    string `json:"full_name" validate:"required,max=100"`
	Email       string `json:"email" validate:"required,email,max=100"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phone_number" validate:"max=15"`
}

// UserLoginRequest represents the request to login a user
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UserLoginResponse represents the response after successful login
type UserLoginResponse struct {
	User  UserProfileResponse `json:"user"`
	Token string              `json:"token"`
}

// UserProfileResponse represents a user's profile information
type UserProfileResponse struct {
	ID           uint      `json:"id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
}

// UserUpdateProfileRequest represents the request to update a user's profile
type UserUpdateProfileRequest struct {
	FullName     string `json:"full_name" validate:"required,max=100"`
	PhoneNumber  string `json:"phone_number" validate:"max=15"`
	ProfileImage string `json:"profile_image"`
}

// UserUpdatePasswordRequest represents the request to update a user's password
type UserUpdatePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
}
