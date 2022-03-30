package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic                         = "swiggy-project-1"
	brokerAddress                 = "localhost:9092"
	REPORTING_SERVICE_SERVER_PORT = 5001
	INVENTORY_SERVICE_SERVER_PORT = "5002"
	ORDER_SERVICE_SERVER_PORT     = "5003"
	USER_SERVICE_SERVER_PORT      = "5004"
)

func main() {
	log.Print("Reporting Service: Starting server at port ", REPORTING_SERVICE_SERVER_PORT)

	go checkServiceHealth()

	// create a new context
	ctx := context.Background()
	consume(ctx)

	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	//go produce(ctx)

}

func Produce(ctx context.Context, key []byte, value []byte) {
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})
	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	//l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "consumer-group-0",
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println(string(msg.Value))
	}
}

func checkServiceHealth() {
	host := "localhost"
	for {
		timeout := time.Duration(1 * time.Second)
		_, err := net.DialTimeout("tcp", host+":"+INVENTORY_SERVICE_SERVER_PORT, timeout)
		if err != nil {
			Produce(context.Background(), nil, []byte("Inventory-Service is down (critical) "+time.Now().Local().String()))
		}

		timeout = time.Duration(1 * time.Second)
		_, err = net.DialTimeout("tcp", host+":"+ORDER_SERVICE_SERVER_PORT, timeout)
		if err != nil {
			Produce(context.Background(), nil, []byte("Order-Service is down (critical) "+time.Now().Local().String()))
		}

		timeout = time.Duration(1 * time.Second)
		_, err = net.DialTimeout("tcp", host+":"+USER_SERVICE_SERVER_PORT, timeout)
		if err != nil {
			Produce(context.Background(), nil, []byte("User-Service is down (critical) "+time.Now().Local().String()))
		}

		time.Sleep(5 * time.Second)
	}
}
