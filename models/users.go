package models

import (
	"time"
)

// User represents a user of the application.
type User struct {
	BaseModel
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	PhoneNumber string `gorm:"unique;not null"`
	Username    string `gorm:"unique;not null"`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Role        string `gorm:"not null"`
}

type SignUpInput struct {
	FirstName       string `json:"firstName" binding:"required,min=3"`
	LastName        string `json:"lastName" binding:"required,min=3"`
	PhoneNumber     string `json:"phoneNumber" binding:"required,min=10"`
	Username        string `json:"username" binding:"required,min=5"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Role            string `json:"role" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
