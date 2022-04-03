package kafka

import (
	"context"
	"paymentService/models"
	"paymentService/services"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "payment"
	brokerAddress = "localhost:9092"
)

var l = services.NewLoggerService(topic)

func PaymentDetails(pd models.Payment) {
	l.Log(pd)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	defer w.Close()

	messsage := "user id: " + pd.UserId.String() + " amount: " + strconv.Itoa(pd.Amount) + " transaction id: " + pd.TransactionId
	messsage = messsage + " status: " + pd.Status

	err := w.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(messsage),
	})

	if err != nil {
		l.Log(err)
	}

}
