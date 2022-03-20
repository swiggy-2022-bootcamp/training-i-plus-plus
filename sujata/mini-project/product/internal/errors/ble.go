package errors

import "net/http"

var (
	UnmarshalError = ServerError{
		ErrorMessage:     "an error occured while unmarshalling the request",
		HttpResponseCode: http.StatusInternalServerError,
	}

	InternalError = ServerError{
		ErrorMessage:     "an error occurred while handling the request",
		HttpResponseCode: http.StatusInternalServerError,
	}
)
