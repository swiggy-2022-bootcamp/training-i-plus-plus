package main

import (
	"log"
	"os"
	"os/signal"
	"fmt"
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/Shopify/sarama"
)

var (
	//List of brokers to connect
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	//Define Topic name
	topic = kingpin.Flag("topic", "topic name").Default("Example1").String()
	//Define Partition number
	partition  = kingpin.Flag("partition", "Partition number").Default("0").String()
	offsetType = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
	//Message counter start from
	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
)

func main() {
	fmt.Println("Consumer started")

	kingpin.Parse()
	//NewConfig returns a new configuration instance with same defaults
	config := sarama.NewConfig()
	// Consumer.Return.Errors setting to true,
	config.Consumer.Return.Errors = true
	// set brokerList in to variable
	brokers := *brokerList
	//NewConsumer creates a new consumer using the given broker addresses and configuration
	master, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()

	//Before you can start consuming a partition, you have to set expectations on it using ExpectConsumePartition.
	consumer, err := master.ConsumePartition(*topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Panic(err)
	}
	//The OS package provides the Signal interface to handle signals and has OS-specific implementations.
	signals := make(chan os.Signal, 1)
	//To notify Signals , we use the Notify function provided by the signal package
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				*messageCountStart++
				log.Println("Received messages :", string(msg.Key), string(msg.Value))
			case <-signals:
				log.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	//print produce messages in important topic
	log.Println("Processed", *messageCountStart, "messages")
}
