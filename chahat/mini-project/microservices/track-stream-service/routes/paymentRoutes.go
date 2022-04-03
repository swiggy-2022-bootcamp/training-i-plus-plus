package routes

import(
    // controller "github.com/bhatiachahat/mongoapi/controllers"
	 //"github.com/bhatiachahat/mongoapi/middleware"
	controller "bhatiachahat/track-stream-service/controller"
	 "github.com/gin-gonic/gin"
)

func TrackStreamRoutes(incomingRoutes *gin.Engine){
	//incomingRoutes.Use(middleware.Authenticate())
	// inventoryRouter.POST("/:inventoryId/product/add", ic.AddProduct)
	// inventoryRouter.GET("/:inventoryId/product/:productId", ic.GetProduct)

	   incomingRoutes.GET("/getAnalytics", controller.GetTrackingData())
	 //  incomingRoutes.GET("/:cart_id/product/:product_id", controller.GetProductFromCart())

    //    incomingRoutes.DELETE("/products/:product_id", controller.DeleteProduct())
    //    incomingRoutes.GET("/products/:product_id", controller.GetProduct())
	//    incomingRoutes.PUT("/products/:product_id", controller.UpdateProduct())
	

}