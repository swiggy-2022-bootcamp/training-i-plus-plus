package services

import (
	"context"
	"errors"

	"github.com/sachinsom93/shopping-cart/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryServiceImpl struct {
	InventoryCollection *mongo.Collection
	Ctx                 context.Context
}

func NewInventoryService(inventoryCollection *mongo.Collection, ctx context.Context) *InventoryServiceImpl {
	return &InventoryServiceImpl{
		InventoryCollection: inventoryCollection,
		Ctx:                 ctx,
	}
}

// Function to register a new inventory
func (is *InventoryServiceImpl) RegisterInventory(inventory *models.Inventory) error {
	if _, err := is.InventoryCollection.InsertOne(is.Ctx, inventory); err != nil {
		return err
	}
	return nil
}

// Function to add a specific product into inventory
func (is *InventoryServiceImpl) AddProduct(product *models.Product) error {
	// matchStage := bson.D{{"$match", bson.D{{"_id", product.InventoryId}}}}
	filter := bson.D{bson.E{Key: "_id", Value: product.InventoryId}}
	pushQuery := bson.D{bson.E{Key: "$push", Value: bson.E{Key: "inventory_products", Value: product}}}
	result, err := is.InventoryCollection.UpdateOne(is.Ctx, filter, pushQuery)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("No inventory found.")
	}
	if result.ModifiedCount != 1 {
		return errors.New("Something went wrong, product not added.")
	}
	return nil
}

// Function to remove a specific product from inventory
func (is *InventoryServiceImpl) RemoveProduct(productID int) error {
	return nil
}

// Function to update product by given product instance
func (is *InventoryServiceImpl) UpdateProduct(product *models.Product) (*models.Product, error) {
	return nil, nil
}

// Function to get a specific product item
func (is *InventoryServiceImpl) GetProduct(productID int) (*models.Product, error) {
	return nil, nil
}

// Function to get all product of an inventory
func (is *InventoryServiceImpl) GetAllProducts(inventoryID int) ([]*models.Product, error) {
	return nil, nil
}
