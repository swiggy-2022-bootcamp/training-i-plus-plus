package main

import (
	// "context"
	// "time"
	// "log"
	//  "fmt"
	//    "go.mongodb.org/mongo-driver/mongo"
	//  "go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	"log"
	"net/http"

	"github.com/bhatiachahat/mongoapi/db"
	"github.com/bhatiachahat/mongoapi/routes/user"
	// "go.mongodb.org/mongo-driver/mongo/driver/mongocrypt/options"
	//  "go.mongodb.org/mongo-driver/mongo/readpref"
)

// const (
// 	// MongoDBUrl is the default mongodb url that will be used to connect to the
// 	// database.
// 	dbName="myFirstDatabase"
// 	colName="users"
// 	MongoDBUrl ="mongodb+srv://chahat:chahat@cluster0.jb4z4.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
// )

// var collection *mongo.Collection

func init(){
	db.Connect()
	// clientOption := options.Client().ApplyURI(MongoDBUrl)
	// client,err := mongo.Connect(context.TODO(),clientOption)
	// if err!=nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connected")
	// collection = client.Database(dbName).Collection(colName)
	// fmt.Println("Instance ready")
}
func main() {
	r := user.Router()
	fmt.Println("Server getting started")
	
	//log.Fatal(http.ListenAndServe(":4000",r))
	router.GET("/new", articles.New)

}