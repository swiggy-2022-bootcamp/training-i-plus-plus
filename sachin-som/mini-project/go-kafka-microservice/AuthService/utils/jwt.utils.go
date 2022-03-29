package utils

import "github.com/go-kafka-microservice/AuthService/models"

type JWTUtils interface {
	GenerateToken(*models.Credentials) string
	ValidateToken(string) error
	RefreshToken(string) (string, error)
}
type JWTUtilsImpl struct{}

func NewJWTUtils() *JWTUtilsImpl {
	return &JWTUtilsImpl{}
}

func (ju *JWTUtilsImpl) GenerateToken(credentials *models.Credentials) string {
	return ""
}

func (ju *JWTUtilsImpl) ValidateToken(token string) error {
	return nil
}

func (ju *JWTUtilsImpl) RefreshToken(token string) (refreshedToken string, err error) {
	return
}
