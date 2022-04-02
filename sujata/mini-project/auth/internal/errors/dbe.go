package errors

import (
	"auth/internal/literals"
	"net/http"
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
	UserAlreadyExists = ServerError{
		ErrorMessage:     literals.AppPrefix + ": " + literals.DBUserAlreadyExists,
		HttpResponseCode: http.StatusConflict,
	}
)
