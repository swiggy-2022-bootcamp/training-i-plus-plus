package main
import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var client *mongo.Client
type Person struct{
	name string
	age int
}
func main(){
	var person=Person{"uday",20}
	fmt.Println("starting the application...")
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	client,_:=mongo.Connect(ctx,"mongodb://localhost:27017")	
	collection :=client.Database("Swiggy").Collection("people")
	result,_:=collection.InsertOne(ctx,person)
	fmt.Println(result)
}