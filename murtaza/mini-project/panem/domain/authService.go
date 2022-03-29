package domain

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

var secret string = "Alohomora"

type AuthService interface {
	AuthenticateUser(string, string) (string, error)
	ParseAuthToken(string) (int, Role, error)
}

type authService struct {
	userMongoRepository UserMongoRepository
}

func (as authService) AuthenticateUser(username, password string) (string, error) {
	user, err := as.userMongoRepository.FindUserByUsername(username)
	if err != nil {
		return "", err
	}
	if user.Password == password {
		return generateJWT(user.Id, user.Role)
	} else {
		return "", errors.New("invalid credentials")
	}
}

func (as authService) ParseAuthToken(tokenString string) (int, Role, error) {
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
	return -1, -1, err
}

func generateJWT(id int, role Role) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role.EnumIndex(),
	}).SignedString([]byte(secret))
}

func NewAuthService(userMongoRepository UserMongoRepository) AuthService {
	return &authService{userMongoRepository: userMongoRepository}
}
