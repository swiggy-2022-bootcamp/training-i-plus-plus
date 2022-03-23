package services

import "github.com/go-kafka-microservice/InventoryService/models"

type InventoryServices interface {
	RegisterInventory(*models.Inventory) error
	AddProduct(*models.Product) error
	GetProduct(int, *models.Product) (*models.Product, error)
}
