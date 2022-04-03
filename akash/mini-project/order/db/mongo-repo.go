package db

import (
	"gopkg.in/mgo.v2"
	"order.akash.com/log"
	"order.akash.com/model"
	"os"
)

const (
	MongoDBUrl = "mongodb://localhost:2719/shopping_cart_dev_order"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

type repo struct{}

func NewMongoRepository() OrderRepository {
	return &repo{}
}

func (*repo) Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		log.Error("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	log.Info("Connected to", uri)
	Session = s
	Mongo = mongo
}

func (*repo) FindAll() []model.Order {

	c := Session.DB("shopping_cart_dev_order").C("order-collection")

	log.Info("Going to find all orders")
	var results []model.Order
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Error("Error while querying all orders ", err)
		panic(err)
	} else {
		log.Info("Found orders : ", results)
	}

	return results
}

func (*repo) SaveOrder(order model.Order) {

	c := Session.DB("shopping_cart_dev_order").C("order-collection")

	log.Info("Trying to save user : ", order)

	if err := c.Insert(order); err != nil {
		log.Error("Error while saving order ", err)
		panic(err)
	}

	log.Info("Order added ")
}
