package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const DBUrl = "mongodb://127.0.0.1:27017/crud_test"

func ConnectDB() {

	fmt.Println("Connecting to ", DBUrl)
	mongo, err := mgo.ParseURL(DBUrl)
	s, err := mgo.Dial(DBUrl)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", DBUrl)
	Session = s
	Mongo = mongo
}
