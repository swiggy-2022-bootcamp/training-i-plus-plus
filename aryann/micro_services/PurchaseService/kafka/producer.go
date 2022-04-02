package kafka

import (
	"PurchaseService/logger"
	"context"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var logger8 = logger.NewLoggerService("Kafka Producer")

func Produce_purchased_ticket(train_id primitive.ObjectID) {
	logger8.Log("Produce_purchased_ticket")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "purchased_ticket",
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(train_id.Hex()),
			Value: []byte(train_id.Hex()),
		},
	)

	if err != nil {
		logger8.Log(err.Error())
	}
}

// func Produce_purchased_ticket(newPurchase models.Purchased) {
// 	logger8.Log("Produce Purchased Ticket")
// 	w := kafka.NewWriter(kafka.WriterConfig{
// 		Brokers:  []string{"localhost:9092"},
// 		Topic:    "purchasedticket",
// 		Balancer: &kafka.LeastBytes{},
// 	})

// 	bytes, _ := json.Marshal(newPurchase)
// 	err := w.WriteMessages(context.Background(), kafka.Message{
// 		Key:   []byte(newPurchase.Train_id.String()),
// 		Value: bytes,
// 	})
// 	if err != nil {
// 		panic("could not write message " + err.Error())
// 	}
// }
