package services

import (
	"context"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

const brokerAddress = "localhost:9092"

func Produce(message string, topic string, ctx context.Context) {
	l := log.New()

	// Intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte("data"),
		Value: []byte(message),
	})

	if err != nil {
		log.Error("could not write message " + err.Error())
	}
}
