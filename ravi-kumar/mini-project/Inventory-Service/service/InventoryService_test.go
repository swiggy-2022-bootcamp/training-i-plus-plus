package service

import (
	repository "Inventory-Service/Repository"
	"Inventory-Service/Repository/mocks"
	errors "Inventory-Service/errors"
	mockdata "Inventory-Service/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShouldCreateProduct(t *testing.T) {
	//making mongoDAO local instead of global to avoid mongoDAO.On() flux between tests
	var mongoDAO *mocks.MongoDAO = &mocks.MongoDAO{}
	newProduct := mockdata.Product{
		Name:        "Airtel Router Dual Band",
		Price:       6100,
		Description: "Dual Band Router Modem",
		Seller:      "Airtel",
		Rating:      4.599999904632568,
		Review: []string{
			"Good quality and service",
			"Decent product",
		},
		QuantityLeft: 22,
	}

	mongoDAO.On("MongoCreateProduct", newProduct).Return("624865bc030ba951de8956aa")

	inventoryService := InitInventoryService(mongoDAO)

	assert.Equal(t, inventoryService.CreateProduct(newProduct), "624865bc030ba951de8956aa")
}

func TestGetCatalogShouldReturnAtleastOneProduct(t *testing.T) {
	inventoryService := InitInventoryService(&repository.MongoDAO{})
	assert.Greater(t, len(inventoryService.GetCatalog()), 0)
}

func TestShouldReturnErrorWhenUserIdIsMalformed(t *testing.T) {
	updatedProduct := mockdata.Product{
		Name:        "Airtel Router Dual Band",
		Price:       6100,
		Description: "Dual Band Router Modem",
		Seller:      "Airtel",
		Rating:      4.599999904632568,
		Review: []string{
			"Good quality and service",
			"Decent product",
		},
		QuantityLeft: 22,
	}

	malformedProductId := "624865bc030ba951de8956bderred23zz"

	inventoryService := InitInventoryService(&repository.MongoDAO{})

	_, err := inventoryService.UpdateProductById(malformedProductId, updatedProduct)
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)

	_, err = inventoryService.GetProductById(malformedProductId)
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)

	_, err = inventoryService.DeleteProductbyId(malformedProductId)
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)

	_, err = inventoryService.UpdateProductQuantity(malformedProductId, 2)
	assert.EqualError(t, err, errors.MalformedIdError().ErrorMessage)
}

func TestShouldReturnErrorWhenUserIdIsNotFound(t *testing.T) {
	var mongoDAO *mocks.MongoDAO = &mocks.MongoDAO{}
	updatedProduct := mockdata.Product{
		Name:        "Airtel Router Dual Band",
		Price:       6100,
		Description: "Dual Band Router Modem",
		Seller:      "Airtel",
		Rating:      4.599999904632568,
		Review: []string{
			"Good quality and service",
			"Decent product",
		},
		QuantityLeft: 22,
	}

	productId, _ := primitive.ObjectIDFromHex("6243296ad6aed7d832e866c1")
	productIdAsStr := "6243296ad6aed7d832e866c1"

	mongoDAO.On("MongoGetProductById", productId).Return(nil, errors.IdNotFoundError())
	mongoDAO.On("MongoUpdateProductById", productId, updatedProduct).Return(nil, errors.IdNotFoundError())
	mongoDAO.On("MongoDeleteProductById", productId).Return(nil, errors.IdNotFoundError())

	inventoryService := InitInventoryService(mongoDAO)

	_, err := inventoryService.GetProductById(productIdAsStr)
	assert.EqualError(t, err, errors.IdNotFoundError().ErrorMessage)

	_, err = inventoryService.UpdateProductById(productIdAsStr, updatedProduct)
	assert.EqualError(t, err, errors.IdNotFoundError().ErrorMessage)

	_, err = inventoryService.DeleteProductbyId(productIdAsStr)
	assert.EqualError(t, err, errors.IdNotFoundError().ErrorMessage)

	_, err = inventoryService.UpdateProductQuantity(productIdAsStr, 2)
	assert.EqualError(t, err, errors.IdNotFoundError().ErrorMessage)
}

func TestShouldThrowOutOfStockError(t *testing.T) {
	var mongoDAO *mocks.MongoDAO = &mocks.MongoDAO{}
	product := mockdata.Product{
		Name:        "Airtel Router Dual Band",
		Price:       6100,
		Description: "Dual Band Router Modem",
		Seller:      "Airtel",
		Rating:      4.599999904632568,
		Review: []string{
			"Good quality and service",
			"Decent product",
		},
		QuantityLeft: 0,
	}

	productId, _ := primitive.ObjectIDFromHex("6243296ad6aed7d832e866c1")
	productIdAsStr := "6243296ad6aed7d832e866c1"

	mongoDAO.On("MongoGetProductById", productId).Return(&product, nil)

	inventoryService := InitInventoryService(mongoDAO)

	_, err := inventoryService.UpdateProductQuantity(productIdAsStr, -1, false)
	assert.EqualError(t, err, errors.OutOfStockError().ErrorMessage)
}

func TestShouldUpdateProductQuantity(t *testing.T) {
	var mongoDAO *mocks.MongoDAO = &mocks.MongoDAO{}
	product := mockdata.Product{
		Name:        "Airtel Router Dual Band",
		Price:       6100,
		Description: "Dual Band Router Modem",
		Seller:      "Airtel",
		Rating:      4.599999904632568,
		Review: []string{
			"Good quality and service",
			"Decent product",
		},
		QuantityLeft: 3,
	}
	updatedProduct := mockdata.Product{
		Name:        "Airtel Router Dual Band",
		Price:       6100,
		Description: "Dual Band Router Modem",
		Seller:      "Airtel",
		Rating:      4.599999904632568,
		Review: []string{
			"Good quality and service",
			"Decent product",
		},
		QuantityLeft: 2,
	}

	productId, _ := primitive.ObjectIDFromHex("6243296ad6aed7d832e866c2")
	productIdAsStr := "6243296ad6aed7d832e866c2"

	mongoDAO.On("MongoGetProductById", productId).Return(&product, nil)
	mongoDAO.On("MongoUpdateProductById", productId, updatedProduct).Return(&updatedProduct, nil)

	inventoryService := InitInventoryService(mongoDAO)

	quantityAfterUpdation, err := inventoryService.UpdateProductQuantity(productIdAsStr, -1, false)
	assert.Nil(t, err)
	assert.Equal(t, *quantityAfterUpdation, 2)
}
