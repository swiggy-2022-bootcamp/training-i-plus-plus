package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	kafka "github.com/segmentio/kafka-go"

	entity "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/doctor_service/api_entities"
)

const (
	brokerAddress = "localhost:29092"
)

type Email struct {
	To      string `json:"to"`
	Subject string `json:"sub"`
	Msg     string `json:"msg"`
}

type Producer struct {
	w      *kafka.Writer
	ctx    context.Context
	cancel context.CancelFunc
}

var (
	Notifier = NewProducer("test_topic")
)

func NewProducer(topic string) *Producer {

	ctx, cancel := context.WithCancel(context.Background())

	l := log.New(os.Stdout, "kafka writer: ", 0)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	return &Producer{
		w,
		ctx,
		cancel,
	}
}

func (p *Producer) SendAppointmentEmail(app entity.AppointmentRequest, user entity.User) {

	msg := fmt.Sprintf(" Doctor = %s, Patient = %s, From = %s, To = %s",
		app.Doctor, user.Name, app.From, app.To)

	toSend := Email{
		user.Email,
		"Appointment created",
		msg,
	}

	marshelled, err := json.Marshal(toSend)

	if err != nil {
		log.Println("Error marshalling email")
		return
	}

	p.send([]byte(user.UserId), marshelled)
}

func (p *Producer) send(key, value []byte) {
	err := p.w.WriteMessages(p.ctx, kafka.Message{
		Key: key,
		// create an arbitrary message payload for the value
		Value: value,
	})
	if err != nil {
		log.Println("Error could not write message " + err.Error())
	}
}
