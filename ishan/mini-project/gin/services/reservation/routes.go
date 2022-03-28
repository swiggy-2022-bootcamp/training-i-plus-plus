package reservation

import (
	middleware "swiggy/gin/lib/middlewares"

	"github.com/gin-gonic/gin"
)

func ReservationRoutes(router *gin.Engine) {
	router.POST("/reservation", middleware.CheckAuthMiddleware, ReserveSeat)
	router.GET("/reservation", middleware.CheckAuthMiddleware, FetchReservations)
	router.DELETE("/reservation/:id", middleware.CheckAuthMiddleware, CancelReservation)
}
