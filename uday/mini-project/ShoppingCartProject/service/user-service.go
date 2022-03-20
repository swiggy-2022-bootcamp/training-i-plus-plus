package service

import (
	"context"
	"fmt"
	"log"

	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "github.com/Udaysonu/SwiggyGoLangProject/config"
	"time"

	"github.com/Udaysonu/SwiggyGoLangProject/database"
)

 
var userCollection *mongo.Collection = database.GetCollection(database.DB, "users")

func GetAllUsers()[]entity.User{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []entity.User
	defer cancel()
	results, _ := userCollection.Find(ctx, bson.M{})
	for results.Next(ctx) {
		var singleUser entity.User
		results.Decode(&singleUser) 
		users = append(users, singleUser)
	}
	return users
}

func SignIn(username string, password string) entity.User{
	var result entity.User
	 
	filter := bson.M{"username": username,"password":password}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
 
	err := userCollection.FindOne(ctx, filter).Decode(&result)
	
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	} 
	return result 
}

func  SignUpUser(username string, password string, email string, phone int,location int) entity.User{
	newUser:=entity.User{Id:primitive.NewObjectID(),Username:username,Password:password,Email:email,Phone:phone,Location:location}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(userCollection)
	_,err:= userCollection.InsertOne(ctx, newUser)
	 if err!=nil{
		 fmt.Println(err)
	 }
	return newUser 
}

func IsUserPresent(username string, password string)bool{
	var isPresent bool=false;
	var result entity.User
	fmt.Println("--------checking is user present",username,password)
	filter := bson.M{"username": username}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("--------checking is user present",username,password)

	err := userCollection.FindOne(ctx, filter).Decode(&result)
	
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	} else{
		isPresent=true
	}
	
	return isPresent
}

func GetUser(username string, password string)entity.User{
 
	var result entity.User
	 
	filter := bson.M{"username": username,"password":password}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
 
	err := userCollection.FindOne(ctx, filter).Decode(&result)
	
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	
	} 
	return result 
}
