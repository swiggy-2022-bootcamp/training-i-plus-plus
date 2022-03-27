package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-mongo-api/logger"
	"gin-mongo-api/models"
	"gin-mongo-api/repository"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	topic1        = "bookedticket"
	topic2        = "availticket"
	topic3        = "train"
	brokerAddress = "localhost:9092"
)

var bookedticketrepo repository.BookedTicketRepository
var availticketrepo repository.AvailTicketRepository

func Produce_booked_ticket_for_avail(trainid primitive.ObjectID, update bool) {
	logger.Klog("kafka producer booking ticket")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic1,
	})
	var updatestring string
	if update {
		updatestring = "increment"
	} else {
		updatestring = "decrement"
	}
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(trainid.String()),
		Value: []byte(updatestring),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func Consume_booked_ticket_for_avail() {
	logger.Klog("kafka producer booking ticket")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic1,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		//fmt.Println("received: ", string(msg.Value))
		trainid, err := primitive.ObjectIDFromHex(string(msg.Key))

		if err != nil {
			panic("incorrect train id " + err.Error())
		}

		availticket, err := availticketrepo.ReadTrainId(trainid)

		if err != nil {
			panic("could not find the train " + err.Error())
		}

		if string(msg.Value) == "increment" {
			availticket.Capacity += 1
		} else {
			availticket.Capacity -= 1
		}

		_, err = availticketrepo.Update(availticket, availticket.Id)
		if err != nil {
			panic("could not update booked ticket " + err.Error())
		}
	}
}

func Produce_avail_ticket(nat models.AvailTicket) {
	logger.Klog("kafka producer avail ticket")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic2,
	})

	bytes, _ := json.Marshal(nat)
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(nat.Id.String()),
		Value: []byte(bytes),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func Consume_avail_ticket() {
	logger.Klog("kafka consumer avail ticket")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic2,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		natr := models.AvailTicket{}
		json.Unmarshal([]byte(msg.Value), &natr)
		res := fmt.Sprintf("Inserted new available ticket %#v", natr)
		logger.ALog(res)
	}
}

func Produce_train(nt models.Train) {
	logger.Klog("kafka producer avail ticket")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic3,
	})

	bytes, _ := json.Marshal(nt)
	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(nt.Id.String()),
		Value: []byte(bytes),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func Consume_train() {
	logger.Klog("kafka consumer avail ticket")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic3,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		ntr := models.Train{}
		json.Unmarshal([]byte(msg.Value), &ntr)
		res := fmt.Sprintf("Inserted new available ticket %#v", ntr)
		logger.ALog(res)
	}
}
