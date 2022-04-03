package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/segmentio/kafka-go"
)

const (
	topic                  = "TrainTicketLelo"
	brokerAddress          = "localhost:9092"
	UserServicePort        = "8001"
	ReservationServicePort = "8002"
	TrainServicePort       = "8003"
	AppHealthPort          = 8004
)

func main() {
	log.Print("App Health Service Running on Port ", AppHealthPort)

	go checkServiceHealth()

	ctx := context.Background()
	consume(ctx)
}

func Produce(ctx context.Context, key []byte, value []byte) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})
	err := w.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "consumer-group-0",
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		msgStr := string(msg.Value)
		if strings.Contains(msgStr, "Service is Down.") {
			color.Red.Println(msgStr)
		} else if strings.Contains(msgStr, "(critical)") || strings.Contains(msgStr, "(auth)") {
			color.Magenta.Println(msgStr)
		} else {
			color.Info.Println(string(msg.Value))
		}
	}
}

func checkServiceHealth() {
	host := "localhost"
	for {
		timeout := time.Duration(1 * time.Second)
		_, err := net.DialTimeout("tcp", host+":"+TrainServicePort, timeout)
		if err != nil {
			Produce(context.Background(), nil, []byte("Trains-Service is Down. "+time.Now().Local().String()))
		}

		timeout = time.Duration(1 * time.Second)
		_, err = net.DialTimeout("tcp", host+":"+ReservationServicePort, timeout)
		if err != nil {
			Produce(context.Background(), nil, []byte("Reservation-Service is Down. "+time.Now().Local().String()))
		}

		timeout = time.Duration(1 * time.Second)
		_, err = net.DialTimeout("tcp", host+":"+UserServicePort, timeout)
		if err != nil {
			Produce(context.Background(), nil, []byte("User-Service is Down. "+time.Now().Local().String()))
		}

		time.Sleep(5 * time.Second)
	}
}
