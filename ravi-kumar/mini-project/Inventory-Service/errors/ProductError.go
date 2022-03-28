package errors

type ProductError struct {
	Status       int
	ErrorMessage string
}

func (productError *ProductError) Error() string {
	return productError.ErrorMessage
}
