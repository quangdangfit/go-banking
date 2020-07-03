package users

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UID      string
	Username string
	Email    string
	Password string
}

type UserResponse struct {
	UID      string
	Username string
	Email    string
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

type LoginRequest struct {
	Username string
	Password string
}
