package services

import (
	"github.com/go-kafka-microservice/ListingService/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListingService interface {
	ShowProducts() ([]*models.Product, error)
	MakeOrder(primitive.ObjectID, primitive.ObjectID) error
	AuthorizeUser(string) (string, error)
}
