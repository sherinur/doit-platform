package dto

import (
	"net/http"
)

type HTTPError struct {
	Code    int
	Message string
}

func FromError(err error) *HTTPError {
	switch {
	// case errors.Is(err, model.ErrInvalidEmail):
	// 	return ErrInvalidEmail
	// case errors.Is(err, model.ErrInvalidPassword):
	// 	return ErrInvalidPassword
	// case errors.Is(err, model.ErrInvalidID):
	// 	return ErrInvalidID
	default:
		return &HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "something went wrong",
		}
	}
}
