package dtos


// Create a new Error.
func NewError(statusCode int, err error, details ...string) HTTPError {
	var errMessage string = "none"
	if err != nil {
		errMessage = err.Error()
	}
	return HTTPError{
		Code:    statusCode,
		Message: errMessage,
		Details: details,
	}
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
	Details []string `json:"details" example:"invalid email"`
}