package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	consumer "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/kafka_notification_service/consumer"
	email "github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/kafka_notification_service/email"
)

func main() {

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
