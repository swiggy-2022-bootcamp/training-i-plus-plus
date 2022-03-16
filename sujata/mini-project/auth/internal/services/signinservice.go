package services

import (
	"auth/internal/errors"
)

type SigninService interface {
	ValidateRequest() *errors.ServerError
	ProcessRequest() *errors.ServerError
}

type Signin struct {
	//config *util.RouterConfig
}

func (service *Signin) ValidateRequest() *errors.ServerError {
	return nil
}

func (service *Signin) ProcessRequest() *errors.ServerError {
	return nil
}
