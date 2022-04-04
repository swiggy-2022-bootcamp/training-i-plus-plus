package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	topic          = "irctc"
	broker1Address = "localhost:9092"
)

func Produce(userName string) {

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	fmt.Println("userName", userName)

	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key: []byte("ticket"),
		// create an arbitrary message payload for the value
		Value: []byte(userName),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	// log a confirmation once the message is written
	fmt.Println("sending ticket book notification:")
	// sleep for a second
	time.Sleep(time.Second)

}
