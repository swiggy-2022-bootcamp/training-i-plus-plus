package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"srctc/middlewares"
	"srctc/models"
	"srctc/repository"
	"srctc/responses"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var registerRepo repository.AuthRepository
var adminRepo repository.AdminRepository
var userRepo repository.UserRepository

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

		// hash password
		hashedPassword, err := middlewares.HashMyPassword(register.Password)

		if err != nil {
			log.Panic(err)
			// panic(err)
		}

		register.Password = string(hashedPassword)

		newSignUp := models.SignUp{
			Username: register.Username,
			Email:    register.Email,
			TypeOf:   register.TypeOf,
			Password: register.Password,
		}

		if register.TypeOf == "admin" || register.TypeOf == "user" {

			result, err := registerRepo.Create(newSignUp)

			if err != nil {
				c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			if register.TypeOf == "admin" {

				_, err := adminRepo.Create(models.Admin{
					Name:  register.Username,
					Email: register.Email,
				})

				if err != nil {
					c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
				}

				// c.JSON(http.StatusOK, responses.LoginResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
				// return

			} else {

				_, err := userRepo.Create(models.User{
					Name:  register.Username,
					Email: register.Email,
				})

				if err != nil {
					c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
				}

				// c.JSON(http.StatusOK, responses.LoginResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
				// return
			}

			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
			return

		} else {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error not valid user group"})
			return
		}
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

		reg, err := registerRepo.Read(register.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in locating user", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		err = middlewares.ComparePassword(reg.Password, register.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in comparing password", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if register.TypeOf == "admin" {

			token, err := middlewares.GetJWT("admin")
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "you have successfully logged in", Token: token})
			return

		} else if register.TypeOf == "user" {

			token, err := middlewares.GetJWT("user")
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "you have successfully logged in", Token: token})
			return

		} else {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error not valid user group"})
			return
		}
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func IsAuthorized(group string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")

		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			respondWithError(c, 401, "No bearer token")
			return
		}

		token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Fatal("Unexpected signing method")
				return nil, fmt.Errorf(("invalid signing method"))
			}

			return middlewares.GetMySigingKey(), nil
		})
		if err != nil {
			respondWithError(c, 501, err.Error())
			return
		}
		if !token.Valid {
			respondWithError(c, 401, "Invalid token")
			return
		}

		if token.Claims.(jwt.MapClaims)["group"] != group {
			respondWithError(c, 401, "unauthorized user")
			return
		}

		c.Next()
	}
}
