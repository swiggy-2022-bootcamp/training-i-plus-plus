package app

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"panem/domain"
	"panem/utils/errs"
	"panem/utils/logger"
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

type userResponseDTO struct {
	Message string `json:"message"`
}

func (h UserHandler) getAllUsers(c *gin.Context) {

}

// @Schemes
// @Description Fetches user details by userId
// @Tags users
// @Param        userId   path      int  true  "User ID"
// @Param        auth-token   header      string  true  "Authentication token"
// @Produce json
// @Success 200 {object} domain.User
// @Failure      403  {object} errs.AppError
// @Router /users/{userId} [get]
func (h UserHandler) getUserByUserId(c *gin.Context) {
	params := c.Params
	param, _ := params.Get("userId")
	userId, err := strconv.ParseInt(param, 10, 0)
	if err != nil {
		logger.Error("Mandatory field userId misisng in request params:")
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Mandatory field userId missing in request params"))
		c.Abort()
		return
	}
	user, err2 := h.userService.GetMongoUserByUserId(int(userId))
	if err2 != nil {
		c.JSON(err2.Code, err2)
		c.Abort()
		return
	} else {
		data, _ := user.MarshalJSON()
		logger.Info(fmt.Sprintf("Sending user details for userId: %d", userId))
		c.Data(http.StatusOK, "application/json", data)
	}
}

// @Schemes
// @Description Creates a user upon signup
// @Tags users
// @Produce json
// @Accept json
// @Param        user  body      userDTO  true  "user signup"
// @Success 200 {object} domain.User
// @Router /signup [post]
func (h UserHandler) createUser(c *gin.Context) {
	var newUser userDTO
	err := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		customErr := errs.NewUnexpectedError("Unable to decode user payload")
		logger.Error(customErr.Message, zap.Error(err))
		c.JSON(http.StatusInternalServerError, customErr)
	} else {
		role, err := domain.GetEnumByIndex(newUser.Role)
		if err != nil {
			logger.Error(err.Message, zap.Int("role", newUser.Role), zap.Error(errors.New(err.Message)))
			c.JSON(http.StatusInternalServerError, err)
		}
		user, err := h.userService.CreateUserInMongo(newUser.FirstName, newUser.LastName, newUser.Username, newUser.Phone, newUser.Email, newUser.Password, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			data, _ := user.MarshalJSON()
			logger.Info(fmt.Sprintf(fmt.Sprintf("user created in DB with Id: %d", user.Id)))
			c.Data(http.StatusCreated, "application/json", data)
		}
	}
}

// @Schemes
// @Description Deletes user by userId
// @Tags users
// @Param        userId   path      int  true  "User ID"
// @Param        auth-token   header      string  true  "Authentication token"
// @Produce json
// @Success 200 {object} userResponseDTO
// @Failure      500  {object} errs.AppError
// @Router /users/{userId} [delete]
func (h UserHandler) deleteUser(c *gin.Context) {
	params := c.Params
	val, err := params.Get("userId")
	userId, _ := strconv.Atoi(val)

	if err == false {
		logger.Error("Mandatory field user id missing in DELETE request")
		c.JSON(http.StatusBadRequest, gin.H{"message": "No user id given"})
	} else {
		err := h.userService.DeleteUserByUserId(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			logger.Error("User with userId deleted successfully", zap.Int("userId", userId))
			userResponse := userResponseDTO{
				Message: fmt.Sprintf("userId: %d deleted successfully", userId),
			}
			c.JSON(http.StatusOK, userResponse)
		}
	}
}

// @Schemes
// @Description Updates user by userId
// @Tags users
// @Param        userId   path      int  true  "User ID"
// @Param        auth-token   header      string  true  "Authentication token"
// @Param        user details   body      userDTO true  "User details"
// @Produce json
// @Success 200 {object} domain.User
// @Failure      500  {object} errs.AppError
// @Router /users/{userId} [put]
func (h UserHandler) updateUser(c *gin.Context) {
	params := c.Params
	userId, err := params.Get("userId")

	if err == false {
		logger.Error("Mandatory field userId missing in request")
		c.JSON(http.StatusBadRequest, "userId missing in request")
	}

	var newUser userDTO
	err2 := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, err2)
	} else {
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, err2)
		}

		userId, _ := strconv.ParseInt(userId, 10, 0)
		user := domain.NewUser(newUser.FirstName, newUser.LastName, newUser.Username, newUser.Phone, newUser.Email, newUser.Password, domain.Role(newUser.Role))
		user.Id = int(userId)
		user, err := h.userService.UpdateUser(*user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err2)
		} else {
			data, _ := user.MarshalJSON()
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

func (h UserHandler) demoHandlerFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}
