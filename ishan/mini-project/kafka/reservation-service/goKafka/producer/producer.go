package goKafka

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	logger "swiggy/gin/logger"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "reservations"
	brokerAddress = "localhost:9092"
)

func WriteMessage(ctx context.Context, topic string, msg interface{}) (bool, error) {
	i := 0
	l := logger.Loggerx() //log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	// Serialize Message
	jsonString, err := json.Marshal(msg)
	msgString := string(jsonString)

	if err != nil {
		return false, err
	}

	// Produce messages to topic (asynchronously)
	for _, word := range []string{string(msgString)} {

		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: []byte(word),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", msgString)
	}

	return true, nil
}
