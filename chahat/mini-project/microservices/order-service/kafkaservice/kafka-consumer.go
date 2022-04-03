package kafkaservice

import (
//	"encoding/json"
	//"fmt"
	//model "bhatiachahat/order-service/model"
	database "bhatiachahat/order-service/db"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	//"context"
	// "fmt"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"time"
//	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const consumerTopic = "Products"
var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

// func StartKafkaConsumer(){
// 	consumer, err := createConsumer()
// 	if err != nil{
// 		fmt.Println("Error in creating kafka-consumer.")
// 	}else{
// 	//	go consumeProduct(consumer,consumerTopic)
// 	}
// }
func createConsumer() (*kafka.Consumer, error){
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group1",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
// func AddProductToDB(product model.Product){
// 	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
// 	defer cancel()
//          product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		product.ID = primitive.NewObjectID()
// 		product.Product_id = product.ID.Hex()
	

// 		_, insertErr := productCollection.InsertOne(ctx, product)
// 		if insertErr !=nil {
		
// 			fmt.Println( insertErr.Error())
// 			return
// 		}
// 		defer cancel()
// 	//	c.JSON(http.StatusOK, product)
// 	// _, err :=  productCollection.InsertOne(ctx, product)
//     // if err != nil {
//     //     fmt.Println(err.Error())
//     // }
// 	fmt.Printf("Saved to MongoDB : %s", product)
// }
// func consumeProduct(consumer *kafka.Consumer, topic string){
// 	consumer.SubscribeTopics([]string{topic}, nil)

// 	for {
// 		msg, err := consumer.ReadMessage(-1)
// 		if err == nil {
// 			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
// 			var product model.Product
// 			data := []byte(msg.Value)
// 			err = json.Unmarshal(data, &product)
// 			if err != nil {
// 				fmt.Println("Error in unmarshalling kafka message into appointment struct")
// 				return
// 			}
// 			//AddProductToDB(product)
// 		} 
// 	}

// }