package main

import (
	db "comms/database"
	emailComms "comms/email"
	kafkaService "comms/kafka"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var emailCommsHandler *emailComms.EmailCommsHandler
var kafkaHandler *kafkaService.KafkaHandler

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
	kafkaHandler = kafkaService.NewKafkaHandler(emailCommsHandler)

}

func main() {
	kafkaHandler.Consume()
}
