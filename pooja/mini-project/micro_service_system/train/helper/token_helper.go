package helper

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// SignedDetails
type SignedDetails struct {
	UserName string `json:"username"`
	EmailId  string `json:"email_id"`
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func CreateToken(username, email_id string) (string, error) {
	claims := &SignedDetails{
		UserName: username,
		EmailId:  email_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	fmt.Print("create token claims", claims)
	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	fmt.Print("create token claims", claims)
	return signedToken, err
}

//ValidateToken validates the jwt token
func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return false, errors.New("Token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return false, errors.New("Session expired. Login again!!!")
	}

	return true, nil
}

func GetClaimsFromToken(tokenString string) (SignedDetails, error) {
	var claims SignedDetails
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})
	if err != nil {
		return claims, err
	}
	if token.Valid {
		return claims, nil
	}
	return claims, errors.New("invalid token")
}
