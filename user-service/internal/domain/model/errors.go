package model

import "errors"

var (
	ErrInvalidName     = errors.New("name must not be empty")
	ErrInvalidPhone    = errors.New("phone must not be empty")
	ErrInvalidEmail    = errors.New("email must be a valid email address")
	ErrInvalidPassword = errors.New("password must not be empty")
)
