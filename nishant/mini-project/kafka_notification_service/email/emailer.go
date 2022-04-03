package email

import (
	"context"
	"log"
	"net/smtp"
)

const (
	from     = "imneo47@yahoo.com"
	password = "Nishant@7"
	smtpHost = "smtp.mail.yahoo.com"
	smtpPort = "587"
)

type Emailer struct {
	in     chan Email
	ctx    context.Context
	cancel context.CancelFunc
	ep     *EmailParser
	auth   smtp.Auth
}

func NewEmailer(in_str chan string) *Emailer {

	ctx, cancel := context.WithCancel(context.Background())
	ep := NewEmailParser(in_str, ctx)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	return &Emailer{
		ep.Out,
		ctx,
		cancel,
		ep,
		auth,
	}
}

func (e *Emailer) Start() {
	go e.listen()
}

func (e *Emailer) Stop() {
	log.Println("Stopping Emailer...")
	e.cancel()
	<-e.ep.Done
	log.Println("Stopped Emailer")
}

func (e *Emailer) listen() {
	for {
		select {
		case email := <-e.in:
			e.sendEmail(email)
		case <-e.ctx.Done():
			return
		}
	}
}

func (e *Emailer) sendEmail(email Email) {
	log.Printf("Sending email %v", email)

	// Sending email.
	// err := smtp.SendMail(smtpHost+":"+smtpPort, e.auth, from, []string{email.To}, []byte(email.Msg))
	// if err != nil {
	// 	log.Println("error while ending email :" + err.Error())
	// 	return
	// }
	// log.Println("Email Sent Successfully!")
}
