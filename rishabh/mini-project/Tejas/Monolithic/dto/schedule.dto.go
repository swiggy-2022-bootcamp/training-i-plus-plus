package dto

import (
	"tejas/models"
	"time"
)

type AddTrainScheduleDTO struct {
	Date       time.Time                 `json:"date"`
	Train      models.TrainsWithSchedule `json:"train"`
	TotalSeats int                       `json:"total_seats"`
}

type ReserveSeatDTO struct {
	From    string    `json:"from"`
	To      string    `json:"to"`
	Date    time.Time `json:"date"`
	TrainId int       `json:"train_id"`
}
