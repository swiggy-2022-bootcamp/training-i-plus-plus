package kafka

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "TrainTicketLelo"
	brokerAddress = "localhost:9092"
)

func Produce(ctx context.Context, key []byte, value []byte) {
	logger := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  logger,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	fmt.Println("writes:", string(value))
}
func Consume(ctx context.Context) (ch chan []byte) {

	logger := log.New(os.Stdout, "kafka reader: ", 0)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "consumer-group-1",
		Logger:  logger,
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
		ch <- msg.Value
	}
}
