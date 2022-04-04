package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
 	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message-log"
	brokerAddress = "localhost:9092"
)
var I=0;
var Channel=make(chan string)
func InitKafka() {
	// create a new context
	fmt.Println("Kafka Started.......")
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	Produce(ctx)
	 
}

func Produce(ctx context.Context) {
	// initialize a counter
	 
	fmt.Println("Kafka message writer Initiated......")
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})
	//var message string
	for message:=range Channel{
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		fmt.Println("writes:", I)
		fmt.Println("message: ",message)
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(I)),
			// create an arbitrary message payload for the value
			
			Value: []byte(message),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		
		I++
		// sleep for a second
	//	time.Sleep(time.Second)
	}
}
