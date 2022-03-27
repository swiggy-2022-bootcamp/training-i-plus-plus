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
)

var registerCollection *mongo.Collection = database.GetCollection(database.DB, "signup")

var (
	mySigningKey = []byte("secret")
)

func GetJWT(group string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["group"] = group //group should be USER or ADMIN
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// claims["exp"] = time.Now().Add(time.Minute * 30000000).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var register models.SignUp
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&register); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newSignUp := models.SignUp{
			Username: register.Username,
			Email:    register.Email,
			TypeOf:   register.TypeOf,
			Password: register.Password,
		}

		if register.TypeOf == "ADMIN" || register.TypeOf == "USER" {
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

		//validate the request body
		if err := c.BindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error in binding", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&register); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error in validate register", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		if register.TypeOf == "ADMIN" {
			var admin_reg models.SignUp
			err := registerCollection.FindOne(ctx, bson.M{"username": register.Username}).Decode(&admin_reg)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in locating user", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			token, err := GetJWT("ADMIN")
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "success", Token: token})
			return
		} else if register.TypeOf == "USER" {
			var user_reg models.SignUp
			err := registerCollection.FindOne(ctx, bson.M{"username": register.Username}).Decode(&user_reg)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in locating user", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			token, err := GetJWT("USER")
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
		//normally Authorization the_token_xxx
		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			respondWithError(c, 401, "No bearer token")
			return
		}

		token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(("invalid Signing Method"))
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
