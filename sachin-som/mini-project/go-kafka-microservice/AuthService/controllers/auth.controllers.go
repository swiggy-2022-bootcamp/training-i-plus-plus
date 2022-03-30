package controllers

import (
	"context"

	"github.com/go-kafka-microservice/AuthService/models"
	"github.com/go-kafka-microservice/AuthService/services"
)

type AuthControllers struct {
	UnimplementedAuthServicesServer
	AuthServices services.AuthServices
}

func NewAuthControllers(authServices services.AuthServices) *AuthControllers {
	return &AuthControllers{
		AuthServices: authServices,
	}
}

func (ac AuthControllers) Authenticate(ctx context.Context, in *Credentials) (*Response, error) {
	credentials := &models.Credentials{
		Email:    in.Email,
		Password: in.Password,
	}
	token, err := ac.AuthServices.Authenticate(credentials)
	res := &Response{
		Token: token,
		Err:   err.Error(),
	}
	return res, err
}

func (ac AuthControllers) Authorize(ctx context.Context, in *TokenRequest) (*Response, error) {
	token, err := ac.AuthServices.Authorize(in.Token)
	return &Response{
		Token: token,
		Err:   err.Error(),
	}, err
}
