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
	default:
		return &HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
}
