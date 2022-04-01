package middleware

import (
	mockdata "User-Service/model"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

const mySigningKey = "secret&$key"

func GenerateJWT(userId string, userRole mockdata.Role) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userId,
		"role": userRole,
	})

	tokenString, err := token.SignedString([]byte(mySigningKey))

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(encodedToken string) (userId string, role mockdata.Role, err error) {
	parsedToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(mySigningKey), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		for key, value := range claims {
			fmt.Printf("%s\t%v\n", key, value)
		}
		if claimedRole := int(claims["role"].(float64)); mockdata.IsValidRole(claimedRole) {
			role = mockdata.Role(claimedRole)
			userId = claims["id"].(string)
			return
		}
	}
	return "", -1, err
}
