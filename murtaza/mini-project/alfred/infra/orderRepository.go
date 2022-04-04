package infra

import (
	"alfred/domain"
	"alfred/utils/errs"
	"alfred/utils/logger"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const ordersDBUrl = "mongodb://127.0.0.1:27019/orders_db"

type OrderRepository struct {
	Session *mgo.Session
	Mongo   *mgo.DialInfo
}

func (or OrderRepository) InsertOrder(order domain.Order) (*domain.Order, *errs.AppError) {
	orders := or.Session.DB(or.Mongo.Database).C(OrderCollectionName)
	persistedOrder := or.toPersistedEntity(order)
	if err := orders.Insert(persistedOrder); err != nil {
		logger.Error("Failed to create Order", zap.Error(err))
		return nil, errs.NewUnexpectedError(err.Error())
	}
	logger.Info(fmt.Sprintf("New Order Created Succesfully with Order Id: %d", persistedOrder.Id))
	return persistedOrder.toDomainEntity(), nil
}

func NewOrderRepository() OrderRepository {

	fmt.Println("Connecting to ", ordersDBUrl)
	mongo, err := mgo.ParseURL(ordersDBUrl)
	s, err := mgo.Dial(ordersDBUrl)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Can't connect to mongo, go error %v\n", err), zap.Error(err))
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	logger.Info(fmt.Sprintf("Connected to: %s", ordersDBUrl))

	return OrderRepository{
		Session: s,
		Mongo:   mongo,
	}
}

func (or OrderRepository) toPersistedEntity(o domain.Order) *Order {
	var nextId = or.getNextSequence("orderId")
	return &Order{
		Id:          nextId,
		UserId:      o.UserId,
		Items:       o.Items,
		OrderAmount: o.OrderAmount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (or OrderRepository) getNextSequence(seqName string) int {

	type sequenceDoc struct {
		Id            string `bson:"_id"`
		SequenceValue int    `bson:"sequence_value"`
	}

	var seq sequenceDoc
	counters := or.Session.DB(or.Mongo.Database).C(CountersCollectionName)

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"sequence_value": 1}},
		ReturnNew: true,
		Upsert:    true,
	}

	_, err := counters.Find(bson.M{"_id": seqName}).Apply(change, &seq)
	if err != nil {
		return 0
	} else {
		return seq.SequenceValue
	}
}
