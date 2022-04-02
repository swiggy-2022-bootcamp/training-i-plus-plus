package email

import (
	"context"
	"encoding/json"
	"log"
)

type EmailParser struct {
	in     chan string
	Out    chan Email
	ctx    context.Context
	cancel context.CancelFunc
	Done   chan bool
}

func NewEmailParser(in chan string, ctx context.Context) *EmailParser {

	ctx, cancel := context.WithCancel(ctx)

	ep := &EmailParser{
		in,
		make(chan Email, 10),
		ctx,
		cancel,
		make(chan bool),
	}

	go ep.loop()
	return ep
}

func (ep *EmailParser) loop() {
	log.Println("Starting EmailParser")
	for {
		select {
		case str := <-ep.in:
			log.Println("recevied str : " + str)
			ep.Out <- parse(str)

		case <-ep.ctx.Done():
			log.Println("Stopping Email Parser ")
			ep.Done <- true
			return
		}
	}

}

func parse(str string) Email {
	var email Email
	if err := json.Unmarshal([]byte(str), &email); err != nil {
		log.Panicf("error while parsing email : %s", err.Error())
	}
	return email
}
