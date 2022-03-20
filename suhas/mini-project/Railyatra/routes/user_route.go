package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//router.Use(controllers.IsAuthorized())
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
	router.PUT("/user/:userId", controllers.EditAUser())
	router.DELETE("/user/:userId", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/user/viewtrains", controllers.GetAllTrains())
	router.POST("/bookedticket", controllers.CreateBookedTicket())
	router.GET("/bookedticket/:bookedticketID", controllers.GetBookedTicket())
	router.DELETE("/bookedticket/:bookedticketID", controllers.DeleteBookedTicket())
}
