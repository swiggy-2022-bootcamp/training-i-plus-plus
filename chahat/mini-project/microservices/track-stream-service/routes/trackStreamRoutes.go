package routes

import(
    // controller "github.com/bhatiachahat/mongoapi/controllers"
	 //"github.com/bhatiachahat/mongoapi/middleware"
	controller "bhatiachahat/track-stream-service/controller"
	 "github.com/gin-gonic/gin"
)

func TrackStreamRoutes(incomingRoutes *gin.Engine){
	
	   incomingRoutes.GET("/getAnalytics", controller.GetTrackingData())

	

}