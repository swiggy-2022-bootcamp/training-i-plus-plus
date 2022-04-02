package routes

import (
	"UserService/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/register", controllers.SignUp())
	router.POST("/login", controllers.Login())

	router.Use(controllers.IsAuthorized("user"))
	router.GET("/user/:userId", controllers.GetUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())

	router.Use(controllers.IsAuthorized("admin"))
	router.GET("/admin/:adminid", controllers.GetAdmin())
	router.DELETE("/admin/:adminid", controllers.DeleteAdmin())
}
