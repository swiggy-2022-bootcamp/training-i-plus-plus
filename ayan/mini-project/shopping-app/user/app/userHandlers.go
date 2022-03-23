package app

import (
	"encoding/json"
	"net/http"
	"user/domain"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	service domain.UserService
}

func (uh *UserHandlers) GetUserByEmail(c *gin.Context) {

	userEmail, ok := c.Params.Get("userEmail")

	if !ok {
		c.JSON(http.StatusBadRequest, nil)

	} else {
		user, err := uh.service.FindUserByEmail(userEmail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {

			data, err := json.Marshal(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			}
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

func (uh *UserHandlers) Register(c *gin.Context) {

	var newUser domain.User
	err := c.Bind(&newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)

	} else {
		regUser, err := uh.service.Register(newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {

			data, err := json.Marshal(regUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			}
			c.Data(http.StatusCreated, "application/json", data)
		}
	}
}

func (uh *UserHandlers) Login(c *gin.Context) {

	var user domain.User
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)

	} else {
		token, err := uh.service.Login(user.Email(), user.Password())
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {

			data, err := json.Marshal(token)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			}
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

func (uh *UserHandlers) UpdateUser(c *gin.Context) {

	var updatedUser domain.User
	err := c.Bind(&updatedUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)

	} else {
		user, err := uh.service.Register(updatedUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {

			data, err := json.Marshal(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			}
			c.Data(http.StatusCreated, "application/json", data)
		}
	}
}

func (uh *UserHandlers) HelloWorldHandler(c *gin.Context) {

	token := "Hello world"
	data, _ := json.Marshal(token)
	c.Data(http.StatusOK, "application/json", data)
}
