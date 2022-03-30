package errors

import "net/http"

type OrderError struct {
	Status       int
	ErrorMessage string
}

func (productError *OrderError) Error() string {
	return productError.ErrorMessage
}

func MalformedIdError() *OrderError {
	return &OrderError{http.StatusBadRequest, "Malformed order id"}
}

func IdNotFoundError() *OrderError {
	return &OrderError{http.StatusNotFound, "order with given id not found"}
}

func OrderAlreadyPaidForError() *OrderError {
	return &OrderError{http.StatusBadRequest, "Order has already been paid for. Aborting this Payment."}
}

func OrderAlreadyDeliveredError() *OrderError {
	return &OrderError{http.StatusBadRequest, "Order has already been delivered. Current delivery aborted."}
}

func PaymentIncompleteError() *OrderError {
	return &OrderError{http.StatusBadRequest, "Order payment not done. Current delivery aborted."}
}

func InternalServerError() *OrderError {
	return &OrderError{http.StatusInternalServerError, "Internal Server Error"}
}

func UserNotFoundError() *OrderError {
	return &OrderError{http.StatusNotFound, "User not found"}
}
