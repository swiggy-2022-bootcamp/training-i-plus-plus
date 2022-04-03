package routes

import (
	"adminService/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	router.Use(controllers.CheckAuthorized("ADMIN"))
	// swagger:route POST /admin admins adminReq
	// Creates a new admin for the currently authenticated user.
	// If admin name is "exists", error conflict (409) will be returned.
	// responses:
	//  200: adminResp
	//  400: badReq
	//  409: conflict
	//  500: internal
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
