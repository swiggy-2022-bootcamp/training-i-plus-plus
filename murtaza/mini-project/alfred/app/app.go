package app

import (
	"alfred/domain"
	"alfred/infra"
	"alfred/utils/logger"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func Start() {

	cartRespository := infra.NewCartRepository()
	orderRepository := infra.NewOrderRepository()

	orderAmountProducer := infra.NewProducer("test_topic")

	orderService := domain.NewOrderService(orderRepository, orderAmountProducer)

	cartHandler := CartHandler{
		cartService: domain.NewCartService(cartRespository, orderService),
	}

	r := Routes{
		router: gin.Default(),
	}

	//docs.SwaggerInfo.BasePath = "/api/v1"
	//r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.router.Group("/api")
	v1 := api.Group("/v1")

	carts := v1.Group("/carts")

	//carts.GET("/", cartHandler.)
	carts.PUT("/add", cartHandler.addToCart)
	carts.POST("/checkout", cartHandler.checkoutCart)
	err := r.router.Run(":8091")
	if err != nil {
		logger.Fatal("Unable to start item service")
	}
}
