package producer

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var brokers = []string{
	"localhost:9092",
}

func getRandomKey() []byte {
	var src = rand.NewSource(time.Now().UnixNano())
	return []byte(string(src.Int63()))
}

func getProducer(ctx context.Context, topic string, brokers []string) *kafka.Writer {
	// intialize the writer with the broker addresses, and the topic
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})
}

func WriteMsgToKafka(topic string, msg interface{}) (bool, error) {
	ctx := context.Background()
	jsonString, err := json.Marshal(msg)
	msgString := string(jsonString)

	if err != nil {
		return false, err
	}

	writer := getProducer(ctx, topic, brokers)
	defer writer.Close()

	err = writer.WriteMessages(ctx, kafka.Message{
		Key:   getRandomKey(),
		Value: []byte(msgString),
	})

	if err != nil {
		log.WithFields(logrus.Fields{"msg": msg}).Error("could not write message ")
		return false, err
	}

	return true, nil
}
