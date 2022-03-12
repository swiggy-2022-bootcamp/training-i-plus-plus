package services

import (
	"github.com/sachinsom93/shopping-cart/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductServices interface {
	RegisterInventory(*models.Inventory) error
	DeactivateInventory(primitive.ObjectID) error
	AddProduct(*models.Product) error
	RemoveProduct(primitive.ObjectID) error
	UpdateProduct(*models.Product) error
	getProduct(primitive.ObjectID) (*models.Product, error)
	getProducts(primitive.ObjectID) ([]*models.Product, error)
}
