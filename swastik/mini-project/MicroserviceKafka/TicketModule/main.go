package main

import (
	"context"
	"log"
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/swastiksahoo153/MicroserviceKafka/TicketModule/controllers"
	"github.com/swastiksahoo153/MicroserviceKafka/TicketModule/services"
	"github.com/swastiksahoo153/MicroserviceKafka/TicketModule/database"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swastiksahoo153/MicroserviceKafka/TicketModule/docs"
)

var (
	
	server      		*gin.Engine
	ctx         		context.Context
	ticketservice   	services.TicketService
	ticketcontroller 	controllers.TicketController
	ticketcollection	*mongo.Collection
	mongoclient 		*mongo.Client
	err         		error
)

func init(){
	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}

	ctx = context.TODO()

	mongoclient = database.GetDatabase(ctx)

	ticketcollection = mongoclient.Database("ticketdb").Collection("tickets")
	ticketservice = services.NewTicketService(ticketcollection, ctx)
	ticketcontroller = controllers.New(ticketservice)
	server = gin.Default()

	docs.SwaggerInfo.Title = "Train Reservation System"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}


// @title          Ticket Module
// @version        1.0
// @description    This microservice is for ticket module.
// @contact.name   Swastik Sahoo
// @contact.email  swastiksahoo22@gmail.com
// @license.name  Apache 2.0
// @host      localhost:8082
// @securityDefinitions.apikey  Bearer Token
// @in                          header
// @name                        Authorization
func main(){
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("")
	ticketcontroller.RegisterTicketRoutes(basepath)

	log.Fatal(server.Run(":8082"))
}