package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"panem/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService domain.UserService
}

type userDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
}

func (h UserHandler) getAllUsers(c *gin.Context) {

}

func (h UserHandler) getUserByUserId(c *gin.Context) {
	params := c.Params
	userId, err := params.Get("userId")

	if err == false {
		c.JSON(http.StatusNotFound, err)
	} else {
		userId, _ := strconv.ParseInt(userId, 10, 0)
		user, err := h.userService.GetMongoUserByUserId(int(userId))
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			data, _ := user.MarshalJSON()
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

func (h UserHandler) createUser(c *gin.Context) {
	var newUser userDTO
	err := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		role, err := domain.GetEnumByIndex(newUser.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		user, err := h.userService.CreateUserInMongo(newUser.FirstName, newUser.LastName, newUser.Username, newUser.Phone, newUser.Email, newUser.Password, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			data, _ := user.MarshalJSON()
			c.Data(http.StatusCreated, "application/json", data)
		}
	}
}

func (h UserHandler) deleteUser(c *gin.Context) {
	params := c.Params
	val, err := params.Get("userId")
	userId, _ := strconv.Atoi(val)

	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No user id given"})
	} else {
		err := h.userService.DeleteUserByUserId(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("userId: %d deleted successfully", userId)})
		}
	}
}

func (h UserHandler) demoHandlerFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}
