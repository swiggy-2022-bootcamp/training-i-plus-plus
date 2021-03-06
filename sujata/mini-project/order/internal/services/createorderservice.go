package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	mongodao "order/internal/dao"
	model "order/internal/dao/models"
	"order/internal/errors"
	"order/util"
	"os"
	"sync"

	l "log"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderService interface {
	ProcessRequest(ctx context.Context, email string, token string) (model.CreateOrderResponse, *errors.ServerError)
}

var createOrderServiceStruct CreateOrderService
var createOrderServiceOnce sync.Once

type createOrderService struct {
	config *util.RouterConfig
	dao    mongodao.MongoDAO
}

func InitCreateOrderService(config *util.RouterConfig, dao mongodao.MongoDAO) CreateOrderService {
	createOrderServiceOnce.Do(func() {
		createOrderServiceStruct = &createOrderService{
			config: config,
			dao:    dao,
		}
	})

	return createOrderServiceStruct
}

func GetCreateOrderService() CreateOrderService {
	if createOrderServiceStruct == nil {
		panic("Create order service not initialised")
	}

	return createOrderServiceStruct
}

// CreateOrderProcessRequest calls CART Service to get product details and then save the details in the order, with
// order status ORDER_PLACED and then publish this status to Kafka.
// Note: multiple Orders allowed for same user.
func (service *createOrderService) ProcessRequest(ctx context.Context, email string, token string) (model.CreateOrderResponse, *errors.ServerError) {

	// call to CART Service to get the product details
	respBytes, err := service.callCartService(token)
	if err != nil {
		return model.CreateOrderResponse{}, err
	}

	var order model.Order
	// unmarshal the cart response of product details to store them in order
	goErr := json.Unmarshal(respBytes, &order)
	if goErr != nil {
		log.WithError(goErr).Error("an error occurred while unmarhsalling the request bytes to order model")
	}

	// update the status of the order to ORDER_PLACED
	order.OrderStatus = model.ORDER_PLACED
	order.Email = email

	// Save the product details to Orders along with name and order Status - "ORDER_PLACED"
	id, err := service.dao.CreateOrder(ctx, order)
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while creating the order")
		return model.CreateOrderResponse{}, err
	}

	log.Info("Inserted Order: ", id)

	response := model.CreateOrderResponse{
		OrderId: id.(primitive.ObjectID).Hex(),
	}

	// Publish the order status to kafka
	err = service.publishToKafka(ctx, order)
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while publishing the message to kafka")
	}

	return response, nil
}

func (service *createOrderService) callCartService(token string) ([]byte, *errors.ServerError) {
	url := service.config.WebServerConfig.ProductBaseUrl
	var bodyBytes []byte

	// Create a new request using http
	req, goErr := http.NewRequest("GET", url, nil)
	if goErr != nil {
		log.WithError(goErr).Error("an error occurred while creating the http request")
		return nil, &errors.InternalError
	}

	// add authorization header to the req
	req.Header.Add("Authorization", token)

	// Send req using http Client
	client := &http.Client{}

	resp, goErr := client.Do(req)
	if goErr != nil || resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"Error": goErr,
		}).Error("an error occurred while calling to service: ", url)
		return bodyBytes, &errors.InternalError
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, goErr = io.ReadAll(resp.Body)
		if goErr != nil {
			log.WithField("Error: ", goErr).Error("an error occurred while reading the body bytes from response body")
			return bodyBytes, &errors.InternalError
		}
		bodyString := string(bodyBytes)
		log.Info(bodyString)
	}

	return bodyBytes, nil
}

// publishToKafka will publish the user email and order status to kafka
func (service *createOrderService) publishToKafka(ctx context.Context, order model.Order) *errors.ServerError {
	logger := l.New(os.Stdout, "kafka writer: ", 0)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "OrderStatus",
		Logger:  logger,
	})

	// Published message will be cosumed by CART Service which allows only one cart per user.
	// So, passing email will serves as a unique key to cart service and thus it can delete the cart
	// once the order is placed.
	err := writer.WriteMessages(ctx, kafka.Message{
		Key: []byte(order.Email),
		// create an arbitrary message payload for the value
		Value: []byte(order.OrderStatus),
	})

	if err != nil {
		log.WithField("Error: ", err).Warn("an error occurred while sending message to Kafka on Topic: ")
	}

	return nil
}
