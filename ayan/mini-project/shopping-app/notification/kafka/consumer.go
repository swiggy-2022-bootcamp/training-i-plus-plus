package kafka

import (
	"fmt"
	"notification/utils/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConsumeOrders() {
	logger.Info("Starting consumer...")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	c.SubscribeTopics([]string{"orders"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			data := string(msg.Value)
			logger.Info(fmt.Sprintf("Message on %s: %s\n", msg.TopicPartition, data))
		} else {
			// The client will automatically try to recover from all errors.
			logger.Error(fmt.Sprintf("Consumer error: %v (%v)\n", err, msg))
		}
	}
}
