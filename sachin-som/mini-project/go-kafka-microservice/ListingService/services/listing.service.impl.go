package services

import (
	"context"
	"errors"

	pb "github.com/go-kafka-microservice/AuthProto"
	gokafkaConsumer "github.com/go-kafka-microservice/ListingService/goKafka/consumer"
	gokafkaProducer "github.com/go-kafka-microservice/ListingService/goKafka/producer"
	"github.com/go-kafka-microservice/ListingService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListingServiceImpl struct {
	KafkaConsumerService gokafkaConsumer.GoKafkaServices
	kafkaProducerService gokafkaProducer.GoKafkaServices
	ProductCollection    *mongo.Collection
	AuthProtoClient      pb.AuthServicesClient
	Ctx                  context.Context
}

func NewListingServiceImpl(kafkaConsumerService gokafkaConsumer.GoKafkaServices, kafkaProducerService gokafkaProducer.GoKafkaServices, productCollection *mongo.Collection, authProtoClient pb.AuthServicesClient, ctx context.Context) *ListingServiceImpl {
	return &ListingServiceImpl{
		kafkaProducerService: kafkaProducerService,
		KafkaConsumerService: kafkaConsumerService,
		ProductCollection:    productCollection,
		AuthProtoClient:      authProtoClient,
		Ctx:                  ctx,
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

// Method to Send Product to ordered_product topic
func (ls *ListingServiceImpl) MakeOrder(productId, ownerId primitive.ObjectID) error {
	var product models.Product
	filter := bson.D{bson.E{Key: "_id", Value: productId}}
	if err := ls.ProductCollection.FindOne(ls.Ctx, filter).Decode(&product); err != nil {
		return err
	}

	// Send product to ordered_product kafka topic
	userProduct := models.UserProduct{
		ID:          primitive.NewObjectID(),
		ProductName: product.ProductName,
		Description: product.Description,
		Ratings:     product.Ratings,
		Price:       product.Price,
		ImageUrl:    product.ImageUrl,
		UserID:      ownerId,
	}
	if _, err := ls.kafkaProducerService.WriteMessage("ordered_products", userProduct); err != nil {
		return err
	}
	return nil
}

func (ls *ListingServiceImpl) AuthorizeUser(tokenStr string) (string, error) {
	res, err := ls.AuthProtoClient.Authorize(ls.Ctx, &pb.TokenRequest{
		Token: tokenStr,
	})
	if err != nil {
		return "", err
	}
	return res.Token, nil
}
