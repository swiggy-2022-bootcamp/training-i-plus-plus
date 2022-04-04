package utils

import (
	"fmt"
	"user/utils/errs"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secret string = "Supercalifragilisticexpialidocious"

func HashPassword(password string) (string, *errs.AppError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return string(bytes), errs.NewUnexpectedError("Unexpected error in password hashing")
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseAuthToken(tokenString string) (string, string, *errs.AppError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		return "", "", errs.NewAuthenticationError("unexpected signing method")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims["email"].(string), claims["role"].(string), nil
	}
	return "", "", errs.NewAuthenticationError("Invalid token")
}

func GenerateJWT(email string, role string) (string, *errs.AppError) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
	}).SignedString([]byte(secret))
	if err != nil {
		return token, errs.NewUnexpectedError("Error generating token")
	}
	return token, nil
}
