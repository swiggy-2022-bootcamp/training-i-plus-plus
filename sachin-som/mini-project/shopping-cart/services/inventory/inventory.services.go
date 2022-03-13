package services

import (
	"github.com/sachinsom93/shopping-cart/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryServices interface {
	RegisterInventory(*models.Inventory) error
	// DeactivateInventory(primitive.ObjectID) error
	AddProduct(*models.Product) error
	RemoveProduct(primitive.ObjectID) error
	UpdateProduct(*models.Product) (*models.Product, error)
	GetProduct(primitive.ObjectID) (*models.Product, error)
	GetAllProducts(primitive.ObjectID) ([]*models.Product, error)
}
