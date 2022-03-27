package routes

import (
	"srctc/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.Use(controllers.IsAuthorized())
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetUser())
	router.PUT("/user/:userId", controllers.EditUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/user/viewtrains", controllers.GetAllTrains())
	router.POST("/bookedticket", controllers.PurchaseTicket())
	router.GET("/bookedticket/:bookedticketID", controllers.GetPurchased())
	router.DELETE("/bookedticket/:bookedticketID", controllers.DeletePurchased())
}
