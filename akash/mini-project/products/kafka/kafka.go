package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"products.akash.com/model"
	//"time"
)

const (
	topic         = "buy-request"
	brokerAddress = "localhost:9092"
)

func CreateComment(buyRequest *model.BuyRequest) {

	l := log.New(os.Stdout, "kafka writer: ", 0)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	j, err := json.Marshal(buyRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(string(j)),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
