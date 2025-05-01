package model

import (
	"time"
)

type User struct {
	ID              int64
	Name            string
	Phone           string
	Email           string
	CurrentPassword string
	NewPassword     string
	PasswordHash    string
	NewPasswordHash string
	CreatedAt       time.Time
	UpdatedAt       time.Time

	IsDeleted bool
}

// func (u *User) Validate() error {
// 	switch {
// 	case u.Name == "" || len(u.Name) < 2:
// 		return ErrInvalidName
// 	case u.Phone == "":
// 		return ErrInvalidPhone
// 	case !utils.ValidateEmail(u.Email):
// 		return ErrInvalidEmail
// 	case !utils.ValidatePassword(u.CurrentPassword):
// 		return ErrInvalidPassword
// 	default:
// 		return nil
// 	}
// }

type UserUpdateData struct {
	ID           *uint64
	Name         *string
	Phone        *string
	Email        *string
	PasswordHash *string
	UpdatedAt    *time.Time

	IsDeleted *bool
}

// func (u *UserUpdateData) Validate() error {
// 	switch {
// 	case *u.Name == "" || len(*u.Name) < 2:
// 		return ErrInvalidName
// 	case *u.Phone == "":
// 		return ErrInvalidPhone
// 	case !utils.ValidateEmail(*u.Email):
// 		return ErrInvalidEmail
// 	case *u.PasswordHash == "":
// 		return ErrInvalidPassword
// 	default:
// 		return nil
// 	}
// }
