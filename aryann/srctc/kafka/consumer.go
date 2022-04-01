package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"srctc/logger"
	"srctc/models"
	"srctc/repository"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	topic1        = "bookedticket"
	topic2        = "ticket"
	topic3        = "train"
	brokerAddress = "localhost:9092"
)

var (
	ticketRepo repository.TicketRepository
	logger9    = logger.NewLoggerService("Kafka Consumer")
)

func Consume_booked_ticket_for_avail() {
	logger9.Log("kafka producer booking ticket")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic1,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		//fmt.Println("received: ", string(msg.Value))
		trainid, err := primitive.ObjectIDFromHex(string(msg.Key))

		if err != nil {
			panic("incorrect train id " + err.Error())
		}

		ticket, err := ticketRepo.ReadTrainId(trainid)

		if err != nil {
			panic("could not find the train " + err.Error())
		}

		if string(msg.Value) == "increment" {
			ticket.Capacity += 1
		} else {
			ticket.Capacity -= 1
		}

		_, err = ticketRepo.Update(ticket, ticket.ID)
		if err != nil {
			panic("could not update booked ticket " + err.Error())
		}
	}
}

func Consume_avail_ticket() {
	logger9.Log("kafka consumer  ticket")
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
		res := fmt.Sprintf("Inserted new able ticket %#v", natr)
		logger9.Log(res)
	}
}

func Consume_train() {
	logger9.Log("kafka consumer  ticket")
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
		res := fmt.Sprintf("Inserted new ticket %#v", ntr)
		logger9.Log(res)
	}
}
