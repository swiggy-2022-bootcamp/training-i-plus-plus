package kafkaservice

import (
	"encoding/json"
	"fmt"
"time"
	 //model "bhatiachahat/payment-service/model"
	// usermodel "bhatiachahat/user-service/model"
	 model "bhatiachahat/payment-service/model"
//	 database "bhatiachahat/payment-service/db"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	//"context"
	// "fmt"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"time"
////	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
)

const consumerTopic = "Orders"
 //var trackstreamCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

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



func consumeOrder(consumer *kafka.Consumer, topic string){
	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println("\n")
			fmt.Printf("\n%s",  time.Now())
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			var order model.Order
			data := []byte(msg.Value)
			err = json.Unmarshal(data, &order)
			if err != nil {
				fmt.Println("Error in unmarshalling kafka message into appointment struct")
				return
			}
			fmt.Println("\n")
			fmt.Printf("\n%s",  time.Now())
			fmt.Println("Order consumed in payment service",order)
			fmt.Println("\n\n");
			fmt.Printf("\n%s",  time.Now())
			fmt.Println("Payment successfull and Email sent to the user with order details.")

			
		} 
	}

}