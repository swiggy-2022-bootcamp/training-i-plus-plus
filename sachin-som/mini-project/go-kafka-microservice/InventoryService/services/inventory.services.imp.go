package services

import (
	"context"

	"github.com/go-kafka-microservice/InventoryService/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryServicesImpl struct {
	InventoryCollection *mongo.Collection
	Ctx                 context.Context
}

func NewInventoryService(inventoryCollection *mongo.Collection, ctx context.Context) *InventoryServicesImpl {
	return &InventoryServicesImpl{
		InventoryCollection: inventoryCollection,
		Ctx:                 ctx,
	}
}

func (is *InventoryServicesImpl) RegisterInventory(proudct *models.Inventory) error {
	return nil
}

func (is *InventoryServicesImpl) AddProduct(inventoryId int, product *models.Product) error {
	return nil
}

func (is *InventoryServicesImpl) GetProduct(inventoryId int, productId int) (*models.Product, error) {
	return nil, nil
}
