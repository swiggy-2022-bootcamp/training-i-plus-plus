package kafka

import (
	"TrainService/logger"
	"TrainService/models"
)

const (
	topic1        = "train"
	brokerAddress = "localhost:9092"
)

var logger8 = logger.NewLoggerService("Kafka Producer")

func Produce_train(nt models.Train) {
	logger8.Log("Produce Train")
	// w := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers: []string{brokerAddress},
	// 	Topic:   topic1,
	// })

	// bytes, _ := json.Marshal(nt)
	// err := w.WriteMessages(context.Background(), kafka.Message{
	// 	Key:   []byte(nt.Destination),
	// 	Value: []byte(bytes),
	// })
	// if err != nil {
	// 	panic("could not write message " + err.Error())
	// }
}
