package main

import (
	"context"
	"encoding/json"
	"strconv"

    "go.mongodb.org/mongo-driver/bson/primitive"

	kafka "github.com/segmentio/kafka-go"
)


const (
	topic          = "UpdateOrderStatus"
	broker1Address = "localhost:9092"

)

type Bird struct {
	OrderId     primitive.ObjectID `json:"orderid"`
	Status string `json:"status"`
}

func produce(ctx context.Context) {
	// initialize a counter
	i := 0

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})
	
	objId, _ := primitive.ObjectIDFromHex("6248aafae1a2394b9d427c89")


	hmm := &Bird{
		OrderId: objId,
		Status: "Out for delivery",
	}
	str, _ :=json.Marshal(hmm)
	
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(i)),
		// create an arbitrary message payload for the value
		
		Value: []byte(str),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

}


func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	produce(ctx)
}