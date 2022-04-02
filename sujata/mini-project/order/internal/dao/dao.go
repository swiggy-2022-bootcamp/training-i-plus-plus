package mongodao

import (
	"context"
	"fmt"
	"order/config"
	model "order/internal/dao/models"
	"order/internal/errors"
	"sync"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDAO interface {
	CreateOrder(ctx context.Context, cartProduct model.Order) (interface{}, *errors.ServerError)
	GetOrders(ctx context.Context, email string) (model.AllOrders, *errors.ServerError)
	SetOrderStatus(ctx context.Context, orderInfo model.OrderInfo) *errors.ServerError
}

type mongoDAO struct {
	client *mongo.Client
	config *config.WebServerConfig
}

var mongoDao MongoDAO
var mongoDaoOnce sync.Once

func InitMongoDAO(client *mongo.Client, config *config.WebServerConfig) MongoDAO {
	mongoDaoOnce.Do(func() {
		mongoDao = &mongoDAO{
			client: client,
			config: config,
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

// AddProduct add product details to the cart of the user
func (dao *mongoDAO) CreateOrder(ctx context.Context, order model.Order) (interface{}, *errors.ServerError) {
	cartCollection := dao.client.Database(dao.config.Db).Collection(dao.config.DbCollection)

	ra, err := cartCollection.InsertOne(ctx, order)
	if err != nil {
		log.WithError(err).Error("an error occured while inserting a new order in database")
		return nil, &errors.DatabaseInsertionError
	}

	if ra == nil {
		log.WithError(err).Error("user inserted: 0, expected 1")
		return nil, &errors.DatabaseNoInsertionError
	}

	log.Info("order created successfully")

	return ra.InsertedID, nil
}

// GetOrders get all the orders related to a particular user
func (dao *mongoDAO) GetOrders(ctx context.Context, email string) (model.AllOrders, *errors.ServerError) {
	cartCollection := dao.client.Database(dao.config.Db).Collection(dao.config.DbCollection)
	allOrders := model.AllOrders{}

	filter := bson.M{"email": email}
	cursor, err := cartCollection.Find(ctx, filter)
	if err != nil {
		log.WithError(err).Error("an error occurred while getting all the orders")
		return allOrders, &errors.InternalError
	}

	for cursor.Next(context.TODO()) {
		var order model.UserOrder
		if err := cursor.Decode(&order); err != nil {
			log.WithError(err).Error("an error occurred while decoding the cursor element")
			return allOrders, &errors.InternalError
		}
		fmt.Println(order)
		allOrders.Orders = append(allOrders.Orders, order)
	}

	if err := cursor.Err(); err != nil {
		log.WithError(err).Error("error from db cursor")
		return allOrders, &errors.InternalError
	}

	return allOrders, nil
}

func (dao *mongoDAO) SetOrderStatus(ctx context.Context, orderInfo model.OrderInfo) *errors.ServerError {
	cartCollection := dao.client.Database(dao.config.Db).Collection(dao.config.DbCollection)

	objID, _ := primitive.ObjectIDFromHex(orderInfo.OrderId)
	update := bson.M{
		"$set": bson.M{
			"orderStatus": orderInfo.OrderStatus,
		},
	}

	ra, err := cartCollection.UpdateByID(ctx, objID, update)
	if err != nil {
		log.WithError(err).Error("an error occurred while updating the order status")
		return &errors.InternalError
	}

	if ra.ModifiedCount == 0 {
		log.Error("Modified count is 0, but expected 1")
		return &errors.InternalError
	}

	return nil
}
