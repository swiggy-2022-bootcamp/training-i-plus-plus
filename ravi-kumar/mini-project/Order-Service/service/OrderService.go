package service

import (
	repository "Order-Service/Repository"
	errors "Order-Service/errors"
	"Order-Service/kafka"
	mockdata "Order-Service/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrderService interface {
	PlaceOrder(orderPlaced mockdata.Order, shouldProduceKafkaMessage ...bool) (orderId *string, err error)
	GetOrders(userId string) (orders []mockdata.Order, err error)
	OrderPayment(orderId string, shouldProduceKafkaMessage ...bool) (successMessage *string, err error)
	DeliverOrder(orderId string, shouldProduceKafkaMessage ...bool) (successMessage *string, err error)
	CancelOrder(orderId string, accessorUserId string) (successMessage *string, err error)
	ReadCloserToString(res *http.Response) (message string)
}

type OrderService struct {
	mongoDAO repository.IMongoDAO
	httpRepo repository.IHttpRepo
}

func InitOrderService(initMongoDAO repository.IMongoDAO, initHttpRepo repository.IHttpRepo) IOrderService {
	orderService := new(OrderService)
	orderService.mongoDAO = initMongoDAO
	orderService.httpRepo = initHttpRepo
	return orderService
}

func (orderService *OrderService) PlaceOrder(orderPlaced mockdata.Order, shouldProduceKafkaMessage ...bool) (orderId *string, err error) {
	//if shouldProduceKafkaMessage argument is absent, canProduceKafkaMessage shall be true by default - for test util
	canProduceKafkaMessage := true
	if len(shouldProduceKafkaMessage) != 0 {
		canProduceKafkaMessage = shouldProduceKafkaMessage[0]
	}

	userId := orderPlaced.UserId
	if !orderService.httpRepo.IsValidUser(userId) {
		return nil, errors.UserNotFoundError()
	}

	success, errorResponse, errorProductIndex := orderService.httpRepo.UpdateProductQuantity(userId, orderPlaced.Items, -1)
	if !success {
		errorMessage := orderService.ReadCloserToString(errorResponse) + ". Product Id: " + orderPlaced.Items[*errorProductIndex] + " (Order rolled back)"
		return nil, &errors.OrderError{Status: http.StatusBadRequest, ErrorMessage: errorMessage}
	}

	orderPlaced.OrderDate = time.Now()
	orderPlaced.DeliveryDate = orderPlaced.OrderDate.AddDate(0, 0, 6)
	orderPlaced.Status = "confirmed"

	returnedOrderId := orderService.mongoDAO.MongoPlaceOrder(orderPlaced)

	if canProduceKafkaMessage {
		ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
		kafka.Produce(ctx, nil, []byte("orderId: "+returnedOrderId+" --- status: "+orderPlaced.Status))
	}

	return &returnedOrderId, nil
}

func (orderService *OrderService) GetOrders(userId string) (orders []mockdata.Order, err error) {
	//TODO: check if user exists
	if !orderService.httpRepo.IsValidUser(userId) {
		return nil, errors.UserNotFoundError()
	}

	return orderService.mongoDAO.MongoGetOrderByUserId(userId)
}

func (orderService *OrderService) OrderPayment(orderId string, shouldProduceKafkaMessage ...bool) (successMessage *string, err error) {
	//if shouldProduceKafkaMessage argument is absent, canProduceKafkaMessage shall be true by default - for test util
	canProduceKafkaMessage := true
	if len(shouldProduceKafkaMessage) != 0 {
		canProduceKafkaMessage = shouldProduceKafkaMessage[0]
	}

	//convert orderId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	order, err := orderService.mongoDAO.MongoGetOrderByOrderId(objectId)

	if err != nil {
		return nil, err
	}

	if order.Status == "payment done" || order.Status == "delivered" {
		return nil, errors.OrderAlreadyPaidForError()
	}

	order.Status = "payment done"

	_, error := orderService.mongoDAO.MongoUpdateOrderByOrderId(objectId, *order)

	if error != nil {
		return nil, errors.InternalServerError()
	}

	if canProduceKafkaMessage {
		ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
		kafka.Produce(ctx, nil, []byte("orderId: "+orderId+" --- status: "+order.Status))
	}

	str := "order payment successful"
	successMessage = &str
	return
}

func (orderService *OrderService) DeliverOrder(orderId string, shouldProduceKafkaMessage ...bool) (successMessage *string, err error) {
	//if shouldProduceKafkaMessage argument is absent, canProduceKafkaMessage shall be true by default - for test util
	canProduceKafkaMessage := true
	if len(shouldProduceKafkaMessage) != 0 {
		canProduceKafkaMessage = shouldProduceKafkaMessage[0]
	}

	//convert order id string to objectId type
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	order, err := orderService.mongoDAO.MongoGetOrderByOrderId(objectId)

	if err != nil {
		return nil, err
	}

	if order.Status == "confirmed" {
		return nil, errors.PaymentIncompleteError()
	}

	if order.Status == "delivered" {
		return nil, errors.OrderAlreadyDeliveredError()
	}

	order.Status = "delivered"

	_, error := orderService.mongoDAO.MongoUpdateOrderByOrderId(objectId, *order)
	if error != nil {
		return nil, errors.InternalServerError()
	}

	if canProduceKafkaMessage {
		ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
		kafka.Produce(ctx, nil, []byte("orderId: "+orderId+" --- status: "+order.Status))
	}

	str := "order delivered"
	successMessage = &str
	return
}

func (orderService *OrderService) CancelOrder(orderId string, accessorUserId string) (successMessage *string, err error) {
	//only undelivered order ordered by the current accessor is allowed

	//convert orderId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}
	order, err := orderService.mongoDAO.MongoGetOrderByOrderId(objectId)
	if err != nil {
		return nil, err
	}

	if order.UserId != accessorUserId {
		return nil, errors.AccessDenied()
	}

	if order.Status == "delivered" {
		return nil, errors.OrderAlreadyDeliveredError()
	}

	orderService.mongoDAO.MongoDeleteOrderById(objectId)

	refundStr := "Order cancelled. Refund of " + fmt.Sprint(order.Amount) + " initiated"
	orderCancelledStr := "Order cancelled."
	if order.Status == "payment done" {
		return &refundStr, nil
	} else {
		return &orderCancelledStr, nil
	}
}

func (orderService *OrderService) ReadCloserToString(res *http.Response) (message string) {
	if res == nil {
		return ""
	}
	body := res.Body
	json.NewDecoder(body).Decode(&message)
	return
}
