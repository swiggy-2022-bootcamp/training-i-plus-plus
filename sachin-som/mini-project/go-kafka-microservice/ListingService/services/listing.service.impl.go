package services

import (
	"context"

	goKafka "github.com/go-kafka-microservice/ListingService/goKafka/consumer"
	"github.com/go-kafka-microservice/ListingService/models"
)

type ListingServiceImpl struct {
	KafkaService goKafka.GoKafkaServices
	Ctx          context.Context
}

func NewListingServiceImpl(kafkaServices goKafka.GoKafkaServices, ctx context.Context) *ListingServiceImpl {
	return &ListingServiceImpl{
		KafkaService: kafkaServices,
		Ctx:          ctx,
	}
}

// Method to show all inventory products
func (ls *ListingServiceImpl) ShowProducts(topic string) ([]models.Product, error) {
	products, err := ls.KafkaService.ReadMessage(topic)
	if err != nil {
		return nil, err
	}
	return products.([]models.Product), err
}
