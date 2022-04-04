package routes

import(
  
	controller "bhatiachahat/payment-service/controller"
	 "github.com/gin-gonic/gin"
)

func PaymentRoutes(incomingRoutes *gin.Engine){
	

	  incomingRoutes.POST("/payment/:order_id", controller.DoPayment())
	
	

}