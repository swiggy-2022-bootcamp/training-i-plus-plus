package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dhi13man/healthcare-app/users_service/configs"
	"github.com/dhi13man/healthcare-app/users_service/models"
	"github.com/dhi13man/healthcare-app/users_service/models/dtos"
	"github.com/dhi13man/healthcare-app/users_service/repositories"
	"github.com/dhi13man/healthcare-app/users_service/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

// Used for validating request data
var validate = validator.New()

// Constant Error message when data binding fails
const bindingRequestBodyFailedMessage string = "Binding Request Body Failed: "

// Constant Error message when data validation fails
const validatingRequestBodyFailedMessage string = "Validating Request Body Failed: "

// @Summary      Create a Client
// @Description  Create a Client in the Database using the data sent by them (REGISTER)
// @Tags         /users/clients
// @Accept       json
// @Produce      json
// @Param        clientDTO        body      models.Client  true  "User DTO"
// @Success      200              {object}  interface{}
// @Failure      400              {object}  dtos.HTTPError
// @Failure      404              {object}  dtos.HTTPError
// @Failure      500              {object}  dtos.HTTPError
// @Router       /users/clients [post]
func CreateClient(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var requestBody models.Client
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Error(bindingRequestBodyFailedMessage, err)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, err))
	} else if validationErr := validate.Struct(&requestBody); validationErr != nil { // Validation of required fields
		log.Error(validatingRequestBodyFailedMessage, validationErr)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, validationErr))
	} else {
		// Hash Client Password
		hashedPass, err := services.HashPassword(requestBody.Password)
		if err != nil {
			const errMsg string = "Error hashing password."
			log.Error(errMsg, err)
		}
		// Set Client data
		newUser := models.NewClient(requestBody.Name, requestBody.Email, hashedPass, []models.ClientSubscription{})
		out, err := repositories.CreateClient(newUser, ctx)
		if err != nil {
			const errMsg string = "Error inserting user."
			log.Error(errMsg, err)
			c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
		} else {
			c.JSON(http.StatusCreated, out)
		}
	}
}

// @Summary      Get a Client from Database.
// @Description  Get a Client from the Database using their email.
// @Tags         /users/clients
// @Accept       json
// @Produce      json
// @Param        email            path      string  true  "User Email"
// @Param        X-USER-PASSWORD  header    string         true  "User Password"
// @Success      200              {object}  models.Client
// @Failure      400              {object}  dtos.HTTPError
// @Failure      404              {object}  dtos.HTTPError
// @Failure      500              {object}  dtos.HTTPError
// @Router       /users/clients/{email} [get]
func GetClient(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	email := c.Param("email")
	if email == "" {
		const errMsg string = "Missing email in parameters."
		log.Error(errMsg)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, nil, errMsg))
	} else {
		user, err := repositories.GetClient(models.Client{Email: email}, ctx)
		if err != nil {
			const errMsg string = "Error finding user: "
			log.Error(errMsg, err)
			c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
		} else if !services.ComparePassword(user.Password, c.GetHeader("X-USER-PASSWORD")) {
			const errMsg string = "Invalid password sent in Header."
			log.Error(errMsg)
			c.JSON(http.StatusUnauthorized, dtos.NewError(http.StatusUnauthorized, nil, errMsg))
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

// @Summary      Updates Clients in the Database.
// @Description  Updates the Client in the Database using their email.
// @Tags         /users/clients
// @Accept       json
// @Produce      json
// @Param        clientDTO  body      models.Client  true  "User DTO"
// @Param        X-USER-PASSWORD  header    string  true  "User Password"
// @Success      200        {object}  interface{}
// @Failure      400        {object}  dtos.HTTPError
// @Failure      404        {object}  dtos.HTTPError
// @Failure      500        {object}  dtos.HTTPError
// @Router       /users/clients/{email} [put]
func UpdateClients(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var requestBody models.Client
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Error(bindingRequestBodyFailedMessage, err)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, err))
	} else if validationErr := validate.Struct(&requestBody); validationErr != nil { // Validation of required fields
		log.Error(validatingRequestBodyFailedMessage, validationErr)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, validationErr))
	} else {
		// Load current user data
		user, err := repositories.GetClient(models.Client{Email: requestBody.Email}, ctx)
		if err != nil {
			const errMsg string = "Error finding Client."
			log.Error(errMsg, err)
			c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, err, errMsg))
		} else if !services.ComparePassword(user.Password, c.GetHeader("X-USER-PASSWORD")) {
			const errMsg string = "Invalid password sent in Header."
			log.Error(errMsg)
			c.JSON(http.StatusUnauthorized, dtos.NewError(http.StatusUnauthorized, nil, errMsg))
		}  else {
			// Hash Client Password
			hashedPass, err := services.HashPassword(requestBody.Password)
			if err != nil {
				const errMsg string = "Error hashing password: "
				log.Error(errMsg, err)
			}
			user := models.NewClient(requestBody.Name, requestBody.Email, hashedPass, requestBody.Subscriptions)
			out, err := repositories.UpdateClient(user, ctx)
			if err != nil {
				const errMsg string = "Error updating user: "
				log.Error(errMsg, err)
				c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
			} else {
				c.JSON(http.StatusOK, out)
			}
		}
	}
}

