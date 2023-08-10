package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct{

	ID        			uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      			string    `gorm:"type:varchar(255);not null"`
	Email     			string    `gorm:"uniqueIndex;not null"`
	Password  			string    `gorm:"not null"`
	PasswordResetToken	string	  
	PasswordResetAt		time.Time
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
}

type RegisterInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	Photo           string `json:"photo"`
}
// type SignUpInput struct {
// 	Name            string `json:"name" binding:"required"`
// 	Email           string `json:"email" binding:"required"`
// 	Password        string `json:"password" binding:"required,min=8"`
// 	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
// 	Photo           string `json:"photo" binding:"required"`
// }

type LoginInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type ForgotPasswordInput struct{
	Email string `json:"email" binding:"required"`
}

type ResetPassword struct {
	Password 		string `json:"password"  binding:"required"`
	PasswordConfirm	string `json:"passwordConfirm" binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}