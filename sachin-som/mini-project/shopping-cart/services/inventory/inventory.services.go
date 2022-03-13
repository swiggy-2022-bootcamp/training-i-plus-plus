package services

import (
	"github.com/sachinsom93/shopping-cart/models"
)

type InventoryServices interface {
	RegisterInventory(*models.Inventory) error
	// DeactivateInventory(primitive.ObjectID) error
	AddProduct(*models.Product) error
	RemoveProduct(int) error
	UpdateProduct(*models.Product) (*models.Product, error)
	GetProduct(int) (*models.Product, error)
	GetAllProducts(int) ([]*models.Product, error)
}
