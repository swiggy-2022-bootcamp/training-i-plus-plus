package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
	"context"
)

func StartKafka(){
	conf := kafka.ReaderConfig{
		Brokers: []string {"localhost:9092"},
		Topic: "mytopic",
		GroupID: "g1",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil{
			fmt.Println("Some err occured", err)
			continue
		}
		fmt.Println("Message is : ", string(m.Value))
	}
}

func main(){
	fmt.Println("working...")
	go StartKafka()

	fmt.Println("kafka has been started...")

	time.Sleep(10 * time.Minute)
}