package middlewares

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	mySigningKey = []byte("secret")
)

func GetJWT(group string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["group"] = group
	claims["exp"] = time.Now().Add(time.Hour * 500).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetMySigingKey() []byte {
	return mySigningKey
}

func HashMyPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
