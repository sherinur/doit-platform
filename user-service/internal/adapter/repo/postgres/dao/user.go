package dao

import (
	"time"

	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

type User struct {
	ID           int64     `db:"id"`
	Name         string    `db:"name"`
	Phone        string    `db:"phone"`
	Email        string    `db:"email"`
	Role         string    `db:"role"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`

	IsemailVerified bool `db:"is_email_verified"`
	IsDeleted       bool `db:"is_deleted"`
}

func FromDomain(user *model.User) User {
	return User{
		ID:              user.ID,
		Name:            user.Name,
		Phone:           user.Phone,
		Email:           user.Email,
		Role:            user.Role,
		PasswordHash:    user.PasswordHash,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		IsemailVerified: user.ISemailVerified,
		IsDeleted:       user.IsDeleted,
	}
}

func ToDomain(user User) *model.User {
	return &model.User{
		ID:              user.ID,
		Name:            user.Name,
		Phone:           user.Phone,
		Email:           user.Email,
		Role:            user.Role,
		PasswordHash:    user.PasswordHash,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		ISemailVerified: user.IsemailVerified,
		IsDeleted:       user.IsDeleted,
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

func FromUserUpdateData(userUpdateData *model.UserUpdateData) UserUpdateData {
	return UserUpdateData{
		ID:              userUpdateData.ID,
		Name:            userUpdateData.Name,
		Phone:           userUpdateData.Phone,
		Email:           userUpdateData.Email,
		Role:            userUpdateData.Role,
		NewPassword:     userUpdateData.NewPassword,
		NewPasswordHash: userUpdateData.NewPasswordHash,
		UpdatedAt:       userUpdateData.UpdatedAt,
	}
}

func ToUserUpdateData(userUpdateData UserUpdateData) *model.UserUpdateData {
	return &model.UserUpdateData{
		ID:              userUpdateData.ID,
		Name:            userUpdateData.Name,
		Phone:           userUpdateData.Phone,
		Email:           userUpdateData.Email,
		Role:            userUpdateData.Role,
		NewPassword:     userUpdateData.NewPassword,
		NewPasswordHash: userUpdateData.NewPasswordHash,
		UpdatedAt:       userUpdateData.UpdatedAt,
	}
}
