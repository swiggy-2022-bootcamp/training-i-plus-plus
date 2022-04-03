package repository

import (
	"context"
	"time"
	"userService/config"
	log "userService/logger"
	"userService/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionBookedTicketName = "bookedtickets"
	collectionBookedTicket     = new(mongo.Collection)
	errLog                     = log.ErrorLogger.Println
	warLog                     = log.WarningLogger.Println
	infLog                     = log.InfoLogger.Println
)

func init() {
	var DB *mongo.Client = config.ConnectDB()
	collectionBookedTicket = DB.Database("golangAPI").Collection(collectionBookedTicketName)
}

type BookedTicketRepository struct{}

func (btk BookedTicketRepository) Insert(newBookedTicket models.BookedTicket) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionBookedTicket.InsertOne(ctx, &newBookedTicket)
	if err != nil {
		errLog(err)
	}
	return result.InsertedID, err
}

func (btk BookedTicketRepository) Read(objId primitive.ObjectID) (models.BookedTicket, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.BookedTicket
	err := collectionBookedTicket.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		errLog(err)
	}
	return user, err
}

func (btk BookedTicketRepository) Delete(objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionBookedTicket.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		errLog(err)
	}
	return result.DeletedCount, err
}

func (btk BookedTicketRepository) ReadAll() ([]models.BookedTicket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var users []models.BookedTicket
	defer cancel()
	results, err := collectionBookedTicket.Find(ctx, bson.M{})
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleBookedTicket models.BookedTicket
		if err = results.Decode(&singleBookedTicket); err != nil {
			errLog(err)
			return users, err
		}

		users = append(users, singleBookedTicket)
	}
	return users, nil
}
