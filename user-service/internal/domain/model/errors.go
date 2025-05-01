package model

import "errors"

var (
	ErrInvalidName     = errors.New("name must not be empty")
	ErrInvalidPhone    = errors.New("phone must not be empty")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidID       = errors.New("invalid id")
	ErrUserExists      = errors.New("the user is already exists")
	ErrInvalidPassword = errors.New("invalid password")
)
