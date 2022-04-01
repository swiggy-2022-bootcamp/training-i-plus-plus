package services

import "github.com/swastiksahoo153/MicroserviceKafka/TrainModule/models"

type TrainService interface {
	CreateTrain(*models.Train) error
	GetTrain(*string) (*models.Train, error)
	GetAll() ([]*models.Train, error)
	UpdateTrain(*models.Train) error
	DeleteTrain(*string) error
}