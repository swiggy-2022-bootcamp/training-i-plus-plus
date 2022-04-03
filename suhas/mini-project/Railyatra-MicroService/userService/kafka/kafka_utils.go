package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	log "userService/logger"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	topic1        = "bookedticket"
	topic2        = "availticket"
	topic3        = "train"
	brokerAddress = "localhost:9092"
)

var (
	errLog = log.ErrorLogger.Println
	infLog = log.InfoLogger.Println
)

type Train struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Station1 string             `json:"station1,omitempty" validate:"required"`
	Station2 string             `json:"station2,omitempty" validate:"required"`
}

type AvailTicket struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	Train_id       primitive.ObjectID `json:"trainid,omitempty"`
	Capacity       int                `json:"capacity"`
	Price          int                `json:"price"`
	Departure      string             `json:"departure"`
	Arrival        string             `json:"arrival"`
	Departure_time string             `json:"departure_time"`
	Arrival_time   string             `json:"arrival_time"`
}

//var bookedticketrepo repository.BookedTicketRepository
//var availticketrepo repository.AvailTicketRepository

func Produce_booked_ticket_for_avail(trainid primitive.ObjectID, update bool) {
	infLog("kafka producer booking ticket")
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
		errLog("could not write message " + err.Error())
		panic("could not write message " + err.Error())
	}
}

func Consume_avail_ticket() {
	errLog("kafka consumer avail ticket")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic2,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			errLog("could not read message " + err.Error())
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		natr := AvailTicket{}
		json.Unmarshal([]byte(msg.Value), &natr)
		res := fmt.Sprintf("Inserted new available ticket %#v", natr)
		infLog(res)
	}
}

func Consume_train() {
	errLog("kafka consumer avail ticket")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic3,
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			errLog("could not read message " + err.Error())
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		ntr := Train{}
		json.Unmarshal([]byte(msg.Value), &ntr)
		res := fmt.Sprintf("Inserted new available ticket %#v", ntr)
		infLog(res)
	}
}
