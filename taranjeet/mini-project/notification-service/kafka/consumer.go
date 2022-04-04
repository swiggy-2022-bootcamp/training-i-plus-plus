package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/taran1515/crud/models"
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

	// the `ReadMessage` method blocks until we receive the next event
	msg, err := r.ReadMessage(context.Background())

	email := models.Email{
		To:      "test@gmail.com",
		Subject: "Ticket Booking",
		Message: string(msg.Value),
	}

	if err != nil {
		panic("could not read message " + err.Error())
	}
	fmt.Println("Sending message !!", email)
}
