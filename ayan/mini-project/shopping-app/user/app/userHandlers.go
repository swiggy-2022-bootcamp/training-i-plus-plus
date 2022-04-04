package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/domain"
	"user/utils"
	"user/utils/errs"
	"user/utils/logger"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	service domain.UserService
}

type UserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Zipcode  int32  `json:"zipcode"`
	MobileNo string `json:"mobile_no"`
	Role     string `json:"role"`
}

type UserResponseDTO struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthDTO struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

// @Schemes
// @Description Fetches user details by email
// @Tags users
// @Param        userEmail   path      int  true  "Email"
// @Param        token   header      string  true  "Auth Token"
// @Produce json
// @Success 200 {object} domain.User
// @Failure      403  {object} errs.AppError
// @Router /users/{userEmail} [get]
func (uh *UserHandlers) GetUserByEmail(c *gin.Context) {

	userEmail, ok := c.Params.Get("userEmail")

	if !ok {
		logger.Error("User email not present in request params")
		err := errs.NewValidationError("User email not present in request params")
		c.JSON(err.Code, err.AsMessage())

	} else {
		user, err := uh.service.FindByEmail(userEmail)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			data, err := json.Marshal(user)
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

// @Schemes
// @Description Creates a user on registration
// @Tags users
// @Produce json
// @Accept json
// @Param        user  body      UserDTO  true  "User Registration"
// @Success 201 {object} domain.User
// @Router /users/register [post]
func (uh *UserHandlers) Register(c *gin.Context) {

	var newUser domain.User
	err := c.Bind(&newUser)
	fmt.Println(newUser, err)

	if err != nil {
		logger.Error("Invalid request body")
		err := errs.NewValidationError("Invalid request body")
		c.JSON(err.Code, err.AsMessage())

	} else {
		regUser, err := uh.service.Register(newUser)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			data, err := json.Marshal(regUser)
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			c.Data(http.StatusCreated, "application/json", data)
		}
	}
}

// @Schemes
// @Description Logs in a user
// @Tags users
// @Produce json
// @Accept json
// @Param        LoginDTO  body LoginDTO true "User login"
// @Success 200 {object} UserResponseDTO
// @Router /login [post]
func (uh *UserHandlers) Login(c *gin.Context) {

	var loginDto LoginDTO
	err := c.Bind(&loginDto)

	if err != nil {
		logger.Error("Invalid request body")
		err := errs.NewValidationError("Invalid request body")
		c.JSON(err.Code, err.AsMessage())

	} else {
		token, err := uh.service.Login(loginDto.Email, loginDto.Password)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			c.SetCookie("token", token, 60*1000, "/", "localhost", false, true)
			data, err := json.Marshal(UserResponseDTO{
				Message: "login successful",
				Token:   token,
			})
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

// @Schemes
// @Description Updates user by userEmail
// @Tags users
// @Param        token   header      string  true  "Auth token"
// @Param        user details   body      UserDTO true  "User details"
// @Produce json
// @Success 200 {object} domain.User
// @Failure      500  {object} errs.AppError
// @Router /users/{userId} [put]
func (uh *UserHandlers) UpdateUser(c *gin.Context) {

	var updatedUser domain.User
	err := c.Bind(&updatedUser)

	if err != nil {
		logger.Error("Invalid request body")
		err := errs.NewValidationError("Invalid request body")
		c.JSON(err.Code, err.AsMessage())

	} else {
		user, err := uh.service.Update(updatedUser)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			data, err := json.Marshal(user)
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			c.Data(http.StatusCreated, "application/json", data)
		}
	}
}

// @Schemes
// @Description Deletes user by email
// @Tags users
// @Param        userEmail   path      int  true  "User Email"
// @Param        token   header      string  true  "Auth Token"
// @Produce json
// @Success 200 {object} UserDTO
// @Failure      500  {object} errs.AppError
// @Router /users/{userId} [delete]
func (uh *UserHandlers) DeleteUserByEmail(c *gin.Context) {

	userEmail, ok := c.Params.Get("userEmail")

	if !ok {
		logger.Error("User email not present in request params")
		err := errs.NewValidationError("User email not present in request params")
		c.JSON(err.Code, err.AsMessage())

	} else {
		user, err := uh.service.DeleteByEmail(userEmail)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			data, err := json.Marshal(user)
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

func (uh *UserHandlers) VerifyUserToken(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		token = c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Invalid / missing Auth token"})
			c.Abort()
			return
		}
	}

	userEmail, role, err1 := utils.ParseAuthToken(token)
	if err1 != nil {
		c.JSON(err1.Code, err1.AsMessage())
		c.Abort()
		return
	}

	authDto := AuthDTO{
		Email: userEmail,
		Role:  role,
	}

	c.JSON(http.StatusOK, authDto)
}

func (uh *UserHandlers) HelloWorldHandler(c *gin.Context) {

	token := "Hello world"
	data, _ := json.Marshal(token)
	c.Data(http.StatusOK, "application/json", data)
}
