package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/dhi13man/healthcare-app/users_service/configs"
	"github.com/dhi13man/healthcare-app/users_service/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateUser(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var requestBody models.Client
    err := c.BindJSON(&requestBody)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

    // Use the validator library to validate required fields
    if validationErr := validate.Struct(&requestBody); validationErr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
    }
    newUser := models.NewClient(requestBody.Email, requestBody.Name)

    result, err := configs.UsersCollection.InsertOne(ctx, newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error inserting user", "error": err.Error()})
		return
    }
  
    c.JSON(http.StatusCreated, gin.H{"message": "success", "data": result})
}

func GetUser(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    email := c.Param("email")
    if email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
        return
    }

    var user models.Client
    err := configs.UsersCollection.FindOne(ctx, models.Client{Email: email}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func UpdateUser() {
    
}