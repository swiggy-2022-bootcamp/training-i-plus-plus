package main

import (
	db "comms/database"
	emailComms "comms/email"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var emailCommsHandler *emailComms.EmailCommsHandler

func init() {
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
	os.Setenv("SENDERS_EMAIL", "no-reply@email.com")
	os.Setenv("MONGO_DATABASE", "comms")
	os.Setenv("ADMIN_EMAIL", "admin@admin.com")

	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("emails")
	log.Println("Connected to MongoDB")
	dbHandler := db.NewDbHandler(collection, ctx)
	emailCommsHandler = emailComms.NewEmailCommsHandler(dbHandler)

}

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		"sasl.username":     os.Getenv("SASL_USERNAME"),
		"sasl.password":     os.Getenv("SASL_PASSWORD"),
		"sasl.mechanism":    os.Getenv("SASL_MECHANISM"),
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
		"ssl.ca.location":   "/usr/local/etc/openssl@1.1/cert.pem",
	})

	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{os.Getenv("KAFKA_TOPIC")}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			return
		}
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		var message map[string]interface{}
		json.Unmarshal(msg.Value, &message)
		ambassador_msg := fmt.Sprintf("You earned $%f from the link #%s", message["ambassador_revenue"].(float64), message["code"])
		adminMessage := fmt.Sprintf("Order #%f with a total of $%f has been completed", message["id"].(float64), message["admin_revenue"].(float64))
		emailCommsHandler.SendEmailComms(ambassador_msg, message["ambassador_email"].(string))
		emailCommsHandler.SendEmailComms(adminMessage, os.Getenv("ADMIN_EMAIL"))
	}

}
