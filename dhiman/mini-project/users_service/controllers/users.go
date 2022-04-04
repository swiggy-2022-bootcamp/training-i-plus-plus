package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/dhi13man/healthcare-app/users_service/models"
	"github.com/dhi13man/healthcare-app/users_service/models/dtos"
	"github.com/dhi13man/healthcare-app/users_service/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

var validate = validator.New()

// @Summary      Create a Client
// @Description  Create a Client in the Database using the data sent by them (REGISTER)
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        email     body      string  true  "User Email"
// @Param        name      body      string  true  "User Name"
// @Param        password  body      string  true  "User Password"
// @Success      200       {object}  interface{}
// @Failure      400       {object}  dtos.HTTPError
// @Failure      404       {object}  dtos.HTTPError
// @Failure      500       {object}  dtos.HTTPError
// @Router       /users/clients [post]
func CreateClient(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var requestBody models.Client
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Error("Binding Request Body Failed: ", err)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, err))
	} else if validationErr := validate.Struct(&requestBody); validationErr != nil { // Validation of required fields
		log.Error("Validating Request Body Failed: ", validationErr)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, validationErr))
	} else {
		newUser := models.NewClient(requestBody.Email, requestBody.Name)
		out, err := repositories.CreateClient(*newUser, ctx)
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
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        email  path      string  true  "User Email"
// @Success      200    {object}  models.Client
// @Failure      400    {object}  dtos.HTTPError
// @Failure      404    {object}  dtos.HTTPError
// @Failure      500    {object}  dtos.HTTPError
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
		user, err := repositories.GetClient(models.Client{User: models.User{Email: email}}, ctx)
		if err != nil {
			const errMsg string = "Error finding user."
			log.Error(errMsg, err)
			c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

// @Summary      Updates Clients in the Database.
// @Description  Updates the Client in the Database using their email.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        email     body      string  true  "User Email"
// @Param        name      body      string  true  "User Name"
// @Param        password  body      string  true  "User Password"
// @Success      200       {object}  models.Client
// @Failure      400       {object}  dtos.HTTPError
// @Failure      404       {object}  dtos.HTTPError
// @Failure      500       {object}  dtos.HTTPError
// @Router       /users/clients/{email} [put]
func UpdateClients(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var requestBody models.Client
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Error("Binding Request Body Failed: ", err)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, err))
	} else if validationErr := validate.Struct(&requestBody); validationErr != nil { // Validation of required fields
		log.Error("Validating Request Body Failed: ", validationErr)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, validationErr))
	} else {
		user := models.NewClient(requestBody.Email, requestBody.Name)
		out, err := repositories.UpdateClient(*user, ctx)
		if err != nil {
			const errMsg string = "Error updating user."
			log.Error(errMsg, err)
			c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
		} else {
			c.JSON(http.StatusOK, out)
		}
	}
}

 // @Summary      Deletes Clients in the Database.
 // @Description  Deletes the Clients in the Database using their email.
 // @Tags         accounts
 // @Accept       json
 // @Produce      json
 // @Param        email     body      string  true  "User Email"
 // @Param        name      body      string  true  "User Name"
 // @Param        password  body      string  true  "User Password"
 // @Success      200       {object}  models.Client
 // @Failure      400       {object}  dtos.HTTPError
 // @Failure      404       {object}  dtos.HTTPError
 // @Failure      500       {object}  dtos.HTTPError
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
	 
	out, err := repositories.DeleteClient(models.Client{User: models.User{Email: email}}, ctx)
	if err != nil {
		const errMsg string = "Error deleting user."
		log.Error(errMsg, err)
		c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
	} else {
		c.JSON(http.StatusOK, out)
	}
}
