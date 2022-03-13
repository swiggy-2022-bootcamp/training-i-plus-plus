package app

import (
	"net/http"
	"panem/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService domain.UserService
}

func (h UserHandler) getAllUsers(c *gin.Context) {
	params := c.Params
	userId, err := params.Get("userId")

	if err == false {
		c.JSON(404, err)

	} else {
		userId, _ := strconv.ParseInt(userId, 10, 0)
		user, err := h.userService.GetUserByUserId(int(userId))
		if err != nil {
			c.JSON(500, err)
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
