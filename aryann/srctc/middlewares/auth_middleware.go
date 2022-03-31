package middlewares

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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
