package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

const (
	topic                         = "swiggy-project-1"
	brokerAddress                 = "localhost:9092"
	REPORTING_SERVICE_SERVER_PORT = 5001
)

func main() {
	router := mux.NewRouter()
	log.Print("Reporting Service: Starting server at port ", REPORTING_SERVICE_SERVER_PORT)
	http.ListenAndServe(":5001", router)

	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	//go produce(ctx)
	consume(ctx)
}

func consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "consumer-group-0",
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
		fmt.Println(string(msg.Value))
	}
}
