package kafka

import (
	"context"
	"reservationService/models"
	"reservationService/services"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "ticket"
	brokerAddress = "localhost:9092"
)

var l = services.NewLoggerService(topic)

func TicketDetails(td models.Reservation) {
	l.Log(td)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	defer w.Close()

	message := "user id: " + td.UserId.String() + "has successfully booked seat number: " + strconv.Itoa(td.SeatNumber)
	message = message + " transaction id: " + td.TransactionId + " status: " + td.Status + " from: " + td.FromStationCode + " to: " + td.ToStationCode

	err := w.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(message),
	})

	if err != nil {
		l.Log(err)
	}

}
