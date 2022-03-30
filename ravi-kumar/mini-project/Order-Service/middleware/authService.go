package middleware

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type authCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

const mySigningKey = "secret&$key"

func GenerateJWT(userName string) (string, error) {
	claims := &authCustomClaims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    "Admin",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(mySigningKey))

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(mySigningKey), nil
	})
}
