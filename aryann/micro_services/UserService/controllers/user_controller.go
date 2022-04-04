package controllers

import (
	"UserService/models"
	"UserService/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUser godoc
// @Summary      Fetch A User
// @Description  Get User details by providing the userid
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        UserId 		body	string  true  "user unique id"
// @Success      200  {object}  models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/:userId [get]
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		user, err := userRepo.Read(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

// DelUser godoc
// @Summary      Delete A User
// @Description  Delete a User by providing the userid
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        UserId 		body	string  true  "user unique id"
// @Success      200  {object}  models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/:userId [delete]
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		_, err := userRepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		signup, err := userRepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!", "result": signup}},
		)
	}
}
