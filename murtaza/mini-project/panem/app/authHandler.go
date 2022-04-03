package app

import (
	"encoding/json"
	"net/http"
	"panem/domain"
	"panem/utils/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService domain.AuthService
}

type loginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ah AuthHandler) handleLogin(c *gin.Context) {
	var credentials loginDTO
	err := json.NewDecoder(c.Request.Body).Decode(&credentials)

	if err != nil {
		customErr := errs.NewValidationError("Invalid request paylaod")
		c.JSON(http.StatusBadRequest, customErr)
	}
	var username = credentials.Username
	var password = credentials.Password

	jwtToken, err2 := ah.authService.AuthenticateUser(username, password)
	if err2 != nil {
		c.JSON(err2.Code, err2)
	}
	c.SetCookie("auth-token", jwtToken, 60*1000, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully !"})
}

func (ah AuthHandler) authMiddleware(c *gin.Context) {

	cookie, err := c.Cookie("auth-token")

	if err != nil {
		cookie = c.Request.Header.Get("auth-token")
		if cookie == "" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Invalid / missing Auth token"})
			c.Abort()
			return
		}
	}

	userId, _, err2 := ah.authService.ParseAuthToken(cookie)
	if err2 != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid / missing Auth token"})
		c.Abort()
		return
	}

	userIdStr := c.Params.ByName("userId")
	if userIdStr != "" {
		userIdFromParams, _ := strconv.Atoi(userIdStr)
		if userIdFromParams != userId {
			c.JSON(http.StatusForbidden, gin.H{"message": "Invalid / missing Auth token"})
			c.Abort()
			return
		}
	}
	c.Params = append(c.Params, gin.Param{
		Key:   "userId",
		Value: strconv.Itoa(userId),
	})
}

func (ah AuthHandler) isTokenValid(c *gin.Context) {

	type authDTO struct {
		UserId int         `json:"user_id"`
		Role   domain.Role `json:"role"`
	}

	cookie, err := c.Cookie("auth-token")
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid / missing Auth token"})
		c.Abort()
		return
	}

	userId, role, err2 := ah.authService.ParseAuthToken(cookie)
	if err2 != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid / missing Auth token"})
		c.Abort()
		return
	}

	authDto := authDTO{
		UserId: userId,
		Role:   role,
	}

	c.JSON(http.StatusOK, authDto)
}
