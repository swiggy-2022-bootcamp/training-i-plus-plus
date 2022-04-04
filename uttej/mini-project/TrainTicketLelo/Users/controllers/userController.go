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

// Login godoc
// @Summary      Login
// @Description  Sign in a user/admin by providing username, password
// @Tags         User
// @Accept       json
// @Produce      json
// @Param		Login	body	models.Login	true	"username and password"
// @Success      200  {string}  jwtTokenResponse
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /users/login [post]
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

// Signup godoc
// @Summary      Signup
// @Description  Register a user/admin by providing fullname, username, email, password and role
// @Tags         User
// @Accept       json
// @Produce      json
// @Param		Signup	body	models.User	true	"id will be populated automatically"
// @Success      200  {string}  responseBody
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /signup [post]
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

// GetAllUsers godoc
// @Summary      Fetch All Users
// @Description  Get All Users details
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /users [get]
func GetAllUsers(c *gin.Context) {
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	if models.Role(acessorUserRole) != models.Admin {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary      Fetch A User
// @Description  Get User details by providing the userid
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        UserId 		body	string  true  "unique user id"
// @Success      200  {object}  models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/:userId [get]
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

// UpdateUser godoc
// @Summary      Update A User
// @Description  Update a User's details by providing the userid
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        UserId 		body	string  true  "unique user id"
// @Success      200  {object}  models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/:userId [put]
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

//DeleteUser godoc
// @Summary      Delete A User
// @Description  Delete a User by providing the userid
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        UserId 		body	string  true  "unique user id"
// @Success      200  {string}  successMessage
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/:userId [delete]
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
