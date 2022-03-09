package routes

import (
	"github.com/meyash/shop-cart/controllers",
	"gitub.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup",controllers.SignUp)
	incomingRoutes.POST("/users/login",controllers.Login)
	incomingRoutes.POST("/admin/addproduct",controllers.AddProduct)
	incomingRoutes.GET("/users/productview",controllers.ProductView)
	incomingRoutes.GET("/users/search",controllers.Search)
}