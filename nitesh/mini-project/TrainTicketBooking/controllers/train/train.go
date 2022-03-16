package controller

import (
	"context"
	"net/http"
	"rail/database"
	helper "rail/helpers/train"
	models "rail/models/train"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SearchQuery struct {
	Source      string
	Destination string
	Date        string
}

var trainCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "train")

func SearchRoute() gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		var search SearchQuery

		if err := g.BindJSON(&search); err != nil {
			g.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		cursor, err := trainCollection.Find(
			ctx,
			bson.D{{"Stations", bson.D{{"$all", bson.A{search.Source, search.Destination}}}}},
		)
		// Decode(&trainDetails)

		if err != nil {
			g.JSON(http.StatusOK, gin.H{"error": "no routes"})
			return
		}

		var trainDetails []models.Train

		if err := cursor.All(ctx, &trainDetails); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//-------------check for date availability
		layout := "01/02/06"
		parseDate, err := time.Parse(layout, search.Date)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if time.Now().After(parseDate) {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid date"})
			return
		}

		weekday := parseDate.Weekday().String()
		responseTrainDetails := helper.FilterDetailsOnWeekdayAwailability(weekday, trainDetails)

		responseTrainDetails = helper.CalculatePrice(
			search.Source,
			search.Destination,
			trainDetails,
		)

		g.JSON(http.StatusOK, gin.H{"trainDetails": responseTrainDetails})
	}
}

func TrainDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel()

		var body struct {
			TrainNumber string
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var TrainDetails models.Train
		err := trainCollection.FindOne(ctx, bson.M{"TrainNumber": body.TrainNumber}).
			Decode(&TrainDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"trainDetails": TrainDetails})
	}
}
