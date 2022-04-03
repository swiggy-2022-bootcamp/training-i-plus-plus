package notification

import (
	"fmt"
	"time"
)

type SMS struct {
	Body string
	To   int
}

const smsBuffer = 10
const numSmsSenderWorkers = 10

func SmsSender() chan<- SMS {

	chSms := make(chan SMS, smsBuffer)

	// start workers
	for i := 0; i < numSmsSenderWorkers; i++ {
		go smsSenderWorker(uint(i+1), chSms)
	}

	return chSms
}

func smsSenderWorker(id uint, chSms <-chan SMS) {

	for e := range chSms {
		fmt.Println("Sms Worker ", id, " sending sms '", e.Body, "'")
		//   simulate sending sms
		time.Sleep(500 * time.Millisecond)

	}
}
