package models

type Charge struct {
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receiptMail"`
	TicketID     string `json:"ticketid"`
}
