package controllers

import (
	"context"
	"net/http"
	"time"
	"trainService/database"
	helper "trainService/helpers"
	models "trainService/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var trainCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "train")

// ShowAccount godoc
// @Summary      Check Availability
// @Description  check whether train is available on particular date or not
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        TrainID  body 	string  true  "unique train id"
// @Param        Date 	  body	string  true  "date of booking"
// @Success      200  {object} 	models.Train
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /train/checkAvailability [post]
func CheckAvailability() gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
		defer cancel()

		var search struct {
			TrainID string
			Date    string
		}

		if err := g.BindJSON(&search); err != nil {
			g.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		var trainDetails models.Train

		err := trainCollection.FindOne(ctx, bson.M{"TrainID": search.TrainID}).Decode(&trainDetails)

		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "no data found"})
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
		isAvailable := helper.IsTrainAvailableOnGivenWeekDay(weekday, trainDetails)
		if !isAvailable {
			g.JSON(http.StatusBadRequest, gin.H{"msg": "train not available"})
			return
		}
		responseTrainDetails := helper.CalculatePriceOne(
			trainDetails.FromStationCode,
			trainDetails.ToStationCode,
			trainDetails,
		)

		g.JSON(http.StatusOK, gin.H{"trainDetails": responseTrainDetails})
	}
}

type SearchQuery struct {
	Source      string
	Destination string
	Date        string
}

// ShowAccount godoc
// @Summary      Search route from source to destination stations
// @Description  search route from source to destination stations on specific date
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        Source  		body 	string  true  "source station code"
// @Param        Destination 	body	string  true  "destination station code"
// @Param        Date 			body	string  true  "date of booking"
// @Success      200  {array} 	models.Train
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /train/searchRoute [post]
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

		responseTrainDetails = helper.CalculatePriceMany(
			search.Source,
			search.Destination,
			trainDetails,
		)

		g.JSON(http.StatusOK, gin.H{"trainDetails": responseTrainDetails})
	}
}

// ShowAccount godoc
// @Summary      Get details of sepecific train with its number
// @Description  Get details of sepecific train with its number
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        TrainNumber  body 	string  true  "unique train number"
// @Success      200  {object} 	models.Train
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /train/trainDetails [post]
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
