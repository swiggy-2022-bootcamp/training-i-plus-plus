package services

import (
	"github.com/sachinsom93/shopping-cart/models"
)

type InventoryServices interface {
	RegisterInventory(*models.Inventory) error
	// DeactivateInventory(primitive.ObjectID) error
	AddProduct(*models.Product) error
	RemoveProduct(string, string) error
	UpdateProduct(*models.Product) (*models.Product, error)
	GetProduct(string, string) (*models.Product, error)
	GetAllProducts(string) ([]*models.Product, error)
}
