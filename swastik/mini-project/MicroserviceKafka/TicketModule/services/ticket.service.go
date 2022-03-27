package services

import "github.com/swastiksahoo153/ticket-module/models"

type TicketService interface {
	CreateTicket(*models.Ticket) error
	GetTicket(*string) (*models.Ticket, error)
	GetAll() ([]*models.Ticket, error)
	UpdateTicket(*models.Ticket) error
	DeleteTicket(*string) error
}