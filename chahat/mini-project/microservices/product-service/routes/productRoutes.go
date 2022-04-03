package routes

import(
    // controller "github.com/bhatiachahat/mongoapi/controllers"
	 //"github.com/bhatiachahat/mongoapi/middleware"
	controller "bhatiachahat/product-service/controller"
	 "github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine){
	//incomingRoutes.Use(middleware.Authenticate())
	   incomingRoutes.GET("/products", controller.GetAllProducts())
	   incomingRoutes.POST("/products", controller.AddProduct())
       incomingRoutes.DELETE("/products/:product_id", controller.DeleteProduct())
       incomingRoutes.GET("/products/:product_id", controller.GetProduct())
	   incomingRoutes.PUT("/products/:product_id", controller.UpdateProduct())
	

}