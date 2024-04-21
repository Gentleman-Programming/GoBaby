package models

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Err     error
	Message string
	Code    int
}

func (e *AppError) String() string {
	return fmt.Sprintf("error: %s (code: %d)", e.Message, e.Code)
}

// Custom error codes

const (
	ErrBadRequest         = http.StatusBadRequest
	ErrUnauthorized       = http.StatusUnauthorized
	ErrForbidden          = http.StatusForbidden
	ErrNotFound           = http.StatusNotFound
	ErrConflict           = http.StatusConflict
	ErrInternalServer     = http.StatusInternalServerError
	ErrServiceUnavailable = http.StatusServiceUnavailable
	ErrGatewayTimeout     = http.StatusGatewayTimeout
)
