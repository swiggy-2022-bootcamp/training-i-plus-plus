package database
import (
    "context"
	"fmt"
    "time"
     "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/mongo/readpref"
//	log "github.com/Udaysonu/SwiggyGoLangProject/config"
)
var DB *mongo.Client=ConnectDatabase()

func ConnectDatabase() *mongo.Client {
	clientOptions := options.Client().
    ApplyURI("mongodb+srv://uday:uday@cluster0.oxuet.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("ConnectDatabase Error: ",err)
	} else {
		fmt.Println("Database Connected Successfully......")
	}
	return client
}


func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	ConnectDatabase()
	collection := client.Database("maindb").Collection(collectionName)
	return collection
}

