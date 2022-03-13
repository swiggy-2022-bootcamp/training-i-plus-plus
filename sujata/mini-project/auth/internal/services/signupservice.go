package services

import (
	"auth/internal/errors"
)

type SignupService interface {
	ValidateRequest() *errors.ServerError
	ProcessRequest() *errors.ServerError
}

type Signup struct {
	//config *util.RouterConfig
}

func (service *Signup) ValidateRequest() *errors.ServerError {
	return nil
}

func (service *Signup) ProcessRequest() *errors.ServerError {
	return nil
}
