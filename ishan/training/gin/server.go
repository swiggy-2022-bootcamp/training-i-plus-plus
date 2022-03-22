package main

import (
	"fmt"
	"log"
	"net/http"
	db "swiggy/gin/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAlbum struct {
	ID     primitive.ObjectID `bson:"_id",omitempty`
	Title  string             `bson:"title"`
	Artist string             `bson:"artist"`
	Price  string             `bson:"price"`
}

func getAlbums(c *gin.Context) {
	cursor, err := db.DataStore.Collection("albums").Find(c, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var albums []bson.M
	if err = cursor.All(c, &albums); err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	client, ctx, cancel := db.ConnectDB()
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
		fmt.Println("MongoDB Connection Closed")
	}()

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
