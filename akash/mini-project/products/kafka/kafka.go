package kafka

import (
	//"encoding/json"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"products.akash.com/model"
	"time"
)

func CreateComment(buyRequest *model.BuyRequest) {

	topic := "buy-requests"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		fmt.Println("failed to dial leader:", err)
		os.Exit(1)
	}

	readBytes := make([]byte, 200)
	if err != nil {
		log.Fatal("Failed here ", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   []byte("0"),
			Value: readBytes},
	)
	if err != nil {
		fmt.Println("failed to write messages:", err)
	}

	time.Sleep(1 * time.Second)

	if err := conn.Close(); err != nil {
		fmt.Println("failed to close writer:", err)
	}
	fmt.Println("Conn Close")
	fmt.Println(err)

}
