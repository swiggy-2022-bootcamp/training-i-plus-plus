package utils

import (
	"errors"
	"os"
	"time"

	"github.com/go-kafka-microservice/AuthService/models"
	"github.com/golang-jwt/jwt"
)

var (
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
)

type JWTUtils interface {
	GenerateToken(*models.Credentials, time.Time) (string, error)
	ValidateToken(string, time.Time) (string, error)
	RefreshToken(*models.Claims, time.Time) (string, error)
}
type JWTUtilsImpl struct{}

func NewJWTUtils() *JWTUtilsImpl {
	return &JWTUtilsImpl{}
}

func (ju *JWTUtilsImpl) GenerateToken(credentials *models.Credentials, exp time.Time) (string, error) {
	// Create New Claims Instance
	claims := &models.Claims{
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func (ju *JWTUtilsImpl) ValidateToken(tokenStr string, exp time.Time) (string, error) {
	claims := models.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}
	if !tkn.Valid {
		return "", errors.New("")
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return ju.RefreshToken(&claims, exp)
	}
	return "", nil
}

func (ju *JWTUtilsImpl) RefreshToken(claims *models.Claims, exp time.Time) (string, error) {
	claims.ExpiresAt = exp.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshedToken, err := token.SignedString(jwtKey)
	return refreshedToken, err
}
