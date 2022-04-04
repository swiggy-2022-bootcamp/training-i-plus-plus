package controllers

import (
	"auth/configs"
	"auth/models"
	"auth/responses"
	"context"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
//var sellerCollection *mongo.Collection = configs.GetCollection(configs.DB, "seller")
var validate = validator.New()

var (
	mySigningKey = []byte(configs.EnvSecretKeyJWT())
)

func GetJWT(role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()


		//validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.AuthResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.AuthResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }
		user.UserId = primitive.NewObjectID()

		hashedPassword, err := configs.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
		}

		user.Password = string(hashedPassword)

		result, err := userCollection.InsertOne(ctx, user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AuthResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.AuthResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var userLoginInfo models.Login
		defer cancel()

		if err := c.BindJSON(&userLoginInfo); err != nil {
			c.JSON(http.StatusBadRequest, responses.AuthResponse{Status: http.StatusBadRequest, Message: "error in binding", Data: map[string]interface{}{"data": err.Error()}})
			configs.ErrorLogger.Println(err)
			return
		}

		if validationErr := validate.Struct(&userLoginInfo); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AuthResponse{Status: http.StatusBadRequest, Message: "Login Info Missing", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"email": userLoginInfo.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.AuthResponse{Status: http.StatusBadRequest, Message: "Invalid Email Id", Data: map[string]interface{}{"data": err.Error()}})
			configs.ErrorLogger.Println(err)
			return
		}

		match := configs.CheckPasswordHash(userLoginInfo.Password, user.Password)

		if !match {
			c.JSON(http.StatusUnauthorized, responses.AuthResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data":"Wrong password, Access denied"}})
			configs.ErrorLogger.Println("data:Wrong password, Access denied")
			return
		}
		
		token, err := GetJWT(user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AuthResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
			configs.ErrorLogger.Println(err)
			return
		}
		c.JSON(http.StatusCreated, responses.AuthResponse{Status: http.StatusCreated, Message: "success",Data: map[string]interface{}{"data": user}, JWT: token})
		configs.InfoLogger.Println(user)
	}
}
