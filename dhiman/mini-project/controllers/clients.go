package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/dhi13man/healthcare-app/configs"
	"github.com/dhi13man/healthcare-app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateClient(c *gin.Context) {
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

func GetClient(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
        return
    }

    var user models.Client
    err := configs.UsersCollection.FindOne(ctx, models.Client{ID: id}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
