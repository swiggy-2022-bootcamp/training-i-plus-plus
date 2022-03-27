package goKafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type GoKafkaServiceImpl struct {
	Producer *kafka.Producer
}

func NewKafkaProducer(producer *kafka.Producer) *GoKafkaServiceImpl {
	return &GoKafkaServiceImpl{
		Producer: producer,
	}
}

func (p *GoKafkaServiceImpl) WriteMessage(topic string, msg interface{}) (bool, error) {
	// Serialize Message
	jsonString, err := json.Marshal(msg)
	msgString := string(jsonString)

	if err != nil {
		return false, err
	}

	// Produce messages to topic (asynchronously)
	for _, word := range []string{string(msgString)} {
		p.Producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	return true, nil
}
