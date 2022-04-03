package kafkaservice

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
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

func ProduceProduct(producer *kafka.Producer, topic string, msg interface{}) (bool, error){
	product, err := json.Marshal(msg)
	if err != nil{
		return false, err
	}

	productString := string(product)
	fmt.Print(productString)


	for _, word := range []string{productString}{
		producer.Produce(&kafka.Message{
			TopicPartition : kafka.TopicPartition{Topic : &topic, Partition : kafka.PartitionAny},
			Value : []byte(word), 
		},nil)
	}
	return true, nil
}