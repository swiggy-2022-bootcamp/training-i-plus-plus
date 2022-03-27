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
	collectionAvailTicketName = "availticket"
	collectionAvailTicket     = new(mongo.Collection)
)

func init() {
	var DB *mongo.Client = config.ConnectDB()
	collectionAvailTicket = DB.Database("golangAPI").Collection(collectionAvailTicketName)
}

type AvailTicketRepository struct{}

func (avl AvailTicketRepository) Insert(newAvailTicket models.AvailTicket) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionAvailTicket.InsertOne(ctx, &newAvailTicket)
	if err == nil {
		fmt.Println("Inserted a single document: ", result.InsertedID)
	}
	return result.InsertedID, err
}

func (avl AvailTicketRepository) Read(objId primitive.ObjectID) (models.AvailTicket, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.AvailTicket
	err := collectionAvailTicket.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	return user, err
}

func (avl AvailTicketRepository) ReadTrainId(objId primitive.ObjectID) (models.AvailTicket, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.AvailTicket
	err := collectionAvailTicket.FindOne(ctx, bson.M{"trainid": objId}).Decode(&user)
	return user, err
}

func (avl AvailTicketRepository) Update(updateAvailTicket models.AvailTicket, objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//update := bson.M{"departure": updateAvailTicket.Departure, "arrival": updateAvailTicket.Arrival}
	updatebson := bson.M{}
	if updateAvailTicket.Departure != "" {
		updatebson["departure"] = updateAvailTicket.Departure
	}
	if updateAvailTicket.Arrival != "" {
		updatebson["arrival"] = updateAvailTicket.Arrival
	}
	if updateAvailTicket.Capacity != 0 {
		updatebson["capacity"] = updateAvailTicket.Capacity
	}
	result, err := collectionAvailTicket.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updatebson})
	if err == nil {
		fmt.Println("Updated a single document: ", result.UpsertedID)
	}
	return result.UpsertedID, err
}

func (avl AvailTicketRepository) Delete(objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionAvailTicket.DeleteOne(ctx, bson.M{"_id": objId})
	if err == nil {
		fmt.Println("Updated a single document: ", result.DeletedCount)
	}
	return result.DeletedCount, err
}

func (avl AvailTicketRepository) ReadAll() ([]models.AvailTicket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var users []models.AvailTicket
	defer cancel()
	results, err := collectionAvailTicket.Find(ctx, bson.M{})
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleAvailTicket models.AvailTicket
		if err = results.Decode(&singleAvailTicket); err != nil {
			return users, err
		}

		users = append(users, singleAvailTicket)
	}
	return users, nil
}
