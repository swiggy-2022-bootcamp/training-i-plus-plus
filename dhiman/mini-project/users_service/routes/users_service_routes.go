package routes

import (
	"github.com/dhi13man/healthcare-app/users_service/controllers"
	"github.com/gin-gonic/gin"
)

const BaseURL string = "/users";

func GenerateUsersServiceRoutes(router *gin.Engine) {
	userRouter := router.Group(BaseURL)
	// Create
	userRouter.POST("clients", controllers.CreateClient)
	userRouter.POST("experts/diagnose", controllers.DiagnoseDisease)
	// Read
	userRouter.GET("clients/:email", controllers.GetClient)
	// Update
	userRouter.PUT("clients/:email", controllers.UpdateClients)
	// Delete
	userRouter.DELETE("clients/:email", controllers.DeleteClients)
}