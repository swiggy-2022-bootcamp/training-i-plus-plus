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
	return &kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "order-1",
		// assign the logger to the reader
		Logger: logger,
	}
}

// Function to create new kafka consumer
func CreateKafkaConsumer(conf *kafka.ReaderConfig) *kafka.Reader {
	c := kafka.NewReader(*conf)
	return c
}
