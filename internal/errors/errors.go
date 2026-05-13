package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidRequest  = errors.New("invalid request")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrExternalService = errors.New("external service error")
	ErrRequestCreation = errors.New("request creation error")
)

type AppError struct {
	Code       string
	Message    string
	StatusCode int
	Err        error
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrDatabaseDown = &AppError{
		Code:       "DATABASE_DOWN",
		Message:    "database is unavailable",
		StatusCode: http.StatusServiceUnavailable,
	}

	ErrInternal = &AppError{
		Code:       "INTERNAL_ERROR",
		Message:    "internal server error",
		StatusCode: http.StatusInternalServerError,
	}
)

func HandleHTTPError(w http.ResponseWriter, err error) {
	appErr, ok := err.(*AppError)

	if !ok {
		appErr = ErrInternal
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.StatusCode)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"code":    appErr.Code,
		"message": appErr.Message,
	})
}
