package goKafka

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-kafka-microservice/ListingService/models"
)

type GoKafkaServicesImpl struct {
	Consumer *kafka.Consumer
}

func NewGokafkaServiceImpl(consumer *kafka.Consumer) *GoKafkaServicesImpl {
	return &GoKafkaServicesImpl{
		Consumer: consumer,
	}
}
func (ks *GoKafkaServicesImpl) ReadMessage(topic string) (interface{}, error) {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Subscribe to given topic
	ks.Consumer.SubscribeTopics([]string{topic}, nil)

	// Read Message
	var products []models.Product
	for {
		msg, err := ks.Consumer.ReadMessage(100 * time.Millisecond)
		// ks.Consumer.Re

		if err == nil {
			var _product models.Product
			p := []byte(msg.Value)
			fmt.Println(p)
			err := json.Unmarshal(p, &_product)
			if err != nil {
				break
			}
			products = append(products, _product)
		} else {
			break
		}
	}

	ks.Consumer.Close()

	return products, nil
}
