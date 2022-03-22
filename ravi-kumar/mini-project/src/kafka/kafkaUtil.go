package kafka

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "mini-project"
	brokerAddress = "localhost:9092"
)

func Produce(ctx context.Context, key []byte, value []byte) {
	logger := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: logger,
	})
	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
	// log a confirmation once the message is written
	//fmt.Println("writes:", value)
}
func Consume(ctx context.Context) (ch chan []byte) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	logger := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		// assign the logger to the reader
		Logger: logger,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
		ch <- msg.Value
	}
}
