package services

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func CreateProducer() (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err!=nil{
		return nil,err
	}
	return producer, nil
}

func ProduceAppointment(producer *kafka.Producer, topic string, msg interface{}) (bool, error){
	appointment, err := json.Marshal(msg)
	if err != nil{
		return false, err
	}

	appointmentString := string(appointment)

	for _, word := range []string{appointmentString}{
		producer.Produce(&kafka.Message{
			TopicPartition : kafka.TopicPartition{Topic : &topic, Partition : kafka.PartitionAny},
			Value : []byte(word), 
		},nil)
	}
	return true, nil
}