package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	os.Setenv("BOOTSTRAP_SERVERS", "pkc-xrnwx.asia-south2.gcp.confluent.cloud:9092")
	os.Setenv("SECURITY_PROTOCOL", "SASL_SSL")
	os.Setenv("SASL_USERNAME", "7ABBHKKEMPQ3SDID")
	os.Setenv("SASL_PASSWORD", "9diVkQpVROzNJroxZW/gPcE+EHWJ+vYO0vd6XFmHAFO/EzaKtw5y8e7CYG31iYIj")
	os.Setenv("SASL_MECHANISM", "PLAIN")
	os.Setenv("KAFKA_TOPIC", "email_topic")
	os.Setenv("EMAIL_HOST", "smtp.mailtrap.io")
	os.Setenv("EMAIL_PORT", "2525")
	os.Setenv("EMAIL_USERNAME", "a09d873245007f")
	os.Setenv("EMAIL_PASSWORD", "47971650eb854b")
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		"sasl.username":     os.Getenv("SASL_USERNAME"),
		"sasl.password":     os.Getenv("SASL_PASSWORD"),
		"sasl.mechanism":    os.Getenv("SASL_MECHANISM"),
		"ssl.ca.location":   "/usr/local/etc/openssl@1.1/cert.pem",
	})
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	message := map[string]interface{}{
		"id":                 23,
		"ambassador_revenue": 230,
		"ambassador_email":   "pswaldia1@gmail.com",
		"admin_revenue":      500,
		"code":               "asaPQXd",
	}
	value, _ := json.Marshal(message)
	// Produce messages to topic (asynchronously)
	topic := os.Getenv("KAFKA_TOPIC")
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	// Wait for message deliveries before shutting down
	producer.Flush(15 * 1000)
}
