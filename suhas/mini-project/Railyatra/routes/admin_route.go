package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	//router.Use(controllers.IsAuthorized())
	router.POST("/admin", controllers.CreateAdmin())
	router.GET("/admin/:adminid", controllers.GetAdmin())
	router.PUT("/admin/:adminid", controllers.EditAdmin())
	router.DELETE("/admin/:adminid", controllers.DeleteAdmin())
	router.GET("/admins", controllers.GetAllAdmins())
	router.POST("/train/", controllers.CreateTrain())
	router.GET("/train/:trainid", controllers.GetTrain())
	router.PUT("/train/:trainid", controllers.EditTrain())
	router.DELETE("/train/:trainid", controllers.DeleteTrain())
	router.GET("/admin/viewtrains", controllers.GetAllTrains())
	router.POST("/availticket/", controllers.CreateAvailTicket())
	router.GET("/availticket/:availticketid", controllers.GetAvailTicket())
	router.PUT("/availticket/:availticketid", controllers.EditAvailTicket())
	router.DELETE("/availticket/:availticketid", controllers.DeleteAvailTicket())
}
