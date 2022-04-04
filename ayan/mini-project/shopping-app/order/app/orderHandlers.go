package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"order/domain"
	kfka "order/kafka"
	"order/utils/errs"
	"order/utils/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

type OrderHandlers struct {
	Service       domain.OrderService
	KafkaProducer *kafka.Producer
}

type OrderItemDTO struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type OrderDTO struct {
	Id        string         `json:"id,omitempty"`
	ItemList  []OrderItemDTO `json:"item_list"`
	Amount    int            `json:"amount,omitempty"`
	UserEmail string         `json:"user_email"`
}

type OrderResponseDTO struct {
	Message string `json:"message"`
}

// @Schemes
// @Description Fetches order details by id
// @Tags orders
// @Param        orderId   path      string  true  "Id"
// @Produce json
// @Success 200 {object} domain.Order
// @Failure      403  {object} errs.AppError
// @Router /orders/{orderId} [get]
func (uh *OrderHandlers) GetOrderById(c *gin.Context) {

	orderId, ok := c.Params.Get("orderId")

	if !ok {
		logger.Error("Order Id not present in request params")
		err := errs.NewValidationError("Order Id not present in request params")
		c.JSON(err.Code, err.AsMessage())

	} else {
		order, err := uh.Service.FindById(orderId)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			data, err := json.Marshal(order)
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

// @Schemes
// @Description Creates a order
// @Tags orders
// @Produce json
// @Accept json
// @Param        order  body      OrderDTO  true  "Order Creation"
// @Success 201 {object} domain.Order
// @Router /orders/ [post]
func (uh *OrderHandlers) PlaceOrder(c *gin.Context) {

	var newOrder domain.Order
	err := c.Bind(&newOrder)
	fmt.Println(newOrder, err)

	if err != nil {
		logger.Error("Invalid request body")
		err := errs.NewValidationError("Invalid request body")
		c.JSON(err.Code, err.AsMessage())

	} else {
		regOrder, err := uh.Service.CreateOrder(newOrder)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			data, err := json.Marshal(regOrder)
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			go kfka.Produce(uh.KafkaProducer, string(data), "orders")
			c.Data(http.StatusCreated, "application/json", data)
		}
	}
}

func (uh *OrderHandlers) HelloWorldHandler(c *gin.Context) {

	token := "Hello Order World!"
	data, _ := json.Marshal(token)
	c.Data(http.StatusOK, "application/json", data)
}
