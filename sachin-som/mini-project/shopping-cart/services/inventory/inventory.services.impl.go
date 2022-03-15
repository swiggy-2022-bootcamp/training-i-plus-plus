package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/sachinsom93/shopping-cart/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (is *InventoryServiceImpl) RemoveProduct(inventoryID, productID string) error {
	inventory_ID, _ := strconv.Atoi(inventoryID)
	product_ID, _ := strconv.Atoi(productID)
	inventoryFilter := bson.D{bson.E{Key: "_id", Value: inventory_ID}}
	productFilterDelete := bson.D{bson.E{Key: "$pull", Value: bson.E{Key: "$pull", bson.E{Key: "inventory_products", bson.E{Key: "_id", Value: product_ID}}}}}}
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
func (is *InventoryServiceImpl) GetAllProducts(inventoryID string) ([]*models.Product, error) {
	inID, err := strconv.Atoi(inventoryID)
	filter := bson.D{bson.E{Key: "_id", Value: inID}}
	fieldsOpts := options.Find().SetProjection(bson.D{bson.E{Key: "inventory_products", Value: 1}}) // TODO: use aggregate
	cursor, err := is.InventoryCollection.Find(is.Ctx, filter, fieldsOpts)
	if err != nil {
		return nil, err
	}
	var products []*models.Product
	fmt.Println(products)
	for cursor.Next(is.Ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(is.Ctx)
	if len(products) == 0 {
		return nil, errors.New("producst not fuond.")
	}
	return products, err
}
