package routes

import (
	"github.com/dhi13man/healthcare-app/users_service/controllers"
	"github.com/gin-gonic/gin"
)

const baseURL string = "/users";

func GenerateUsersServiceRoutes(router *gin.Engine) {
	userRouter := router.Group(baseURL)
	userRouter.POST("clients", controllers.CreateClient)
	userRouter.GET("clients/:email", controllers.GetClient)
	userRouter.PUT("clients/:email", controllers.UpdateClients)
}