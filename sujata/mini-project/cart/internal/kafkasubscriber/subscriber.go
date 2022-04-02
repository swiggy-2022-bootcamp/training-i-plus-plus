package subscriber

import (
	mongodao "cart/internal/dao"
	"context"
	l "log"
	"os"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

const (
	ORDER_PLACED string = "ORDER_PLACED"
)

func KafkaSubscriberInit() {
	logger := l.New(os.Stdout, "kafka reader: ", 0)

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "OrderStatus",
		GroupID: "cart-ms-group",
		// assign the logger to the reader
		Logger: logger,
	})

	go readMessage(reader)
}

func readMessage(reader *kafka.Reader) {
	dao := mongodao.GetMongoDAO()

	ctx := context.Background()
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.WithError(err).Error("an error occurred while reading the kafka message")
		}

		userEmail := string(msg.Key)
		orderStatus := string(msg.Value)

		if orderStatus == ORDER_PLACED {
			err := dao.DeleteCart(ctx, userEmail)
			if err != nil {
				log.WithField("Error: ", err).Error("an error occurred while deleting the cart")
			}
		}
	}
}
