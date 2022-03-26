package goKafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Function to create ConfigMap for kafka
func Cfg() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}
}

// Function to create New Producer
func CreateProducer(conf *kafka.ConfigMap) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(conf)
	if err != nil {
		return nil, err
	}
	return p, nil
}
