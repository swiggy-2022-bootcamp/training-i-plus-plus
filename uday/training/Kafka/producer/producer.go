package main

import (
	"fmt"
	"log"
	"bufio"
	"os"  
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/Shopify/sarama"
)

var (
	//List of brokers to connect
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	//create a Topic
	topic = kingpin.Flag("topic", "Example1").Default("Example1").String()
	//retry limit
	maxRetry = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
)

func main() {
	fmt.Println("Producer started...")
	kingpin.Parse()
	//NewConfig returns a new configuration instance with same defaults
	config := sarama.NewConfig()
	//Wait for the response after all copies of the server are successfully saved.
	config.Producer.RequiredAcks = sarama.WaitForAll
	//Allowing retries without setting max.in.flight.requests.per.connection to 1
	config.Producer.Retry.Max = *maxRetry
	//Return.Successes must be true to be used in a SyncProducer's config
	config.Producer.Return.Successes = true
	// creates a new SyncProducer using the given broker addresses and configuration.
	producer, err := sarama.NewSyncProducer(*brokerList, config)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()
	//set message for produce to topic
	msg := &sarama.ProducerMessage{
		Topic: *topic,
		Value: sarama.StringEncoder("Something Cool"),
	}
	// Produce message in to a topic
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	var message string
	for{
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		message = scanner.Text()
	

		msg = &sarama.ProducerMessage{
			Topic: *topic,
			Value: sarama.StringEncoder(message),
		}
		// Produce message in to a topic
		partition, offset, err = producer.SendMessage(msg)
		if err != nil {
			log.Panic(err)
		}
	}
	
	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)
}
