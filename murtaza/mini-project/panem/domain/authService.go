package domain

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"panem/utils/errs"

	"github.com/golang-jwt/jwt"
)

var secret = "Alohomora"

type AuthService interface {
	AuthenticateUser(string, string) (string, *errs.AppError)
	ParseAuthToken(string) (int, Role, *errs.AppError)
}

type authService struct {
	userMongoRepository UserMongoRepository
}

func (as authService) AuthenticateUser(username, password string) (string, *errs.AppError) {
	user, err := as.userMongoRepository.FindUserByUsername(username)
	if err != nil {
		return "", err
	}
	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err2 == nil {
		return generateJWT(user.Id, user.Role)
	} else {
		return "", errs.NewAuthenticationError("invalid credentials")
	}
}

func (as authService) ParseAuthToken(tokenString string) (int, Role, *errs.AppError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if role, err := GetEnumByIndex(int(claims["role"].(float64))); err == nil {
			return int(claims["id"].(float64)), role, nil
		}
	}
	return -1, -1, errs.NewUnexpectedError(err.Error())
}

func generateJWT(id int, role Role) (string, *errs.AppError) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role.EnumIndex(),
	}).SignedString([]byte(secret))

	if err != nil {
		return "", errs.NewValidationError(err.Error())
	}
	return token, nil
}

func NewAuthService(userMongoRepository UserMongoRepository) AuthService {
	return &authService{userMongoRepository: userMongoRepository}
}
