package kafka

import (
	"context"
	"fmt"
	"strconv"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic         = "test"
	brokerAddress = "localhost:9092"
)

func UpdateSeatsProducer(trainId string, seatsBooked int32, isBook bool) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:   topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	var msg string
	if isBook {
		msg = "Book"+" "+trainId+" "+strconv.FormatInt(int64(seatsBooked), 10)
	} else {
		msg = "Cancel"+" "+trainId+" "+strconv.FormatInt(int64(seatsBooked), 10)
	}
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("traindId"),
			Value: []byte(msg),
		},
	)

	if err != nil {
		fmt.Print("Error Writing Kafka Msg")
	}
}