package db

import (
	"gopkg.in/mgo.v2"
	"os"
	"sample.akash.com/log"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the database.
	MongoDBUrl = "mongodb://localhost:2717/shopping_cart_dev"
)

// Connect connects to mongodb
func Connect() {
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

// docker run --rm -d -p 2717:27017 -v /Users/aky/Dev/swiggy/training-i-plus-plus/akash/mini-project/customer/db/mongo-files:/data/db --name mymongo mongo:latest

// docker run --rm -d -p 2717:27017 -v /Users/aky/Dev/swiggy/training-i-plus-plus/akash/training/gin-demo/mongodb-docker:/data/db --name mymongo mongo:latest
// https://medium.com/@arkamukherjee/a-guide-on-migrating-mgo-apis-to-mongo-go-driver-in-a-golang-server-32a4e6f0fc5e
// https://riptutorial.com/go/example/27707/example
