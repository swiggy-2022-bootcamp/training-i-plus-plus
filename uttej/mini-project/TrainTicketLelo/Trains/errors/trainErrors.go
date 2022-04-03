package errors

import "net/http"

type TrainError struct {
	Status       int
	ErrorMessage string
}

func AccessDenied() *TrainError {
	return &TrainError{http.StatusUnauthorized, "Access Denied"}
}

func (trainError *TrainError) Error() string {
	return trainError.ErrorMessage
}

func OutOfStockError() *TrainError {
	return &TrainError{http.StatusBadRequest, "No Tickets Available For The Train"}
}

func MalformedIdError() *TrainError {
	return &TrainError{http.StatusBadRequest, "Incorrect Train ID"}
}

func IdNotFoundError() *TrainError {
	return &TrainError{http.StatusNotFound, "Train ID not Found"}
}

func InternalServerError() *TrainError {
	return &TrainError{http.StatusInternalServerError, "Internal Server Error"}
}

func UnmarshallError() *TrainError {
	return &TrainError{http.StatusBadRequest, "Couldn't unmarshall product body in request"}
}
