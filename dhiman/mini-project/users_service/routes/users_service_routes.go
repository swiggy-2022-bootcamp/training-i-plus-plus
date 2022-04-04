package routes

import (
	"github.com/dhi13man/healthcare-app/users_service/controllers"
	"github.com/gin-gonic/gin"
)

const BaseURL string = "/users";

func GenerateUsersServiceRoutes(router *gin.Engine) {
	usersRouter := router.Group(BaseURL)
	clientsUsersRouter := usersRouter.Group("/clients")

	// Create
	clientsUsersRouter.POST("/", controllers.CreateClient)
	usersRouter.POST("experts/diagnose", controllers.DiagnoseDisease)
	// Read
	clientsUsersRouter.GET("/:email", controllers.GetClient)
	// Update
	clientsUsersRouter.PUT("/:email", controllers.UpdateClients)
	// Delete
	clientsUsersRouter.DELETE("/:email", controllers.DeleteClients)
}