package repository

import (
	"context"
	"fmt"
	"srctc/database"
	"srctc/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionTicketName = "ticket"
	collectionTicket     = new(mongo.Collection)
)

func init() {
	var DB *mongo.Client = database.ConnectDB()
	collectionTicket = DB.Database("golangAPI").Collection(collectionTicketName)
}

type TicketRepository struct{}

func (avl TicketRepository) Insert(newTicket models.Ticket) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionTicket.InsertOne(ctx, &newTicket)
	if err == nil {
		fmt.Println("Inserted a single document: ", result.InsertedID)
	}
	return result.InsertedID, err
}

func (avl TicketRepository) Read(objId primitive.ObjectID) (models.Ticket, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.Ticket
	err := collectionTicket.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	return user, err
}

func (avl TicketRepository) ReadTrainId(objId primitive.ObjectID) (models.Ticket, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.Ticket
	err := collectionTicket.FindOne(ctx, bson.M{"trainid": objId}).Decode(&user)
	return user, err
}

func (avl TicketRepository) Update(updateTicket models.Ticket, objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//update := bson.M{"departure": updateTicket.Departure, "arrival": updateTicket.Arrival}
	updatebson := bson.M{}
	if updateTicket.Departure != "" {
		updatebson["departure"] = updateTicket.Departure
	}
	if updateTicket.Arrival != "" {
		updatebson["arrival"] = updateTicket.Arrival
	}
	if updateTicket.Capacity != 0 {
		updatebson["capacity"] = updateTicket.Capacity
	}
	result, err := collectionTicket.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updatebson})
	if err == nil {
		fmt.Println("Updated a single document: ", result.UpsertedID)
	}
	return result.UpsertedID, err
}

func (avl TicketRepository) Delete(objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionTicket.DeleteOne(ctx, bson.M{"_id": objId})
	if err == nil {
		fmt.Println("Updated a single document: ", result.DeletedCount)
	}
	return result.DeletedCount, err
}

func (avl TicketRepository) ReadAll() ([]models.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var users []models.Ticket
	defer cancel()

	results, err := collectionTicket.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleTicket models.Ticket
		if err = results.Decode(&singleTicket); err != nil {
			return users, err
		}

		users = append(users, singleTicket)
	}
	return users, nil
}
