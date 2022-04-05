package infra

import (
	"context"
	"encoding/json"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"os"
	"panem/utils/logger"
)

const (
	brokerAddress = "localhost:29092"
)

type Consumer struct {
	reader *kafka.Reader
	ctx    context.Context
	cancel context.CancelFunc
	ack    chan bool
	Output chan orderInfo
}

type orderInfo struct {
	UserId      int     `json:"user_id"`
	OrderAmount float64 `json:"order_amount"`
	OrderId     int     `json:"order_id"`
}

func NewConsumer(topic string) *Consumer {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{brokerAddress},
		Topic:       topic,
		GroupID:     "my-group",
		StartOffset: kafka.LastOffset,
		//MinBytes:    1,
		// assign the logger to the reader
		Logger: l,
	})

	//create context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	return &Consumer{
		r,
		ctx,
		cancel,
		make(chan bool),
		make(chan orderInfo, 10),
	}
}

func (c Consumer) Start() {
	go c.poll()
}

func (c Consumer) poll() {
	for {
		msg, err := c.reader.ReadMessage(c.ctx)
		if err != nil {
			log.Println("could not read message " + err.Error())
			break
		}

		var oi orderInfo
		err = json.Unmarshal(msg.Value, &oi)
		if err != nil {
			logger.Error("Unable to parse payload")
			break
		}

		logger.Info(fmt.Sprintf("received: %v", oi))
		c.Output <- oi
	}
	c.reader.Close()
	c.ack <- true
}

func (c Consumer) Stop() {
	c.cancel()
	<-c.ack
	log.Println("Kafka Consumer Stopped")
}
