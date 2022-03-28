package routes

import (
	"aman-swiggy-mini-project/controllers"
	"aman-swiggy-mini-project/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts())
	r.Use(middleware.Authentication())
	r.GET("/products/:product_id", controllers.GetProduct())
	r.POST("/products", controllers.CreateProduct())
	r.PATCH("/products/:product_id", controllers.UpdateProduct())
}
