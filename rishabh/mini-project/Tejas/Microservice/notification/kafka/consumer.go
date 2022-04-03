package kafka

import (
	"context"
	"notificationService/services"

	"github.com/segmentio/kafka-go"
)

const (
	paymentTopic  = "payment"
	ticketTopic   = "ticket"
	brokerAddress = "localhost:9092"
)

var l = services.NewLoggerService("notification")

func PaymentDetails() {
	l.Log("Payment Details")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   paymentTopic,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			l.Log(err)
		}
		l.Log(string(m.Value))
	}
}

func TicketDetails() {
	l.Log("Ticket Details")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   ticketTopic,
	})
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			l.Log(err)
		}
		l.Log(string(m.Value))
	}
}
