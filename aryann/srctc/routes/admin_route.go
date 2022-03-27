package routes

import (
	"srctc/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	// router.Use(controllers.IsAuthorized())
	router.POST("/admin", controllers.CreateAdmin())
	router.GET("/admin/:adminid", controllers.GetAdmin())
	// router.PUT("/admin/:adminid", controllers.EditAdmin())
	router.DELETE("/admin/:adminid", controllers.DeleteAdmin())
	// router.GET("/admins", controllers.GetAllAdmins())
	router.POST("/train", controllers.CreateTrain())
	router.GET("/train/:trainid", controllers.GetTrain())
	router.PUT("/train/:trainid", controllers.EditTrain())
	router.DELETE("/train/:trainid", controllers.DeleteTrain())
	// router.GET("/admin/viewtrains", controllers.GetAllTrains())
	router.POST("/availticket/", controllers.CreateTicket())
	router.GET("/availticket/:availticketid", controllers.GetTicket())
	// router.PUT("/availticket/:availticketid", controllers.EditTicket())
	router.DELETE("/availticket/:availticketid", controllers.DeleteTicket())
}
