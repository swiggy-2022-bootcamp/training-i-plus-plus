package train

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

type TrainBody struct {
	Name        string   `json:"name"`
	Number      string   `json:"number"`
	Destination string   `json:"destination"`
	Source      string   `json:"source"`
	Stations    []string `json:"stations"`
}

func createNewTrain(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body := TrainBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect Body"})
		return
	}
	train := Train{
		ID:          primitive.NewObjectID(),
		Destination: body.Destination,
		Source:      body.Source,
		Number:      body.Number,
		Stations:    body.Stations,
		Name:        body.Name,
	}
	train.ID = primitive.NewObjectID()
	res, err := db.DataStore.Collection("train").InsertOne(ctx, train)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("can not convert to oid %v", err)})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"name": train.Name, "Id": oid.Hex()})
}

func getTrainsByCondition(condition bson.D) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	matchStage := bson.D{{"$match", condition}}
	cursor, err := db.DataStore.Collection("train").Aggregate(ctx, mongo.Pipeline{matchStage})
	if err != nil {
		return nil, err
	}
	var trains []bson.M

	if err = cursor.All(ctx, &trains); err != nil {
		return nil, err
	}

	return trains, nil
}

func FetchTrains(c *gin.Context) {
	name := c.Query("name")
	number := c.Query("number")
	var condition bson.D

	if name != "" {
		condition = append(condition, bson.E{Key: "name", Value: name})
	}

	if number != "" {
		condition = append(condition, bson.E{Key: "number", Value: number})
	}

	res, err := getTrainsByCondition(condition)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"payload": res})
}
