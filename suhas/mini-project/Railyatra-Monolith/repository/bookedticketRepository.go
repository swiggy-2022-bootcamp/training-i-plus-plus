package repository

import (
	"context"
	"fmt"
	"gin-mongo-api/config"
	"gin-mongo-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionBookedTicketName = "availticket"
	collectionBookedTicket     = new(mongo.Collection)
)

func init() {
	var DB *mongo.Client = config.ConnectDB()
	collectionBookedTicket = DB.Database("golangAPI").Collection(collectionBookedTicketName)
}

type BookedTicketRepository struct{}

func (btk BookedTicketRepository) Insert(newBookedTicket models.BookedTicket) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionBookedTicket.InsertOne(ctx, &newBookedTicket)
	if err == nil {
		fmt.Println("Inserted a single document: ", result.InsertedID)
	}
	return result.InsertedID, err
}

func (btk BookedTicketRepository) Read(objId primitive.ObjectID) (models.BookedTicket, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.BookedTicket
	err := collectionBookedTicket.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	return user, err
}

func (btk BookedTicketRepository) Delete(objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionBookedTicket.DeleteOne(ctx, bson.M{"_id": objId})
	if err == nil {
		fmt.Println("Updated a single document: ", result.DeletedCount)
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
			return users, err
		}

		users = append(users, singleBookedTicket)
	}
	return users, nil
}
