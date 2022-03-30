package models

import (
	"swiggy/gin/services/train"
	"swiggy/gin/services/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReserveType struct {
	ID              primitive.ObjectID `bson:"_id",omitempty`
	Train           primitive.ObjectID `bson:"train"`
	User            primitive.ObjectID `bson:"user"`
	Cost            int64              `bson:"cost"`
	DateOfJourney   string             `bson:"dateOfJourney"`
	BoardingStation string             `bson:"boardingStation"`
	Destination     string             `bson:"destination"`
	Seat            int32              `bson:"seat"`
}

type ReservationInfoType struct {
	ID              primitive.ObjectID `bson:"_id",omitempty`
	Cost            int64              `bson:"cost"`
	DateOfJourney   string             `bson:"dateOfJourney"`
	BoardingStation string             `bson:"boardingStation"`
	Destination     string             `bson:"destination"`
	TrainInfo       train.Train        `bson:"trainInfo"`
	UserInfo        user.UserPublic    `bson:"userInfo"`
}
