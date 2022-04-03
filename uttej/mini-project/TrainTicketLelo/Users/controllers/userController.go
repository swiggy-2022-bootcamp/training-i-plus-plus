package controller

import (
	errors "Users/errors"
	models "Users/model"
	service "Users/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LogInUser(c *gin.Context) {
	var Login models.Login
	json.NewDecoder(c.Request.Body).Decode(&Login)
	jwtToken, error := service.LogInUser(Login)

	if error != nil {
		userError, ok := error.(*errors.UserError)
		if ok {
			c.JSON(userError.Status, userError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Let The User Login")
			return
		}
	}
	type LoginResponse struct {
		JwtToken string
	}
	var jwtTokenResponse LoginResponse = LoginResponse{jwtToken}
	c.JSON(http.StatusOK, jwtTokenResponse)
}

func CreateUser(c *gin.Context) {
	insertedId, err := service.CreateUser(&c.Request.Body)
	if err != nil {
		c.JSON(http.StatusFailedDependency, err)
		return
	}

	type ResponseBody struct {
		InsertId string
	}

	var responseBody ResponseBody = ResponseBody{insertedId}

	c.JSON(http.StatusOK, responseBody)
}

func GetAllUsers(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	if models.Role(acessorUserRole) != models.Admin {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	acessorUserId := c.Param("acessorUserId")
	userId := c.Param("userId")

	if !(models.Role(acessorUserRole) == models.Admin || acessorUserId == userId) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	userRetrieved, error := service.GetUserById(userId)

	if error != nil {
		userError, ok := error.(*errors.UserError)
		if ok {
			c.JSON(userError.Status, userError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Get User")
			return
		}
	}
	c.JSON(http.StatusOK, userRetrieved)
}

func UpdateUserById(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	acessorUserId := c.Param("acessorUserId")
	userId := c.Param("userId")

	if !(models.Role(acessorUserRole) == models.Admin || acessorUserId == userId) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	userRetrieved, error := service.UpdateUserById(userId, &c.Request.Body)

	if error != nil {
		userError, ok := error.(*errors.UserError)
		if ok {
			c.JSON(userError.Status, userError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Update User")
			return
		}
	}
	c.JSON(http.StatusOK, userRetrieved)

}

func DeleteUserbyId(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	acessorUserId := c.Param("acessorUserId")
	userId := c.Param("userId")

	if !(models.Role(acessorUserRole) == models.Admin || acessorUserId == userId) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	successMessage, error := service.DeleteUserbyId(userId)

	if error != nil {
		userError, ok := error.(*errors.UserError)
		if ok {
			c.JSON(userError.Status, userError.ErrorMessage)
			return
		} else {
			fmt.Println("Couldn't Delete User")
			return
		}
	}
	c.JSON(http.StatusOK, successMessage)
}
