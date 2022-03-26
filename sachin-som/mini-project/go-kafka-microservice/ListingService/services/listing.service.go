package services

import (
	"github.com/go-kafka-microservice/ListingService/models"
)

type ListingService interface {
	ShowProducts(string) ([]models.Product, error)
}
