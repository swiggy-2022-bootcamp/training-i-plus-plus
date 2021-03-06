package errors

import (
	"auth/internal/literals"
	"net/http"
)

// ServerError defines the error message and related http response code for any
// error return by the microservice.
type ServerError struct {
	ErrorMessage     string
	HttpResponseCode int
}

// Validation errors
var (
	ParametersMissingError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.ParametersMissing,
		HttpResponseCode: http.StatusBadRequest,
	}

	InvalidEmailFormatError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.InvalidEmailFormat,
		HttpResponseCode: http.StatusBadRequest,
	}

	WeakPasswordError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.WeakPasswordError,
		HttpResponseCode: http.StatusBadRequest,
	}

	IncorrectUserPasswordError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.IncorrectUserPassword,
		HttpResponseCode: http.StatusUnauthorized,
	}

	IncorrectUserRoleError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.IncorrectUserRole,
		HttpResponseCode: http.StatusBadRequest,
	}
)
