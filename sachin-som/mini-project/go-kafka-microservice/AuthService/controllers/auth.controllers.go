package controllers

import (
	"context"

	pb "github.com/go-kafka-microservice/AuthProto"
	"github.com/go-kafka-microservice/AuthService/models"
	"github.com/go-kafka-microservice/AuthService/services"
)

type AuthControllers struct {
	pb.UnimplementedAuthServicesServer
	AuthServices services.AuthServices
}

func NewAuthControllers(authServices services.AuthServices) *AuthControllers {
	return &AuthControllers{
		AuthServices: authServices,
	}
}

func (ac AuthControllers) Authenticate(ctx context.Context, in *pb.Credentials) (*pb.Response, error) {
	credentials := models.Credentials{
		Email:    in.Email,
		Password: in.Password,
	}
	token, err := ac.AuthServices.Authenticate(&credentials)

	var res *pb.Response
	if err != nil {
		res = &pb.Response{
			Token: "",
			Err:   err.Error(),
		}
		return res, err
	}
	res = &pb.Response{
		Token: token,
		Err:   "",
	}
	return res, nil
}

func (ac AuthControllers) Authorize(ctx context.Context, in *pb.TokenRequest) (*pb.Response, error) {
	token, err := ac.AuthServices.Authorize(in.Token)
	if err != nil {
		return &pb.Response{
			Token: "",
			Err:   err.Error(),
		}, err
	}
	return &pb.Response{
		Token: token,
		Err:   "",
	}, nil
}
