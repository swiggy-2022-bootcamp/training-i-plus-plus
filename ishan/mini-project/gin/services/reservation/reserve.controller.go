package reservation

import (
	"context"
	"fmt"
	"net/http"
	db "swiggy/gin/lib/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReserveBody struct {
	Train           string `json:"train"`
	Cost            int64  `json:"cost"`
	DateOfJourney   string `json:"dateOfJourney"`
	BoardingStation string `json:"boardingStation"`
	Destination     string `json:"destination"`
}

func ReserveSeat(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	userId := c.GetString("User")
	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	body := ReserveBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect Body"})
		return
	}

	TrainID, err := primitive.ObjectIDFromHex(body.Train)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	reservation := ReserveType{
		ID:              primitive.NewObjectID(),
		Destination:     body.Destination,
		Train:           TrainID,
		User:            objID,
		Cost:            body.Cost,
		DateOfJourney:   body.DateOfJourney,
		BoardingStation: body.BoardingStation,
	}
	res, err := db.DataStore.Collection("reservation").InsertOne(ctx, reservation)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("can not convert to oid %v", err)})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Id": oid.Hex()})

}

func FetchReservations(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	userId := c.GetString("User")

	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	reservationMatchStage := bson.D{{"$match", bson.D{{"user", objID}}}}
	userLookupStage := bson.D{{"$lookup", bson.D{{"from", "user"}, {"localField", "user"}, {"foreignField", "_id"}, {"as", "userInfo"}}}}
	userUnwindStage := bson.D{{"$unwind", bson.D{{"path", "$userInfo"}, {"preserveNullAndEmptyArrays", false}}}}
	trainLookupStage := bson.D{{"$lookup", bson.D{{"from", "train"}, {"localField", "train"}, {"foreignField", "_id"}, {"as", "trainInfo"}}}}
	trainUnwindStage := bson.D{{"$unwind", bson.D{{"path", "$trainInfo"}, {"preserveNullAndEmptyArrays", false}}}}

	cursor, err := db.DataStore.Collection("reservation").Aggregate(ctx, mongo.Pipeline{reservationMatchStage, userLookupStage, userUnwindStage, trainLookupStage, trainUnwindStage})

	var reservations []ReservationInfoType

	if err = cursor.All(ctx, &reservations); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"payload": reservations})
}
