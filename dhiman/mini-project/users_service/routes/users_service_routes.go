package routes

import (
	"github.com/dhi13man/healthcare-app/users_service/controllers"
	"github.com/gin-gonic/gin"
)

const baseURL string = "/users/";

func GenerateUsersServiceRoutes(router *gin.Engine) {
	router.POST(baseURL + "clients/", controllers.CreateClient)
	router.GET(baseURL + "clients/:email/", controllers.GetClient)
	router.PUT(baseURL + "clients/:email/", controllers.UpdateClients)
}