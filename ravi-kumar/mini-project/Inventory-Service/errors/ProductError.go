package errors

import "net/http"

type ProductError struct {
	Status       int
	ErrorMessage string
}

func (productError *ProductError) Error() string {
	return productError.ErrorMessage
}

func OutOfStockError() *ProductError {
	return &ProductError{http.StatusBadRequest, "product out of stock"}
}

func MalformedIdError() *ProductError {
	return &ProductError{http.StatusBadRequest, "Malformed product id"}
}

func IdNotFoundError() *ProductError {
	return &ProductError{http.StatusNotFound, "product with given id not found"}
}

func InternalServerError() *ProductError {
	return &ProductError{http.StatusInternalServerError, "Internal Server Error"}
}

func UnmarshallError() *ProductError {
	return &ProductError{http.StatusBadRequest, "Couldn't unmarshall product body in request"}
}
