package controllers

import (
	"context"
	"net/http"
	"tejas/models"
	"tejas/services"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const requestTimeout = 10 * time.Second

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var user models.User
		c.BindJSON(&user)
		fillUserDefaults(&user)
		hashPassword := services.HashPassword(user.Password)
		user.Password = hashPassword

		result, err := models.UserCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := models.UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User with this email doesn't exists "})
			return
		}

		isValidPassword, msg := services.VerifyPassword(user.Password, foundUser.Password)

		if !isValidPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}
		token, err := services.CreateToken(foundUser.Id, foundUser.Email, foundUser.Name, foundUser.IsAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func fillUserDefaults(user *models.User) {
	user.IsAdmin = false
	user.Id = primitive.NewObjectID()
}
