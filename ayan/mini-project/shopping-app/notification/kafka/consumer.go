package kafka

import (
	"context"
	"fmt"
	"log"
	"notification/utils/logger"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "orders"
	brokerAddress = "localhost:9092"
)

func ConsumeOrders(ctx context.Context) {

	logger.Info("Starting consumer...")

	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		// assign the logger to the reader
		Logger: l,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err == nil {
			data := string(msg.Value)
			logger.Info(fmt.Sprintf("Message on %s: %s\n", topic, data))
		} else {
			// The client will automatically try to recover from all errors.
			logger.Error(fmt.Sprintf("Consumer error: %v (%v)\n", err, msg))
		}
	}
}
