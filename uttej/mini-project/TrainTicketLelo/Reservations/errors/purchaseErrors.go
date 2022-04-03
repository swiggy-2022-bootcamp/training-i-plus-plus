package errors

import "net/http"

type PurchaseError struct {
	Status       int
	ErrorMessage string
}

func (purchaseError *PurchaseError) Error() string {
	return purchaseError.ErrorMessage
}

func AccessDenied() *PurchaseError {
	return &PurchaseError{http.StatusUnauthorized, "Access Denied"}
}

func MalformedIdError() *PurchaseError {
	return &PurchaseError{http.StatusBadRequest, "Incorrect User ID"}
}

func IdNotFoundError() *PurchaseError {
	return &PurchaseError{http.StatusNotFound, "Ticket ID not Found"}
}

func PaymentAlreadyDoneError() *PurchaseError {
	return &PurchaseError{http.StatusBadRequest, "Payment for the Ticket is already done, current payment will not proceed"}
}
func TicketAlreadyCancelledError() *PurchaseError {
	return &PurchaseError{http.StatusBadRequest, "Ticket Already Cancelled"}
}

func InternalServerError() *PurchaseError {
	return &PurchaseError{http.StatusInternalServerError, "Internal Server Error"}
}

func UserNotFoundError() *PurchaseError {
	return &PurchaseError{http.StatusNotFound, "User not Found"}
}
