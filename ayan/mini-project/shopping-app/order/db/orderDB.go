package db

import (
	"context"
	"time"

	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/domain"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/utils/errs"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/utils/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderRepositoryDB struct {
	dbClient *mongo.Client
}

func NewOrderRepositoryDB(dbClient *mongo.Client) domain.OrderRepositoryDB {
	return &orderRepositoryDB{
		dbClient: dbClient,
	}
}

func (pdb orderRepositoryDB) Save(u domain.Order) (*domain.Order, *errs.AppError) {

	newOrder := NewOrder(
		u.ItemList,
		u.Amount,
	)
	newOrder.Id = primitive.NewObjectID().Hex()

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	orderCollection := Collection(pdb.dbClient, "orders")
	_, err := orderCollection.InsertOne(ctx, newOrder)

	if err != nil {
		logger.Error("Error while inserting Order into DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}
	u.Id = newOrder.Id

	return &u, nil
}

func (pdb orderRepositoryDB) FetchOrderById(id string) (*domain.Order, *errs.AppError) {

	dbOrder := Order{}

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	orderCollection := Collection(pdb.dbClient, "orders")

	err := orderCollection.FindOne(ctx, bson.M{"id": id}).Decode(&dbOrder)

	if err != nil {
		logger.Error("Error while fetching Order from DB : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	domainOrder := domain.NewOrder(dbOrder.ItemList, dbOrder.Amount)
	domainOrder.Id = dbOrder.Id

	return domainOrder, nil
}
