package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"order.akash.com/db"
	"order.akash.com/model"
	"os"
)

const (
	topic         = "buy-request"
	brokerAddress = "localhost:9092"
	groupId       = "my-group"
)

var (
	repo = db.NewMongoRepository()
)

func StartOrderListener(repository db.OrderRepository) {
	repo = repository
	ctx := context.Background()
	l := log.New(os.Stdout, "kafka reader: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupId,
		Logger:  l,
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		var data model.Order
		err = json.Unmarshal(msg.Value, &data)

		fmt.Println("received order: ", data)
		repo.SaveOrder(data)
	}
}
