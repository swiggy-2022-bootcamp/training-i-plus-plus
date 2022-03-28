package controllers

import (
	"context"
	"fmt"
	"net/http"
	"srctc/database"
	"srctc/models"
	"srctc/responses"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var registerCollection *mongo.Collection = database.GetCollection(database.DB, "signup")

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

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var register models.SignUp
		defer cancel()

		if err := c.BindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := avalidate.Struct(&register); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)

		if err != nil {
			panic(err)
		}

		register.Password = string(hashedPassword)

		newSignUp := models.SignUp{
			Username: register.Username,
			Email:    register.Email,
			TypeOf:   register.TypeOf,
			Password: register.Password,
		}

		if register.TypeOf == "admin" || register.TypeOf == "user" {
			result, err := registerCollection.InsertOne(ctx, newSignUp)
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
			return
		} else {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error not valid group"})
			return
		}

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var register models.SignUp
		defer cancel()

		if err := c.BindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error in binding", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := avalidate.Struct(&register); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error in validate register", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		if register.TypeOf == "admin" {
			var admin_reg models.SignUp
			err := registerCollection.FindOne(ctx, bson.M{"username": register.Username}).Decode(&admin_reg)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in locating user", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			token, err := GetJWT("admin")
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "success", Token: token})
			return
		} else if register.TypeOf == "user" {
			var user_reg models.SignUp
			err := registerCollection.FindOne(ctx, bson.M{"username": register.Username}).Decode(&user_reg)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in locating user", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			token, err := GetJWT("user")
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "success", Token: token})
			return
		} else {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error not valid group"})
			return
		}

	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")

		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			respondWithError(c, 401, "No bearer token")
			return
		}

		token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(("invalid signing method"))
			}

			return mySigningKey, nil
		})
		if err != nil {
			respondWithError(c, 501, err.Error())
			return
		}
		if !token.Valid {
			respondWithError(c, 401, "Invalid token")
			return
		}
		c.Next()
	}
}
