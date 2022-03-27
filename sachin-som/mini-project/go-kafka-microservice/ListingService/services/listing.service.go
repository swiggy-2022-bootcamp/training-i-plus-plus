package services

import "github.com/go-kafka-microservice/ListingService/models"

type ListingService interface {
	ShowProducts() ([]*models.Product, error)
}
