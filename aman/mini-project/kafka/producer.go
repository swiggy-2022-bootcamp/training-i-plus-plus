package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "shopTopic"
	brokerAddress = "localhost:9092"
)

func Produce(ctx context.Context, msg string) {
	i := 0

	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(i)),
		Value: []byte(msg),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
	fmt.Println("Writes:", msg)
	i++
}
