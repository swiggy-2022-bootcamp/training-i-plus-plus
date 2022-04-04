package kafka

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

const (
	paymentTopic  = "payment"
	ticketTopic   = "booking"
	brokerAddress = "localhost:9092"
)

func PaymentDetails() {
	log.Info("Payment Details")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   paymentTopic,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Error(err)
		}
		log.Info(string(m.Value))
	}
}

func TicketDetails() {
	log.Info("Ticket Details")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   ticketTopic,
	})
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Error(err)
		}
		log.Info(string(m.Value))
	}
}
