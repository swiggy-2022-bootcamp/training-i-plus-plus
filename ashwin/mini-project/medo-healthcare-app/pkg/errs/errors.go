package errs

import (
	"net/http"

	"github.com/pkg/errors"
)

//AppError ..
type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

//Error ..
func (e AppError) Error() error {
	return errors.New(e.Message)
}

//AsMessage ..
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

//NewNotFoundError ..
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

//NewUnexpectedError ..
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

//NewValidationError ..
func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

//NewBadRequest ..
func NewBadRequest(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}
