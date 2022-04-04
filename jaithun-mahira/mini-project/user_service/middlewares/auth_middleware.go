package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"user_service/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


func CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}


var (
	mySigningKey = []byte("fdsgdgjhgfj")
)

func GetJWT(role string, username string, emailId string, userId primitive.ObjectID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["emailId"] = emailId
	claims["userId"] = userId
	claims["authorized"] = true
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 500).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func GetKey() []byte {
	return mySigningKey
}

func Authenticate(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")

		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusInternalServerError, Message: "No token present"})
			return
		}

		token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(("invalid token"))
			}

			return GetKey(), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Internal Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Invalid token"})
			return
		}

		if token.Claims.(jwt.MapClaims)["role"] != role {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Unauthorized User"})
			return
		}
		c.Next()
	}
}