package services

import (
	"context"

	"github.com/dhi13man/healthcare-app/users_service/configs"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var brokerAddress string = configs.KafkaBrokerAddress()

func Produce(message string, topic string, ctx context.Context) error {
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
	return err
}
