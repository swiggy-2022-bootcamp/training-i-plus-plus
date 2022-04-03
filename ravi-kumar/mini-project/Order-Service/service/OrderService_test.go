package service

import (
	repository "Order-Service/Repository"
	"Order-Service/Repository/mocks"
	"Order-Service/errors"
	mockdata "Order-Service/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShouldThrowUserNotFoundErrorForInvalidUserId(t *testing.T) {
	order := mockdata.Order{
		UserId: "InvalidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c1",
		},
		Amount: 5400,
	}

	orderService := InitOrderService(&repository.MongoDAO{}, &repository.HttpRepo{})

	_, err := orderService.PlaceOrder(order)
	assert.EqualError(t, err, errors.UserNotFoundError().ErrorMessage)

	_, err = orderService.GetOrders(order.UserId)
	assert.EqualError(t, err, errors.UserNotFoundError().ErrorMessage)
}

func TestShouldThrowMalformedIdError(t *testing.T) {
	malformedOrderId := "MalformedOrderId"

	orderService := InitOrderService(&repository.MongoDAO{}, &repository.HttpRepo{})

	_, err := orderService.OrderPayment(malformedOrderId)
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)

	_, err = orderService.DeliverOrder(malformedOrderId)
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)

	_, err = orderService.CancelOrder(malformedOrderId, "validAccessorUserId")
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)
}

func TestShouldThrowErrorOnRollBackDueToProductQuantity(t *testing.T) {
	order := mockdata.Order{
		UserId: "SomeValidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c4",
		},
		Amount: 5400,
	}

	httpRepo := &mocks.IHttpRepo{}
	httpRepo.On("IsValidUser", order.UserId).Return(true)
	num := 1
	httpRepo.On("UpdateProductQuantity", order.UserId, order.Items, -1).Return(false, nil, &num)

	orderService := InitOrderService(&repository.MongoDAO{}, httpRepo)

	_, err := orderService.PlaceOrder(order)
	assert.NotNil(t, err)
}

func TestShouldReturnPlacedOrdersIdWithoutAnyError(t *testing.T) {
	order := mockdata.Order{
		UserId: "SomeValidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c4",
		},
		Amount: 5400,
	}

	httpRepo := &mocks.IHttpRepo{}
	httpRepo.On("IsValidUser", order.UserId).Return(true)
	num := 1
	httpRepo.On("UpdateProductQuantity", order.UserId, order.Items, -1).Return(true, nil, &num)

	mongoDAO := &mocks.IMongoDAO{}
	mongoDAO.On("MongoPlaceOrder", mock.Anything).Return("ReturnedOrderId")
	orderService := InitOrderService(mongoDAO, httpRepo)

	orderId, err := orderService.PlaceOrder(order, false)
	assert.Nil(t, err)
	assert.Equal(t, *orderId, "ReturnedOrderId")
}

func TestShouldReturnAtleastOneOrder(t *testing.T) {
	order := mockdata.Order{
		UserId: "SomeValidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c4",
		},
		Amount: 5400,
	}
	mongoDAO := &mocks.IMongoDAO{}
	mongoDAO.On("MongoGetOrderByUserId", order.UserId).Return([]mockdata.Order{order}, nil)

	httpRepo := &mocks.IHttpRepo{}
	httpRepo.On("IsValidUser", order.UserId).Return(true)

	orderService := InitOrderService(mongoDAO, httpRepo)

	orders, err := orderService.GetOrders(order.UserId)
	assert.Nil(t, err)
	assert.Equal(t, orders[0], order)
}

