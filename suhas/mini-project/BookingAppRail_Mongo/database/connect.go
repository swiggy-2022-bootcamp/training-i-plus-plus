package database

import (
	"fmt"
	"log"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	"BookingAppMongo/models"
)


var (
	client     *mongo.Client
	mongoURL = "mongodb://localhost:27017"
	ctx context.Context
	err error
)


func init() {
	ctx,_ := context.WithTimeout(context.Background(),10 *time.Second)
	client,err = mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017/"))	
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx,readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}

func InsertAdmin(newAdmin models.Admin) (err error) {
	database := client.Database("bookapp")
	adminDB := database.Collection("admin")
	insertResult,err := adminDB.InsertOne(context.TODO(),newAdmin)
	
	if err != nil {
		return err
	}
	fmt.Println(insertResult)
	return nil
} 

func ReadAdmin(adminid string) []models.Admin{
	database := client.Database("bookapp")
	adminDB := database.Collection("admin")
	
	filterCursor,err := adminDB.Find(ctx, bson.M{"adminid":adminid})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []models.Admin

	if err = filterCursor.All(ctx,&episodesFiltered); err != nil {
		log.Fatal(err)
	}
	return episodesFiltered
}


func ReadAllAdmin() []models.Admin{
	database := client.Database("bookapp")
	adminDB := database.Collection("admin")
	
	filterCursor,err := adminDB.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []models.Admin

	if err = filterCursor.All(ctx,&episodesFiltered); err != nil {
		log.Fatal(err)
	}
	return episodesFiltered
}

func UpdateAdmin(oldadminid string,newAdmin models.Admin) (err error) {
	database := client.Database("bookapp")
	adminDB := database.Collection("admin")
	result,err = adminDB.ReplaceOne(
		ctx,
		bson.M{"adminid":oldadminid},
		newAdmin
	)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func DeleteAdmin(adminid string) {
	database := client.Database("bookapp")
	adminDB := database.Collection("admin")
	result,err = adminDB.DeleteOne(ctx,bson.M{"adminid":adminid})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}