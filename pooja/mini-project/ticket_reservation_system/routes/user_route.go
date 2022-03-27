package routes

import (
	"ticket_reservation_system/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controller.AddUser())
	router.GET("/user/:username", controller.GetUserByUsername())
	router.PUT("/user/:username", controller.UpdateUserPassword())
	router.DELETE("/user/:username", controller.DeleteUserByUsername())

	router.GET("/users", controller.GetUsers())

	router.POST("/signup", controller.Signup())
	router.POST("/login", controller.Login())

	router.POST("/train", controller.AddTrain())
	router.GET("/train/:trainnumber", controller.GetTrainByTrainNumber())
	router.GET("/trains", controller.GetAllTrains())
	router.PUT("/train/:train_number", controller.UpdateTrainDetails())
	router.DELETE("/train/:train_number", controller.DeleteTrain())

	router.GET("/search_trains", controller.SearchTrains())
	router.POST("/booking/book_ticket", controller.BookTickets())
	router.PUT("/booking/cancel_booking", controller.CancelBooking())
	router.GET("/bookings", controller.ViewBookings())

}
