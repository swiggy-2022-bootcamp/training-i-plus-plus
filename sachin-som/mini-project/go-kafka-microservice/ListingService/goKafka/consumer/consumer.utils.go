package goKafka

import (
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

var (
	logger = log.New(os.Stdout, "kafka reader: ", 0)
)

const (
	topic         = "products"
	brokerAddress = "localhost:9092"
)

// Function to create configuration instance
func ConsumerConfig() *kafka.ReaderConfig {
	// return &kafka.ConfigMap{
	// 	"bootstrap.servers":               "localhost:9092",
	// 	"group.id":                        "group-id-1",
	// 	"go.application.rebalance.enable": true, // delegate Assign() responsibility to app
	// 	"session.timeout.ms":              6000,
	// 	"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
	// }
	return &kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "consumer-group-1",
		// assign the logger to the reader
		Logger: logger,
	}
}

// Function to create new kafka consumer
func CreateKafkaConsumer(conf *kafka.ReaderConfig) *kafka.Reader {
	c := kafka.NewReader(*conf)
	return c
}
