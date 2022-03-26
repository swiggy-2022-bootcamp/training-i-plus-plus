package goKafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

// Function to create configuration instance
func ConsumerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers":               "localhost:9092",
		"group.id":                        "group-id-1",
		"go.application.rebalance.enable": true, // delegate Assign() responsibility to app
		"session.timeout.ms":              6000,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
	}
}

// Function to create new kafka consumer
func CreateKafkaConsumer(conf *kafka.ConfigMap) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(conf)
	if err != nil {
		return nil, err
	}
	return c, nil
}
