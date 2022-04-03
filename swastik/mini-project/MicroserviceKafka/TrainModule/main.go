package main

import (
	"context"
	"log"
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/swastiksahoo153/MicroserviceKafka/TrainModule/controllers"
	"github.com/swastiksahoo153/MicroserviceKafka/TrainModule/services"
	"github.com/swastiksahoo153/MicroserviceKafka/TrainModule/database"
	kf "github.com/swastiksahoo153/MicroserviceKafka/TrainModule/kafkaConsumer"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/swastiksahoo153/MicroserviceKafka/TrainModule/docs"	
)

var (
	
	server      	*gin.Engine
	ctx         	context.Context
	trainservice    services.TrainService
	traincontroller controllers.TrainController
	Traincollection	*mongo.Collection
	mongoclient 	*mongo.Client
	err         	error
)

func init(){
	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}

	ctx = context.TODO()

	mongoclient = database.GetDatabase(ctx)

	Traincollection = mongoclient.Database("traindb").Collection("trains")
	trainservice = services.NewTrainService(Traincollection, ctx)
	traincontroller = controllers.New(trainservice)
	server = gin.Default()

	docs.SwaggerInfo.Title = "Train Reservation System"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @title          Train Module
// @version        1.0
// @description    This microservice is for train module.
// @contact.name   Swastik Sahoo
// @contact.email  swastiksahoo22@gmail.com
// @license.name  Apache 2.0
// @host      localhost:8081
// @securityDefinitions.apikey  Bearer Token
// @in                          header
// @name                        Authorization
func main(){
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("")
	traincontroller.RegisterTrainRoutes(basepath)

	go kf.Consume(ctx)

	log.Fatal(server.Run(":8081"))
}