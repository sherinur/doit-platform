package model

import (
	"time"

	"github.com/sherinur/doit-platform/user-service/pkg/utils"
)

type User struct {
	ID              int64
	Name            string
	Phone           string
	Email           string
	Role            string
	CurrentPassword string
	NewPassword     string
	PasswordHash    string
	NewPasswordHash string
	CreatedAt       time.Time
	UpdatedAt       time.Time

	ISemailVerified bool
	IsDeleted       bool
}

func (u *User) Validate() error {
	switch {
	case !utils.ValidateEmail(u.Email):
		return ErrInvalidEmail
	case !utils.ValidatePassword(u.CurrentPassword):
		return ErrInvalidPassword
	default:
		return nil
	}
}

type UserUpdateData struct {
	ID              int64
	Name            string
	Phone           string
	Email           string
	Role            string
	NewPassword     string
	NewPasswordHash string
	UpdatedAt       time.Time
}

func (u *UserUpdateData) Validate() error {
	switch {
	case u.Name == "":
		return ErrInvalidName
	case u.Phone == "":
		return ErrInvalidPhone
	case !utils.ValidateEmail(u.Email):
		return ErrInvalidEmail
	case u.Role == "":
		return ErrInvalidRole
	default:
		return nil
	}
}