// @Summary      Deletes Clients in the Database.
// @Description  Deletes the Clients in the Database using their email.
// @Tags         /users/clients
// @Accept       json
// @Produce      json
// @Param        email            path      string  true  "User Email"
// @Param        X-USER-PASSWORD  header    string  true  "User Password"
// @Success      200              {object}  int64
// @Failure      400              {object}  dtos.HTTPError
// @Failure      404              {object}  dtos.HTTPError
// @Failure      500              {object}  dtos.HTTPError
// @Router       /users/clients/{email} [delete]
func DeleteClients(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	email := c.Param("email")
	if email == "" {
		const errMsg string = "Client email is required."
		log.Error(errMsg)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, nil, errMsg))
		return
	}

	// Load current user data
	user, err := repositories.GetClient(models.Client{Email: email}, ctx)
	if err != nil {
		const errMsg string = "Error finding Client."
		log.Error(errMsg, err)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, err, errMsg))
	} else if !services.ComparePassword(user.Password, c.GetHeader("X-USER-PASSWORD")) {
		const errMsg string = "Invalid password sent in Header."
		log.Error(errMsg)
		c.JSON(http.StatusUnauthorized, dtos.NewError(http.StatusUnauthorized, nil, errMsg))
	}  else {
		out, err := repositories.DeleteClient(models.Client{Email: email}, ctx)
		if err != nil {
			const errMsg string = "Error deleting user."
			log.Error(errMsg, err)
			c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
		} else {
			c.JSON(http.StatusOK, out)
		}
	}
}

// @Summary      Make Disease Diagnosis by expert to the bookkeeping_service.
// @Description  Sends a Disease Diagnosis using Kafka to the bookkeeping_service
// @Tags         /users/experts/diagnose
// @Accept       json
// @Produce      json
// @Param        disease  body      models.Disease  true  "The Diagnosed Disease"
// @Success      200      {object}  models.Client
// @Failure      400      {object}  dtos.HTTPError
// @Failure      404      {object}  dtos.HTTPError
// @Failure      500      {object}  dtos.HTTPError
// @Router       /users/experts/diagnose [post]
func DiagnoseDisease(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var requestBody models.Disease
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Error(bindingRequestBodyFailedMessage, err)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, err))
	} else if validationErr := validate.Struct(&requestBody); validationErr != nil { // Validation of required fields
		log.Error(validatingRequestBodyFailedMessage, validationErr)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, validationErr))
	} else {
		// Serialize the request body
		requestBodyBytes, serialErr := json.Marshal(requestBody)
		if serialErr != nil {
			const errMsg string = "Error marshalling request body."
			log.Error(errMsg, serialErr)
			c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, serialErr, errMsg))
			return
		}
		// Send the disease diagnosis to the bookkeeping service
		err := services.Produce(string(requestBodyBytes), configs.KafkaDiagnosisTopic(), ctx)
		if err != nil {
			const errMsg string = "Error sending disease diagnosis to bookkeeping."
			log.Error(errMsg, err)
			c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Disease Diagnosis Sent"})
		}
	}
}