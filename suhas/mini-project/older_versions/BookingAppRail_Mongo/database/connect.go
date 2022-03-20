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
	_,err = adminDB.ReplaceOne(
		ctx,bson.M{"adminid":oldadminid},newAdmin,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAdmin(adminid string)(err error){
	database := client.Database("bookapp")
	adminDB := database.Collection("admin")
	_,err = adminDB.DeleteOne(ctx,bson.M{"adminid":adminid})
	if err != nil {
		return err
	}
	return nil
}



func InsertCustomer(newCustomer models.Customer) (err error) {
	database := client.Database("bookapp")
	CustomerDB := database.Collection("Customer")
	insertResult,err := CustomerDB.InsertOne(context.TODO(),newCustomer)
	
	if err != nil {
		return err
	}
	fmt.Println(insertResult)
	return nil
} 

func ReadCustomer(Customerid string) []models.Customer{
	database := client.Database("bookapp")
	CustomerDB := database.Collection("Customer")
	
	filterCursor,err := CustomerDB.Find(ctx, bson.M{"Customerid":Customerid})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []models.Customer

	if err = filterCursor.All(ctx,&episodesFiltered); err != nil {
		log.Fatal(err)
	}
	return episodesFiltered
}


func ReadAllCustomer() []models.Customer{
	database := client.Database("bookapp")
	CustomerDB := database.Collection("Customer")
	
	filterCursor,err := CustomerDB.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []models.Customer

	if err = filterCursor.All(ctx,&episodesFiltered); err != nil {
		log.Fatal(err)
	}
	return episodesFiltered
}

func UpdateCustomer(oldCustomerid string,newCustomer models.Customer) (err error) {
	database := client.Database("bookapp")
	CustomerDB := database.Collection("Customer")
	_,err = CustomerDB.ReplaceOne(
		ctx,
		bson.M{"Customerid":oldCustomerid},
		newCustomer,
	)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func DeleteCustomer(Customerid string) (err error) {
	database := client.Database("bookapp")
	CustomerDB := database.Collection("Customer")
	_,err = CustomerDB.DeleteOne(ctx,bson.M{"Customerid":Customerid})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}