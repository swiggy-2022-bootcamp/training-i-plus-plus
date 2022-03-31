package controller

import (
	"context"
	"fmt"
	"net/http"
	"ticket_reservation_system/helper"
	"ticket_reservation_system/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()
		var user model.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": validationErr.Error()})
			return
		}
		var userFromDB model.User
		count, err := userCollection.CountDocuments(ctx, bson.M{"username": userFromDB.UserName})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error! occured while checking for email"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists."})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		newUser := model.User{
			ID:       primitive.NewObjectID(),
			UserName: user.UserName,
			EmailId:  user.EmailId,
			Password: string(hashedPassword),
		}
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "success", "data": result})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var login model.LoginDTO
		var user model.User
		defer cancel()

		if err := c.BindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		if err := userCollection.FindOne(ctx, bson.M{"username": login.Username}).Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "user not registered", "error": err.Error()})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "incorrect password", "details": err.Error()})
			return
		}
		//token, refreshToken, _ := helper.GenerateAllTokens(user.UserName)

		// helper.UpdateAllTokens(token, refreshToken, user.UserName)

		// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": gin.H{
		// 	"token":         user.Token,
		// 	"refresh_token": user.Refresh_token}})
		fmt.Print("before token")
		token, err := helper.CreateToken(user.UserName, user.EmailId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Print("after token")
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
