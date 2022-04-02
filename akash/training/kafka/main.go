package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "test-topic"
	brokerAddress = "localhost:9092"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go produce(ctx)
	consume(ctx)
}

func produce(ctx context.Context) {

	l := log.New(os.Stdout, "kafka writer: ", 0)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Value: []byte("Kafka is working!!!!!!!!"),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	// sleep for a second
	time.Sleep(time.Second)
}

func consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the reader
		Logger: l,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
