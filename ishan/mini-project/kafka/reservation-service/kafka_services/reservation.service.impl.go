package services

import (
	"context"
	kf "swiggy/gin/goKafka/producer"
	db "swiggy/gin/lib/utils"
	rsv "swiggy/gin/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	db.ConnectDB()
}

// Method to update Seat Matrix depending upon reservation and cancelation
func UpdateSeatMatrix(reservationId, userId primitive.ObjectID, seat int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	reservation := &rsv.ReserveType{}
	filter := bson.D{bson.E{Key: "_id", Value: reservationId}}
	if err := db.DataStore.Collection("reservation").FindOne(ctx, filter).Decode(&reservation); err != nil {
		return err
	}

	// Send reservations details to train service
	reservationBody := rsv.ReserveType{
		ID:              primitive.NewObjectID(),
		Destination:     reservation.Destination,
		Train:           reservation.Train,
		User:            reservation.User,
		Cost:            reservation.Cost,
		DateOfJourney:   reservation.DateOfJourney,
		BoardingStation: reservation.BoardingStation,
		Seat:            seat,
	}
	if _, err := kf.WriteMessage(ctx, "reservations", reservationBody); err != nil {
		return err
	}
	return nil
}
