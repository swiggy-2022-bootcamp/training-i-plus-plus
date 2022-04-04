package dtos


// Create a new Error.
func NewError(statusCode int, err error, details ...string) *HTTPError {
	return &HTTPError{
		Code:    statusCode,
		Message: err.Error(),
		Details: details,
	}
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
	Details []string `json:"details" example:"invalid email"`
}