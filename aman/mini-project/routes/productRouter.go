package routes

import (
	"aman-swiggy-mini-project/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(r *gin.Engine) {
	r.GET("/foods", controllers.GetProducts())
	r.GET("/foods/:food_id", controllers.GetProduct())
	r.POST("/foods", controllers.CreateProduct())
	r.PATCH("/foods/:food_id", controllers.UpdateProduct())
}
