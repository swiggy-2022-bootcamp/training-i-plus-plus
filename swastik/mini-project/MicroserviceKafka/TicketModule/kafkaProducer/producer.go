package kafkaProducer

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
    "fmt"
	"strconv"
    "context"
	"log"
	"os"
)
const (
    brokerAddress = "localhost:9092"
)


func Produce(ctx context.Context, topic string, msg interface{}) (bool, error) {
    i := 0
    l := log.New(os.Stdout, "kafka writer: ", 0)
    w := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{brokerAddress},
        Topic:   topic,
        Logger: l,
    })

    jsonString, err := json.Marshal(msg)
	msgString := string(jsonString)

    if err != nil {
		return false, err
	}

    for _, word := range []string{string(msgString)} {

		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			Value: []byte(word),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		fmt.Println("writes:", msgString)
        i++
	}

	return true, nil
}
