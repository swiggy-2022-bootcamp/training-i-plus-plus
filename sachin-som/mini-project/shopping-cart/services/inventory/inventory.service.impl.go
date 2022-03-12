package services

import (
	"context"

	"github.com/sachinsom93/shopping-cart/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryServiceImpl struct {
	inventoryCollection *mongo.Collection
	ctx                 context.Context
}

func NewInventoryService(inventoryCollection *mongo.Collection, ctx context.Context) *InventoryServiceImpl {
	return &InventoryServiceImpl{
		inventoryCollection: inventoryCollection,
		ctx:                 ctx,
	}
}

// Function to register a new inventory
func (is *InventoryServiceImpl) RegisterInventory(inventory *models.Inventory) error {
	return nil
}

// Function to add a specific product into inventory
func (is *InventoryServiceImpl) AddProduct(product *models.Product) error {
	return nil
}

// Function to remove a specific product from inventory
func (is *InventoryServiceImpl) RemoveProduct(productID primitive.ObjectID) error {
	return nil
}

// Function to update product by given product instance
func (is *InventoryServiceImpl) UpdateProduct(product *models.Product) (*models.Product, error) {
	return nil, nil
}

// Function to get a specific product item
func (is *InventoryServiceImpl) GetProduct(productID primitive.ObjectID) (*models.Product, error) {
	return nil, nil
}

// Function to get all product of an inventory
func (is *InventoryServiceImpl) GetAllProducts(inventoryID primitive.ObjectID) ([]*models.Product, error) {
	return nil, nil
}
