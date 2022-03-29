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
