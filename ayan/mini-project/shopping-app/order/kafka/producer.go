package kafka

import (
	"context"
	"log"
	"order/utils/logger"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "orders"
	brokerAddress = "localhost:9092"
)

func KafkaWriter() *kafka.Writer {

	logger.Info("Starting producer...")

	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	return w
}

func Produce(ctx context.Context, w *kafka.Writer, key string, message string) {

	// Produce messages to topic (asynchronously)
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(key),
		// create an arbitrary message payload for the value
		Value: []byte(message),
	})
	if err != nil {
		logger.Fatal(err.Error())
		panic("could not write message " + err.Error())
	}

}
