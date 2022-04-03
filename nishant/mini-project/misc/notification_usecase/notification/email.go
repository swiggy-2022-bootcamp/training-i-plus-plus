package notification

import (
	"fmt"
	"time"
)

type Email struct {
	Title string
	Body  string
	To    string
}

const emailBuffer = 10
const numEmailSenderWorkers = 10

func EmailSender() chan<- Email {

	chEmail := make(chan Email, emailBuffer)

	// start workers
	for i := 0; i < numEmailSenderWorkers; i++ {
		go emailSenderWorker(uint(i+1), chEmail)
	}

	return chEmail
}

func emailSenderWorker(id uint, chEmail <-chan Email) {

	for e := range chEmail {
		fmt.Println("Email Worker ", id, " sending email '", e.Body, "'")
		//   simulate sending email
		time.Sleep(500 * time.Millisecond)

	}
}
