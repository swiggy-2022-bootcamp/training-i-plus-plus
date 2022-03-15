package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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
	pushQuery := bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "inventory_products", Value: product}}}}
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
func (is *InventoryServiceImpl) RemoveProduct(inventoryID, productID string) error {
	inventory_ID, _ := strconv.Atoi(inventoryID)
	product_ID, _ := strconv.Atoi(productID)
	inventoryFilter := bson.D{bson.E{Key: "_id", Value: inventory_ID}}
	productFilterDelete := bson.D{bson.E{Key: "$pull", Value: bson.E{Key: "inventory_products", Value: bson.E{Key: "_id", Value: product_ID}}}}
	result, err := is.InventoryCollection.UpdateOne(is.Ctx, inventoryFilter, productFilterDelete)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 1 {
		return nil
	}
	return errors.New("No Product Found.")
}

// Function to update product by given product instance
func (is *InventoryServiceImpl) UpdateProduct(product *models.Product) (*models.Product, error) {
	return nil, nil
}

// Function to get a specific product item
func (is *InventoryServiceImpl) GetProduct(inventoryID, productID string) (*models.Product, error) {
	inventory_ID, _ := strconv.Atoi(inventoryID)
	product_ID, _ := strconv.Atoi(productID)
	filter := bson.D{bson.E{Key: "_id", Value: inventory_ID}}
	var inventory *models.Inventory
	err := is.InventoryCollection.FindOne(is.Ctx, filter).Decode(&inventory)
	if err != nil {
		return nil, err
	}
	for _, p := range inventory.InventoryProducts {
		if p.ProductID == product_ID {
			fmt.Println(p)
			return &p, nil
		}
	}
	return nil, errors.New("No Product Found.")
}

// Function to get all product of an inventory
func (is *InventoryServiceImpl) GetAllProducts(inventoryID string) ([]models.Product, error) {
	var inventory *models.Inventory
	var products []models.Product
	var err error

	var inventory_ID int
	inventory_ID, err = strconv.Atoi(inventoryID)
	filter := bson.D{bson.E{Key: "_id", Value: inventory_ID}}

	err = is.InventoryCollection.FindOne(is.Ctx, filter).Decode(&inventory)
	if err != nil {
		return nil, err
	}
	for _, p := range inventory.InventoryProducts {
		products = append(products, p)
	}
	return products, nil
}

// [[
// 	{
// 		inventory_products
// 		[
// 			[{_id 2} {product_name test} {description } {price 10000} {ratings 5} {image_url } {upload_by 1} {inventory_id 5} {uploaded_at <nil>} {updated_at <nil>}]
// 			[{_id 3} {product_name test} {description } {price 10000} {ratings 5} {image_url } {upload_by 1} {inventory_id 5} {uploaded_at <nil>} {updated_at <nil>}]
// 			[{_id 7} {product_name iphone 12} {description A branch new iphone} {price 100000} {ratings 5} {image_url www.google.com/iphone} {upload_by 0} {inventory_id 5} {uploaded_at -62135596800000} {updated_at -62135596800000}]
// 			[{_id 7} {product_name iphone 12} {description A branch new iphone} {price 100000} {ratings 5} {image_url www.google.com/iphone} {upload_by 0} {inventory_id 5} {uploaded_at -62135596800000} {updated_at -62135596800000}]
// 		]
// 	}
// ]]
