package db
import (
    "context"
   // "time"
   "log"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/driver/mongocrypt/options"
  //  "go.mongodb.org/mongo-driver/mongo/readpref"
)
const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	dbName="myFirstDatabase"
	colName="users"
	MongoDBUrl ="mongodb+srv://chahat:chahat@cluster0.jb4z4.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
)

var Collection *mongo.Collection

func Connect(){
	clientOption := options.Client().ApplyURI(MongoDBUrl)
	client,err := mongo.Connect(context.TODO(),clientOption)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Instance ready")
}