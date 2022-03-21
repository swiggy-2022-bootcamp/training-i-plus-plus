package main

import (
	"context"
	"log"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swastiksahoo153/train-reservation-system/controllers"
	"github.com/swastiksahoo153/train-reservation-system/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/joho/godotenv"
)

var (
	
	server      	*gin.Engine
	ctx         	context.Context
	userservice    	services.UserService
	usercontroller 	controllers.UserController
	usercollection	*mongo.Collection
	mongoclient 	*mongo.Client
	err         	error
)

func init(){
	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}

	ctx = context.TODO()

	mongouri := os.Getenv("MONGO_URI")
	fmt.Println(mongouri)

	mongoconn := options.Client().ApplyURI(mongouri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")

	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()
}

func main(){
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}