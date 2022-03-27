package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/OrderService/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderControllers struct {
	OrderService services.OrderServices
}

func NewOrderCollection(orderService services.OrderServices) *OrderControllers {
	return &OrderControllers{
		OrderService: orderService,
	}
}

func (oc *OrderControllers) GetOrders(gctx *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(gctx.Param("userId"))
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	orders, err := oc.OrderService.GetOrders(userId)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, orders)
}
