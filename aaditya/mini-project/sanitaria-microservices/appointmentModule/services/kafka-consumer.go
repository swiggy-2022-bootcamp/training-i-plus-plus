package services

import (
	"encoding/json"
	"fmt"
	"sanitaria-microservices/appointmentModule/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const consumerTopic = "Appointment"

func StartKafkaConsumer(){
	consumer, err := createConsumer()
	if err != nil{
		fmt.Println("Error in creating kafka-consumer.")
	}else{
		go consumeAppointment(consumer,consumerTopic)
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

func consumeAppointment(consumer *kafka.Consumer, topic string){
	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			//fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			var appointment models.Appointment
			data := []byte(msg.Value)
			err = json.Unmarshal(data, &appointment)
			if err != nil {
				fmt.Println("Error in unmarshalling kafka message into appointment struct")
				return
			}
			AddAppointmentToDB(appointment)
		} 
	}

}