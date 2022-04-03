package kafka

import (
	emailComms "comms/email"
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	logger "github.com/sirupsen/logrus"
)

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
}

type KafkaHandler struct {
	EmailCommsHandler *emailComms.EmailCommsHandler
}

func NewKafkaHandler(emailCommsHandler *emailComms.EmailCommsHandler) *KafkaHandler {
	return &KafkaHandler{
		EmailCommsHandler: emailCommsHandler,
	}
}

func (handler *KafkaHandler) Consume() {
	const methodName = "#kafka-consumer.Consume"
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
			logger.Error(fmt.Sprintf("%s : Consumer error: %v (%v)\n", methodName, err, msg))
			return
		}
		logger.Info(fmt.Sprintf("%s : Message on %s: %s\n", methodName, msg.TopicPartition, string(msg.Value)))
		var message map[string]interface{}
		json.Unmarshal(msg.Value, &message)
		ambassador_msg := fmt.Sprintf("You earned $%f from the link #%s", message["ambassador_revenue"].(float64), message["code"])
		adminMessage := fmt.Sprintf("Order #%f with a total of $%f has been completed", message["id"].(float64), message["admin_revenue"].(float64))
		handler.EmailCommsHandler.SendEmailComms(ambassador_msg, message["ambassador_email"].(string))
		handler.EmailCommsHandler.SendEmailComms(adminMessage, os.Getenv("ADMIN_EMAIL"))
	}
}
