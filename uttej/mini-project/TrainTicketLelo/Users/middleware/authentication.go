package middleware

import (
	models "Users/model"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const mySigningKey = "Secret"

func GenerateJWT(userId string, userRole models.Role) (string, error) {
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

func ValidateToken(encodedToken string) (userId string, role models.Role, err error) {
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
		if claimedRole := int(claims["role"].(float64)); models.IsValidRole(claimedRole) {
			role = models.Role(claimedRole)
			userId = claims["id"].(string)
			return
		}
	}
	return "", -1, err
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
