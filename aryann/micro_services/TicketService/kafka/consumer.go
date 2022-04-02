package kafka

import (
	"TicketService/logger"
	"TicketService/models"
	"TicketService/repository"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

var (
	ticketRepo repository.TicketRepository
	logger9    = logger.NewLoggerService("Kafka Consumer")
)

func Consume_purchased_ticket() {
	logger9.Log("Consume Purchased Ticket")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic1,
	})

	for {

		m, err := r.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		// Update ticket capacity
		var ticket models.Ticket
		err = json.Unmarshal(m.Value, &ticket)
		if err != nil {
			panic(err)
		}

		result, err := ticketRepo.ReadTrainId(ticket.Train_id)
		if err != nil {
			panic(err)
		}

		result.Capacity = result.Capacity - 1
		res, err := ticketRepo.Update(result, result.ID)
		if err != nil {
			panic(err)
		}

		logger9.Log("Purchased Ticket Consumed", res)
	}
}
