package services

import (
	"context"
	"errors"

	goKafka "github.com/go-kafka-microservice/ListingService/goKafka/consumer"
	"github.com/go-kafka-microservice/ListingService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListingServiceImpl struct {
	KafkaService      goKafka.GoKafkaServices
	ProductCollection *mongo.Collection
	Ctx               context.Context
}

func NewListingServiceImpl(kafkaServices goKafka.GoKafkaServices, productCollection *mongo.Collection, ctx context.Context) *ListingServiceImpl {
	return &ListingServiceImpl{
		KafkaService:      kafkaServices,
		ProductCollection: productCollection,
		Ctx:               ctx,
	}
}

func (ls *ListingServiceImpl) ShowProducts() ([]*models.Product, error) {
	var products []*models.Product
	cursor, err := ls.ProductCollection.Find(ls.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ls.Ctx) {
		var _product models.Product
		err := cursor.Decode(&_product)
		if err != nil {
			return nil, err
		}
		products = append(products, &_product)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(ls.Ctx)
	if len(products) == 0 {
		return nil, errors.New("Products not fuond.")
	}
	return products, nil
}
