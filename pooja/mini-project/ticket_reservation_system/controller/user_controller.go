package controller

import (
	"context"
	"net/http"
	"ticket_reservation_system/config"
	"ticket_reservation_system/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate *validator.Validate = validator.New()

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user model.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": validationErr.Error()})
			return
		}

		newUser := model.User{
			UserId:   user.UserId,
			UserName: user.UserName,
			EmailId:  user.EmailId,
			Password: user.Password,
		}
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "success", "data": result})
	}
}

func GetUserByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		username := c.Param("username")
		var user model.User
		defer cancel()
		if err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusCreated, "message": "success", "data": user})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []model.User
		defer cancel()
		results, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser model.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
				return
			}
			users = append(users, singleUser)
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": users})
	}
}

func UpdateUserPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		username := c.Param("username")
		var user model.User
		defer cancel()
		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": validationErr.Error()})
			return
		}
		update := bson.M{"user_id": user.UserId, "username": user.UserName, "email_id": user.EmailId, "password": user.Password}
		result, err := userCollection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		//get updated user details
		var updatedUser model.User
		if result.MatchedCount == 1 {
			if err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&updatedUser); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": updatedUser})
	}
}

func DeleteUserByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		username := c.Param("username")
		defer cancel()
		result, err := userCollection.DeleteOne(ctx, bson.M{"username": username})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": "user successfully deleted"})
	}
}
