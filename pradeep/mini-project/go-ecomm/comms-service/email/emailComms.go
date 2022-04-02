package mail

import (
	db "comms/database"
	models "comms/models"
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/rs/xid"
	logger "github.com/sirupsen/logrus"
)

var host = os.Getenv("EMAIL_HOST")
var port = os.Getenv("EMAIL_PORT")
var senders_email = os.Getenv("SENDERS_EMAIL")

type EmailCommsHandler struct {
	db *db.DbHandler
}

func NewEmailCommsHandler(dbInstance *db.DbHandler) *EmailCommsHandler {
	return &EmailCommsHandler{
		db: dbInstance,
	}
}

func (handler *EmailCommsHandler) SendEmailComms(message string, recipient string) {
	const methodName = "#SendEmailComms"
	logger.Info(fmt.Sprintf("%s : Request recieved to send email communication with message %s to recipient %s", methodName, message, recipient))
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"), host)
	err := smtp.SendMail(host+":"+port, auth, senders_email, []string{recipient}, []byte(message))
	if err != nil {
		logger.Error("%s : Error sending email to recipient: %s with error: %s", methodName, recipient, err)
		return
	}
	logger.Info(fmt.Sprintf("%s : Email successfully sent to recipient: %s", methodName, recipient))
	emailEntity := models.Email{
		ID:          xid.New().String(),
		Email:       recipient,
		Message:     message,
		PublishedAt: time.Now(),
	}
	data, _ := json.Marshal(emailEntity)
	_, err = handler.db.Collection.InsertOne(handler.db.Ctx, emailEntity)
	if err != nil {
		logger.Error(fmt.Sprintf("%s : Error occured while saving email entity : %s to the database.", methodName, string(data)))
		return
	}
	logger.Info(fmt.Sprintf("%s : Email entity: %s successfully saved to database", methodName, string(data)))
}
