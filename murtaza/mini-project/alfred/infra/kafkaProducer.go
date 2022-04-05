package infra

import (
	"alfred/utils/logger"
	"context"
	"encoding/json"
	"log"
	"os"

	kafka "github.com/segmentio/kafka-go"
)

const (
	brokerAddress = "localhost:29092"
)

type msgBody struct {
	UserId      int     `json:"user_id"`
	OrderAmount float64 `json:"order_amount"`
	OrderId     int     `json:"order_id"`
}

type Producer struct {
	w      *kafka.Writer
	ctx    context.Context
	cancel context.CancelFunc
}

func NewProducer(topic string) *Producer {

	ctx, cancel := context.WithCancel(context.Background())

	l := log.New(os.Stdout, "kafka writer: ", 0)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	return &Producer{
		w,
		ctx,
		cancel,
	}
}

func (p *Producer) SendOrderAmount(orderId int, userId int, orderAmount float64) {

	toSend := msgBody{
		userId,
		orderAmount,
		orderId,
	}

	marshalled, err := json.Marshal(toSend)

	if err != nil {
		logger.Error(err.Error())
		return
	}

	p.send([]byte(string(rune(userId))), marshalled)
}

func (p *Producer) send(key, value []byte) {
	err := p.w.WriteMessages(p.ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		log.Println("Error could not write message " + err.Error())
	}
}
