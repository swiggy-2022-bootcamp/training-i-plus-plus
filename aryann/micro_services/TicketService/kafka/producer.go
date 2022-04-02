package kafka

import (
	"TicketService/logger"
	"TicketService/models"
)

const (
	topic1        = "purchasedticket"
	topic2        = "ticket"
	topic3        = "train"
	brokerAddress = "localhost:9092"
)

var logger8 = logger.NewLoggerService("Kafka Producer")

func Produce_ticket(nat models.Ticket) {
	logger8.Log("Produce Ticket")
	// w := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers: []string{brokerAddress},
	// 	Topic:   topic2,
	// })

	// bytes, _ := json.Marshal(nat)
	// err := w.WriteMessages(context.Background(), kafka.Message{
	// 	Key:   []byte(nat.ID.String()),
	// 	Value: []byte(bytes),
	// })
	// if err != nil {
	// 	panic("could not write message " + err.Error())
	// }
}
