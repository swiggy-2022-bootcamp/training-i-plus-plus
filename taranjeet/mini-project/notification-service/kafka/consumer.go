package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

const (
	topic          = "irctc"
	broker1Address = "localhost:9092"
)

func Consume() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	fmt.Println("--------------------------------")

	// the `ReadMessage` method blocks until we receive the next event
	msg, err := r.ReadMessage(context.Background())
	if err != nil {
		panic("could not read message " + err.Error())
	}
	fmt.Println("Sending message !!", msg.Value)
}
