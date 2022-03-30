package errors

import "net/http"

type UserError struct {
	Status       int
	ErrorMessage string
}

func (productError *UserError) Error() string {
	return productError.ErrorMessage
}

func UnauthorizedError() *UserError {
	return &UserError{http.StatusUnauthorized, "Incorrect Credentials"}
}

func MalformedIdError() *UserError {
	return &UserError{http.StatusBadRequest, "Malformed user id"}
}

func IdNotFoundError() *UserError {
	return &UserError{http.StatusNotFound, "user with given id not found"}
}

func UnmarshallError() *UserError {
	return &UserError{http.StatusBadRequest, "Couldn't unmarshall user body in request"}
}

func InternalServerError() *UserError {
	return &UserError{http.StatusInternalServerError, "Internal Server Error"}
}
