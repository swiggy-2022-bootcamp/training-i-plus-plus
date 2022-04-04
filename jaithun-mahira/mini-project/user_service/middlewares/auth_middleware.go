package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"user_service/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	zap.L().Info("Inside HashPassword Method")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("Error Hashing Password"+err.Error())
		return "", err
	}
	zap.L().Info("Hashed Password Successfully")
	return string(hashedPassword), nil
}


func CheckPassword(hashedPassword, password string) error {
	zap.L().Info("Inside CheckPassword Method")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		zap.L().Error("Error Comparing Password"+err.Error())
		return err
	}
	return nil
}


var (
	mySigningKey = []byte("fdsgdgjhgfj")
)

func GetJWT(role string, username string, emailId string, userId primitive.ObjectID) (string, error) {
	zap.L().Info("Inside GetJWT Method")
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
		zap.L().Error("Error Generating Token"+err.Error())
		return "", err
	}
	zap.L().Info("Token generated Successfully")
	return tokenString, nil
}


func GetKey() []byte {
	return mySigningKey
}

func Authenticate(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside Authentication Fucntion")
		bearToken := c.Request.Header.Get("Authorization")

		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			zap.L().Error("No token present")
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusInternalServerError, Message: "No token present"})
			return
		}

		token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				zap.L().Error("invalid token")
				return nil, fmt.Errorf(("invalid token"))
			}

			return GetKey(), nil
		})
		if err != nil {
			zap.L().Error("Internal Error while validating Token")
			c.AbortWithStatusJSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Internal Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if !token.Valid {
			zap.L().Error("Invalid token")
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Invalid token"})
			return
		}

		if token.Claims.(jwt.MapClaims)["role"] != role {
			zap.L().Error("Unauthorized User")
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Unauthorized User"})
			return
		}
		zap.L().Info("User Authenticated Successfully")
		c.Next()
	}
}