package errors

import "net/http"

type UserError struct {
	Status       int
	ErrorMessage string
}

func (productError *UserError) Error() string {
	return productError.ErrorMessage
}

func AccessDenied() *UserError {
	return &UserError{http.StatusUnauthorized, "Access Denied"}
}
func UnauthorizedError() *UserError {
	return &UserError{http.StatusUnauthorized, "Credentials Are Incorrect"}
}
func UserAlreadyExists() *UserError {
	return &UserError{http.StatusFailedDependency, "Username Already Exists"}
}

func MalformedIdError() *UserError {
	return &UserError{http.StatusBadRequest, "Incorrect User ID"}
}

func IdNotFoundError() *UserError {
	return &UserError{http.StatusNotFound, "User ID not Found"}
}

func UnmarshallError() *UserError {
	return &UserError{http.StatusBadRequest, "Couldn't Parse User Body In Request"}
}

func InternalServerError() *UserError {
	return &UserError{http.StatusInternalServerError, "Internal Server Error"}
}
