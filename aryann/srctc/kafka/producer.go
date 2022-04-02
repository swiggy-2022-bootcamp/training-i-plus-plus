package kafka

import (
	"context"
	"encoding/json"
	"srctc/logger"
	"srctc/models"

	"github.com/segmentio/kafka-go"
)

var logger8 = logger.NewLoggerService("Kafka Producer")

func Produce_purchased_ticket(newPurchase models.Purchased) {
	logger8.Log("Produce Purchased Ticket")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "purchasedticket",
		Balancer: &kafka.LeastBytes{},
	})

	bytes, _ := json.Marshal(newPurchase)
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(newPurchase.Train_id.String()),
		Value: bytes,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func Produce_ticket(nat models.Ticket) {
	logger8.Log("Produce Ticket")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic2,
	})

	bytes, _ := json.Marshal(nat)
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(nat.ID.String()),
		Value: []byte(bytes),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func Produce_train(nt models.Train) {
	logger8.Log("Produce Train")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic3,
	})

	bytes, _ := json.Marshal(nt)
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(nt.Destination),
		Value: []byte(bytes),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
