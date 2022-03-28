package errors

import (
	"cart/internal/literals"
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
		ErrorMessage:     literals.AppPrefix + literals.ParametersMissing,
		HttpResponseCode: http.StatusBadRequest,
	}

	QueryParamMissingError = ServerError{
		ErrorMessage:     literals.AppPrefix + literals.QueryParamMissing,
		HttpResponseCode: http.StatusBadRequest,
	}

	MalformedQueryParamError = ServerError{
		ErrorMessage:     literals.AppPrefix + literals.MalformedQueryParam,
		HttpResponseCode: http.StatusBadRequest,
	}
)
