package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/dhi13man/healthcare-app/users_service/repositories"
	"github.com/dhi13man/healthcare-app/users_service/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

var validate = validator.New()

func CreateClient(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var requestBody models.Client
    err := c.BindJSON(&requestBody)
    if err != nil {
        log.Error("Binding Request Body Failed: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

    // Use the validator library to validate required fields
    if validationErr := validate.Struct(&requestBody); validationErr != nil {
        log.Error("Validating Request Body Failed: ", validationErr)
        c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
    }

    newUser := models.NewClient(requestBody.Email, requestBody.Name)
    out, err := repositories.CreateClient(*newUser, ctx)
    if err != nil {
        const errMsg string = "Error inserting user."
        log.Error(errMsg, err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": errMsg, "error": err.Error()})
    } else {
        c.JSON(http.StatusCreated, gin.H{"message": "success", "data": out})
    }
}

func GetClient(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    email := c.Param("email")
    if email == "" {
        const errMsg string = "Missing email in parameters."
        log.Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
        return
    }

    user, err := repositories.GetClient(models.Client{User: models.User{Email: email}}, ctx)
    if err != nil {
        const errMsg string = "Error finding user."
        log.Error(errMsg, err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": errMsg, "error": err.Error()})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
    }
}

func UpdateClients(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var requestBody models.Client
    err := c.BindJSON(&requestBody)
    if err != nil {
        log.Error("Binding Request Body Failed: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Use the validator library to validate required fields
    if validationErr := validate.Struct(&requestBody); validationErr != nil {
        log.Error("Validating Request Body Failed: ", validationErr)
        c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
        return
    }

    user := models.NewClient(requestBody.Email, requestBody.Name)
    out, err := repositories.UpdateClient(*user, ctx)
    if err != nil {
        const errMsg string = "Error updating user."
        log.Error(errMsg, err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": errMsg, "error": err.Error()})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "success", "data": out})
    }
}
