package services

import (
	"github.com/go-kafka-microservice/AuthService/models"
	"github.com/go-kafka-microservice/AuthService/utils"
)

type AuthServicesImpl struct {
	JWTUtils utils.JWTUtils
}

func NewAuthServiceImpl(jwtUtils utils.JWTUtils) *AuthServicesImpl {
	return &AuthServicesImpl{
		JWTUtils: jwtUtils,
	}
}

func (as *AuthServicesImpl) Authenticate(credentials *models.Credentials) (string, error) {
	return "", nil
}

// Authorize takes a token as string
// validates the string
// returns nil if case of valid credentials
// else returns error with proper reason
func (as *AuthServicesImpl) Authorize(string) error {
	return nil
}
