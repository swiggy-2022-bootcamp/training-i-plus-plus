package mongodao

import (
	"context"
	model "product/internal/dao/mongodao/models"
	"product/internal/errors"
	"sync"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDAO interface {
	AddProduct(ctx context.Context, user model.Product) *errors.ServerError
	GetProduct(ctx context.Context, productId string) (model.Product, *errors.ServerError)
}

type mongoDAO struct {
	client *mongo.Client
}

var mongoDao MongoDAO
var mongoDaoOnce sync.Once

func InitMongoDAO(client *mongo.Client) MongoDAO {
	mongoDaoOnce.Do(func() {
		mongoDao = &mongoDAO{
			client: client,
		}
	})

	return mongoDao
}

func GetMongoDAO() MongoDAO {
	if mongoDao == nil {
		panic("Mongo DAO not initialised")
	}

	return mongoDao
}

func (dao *mongoDAO) AddProduct(ctx context.Context, product model.Product) *errors.ServerError {
	productCollection := dao.client.Database("shopKart").Collection("products")
	ra, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		log.WithError(err).Error("an error occured while inserting a new product in database")
		return &errors.DatabaseInsertionError
	}

	if ra == nil {
		log.WithError(err).Error("product inserted: 0, expected 1")
		return &errors.DatabaseNoInsertionError
	}

	log.Info("product created successfully")
	return nil
}

func (dao *mongoDAO) GetProduct(ctx context.Context, productId string) (model.Product, *errors.ServerError) {
	productCollection := dao.client.Database("shopKart").Collection("products")

	objID, _ := primitive.ObjectIDFromHex(productId)
	filter := bson.M{"_id": objID}
	result := productCollection.FindOne(ctx, filter)

	var product model.Product
	err := result.Decode(&product)
	if err != nil {
		log.WithError(err).Error("an error occurred while decoding the product")
		return product, &errors.InternalError
	}

	return product, nil
}
