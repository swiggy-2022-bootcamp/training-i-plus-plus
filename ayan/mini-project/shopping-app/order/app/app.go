package app

import (
	"fmt"

	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/db"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/docs"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/domain"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/utils/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {

	dbClient := db.NewDbClient()

	orderRepo := db.NewOrderRepositoryDB(dbClient)
	orderService := domain.NewOrderService(orderRepo)
	orderHandlers := OrderHandlers{service: orderService}

	orderRouter := gin.Default()

	apiRouter := orderRouter.Group("/api")

	docs.SwaggerInfo.BasePath = "/api"
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	orderRoutesGroup := apiRouter.Group("/orders")

	orderRoutesGroup.GET("/", orderHandlers.HelloWorldHandler)
	orderRoutesGroup.POST("/", orderHandlers.PlaceOrder)
	orderRoutesGroup.GET("/:orderId", orderHandlers.GetOrderById)

	orderRouter.Run(":8082")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", "127.0.0.1", "8082"))
}
