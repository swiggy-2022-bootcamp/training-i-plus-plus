package routes

import (
	"userService/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.Use(controllers.CheckAuthorized("USER"))
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userid", controllers.GetAUser())
	router.PUT("/user/:userid", controllers.EditAUser())
	router.DELETE("/user/:userid", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/user/viewtrains", controllers.GetAllTrains())
	router.GET("/user/checktrain", controllers.AvailTicketCheck())
	router.POST("/bookedticket", controllers.CreateBookedTicket())
	router.GET("/bookedticket/:bookedticketid", controllers.GetBookedTicket())
	router.DELETE("/bookedticket/:bookedticketid", controllers.DeleteBookedTicket())
}
