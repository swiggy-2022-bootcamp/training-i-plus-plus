package services

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-kafka-microservice/ListingService/models"
)

type ListingServiceImpl struct {
	KafkaConsumer *kafka.Consumer
	Ctx           context.Context
}

func NewListingServiceImpl(kafkaConsumer *kafka.Consumer, ctx context.Context) *ListingServiceImpl {
	return &ListingServiceImpl{
		KafkaConsumer: kafkaConsumer,
		Ctx:           ctx,
	}
}

// Method to show all inventory products
func (ls *ListingServiceImpl) ShowProducts() ([]*models.Product, error) {
	return nil, nil
}
