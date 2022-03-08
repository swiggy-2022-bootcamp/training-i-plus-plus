package main

import (
	"golang-traintiketlelo/database"
	"golang-traintiketlelo/middleware"
	"golang-traintiketlelo/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var trainCollection *mongo.Collection = database.OpenCollection(database.Client, "train")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.TrainRoutes(router)
	routes.TicketRoutes(router)
	routes.TicketRouter(router)

	router.Run(":" + port)
}
