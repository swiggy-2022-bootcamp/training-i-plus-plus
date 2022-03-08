package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	PlotNumber int    `json:"plotnumber" bson:"plotnumber"`
	LandMark   string `json:"landmark" bson:"landmark"` // TODO: need to make it optional
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	Country    string `json:"country" bson:"country"`
}

type Customer struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstname" bson:"firstname"`
	LastName  string        `json:"lastname" bson:"lastname"`
	Address
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
