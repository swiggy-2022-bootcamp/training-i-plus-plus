package services

import (
	"context"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

const brokerAddress = "localhost:9092"

func Consume(topic string, ctx context.Context) {
	l := log.New()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		Logger: l,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Error("Could not read Kafka Message: " + err.Error())
		}
		// after receiving the message, log its value
		log.Info(string(msg.Key) + ": " + string(msg.Value))
	}
}
