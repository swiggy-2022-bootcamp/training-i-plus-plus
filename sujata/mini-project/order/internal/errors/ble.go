package errors

import (
	"net/http"
	"order/internal/literals"
)

var (
	InternalError = ServerError{
		ErrorMessage:     "an error occurred while handling the request",
		HttpResponseCode: http.StatusInternalServerError,
	}

	BadRequest = ServerError{
		ErrorMessage:     literals.BadRequest,
		HttpResponseCode: http.StatusBadRequest,
	}
)
