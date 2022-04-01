package routes

import (
	"srctc/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.Use(controllers.IsAuthorized("user"))
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.POST("/purchase", controllers.PurchaseTicket())
	router.GET("/purchase/:bookedticketID", controllers.GetPurchased())
	router.DELETE("/purchase/:bookedticketID", controllers.DeletePurchased())
}
