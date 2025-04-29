package dao

import (
	"time"
	"user-services/internal/domain/model"
)

type User struct {
	ID              int64     `db:"id"`
	Name            string    `db:"name"`
	Phone           string    `db:"phone"`
	Email           string    `db:"email"`
	PasswordHash    string    `db:"password_hash"`
	NewPasswordHash string    `db:"new_password_hash"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	IsDeleted       bool      `db:"is_deleted"`
}

func FromDomain(user *model.User) User {
	return User{
		ID:              user.ID,
		Name:            user.Name,
		Phone:           user.Phone,
		Email:           user.Email,
		PasswordHash:    user.PasswordHash,
		NewPasswordHash: user.NewPasswordHash,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		IsDeleted:       user.IsDeleted,
	}
}

func ToDomain(user User) *model.User {
	return &model.User{
		ID:              user.ID,
		Name:            user.Name,
		Phone:           user.Phone,
		Email:           user.Email,
		PasswordHash:    user.PasswordHash,
		NewPasswordHash: user.NewPasswordHash,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		IsDeleted:       user.IsDeleted,
	}
}
