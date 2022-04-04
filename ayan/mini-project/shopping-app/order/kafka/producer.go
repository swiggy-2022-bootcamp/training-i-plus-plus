package kafka

import (
	"fmt"
	"order/utils/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaProducer() *kafka.Producer {
	logger.Info("Starting producer...")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logger.Error(fmt.Sprintf("Delivery failed: %v\n", ev.TopicPartition))
				} else {
					logger.Info(fmt.Sprintf("Delivered message to %v\n", ev.TopicPartition))
				}
			}
		}
	}()

	return p
}

func Produce(p *kafka.Producer, messages string, topicName string) {

	// Produce messages to topic (asynchronously)
	topic := topicName
	data := []byte(messages)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          data,
	}, nil)

}