func TestShouldThrowErrorIfTransactionFlowIsIncorrect(t *testing.T) {
	//transaction flow: order confirmed -> payment -> delivery

	order := mockdata.Order{
		UserId: "SomeValidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c4",
		},
		Amount: 5400,
		Status: "payment done",
	}

	orderId := "62432a48ded825f68d264ed8"
	objectId, _ := primitive.ObjectIDFromHex(orderId)

	mongoDAO := &mocks.IMongoDAO{}
	mongoDAO.On("MongoGetOrderByOrderId", objectId).Return(&order, nil)

	orderService := InitOrderService(mongoDAO, &repository.HttpRepo{})
	_, err := orderService.OrderPayment(orderId)

	assert.EqualError(t, err, errors.OrderAlreadyPaidForError().ErrorMessage)

	////////////////////////////////////////////////////////////////////////////////

	order = mockdata.Order{
		UserId: "SomeValidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c4",
		},
		Amount: 5400,
		Status: "confirmed",
	}

	deliveredOrder := mockdata.Order{
		UserId: "SomeValidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c4",
		},
		Amount: 5400,
		Status: "delivered",
	}

	mongoDAO = &mocks.IMongoDAO{}
	mongoDAO.On("MongoGetOrderByOrderId", objectId).Return(&order, nil)

	_, err = orderService.DeliverOrder(orderId)
	assert.EqualError(t, err, errors.PaymentIncompleteError().ErrorMessage)

	order.Status = "delivered"
	_, err = orderService.DeliverOrder(orderId)
	assert.EqualError(t, err, errors.OrderAlreadyDeliveredError().ErrorMessage)

	orderService = InitOrderService(mongoDAO, &repository.HttpRepo{})
	order.Status = "payment done"
	mongoDAO.On("MongoUpdateOrderByOrderId", objectId, deliveredOrder).Return(&deliveredOrder, nil)
	msg, err := orderService.DeliverOrder(orderId)
	assert.Nil(t, err)
	assert.Equal(t, *msg, "order delivered")
}

func TestForSuccessfullOrderPayment(t *testing.T) {
	order := mockdata.Order{
		UserId: "SomeValidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c4",
		},
		Amount: 5400,
		Status: "confirmed",
	}

	paidOrder := mockdata.Order{
		UserId: "SomeValidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c4",
		},
		Amount: 5400,
		Status: "payment done",
	}

	orderId := "62432a48ded825f68d264ed8"
	objectId, _ := primitive.ObjectIDFromHex(orderId)

	mongoDAO := &mocks.IMongoDAO{}
	mongoDAO.On("MongoGetOrderByOrderId", objectId).Return(&order, nil)
	mongoDAO.On("MongoUpdateOrderByOrderId", objectId, paidOrder).Return(&paidOrder, nil)

	orderService := InitOrderService(mongoDAO, &repository.HttpRepo{})
	msg, err := orderService.OrderPayment(orderId, false)
	assert.Nil(t, err)
	assert.Equal(t, *msg, "order payment successful")
}

func TestShouldDenyAccess(t *testing.T) {
	order := mockdata.Order{
		UserId: "SomeValidUserId",
		Items: []string{
			"6243296ad6aed7d832e866c1",
			"6243296ad6aed7d832e866c4",
		},
		Amount: 5400,
		Status: "delivered",
	}
	orderId := "62432a85ded825f68d264ef1"
	accessorUserId := "62432a85ded825f68d264ef2"
	objectId, _ := primitive.ObjectIDFromHex(orderId)

	mongoDAO := &mocks.IMongoDAO{}
	mongoDAO.On("MongoGetOrderByOrderId", objectId).Return(&order, nil)

	orderService := InitOrderService(mongoDAO, &repository.HttpRepo{})
	_, err := orderService.CancelOrder(orderId, accessorUserId)
	assert.EqualError(t, err, errors.AccessDenied().ErrorMessage)

	order.UserId = accessorUserId
	_, err = orderService.CancelOrder(orderId, accessorUserId)
	assert.EqualError(t, err, errors.OrderAlreadyDeliveredError().ErrorMessage)

	order.Status = "payment done"
	mongoDAO.On("MongoDeleteOrderById", objectId).Return()
	msg, err := orderService.CancelOrder(orderId, accessorUserId)
	assert.Nil(t, err)
	assert.Equal(t, *msg, "Order cancelled. Refund of "+fmt.Sprint(order.Amount)+" initiated")

	order.Status = "confirmed"
	mongoDAO.On("MongoDeleteOrderById", objectId).Return()
	msg, err = orderService.CancelOrder(orderId, accessorUserId)
	assert.Nil(t, err)
	assert.Equal(t, *msg, "Order cancelled.")

}
