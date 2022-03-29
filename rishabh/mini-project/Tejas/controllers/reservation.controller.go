package controllers

import (
	"context"
	"tejas/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TicketGeneration(userID primitive.ObjectID, trainID int, from, to string, date time.Time, transactionID string, seatNumber int) (models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var ticket models.Reservation
	ticket.PNR = primitive.NewObjectID()
	ticket.UserId = userID
	ticket.TrainId = trainID
	ticket.FromStationCode = from
	ticket.ToStationCode = to
	ticket.Date = date
	ticket.TransactionId = transactionID
	ticket.Status = "success"
	ticket.SeatNumber = seatNumber
	_, err := models.ReservationCollection.InsertOne(ctx, ticket)
	if err != nil {
		return models.Reservation{}, err
	}
	return ticket, nil
}
