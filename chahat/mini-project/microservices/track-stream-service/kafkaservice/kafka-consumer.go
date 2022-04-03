package kafkaservice

import (
	"encoding/json"
	"fmt"
	 //model "bhatiachahat/payment-service/model"
	// usermodel "bhatiachahat/user-service/model"
	 model "bhatiachahat/track-stream-service/model"
	 database "bhatiachahat/track-stream-service/db"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"context"
	// "fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
//	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const consumerTopic = "Orders"
 var trackstreamCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

func StartKafkaConsumer(){
	consumer, err := createConsumer()
	if err != nil{
		fmt.Println("Error in creating kafka-consumer.")
	}else{
		go consumeOrder(consumer,consumerTopic)
	}
}
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

func AddDatainDB(order model.Order){

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	var trackstreamobj model.TrackStream


		//product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	trackstreamobj.ID = primitive.NewObjectID()
    trackstreamobj.PaymentType =order.Payment_Method
	//	product.Product_id = product.ID.Hex()
	

		_, insertErr := trackstreamCollection.InsertOne(ctx,trackstreamobj)
		if insertErr !=nil {
		
			fmt.Println( insertErr.Error())
			return
		}
		defer cancel()
	//	c.JSON(http.StatusOK, product)
	// _, err :=  productCollection.InsertOne(ctx, product)
    // if err != nil {
    //     fmt.Println(err.Error())
    // }
	fmt.Printf("Saved to MongoDB : %s\n", trackstreamobj)
}

func consumeOrder(consumer *kafka.Consumer, topic string){
	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println("\n")
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			var order model.Order
			data := []byte(msg.Value)
			err = json.Unmarshal(data, &order)
			if err != nil {
				fmt.Println("Error in unmarshalling kafka message into appointment struct")
				return
			}
			fmt.Println("\n")
			fmt.Println("Order consumed in track stream service",order)
			fmt.Println("\n",order.Payment_Method)
			AddDatainDB(order)
		
			// AddOrderToDB(user)
		} 
	}

}