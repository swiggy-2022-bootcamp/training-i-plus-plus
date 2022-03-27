package routes

import (
	"srctc/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/register", controllers.SignUp())
	router.POST("/login", controllers.Login())
}
