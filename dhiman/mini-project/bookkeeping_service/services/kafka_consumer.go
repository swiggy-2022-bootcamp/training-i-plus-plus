package services

import (
	"context"

	"github.com/dhi13man/healthcare-app/bookkeeping_service/configs"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var brokerAddress string = configs.KafkaBrokerAddress()

func Consume(topic string, callback func(string, context.Context), ctx context.Context) {
	l := log.New()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		Logger: l,
	})

	for {
		// The `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Error("Could not read Kafka Message: " + err.Error())
		}

		// After receiving the message, log its value and apply the callback
		log.Info(string(msg.Key) + ": " + string(msg.Value))
		callback(string(msg.Value), ctx)
	}
}
