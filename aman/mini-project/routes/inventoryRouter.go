package routes

import (
	"aman-swiggy-mini-project/controllers"
	"aman-swiggy-mini-project/middleware"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(r *gin.Engine) {
	r.Use(middleware.Authentication())
	r.GET("/inventory/:inventory_id", controllers.GetInventory())
}
