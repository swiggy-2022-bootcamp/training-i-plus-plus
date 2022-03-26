package services

import (
	"context"
	"errors"

	goKafka "github.com/go-kafka-microservice/InventoryService/goKafka/producer"
	"github.com/go-kafka-microservice/InventoryService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryServicesImpl struct {
	InventoryCollection *mongo.Collection
	ProductCollection   *mongo.Collection
	Ctx                 context.Context
}

func NewInventoryService(inventoryCollection *mongo.Collection, productCollection *mongo.Collection, ctx context.Context) *InventoryServicesImpl {
	return &InventoryServicesImpl{
		InventoryCollection: inventoryCollection,
		ProductCollection:   productCollection,
		Ctx:                 ctx,
	}
}

func (is *InventoryServicesImpl) RegisterInventory(inventory *models.Inventory) (string, error) {
	inventoryId := primitive.NewObjectID()
	inventory.ID = inventoryId
	if _, err := is.InventoryCollection.InsertOne(is.Ctx, inventory); err != nil {
		return "", err
	}
	return inventoryId.Hex(), nil
}

func (is *InventoryServicesImpl) AddProduct(inventoryId primitive.ObjectID, product *models.Product) error {

	// Mongo Queries
	filterInventory := bson.D{bson.E{Key: "_id", Value: inventoryId}}
	pushProuductID := bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "products", Value: product}}}}

	// Craete Product and get prouduct ID
	product.ID = primitive.NewObjectID()
	if _, err := is.ProductCollection.InsertOne(is.Ctx, product); err != nil {
		return err
	}

	// Push ProuductID to inventory
	result, err := is.InventoryCollection.UpdateOne(is.Ctx, filterInventory, pushProuductID)

	// Errors
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("No inventory found.")
	}
	if result.ModifiedCount != 1 {
		return errors.New("Something went wrong, product not added.")
	}

	// Save Product to Kafka - products (topic)
	p, err := goKafka.CreateProducer(goKafka.Cfg())
	if err != nil {
		return err
	}
	kp := goKafka.NewKafkaProducer(p)
	_, err = kp.WriteMessage("products", product)
	if err != nil {
		return nil
	}
	return nil
}

func (is *InventoryServicesImpl) GetProduct(inventoryId, productId primitive.ObjectID) (*models.Product, error) {
	// Mongo Queries
	filterInventory := bson.D{bson.E{Key: "_id", Value: inventoryId}}

	// Find inventory
	var inventory models.Inventory
	if err := is.InventoryCollection.FindOne(is.Ctx, filterInventory).Decode(&inventory); err != nil {
		return nil, err
	}

	// Search for Product
	for _, p := range inventory.Products {
		if p.Hex() == productId.Hex() {
			filterProduct := bson.D{bson.E{Key: "_id", Value: p}}
			var product models.Product
			if err := is.ProductCollection.FindOne(is.Ctx, filterProduct).Decode(&product); err != nil {
				return nil, err
			}
			return &product, nil
		}
	}

	return nil, errors.New("No Product Found.")
}
