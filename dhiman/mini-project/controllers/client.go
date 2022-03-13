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

    client, exists := c.Get("client")
    if !exists {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Error Validating Request",  "error": "Client data not sent"})
		return
    }
    // Use the validator library to validate required fields
    if validationErr := validate.Struct(&client); validationErr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Error Validating Request",  "error": validationErr.Error()})
		return
    }

    newUser := models.NewClient(client.(models.Client).Email, client.(models.Client).Name)

    result, err := configs.UsersCollection.InsertOne(ctx, newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error inserting user", "error": err.Error()})
		return
    }
  
    c.JSON(http.StatusCreated, gin.H{"message": "success", "data": result})
}