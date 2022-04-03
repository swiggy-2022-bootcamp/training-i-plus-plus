package routes

import(
    // controller "github.com/bhatiachahat/mongoapi/controllers"
	 //"github.com/bhatiachahat/mongoapi/middleware"
	controller "bhatiachahat/order-service/controller"
	 "github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	//incomingRoutes.Use(middleware.Authenticate())
	  // incomingRoutes.GET("/products", controller.GetAllProducts())
	   incomingRoutes.POST("/orders/place_order/:user_id", controller.PlaceOrder())
     //  incomingRoutes.DELETE("/products/:product_id", controller.DeleteProduct())
      // incomingRoutes.GET("/products/:product_id", controller.GetProduct())
	  // incomingRoutes.PUT("/products/:product_id", controller.UpdateProduct())
	

}