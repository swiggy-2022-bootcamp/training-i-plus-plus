package repository

import (
	"UserService/database"
	"UserService/logger"
	"UserService/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionUserName = "users"
	collectionUser     = new(mongo.Collection)
	logger7            = logger.NewLoggerService("purchased_repository")
)

func init() {
	var DB *mongo.Client = database.ConnectDB()
	collectionUser = DB.Database("golangAPI").Collection(collectionUserName)
}

type UserRepository struct{}

func (usr UserRepository) Create(newUser models.User) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionUser.InsertOne(ctx, &newUser)
	if err == nil {
		logger7.Log("User created successfully", newUser.Name)
		// fmt.Println("Inserted a single document: ", result.InsertedID)
	}
	return result.InsertedID, err
}

func (usr UserRepository) Read(objId primitive.ObjectID) (models.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.User
	err := collectionUser.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	return user, err
}

func (usr UserRepository) Update(updateUser models.User, objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//update := bson.M{"name": updateUser.Name, "email": updateUser.Email}
	updatebson := bson.M{}
	if updateUser.Name != "" {
		updatebson["name"] = updateUser.Name
	}
	if updateUser.Email != "" {
		updatebson["email"] = updateUser.Email
	}
	result, err := collectionUser.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updatebson})
	//result, err := collectionUser.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err == nil {
		logger7.Log("User updated successfully", updateUser.Name)
		// fmt.Println("Updated a single document: ", result.UpsertedID)
	}
	return result.UpsertedID, err
}

func (usr UserRepository) Delete(objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionUser.DeleteOne(ctx, bson.M{"_id": objId})
	if err == nil {
		logger7.Log("User deleted successfully", objId)
		// fmt.Println("Updated a single document: ", result.DeletedCount)
	}
	return result.DeletedCount, err
}

func (usr UserRepository) ReadAll() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var users []models.User
	defer cancel()
	results, err := collectionUser.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return users, err
		}

		users = append(users, singleUser)
	}
	return users, nil
}
