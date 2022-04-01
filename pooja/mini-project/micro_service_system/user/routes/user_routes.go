package routes

import (
	"user/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/user")

	userRouter.POST("/signup", service.Signup())
	userRouter.POST("/login", service.Login())

}
