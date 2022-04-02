package main

import (
	"context"
	"invoiceService/controllers"
	consumer "invoiceService/kafkaconsumer"
	"invoiceService/logger"
	"time"

	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

func main() {
	for {
		ctx := context.Background()

		// the `ReadMessage` method blocks until we receive the next event
		reader := consumer.GetKafkaReader()
		msg, err := reader.ReadMessage(ctx)
		invoice := controllers.Invoice{
			msg.ID, msg.UserID, msg.BookingID, msg.TransactionID, time.Now().Local(),
		}

		if err != nil {
			panic("could not read message " + err.Error())
		}

		log.WithFields(logrus.Fields{
			"inv": invoice,
		}).Info("invoice fetched")

		go invoice.SendEmailInvoice()
		go invoice.SendSMSInvoice()
		log.Info("invoice sent both via email and sms.")
	}
}
