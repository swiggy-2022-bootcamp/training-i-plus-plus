package service

import (
	"errors"
	"mini-project/bff/model"

	"github.com/google/uuid"
)

////////////////////////////////////////////////////////////////////////////////////

// StationService provides CRUD operations for managing stations.
type StationService interface {
	AddStation(model.Station) (model.Station, error)
	// RetrieveStationByID(string) (string, error)
	RetrieveAllStations() ([]model.Station, error)
	// UpdateStation(string) (string, error)
	// DeleteStation(string) (string, error)
}

////////////////////////////////////////////////////////////////////////////////////

// stationService is a concrete implementation of StationService
type stationService struct {
	stations []model.Station
}

func New() StationService {
	return &stationService{
		stations: []model.Station{},
	}
}

func (svc *stationService) AddStation(in model.Station) (model.Station, error) {
	// TODO add db implementation
	in.ID = uuid.New().String()
	in.Address.ID = uuid.New().String()

	svc.stations = append(svc.stations, in)
	return in, nil
}

func (svc *stationService) RetrieveAllStations() ([]model.Station, error) {
	// TODO add implementation
	return svc.stations, nil
}

// func (StationServiceImpl) UpdateStation(s string) (string, error) {
// 	// TODO add implementation
// 	return "", ErrEmpty
// }

// func (StationServiceImpl) DeleteStation(s string) (string, error) {
// 	// TODO add implementation
// 	return "", ErrEmpty
// }

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

////////////////////////////////////////////////////////////////////////////////////
