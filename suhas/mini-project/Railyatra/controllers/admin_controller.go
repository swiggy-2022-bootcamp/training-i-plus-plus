package controllers

import (
	"context"
	"gin-mongo-api/config"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var adminCollection *mongo.Collection = config.GetCollection(config.DB, "admins")
var trainCollection *mongo.Collection = config.GetCollection(config.DB, "trains")
var availticketCollection *mongo.Collection = config.GetCollection(config.DB, "availtickets")
var avalidate = validator.New()

const layout = "Jan 2, 2006 at 3:04pm (MST)"

func CreateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var admin models.Admin
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&admin); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newAdmin := models.Admin{
			Name:  admin.Name,
			Email: admin.Email,
		}

		result, err := adminCollection.InsertOne(ctx, newAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AdminResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		var admin models.Admin
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		err := adminCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&admin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admin}})
	}
}

func EditAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		var admin models.Admin
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		//validate the request body
		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&admin); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"name": admin.Name, "email": admin.Email}
		result, err := adminCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated admin details
		var updatedAdmin models.Admin
		if result.MatchedCount == 1 {
			err := adminCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedAdmin)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedAdmin}})
	}
}

func DeleteAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		result, err := adminCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.AdminResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Admin with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Admin successfully deleted!"}},
		)
	}
}

func GetAllAdmins() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var admins []models.Admin
		defer cancel()

		results, err := adminCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleAdmin models.Admin
			if err = results.Decode(&singleAdmin); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			admins = append(admins, singleAdmin)
		}

		c.JSON(http.StatusOK,
			responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admins}},
		)
	}
}

func CreateTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var train models.Train
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newTrain := models.Train{
			Station1: train.Station1,
			Station2: train.Station2,
		}

		result, err := trainCollection.InsertOne(ctx, newTrain)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.TrainResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		var train models.Train
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		err := trainCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&train)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": train}})
	}
}

func EditTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		var train models.Train
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"station1": train.Station1, "station2": train.Station2}
		result, err := trainCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated train details
		var updatedTrain models.Train
		if result.MatchedCount == 1 {
			err := trainCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedTrain)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedTrain}})
	}
}

func DeleteTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		result, err := trainCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.TrainResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Train with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Train successfully deleted!"}},
		)
	}
}

func GetAllTrains() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var trains []models.Train
		defer cancel()

		results, err := trainCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTrain models.Train
			if err = results.Decode(&singleTrain); err != nil {
				c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			trains = append(trains, singleTrain)
		}

		c.JSON(http.StatusOK,
			responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": trains}},
		)
	}
}

func CreateAvailTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var availticket models.AvailTicket
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&availticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&availticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		_, err1 := time.Parse(layout, availticket.Departure_time)
		_, err := time.Parse(layout, availticket.Arrival_time)

		if err != nil || err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "time not in correct format", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		newAvailTicket := models.AvailTicket{
			Train_id:       availticket.Train_id,
			Capacity:       availticket.Capacity,
			Price:          availticket.Price,
			Departure:      availticket.Departure,
			Arrival:        availticket.Arrival,
			Departure_time: availticket.Departure_time,
			Arrival_time:   availticket.Arrival_time,
		}

		result, err := availticketCollection.InsertOne(ctx, newAvailTicket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AvailTicketResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAvailTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		availticketId := c.Param("availticketid")
		var availticket models.AvailTicket
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(availticketId)

		err := availticketCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&availticket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AvailTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": availticket}})
	}
}

func EditAvailTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		availticketId := c.Param("availticketid")
		var availticket models.AvailTicket
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(availticketId)

		//validate the request body
		if err := c.BindJSON(&availticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.AvailTicketResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&availticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AvailTicketResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"departure": availticket.Departure, "arrival": availticket.Arrival}
		result, err := availticketCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated availticket details
		var updatedAvailTicket models.AvailTicket
		if result.MatchedCount == 1 {
			err := availticketCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedAvailTicket)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.AvailTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedAvailTicket}})
	}
}

func DeleteAvailTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		availticketId := c.Param("availticketid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(availticketId)

		result, err := availticketCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.AvailTicketResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "AvailTicket with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.AvailTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "AvailTicket successfully deleted!"}},
		)
	}
}

func GetAllAvailTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var availtickets []models.AvailTicket
		defer cancel()

		results, err := availticketCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleAvailTicket models.AvailTicket
			if err = results.Decode(&singleAvailTicket); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AvailTicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			availtickets = append(availtickets, singleAvailTicket)
		}

		c.JSON(http.StatusOK,
			responses.AvailTicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": availtickets}},
		)
	}
}
