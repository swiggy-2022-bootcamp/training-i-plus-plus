package service

import (
	repository "Inventory-Service/Repository"
	errors "Inventory-Service/errors"
	kafka "Inventory-Service/kafka"
	mockdata "Inventory-Service/model"
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IInventoryService interface {
	CreateProduct(newProduct mockdata.Product) (InsertedID string)
	GetCatalog() (allProducts []mockdata.Product)
	GetProductById(productId string) (productRetrieved *mockdata.Product, err error)
	UpdateProductById(productId string, updatedProduct mockdata.Product) (productRetrieved *mockdata.Product, err error)
	DeleteProductbyId(productId string) (successMessage *string, err error)
	UpdateProductQuantity(productId string, updateCount int, shouldProduceKafkaMessages ...bool) (quantityAfterUpdation *int, err error)
}

type InventoryService struct {
	mongoDAO repository.IMongoDAO
}

func InitInventoryService(initMongoDAO repository.IMongoDAO) IInventoryService {
	inventoryService := new(InventoryService)
	inventoryService.mongoDAO = initMongoDAO
	return inventoryService
}

func (inventoryService *InventoryService) CreateProduct(newProduct mockdata.Product) (InsertedID string) {
	result := inventoryService.mongoDAO.MongoCreateProduct(newProduct)
	return result
}

func (inventoryService *InventoryService) GetCatalog() (allProducts []mockdata.Product) {
	return inventoryService.mongoDAO.MongoGetCatalog()
}

func (inventoryService *InventoryService) GetProductById(productId string) (productRetrieved *mockdata.Product, err error) {
	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)

	if err != nil {
		return nil, errors.MalformedIdError()
	}

	return inventoryService.mongoDAO.MongoGetProductById(objectId)
}

func (inventoryService *InventoryService) UpdateProductById(productId string, updatedProduct mockdata.Product) (productRetrieved *mockdata.Product, err error) {
	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	return inventoryService.mongoDAO.MongoUpdateProductById(objectId, updatedProduct)
}

func (inventoryService *InventoryService) DeleteProductbyId(productId string) (successMessage *string, err error) {
	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	return inventoryService.mongoDAO.MongoDeleteProductById(objectId)
}

func (inventoryService *InventoryService) UpdateProductQuantity(productId string, updateCount int, shouldProduceKafkaMessages ...bool) (quantityAfterUpdation *int, err error) {
	productRetrieved, error := inventoryService.GetProductById(productId)

	if error != nil {
		productError, ok := error.(*errors.ProductError)
		if ok {
			return nil, productError
		} else {
			fmt.Println("productError casting error in UpdateProductQuantity")
			return
		}
	}

	productRetrieved.QuantityLeft += updateCount

	//if shouldProduceKafkaMessages parameter is absent, do produce kafka messages by default
	canProduceKafkaMessages := true
	//else
	if len(shouldProduceKafkaMessages) != 0 {
		canProduceKafkaMessages = shouldProduceKafkaMessages[0]
	}

	if productRetrieved.QuantityLeft < 0 {
		if canProduceKafkaMessages {
			ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
			kafka.Produce(ctx, nil, []byte("productId: "+productId+" --- status: out of stock (critical)"))
		}

		return nil, errors.OutOfStockError()
	}

	//if quantity below threshold, notify monitoring service
	if canProduceKafkaMessages && productRetrieved.QuantityLeft < 20 {
		ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
		kafka.Produce(ctx, nil, []byte("productId: "+productId+" --- status: quantity below threshold ("+strconv.Itoa(productRetrieved.QuantityLeft)+") (critical)"))
	}

	_, err = inventoryService.UpdateProductById(productId, *productRetrieved)
	if err != nil {
		return nil, err
	}

	return &productRetrieved.QuantityLeft, nil
}
