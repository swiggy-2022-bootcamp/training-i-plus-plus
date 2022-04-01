package repository

import (
	"context"
	"srctc/database"
	"srctc/logger"
	"srctc/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionAdminName = "admins"
	collectionAdmin     = new(mongo.Collection)
	logger2             = logger.NewLoggerService("admin_repository")
)

func init() {
	var DB *mongo.Client = database.ConnectDB()
	collectionAdmin = DB.Database("golangAPI").Collection(collectionAdminName)
}

type AdminRepository struct{}

func (adm AdminRepository) Create(newAdmin models.Admin) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionAdmin.InsertOne(ctx, &newAdmin)
	if err == nil {
		logger2.Log("Created a new admin: ", newAdmin.Name)
	}
	return result.InsertedID, err
}

func (adm AdminRepository) Read(objId primitive.ObjectID) (models.Admin, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.Admin
	err := collectionAdmin.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	return user, err
}

func (adm AdminRepository) Update(updateAdmin models.Admin, objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	updatebson := bson.M{}
	if updateAdmin.Name != "" {
		updatebson["name"] = updateAdmin.Name
	}
	if updateAdmin.Email != "" {
		updatebson["email"] = updateAdmin.Email
	}
	result, err := collectionAdmin.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updatebson})
	if err == nil {
		logger2.Log("Updated Admin details: ", result.UpsertedID)
		// fmt.Println("Updated a single document: ", result.UpsertedID)
	}
	return result.UpsertedID, err
}

func (adm AdminRepository) Delete(objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionAdmin.DeleteOne(ctx, bson.M{"_id": objId})
	if err == nil {
		logger2.Log("Deleted Admin: ", objId)
		// fmt.Println("Updated a single document: ", result.DeletedCount)
	}
	return result.DeletedCount, err
}

func (adm AdminRepository) ReadAll() ([]models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var users []models.Admin
	defer cancel()
	results, err := collectionAdmin.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleAdmin models.Admin
		if err = results.Decode(&singleAdmin); err != nil {
			return users, err
		}

		users = append(users, singleAdmin)
	}
	return users, nil
}
