package kafka

import (
	"context"
	"encoding/json"
	"srctc/logger"
	"srctc/models"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var logger8 = logger.NewLoggerService("Kafka Producer")

func Produce_booked_ticket_for_avail(trainid primitive.ObjectID, update bool) {
	logger8.Log("kafka producer booking ticket")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic1,
	})
	var updatestring string
	if update {
		updatestring = "increment"
	} else {
		updatestring = "decrement"
	}
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(trainid.String()),
		Value: []byte(updatestring),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func Produce_avail_ticket(nat models.Ticket) {
	logger8.Log("kafka producer avail ticket")
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
	logger8.Log("kafka producer avail ticket")
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
