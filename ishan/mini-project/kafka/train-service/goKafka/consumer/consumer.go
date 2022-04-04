package goKafka

import (
	"context"
	"encoding/json"
	trainservice "swiggy/gin/kafka_services"
	logger "swiggy/gin/logger"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "reservations"
	brokerAddress = "localhost:9092"
)

func Consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := logger.Loggerx() //log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
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
		//fmt.Println("received: ", string(msg.Value))
		// Update seat Matrix
		t := trainservice.TrainData{}
		json.Unmarshal(msg.Value, &t)

		trainservice.UpdateTrainInfo(t.Train, t.Seat)
	}
}
