package config

import (
  "fmt"
  "net/smtp"
)

func SendMail(message []byte) {

  // Sender data.
  from := "udaysonubakka123@gmail.com"
  password := "Udayyadusonu@1"

  // Receiver email address.
  to := []string{
    "udaysonubakka199@gmail.com",
  }

  // smtp server configuration.
  smtpHost := "smtp.gmail.com"
  smtpPort := "587"

  // Message.
  //message := []byte(messageString)
  
  // Authentication.
  auth := smtp.PlainAuth("", from, password, smtpHost)
  
  // Sending email.
  err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("######      Email Sent Successfully!       #######")
}
