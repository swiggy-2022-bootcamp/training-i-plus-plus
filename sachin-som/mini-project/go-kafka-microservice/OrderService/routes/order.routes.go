package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/OrderService/controllers"
)

type OrderRoutes struct {
	OrderControllers *controllers.OrderControllers
}

func NewListingRoutes(orderControllers *controllers.OrderControllers) *OrderRoutes {
	return &OrderRoutes{
		OrderControllers: orderControllers,
	}
}

func (or *OrderRoutes) RegisterOrderRoutes(rg *gin.RouterGroup) {
	listingRouter := rg.Group("/")
	listingRouter.GET("/get/:userId", or.OrderControllers.GetOrders)
}
