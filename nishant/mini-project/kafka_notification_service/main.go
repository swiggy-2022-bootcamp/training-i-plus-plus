package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	consumer "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/kafka_notification_service/consumer"
	email "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/kafka_notification_service/email"
)

func main() {
	f, err := os.OpenFile("notification_service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, f))
	log.Println("Logger setup!")

	con := consumer.NewConsumer("test_topic")
	con.Start()

	em := email.NewEmailer(con.Output)
	em.Start()

	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Listing to STDIN")
	for scanner.Scan() {
		s := strings.TrimRight(scanner.Text(), "\n")
		log.Println("From STDIN : " + s)

		if s == "exit" {
			log.Println("Exiting Program...")
			con.Stop()
			em.Stop()
			break
		}
	}

	if scanner.Err() != nil {
		// Handle error.
		con.Stop()
		em.Stop()
	}

}
