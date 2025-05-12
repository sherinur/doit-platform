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

	IsDeleted bool `db:"is_deleted"`
}

func FromDomain(user *model.User) User {
	return User{
		ID:           user.ID,
		Name:         user.Name,
		Phone:        user.Phone,
		Email:        user.Email,
		Role:         user.Role,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		IsDeleted:    user.IsDeleted,
	}
}

func ToDomain(user User) *model.User {
	return &model.User{
		ID:           user.ID,
		Name:         user.Name,
		Phone:        user.Phone,
		Email:        user.Email,
		Role:         user.Role,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		IsDeleted:    user.IsDeleted,
	}
}

type UserUpdateData struct {
	Name      string
	Phone     string
	Email     string
	Role      string
	UpdatedAt time.Time
}

func FromUserUpdateData(userUpdateData *model.UserUpdateData) UserUpdateData {
	return UserUpdateData{
		Name:      userUpdateData.Name,
		Phone:     userUpdateData.Phone,
		Email:     userUpdateData.Email,
		Role:      userUpdateData.Role,
		UpdatedAt: userUpdateData.UpdatedAt,
	}
}

func ToUserUpdateData(userUpdateData UserUpdateData) *model.UserUpdateData {
	return &model.UserUpdateData{
		Name:      userUpdateData.Name,
		Phone:     userUpdateData.Phone,
		Email:     userUpdateData.Email,
		Role:      userUpdateData.Role,
		UpdatedAt: userUpdateData.UpdatedAt,
	}
}
