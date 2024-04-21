package models

import (
	"fmt"
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
	ErrBadRequest         = 400
	ErrUnauthorized       = 401
	ErrForbidden          = 403
	ErrNotFound           = 404
	ErrConflict           = 409
	ErrInternalServer     = 500
	ErrServiceUnavailable = 503
	ErrGatewayTimeout     = 504
	ErrUnknown            = 520
	ErrDatabase           = 521
)
