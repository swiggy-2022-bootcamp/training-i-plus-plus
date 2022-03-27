package errors

import (
	"net/http"
	"search/internal/literals"
)

// database error
var (
	DatabaseInsertionError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.DBInsertionError,
		HttpResponseCode: http.StatusInternalServerError,
	}
	DatabaseNoInsertionError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.DBInsertionFail,
		HttpResponseCode: http.StatusInternalServerError,
	}
	UserNotFoundError = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.DBUserNotFound,
		HttpResponseCode: http.StatusNotFound,
	}
)
