package routes

import (
	"srctc/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/login", controllers.Login())
	router.Use(controllers.IsAuthorized("user"))
	router.GET("/user/:userId", controllers.GetUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.POST("/purchase", controllers.PurchaseTicket())
	router.GET("/purchase/:purchasedid", controllers.GetPurchased())
	router.DELETE("/purchase/:purchasedid", controllers.DeletePurchased())
}
