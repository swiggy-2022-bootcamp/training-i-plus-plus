package errors

import (
	"net/http"
	"order/internal/literals"
)

var (
	DatabaseInsertionError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.DBInsertionError,
		HttpResponseCode: http.StatusInternalServerError,
	}
	DatabaseNoInsertionError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.DBInsertionFail,
		HttpResponseCode: http.StatusInternalServerError,
	}
)
