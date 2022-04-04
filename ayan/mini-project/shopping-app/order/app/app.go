package app

import (
	"fmt"

	"order/db"
	"order/docs"
	"order/domain"
	"order/kafka"
	"order/utils/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {

	kafkaProducer := kafka.KafkaProducer()

	defer func() {
		kafkaProducer.Flush(15 * 1000)
		kafkaProducer.Close()
	}()

	dbClient := db.NewDbClient()

	orderRepo := db.NewOrderRepositoryDB(dbClient)
	orderService := domain.NewOrderService(orderRepo)
	orderHandlers := OrderHandlers{Service: orderService, KafkaProducer: kafkaProducer}

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
