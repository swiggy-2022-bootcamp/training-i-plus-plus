package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	router.POST("/admin", controllers.CreateAdmin())
	router.GET("/admin/:adminId", controllers.GetAdmin())
	router.PUT("/admin/:adminId", controllers.EditAdmin())
	router.DELETE("/admin/:adminId", controllers.DeleteAdmin())
	router.GET("/admins", controllers.GetAllAdmins())
	router.POST("/train/", controllers.CreateTrain())
	router.GET("/train/:trainID", controllers.GetTrain())
	router.PUT("/train/:trainID", controllers.EditTrain())
	router.DELETE("/train/:trainID", controllers.DeleteTrain())
	router.POST("/availticket/", controllers.CreateAvailTicket())
	router.GET("/availticket/:availticketID", controllers.GetAvailTicket())
	router.PUT("/availticket/:availticketID", controllers.EditAvailTicket())
	router.DELETE("/availticket/:availticketID", controllers.DeleteAvailTicket())
}
