package routes

import (
	"PurchaseService/controllers"

	"github.com/gin-gonic/gin"
)

func PurchaseRoute(router *gin.Engine) {
	router.Use(controllers.IsAuthorized("user"))
	router.POST("/purchase", controllers.PurchaseTicket())
	router.GET("/purchase/:purchasedid", controllers.GetPurchased())
	router.DELETE("/purchase/:purchasedid", controllers.DeletePurchased())
}
