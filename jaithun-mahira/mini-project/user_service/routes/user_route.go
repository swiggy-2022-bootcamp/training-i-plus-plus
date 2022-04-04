package routes

import (
	"user_service/controllers"
	"user_service/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserRoute(router *gin.Engine)  {
	  zap.L().Info("Initializing Routes")
    router.POST("/signup", controllers.SignUp())
		router.POST("/login", controllers.Login())

		router.Use(middlewares.Authenticate("ADMIN"))
		router.GET("/users", controllers.GetAllUsers())
    router.GET("/users/:userId", controllers.GetUserById())
		router.PUT("/users/:userId", controllers.UpdateUser())
		router.DELETE("/users/:userId", controllers.DeleteUser())	
}