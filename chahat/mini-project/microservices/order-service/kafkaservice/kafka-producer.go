package kafkaservice

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
	"time"
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

func ProduceOrder(producer *kafka.Producer, topic string, msg interface{}) (bool, error){
	order, err := json.Marshal(msg)
	if err != nil{
		return false, err
	}

	orderString := string(order)
	fmt.Printf("\n%s\n",  time.Now())
	fmt.Print("\nOrder produced in kafka\n",orderString)


	for _, word := range []string{orderString}{
		producer.Produce(&kafka.Message{
			TopicPartition : kafka.TopicPartition{Topic : &topic, Partition : kafka.PartitionAny},
			Value : []byte(word), 
		},nil)
	}
	return true, nil
}