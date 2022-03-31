package routes

import (
	"ticket_reservation_system/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/user")

	router.POST("/add", controller.AddUser())
	router.GET("/get/:username", controller.GetUserByUsername())
	router.PUT("/update/:username", controller.UpdateUserPassword())
	router.DELETE("/delete/:username", controller.DeleteUserByUsername())

	router.GET("/users", controller.GetUsers())

	userRouter.POST("/signup", controller.Signup())
	userRouter.POST("/login", controller.Login())

	router.GET("/search_trains", controller.SearchTrains())
	router.POST("/booking/book_ticket", controller.BookTickets())
	router.PUT("/booking/cancel_booking", controller.CancelBooking())
	router.GET("/bookings", controller.ViewBookings())

}
