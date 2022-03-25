package auth

import (
	middleware "swiggy/gin/lib/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", Signup)
	router.POST("/login", Login)
	router.GET("/check", middleware.CheckAuthMiddleware, CheckAuth)
}
