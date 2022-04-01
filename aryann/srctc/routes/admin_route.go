package routes

import (
	"srctc/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	router.POST("/login", controllers.Login())
	router.Use(controllers.IsAuthorized("admin"))
	router.POST("/admin", controllers.CreateAdmin())
	router.GET("/admin/:adminid", controllers.GetAdmin())
	router.DELETE("/admin/:adminid", controllers.DeleteAdmin())
	router.POST("/train", controllers.CreateTrain())
	router.GET("/train/:trainid", controllers.GetTrain())
	router.PUT("/train/:trainid", controllers.EditTrain())
	router.DELETE("/train/:trainid", controllers.DeleteTrain())
	router.POST("/availticket/", controllers.CreateTicket())
	router.GET("/availticket/:availticketid", controllers.GetTicket())
	router.DELETE("/availticket/:availticketid", controllers.DeleteTicket())
}
