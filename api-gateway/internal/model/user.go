package model

import "time"

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
