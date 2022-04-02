package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"srctc/logger"
	"srctc/models"
	"srctc/repository"

	"github.com/segmentio/kafka-go"
)

const (
	topic1        = "bookedticket"
	topic2        = "ticket"
	topic3        = "train"
	brokerAddress = "localhost:9092"
)

var (
	purchasedRepo repository.PurchasedRepository
	logger9       = logger.NewLoggerService("Kafka Consumer")
)

func Consume_purchased_ticket() {
	logger9.Log("Consume Purchased Ticket")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic1,
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}

		purchased := models.Purchased{}
		json.Unmarshal([]byte(msg.Value), &purchased)
		res := fmt.Sprintf("Created new purchased ticket %#v", purchased)
		logger9.Log(res)
		if _, err := purchasedRepo.Create(purchased); err != nil {
			panic("could not create purchased ticket " + err.Error())
		}
	}
}

func Consume_ticket() {
	logger9.Log("Consume Ticket")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic2,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		natr := models.Ticket{}
		json.Unmarshal([]byte(msg.Value), &natr)
		res := fmt.Sprintf("Created new ticket %#v", natr)
		logger9.Log(res)
	}
}

func Consume_train() {
	logger9.Log("Consume Train")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic3,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		ntr := models.Train{}
		json.Unmarshal([]byte(msg.Value), &ntr)
		res := fmt.Sprintf("Inserted new train journey %#v", ntr)
		logger9.Log(res)
	}
}
