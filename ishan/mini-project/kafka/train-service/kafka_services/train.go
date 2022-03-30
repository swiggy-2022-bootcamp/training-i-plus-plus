package services

import (
	"context"
	db "swiggy/gin/lib/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	db.ConnectDB()
}

type TrainData struct {
	Train string `json:"Train"`
	Seat  int32  `json:"Seat"`
}

// Method to Send Product to ordered_product topic
func UpdateTrainInfo(trainId string, seat int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	TrainID, err := primitive.ObjectIDFromHex(trainId)
	if err != nil {
		return err
	}

	if _, err := db.DataStore.Collection("train").UpdateByID(ctx, TrainID, bson.D{{"$inc", bson.D{{"seats", seat}}}}); err != nil {
		return err
	}

	return nil
}
