package controllers

import (
	"time"
)

type Invoice struct {
	ID            string
	UserID        string
	BookingID     string
	TransactionID string
	CreatedAt     time.Time
}

type invoiceMethod interface {
	SendEmailInvoice()
	SendMobileInvoice()
}

func (invoice *Invoice) SendEmailInvoice() {
}

func (invoice *Invoice) SendSMSInvoice() {
}
