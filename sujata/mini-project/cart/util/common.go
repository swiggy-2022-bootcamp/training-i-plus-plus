package util

import (
	"cart/config"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type RouterConfig struct {
	WebServerConfig *config.WebServerConfig
}

func ExtractDetailsFromToken(req *http.Request) (string, string) {
	token, _ := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, keyLookupFunc)

	// extract user id from the JWT token
	claims := token.Claims.(jwt.MapClaims)
	userInfo := claims["CustomUserInfo"].(map[string]interface{})

	fmt.Println("Userinfoo", userInfo)
	email := userInfo["Email"].(string)
	role := userInfo["role"].(string)

	return role, email
}

// keyLookupFunc returns the public key for JWT authentication
func keyLookupFunc(*jwt.Token) (interface{}, error) {
	return VerifyKey, nil
}
