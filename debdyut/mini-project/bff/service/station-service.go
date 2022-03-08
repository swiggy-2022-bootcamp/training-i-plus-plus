package service

import (
	"errors"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////

// StationService provides CRUD operations for managing stations.
type StationService interface {
	AddStation(string) (string, error)
	RetrieveStation(string) (string, error)
	UpdateStation(string) (string, error)
	DeleteStation(string) (string, error)
}

////////////////////////////////////////////////////////////////////////////////////

// stationService is a concrete implementation of StationService
type StationServiceImpl struct{}

func (StationServiceImpl) AddStation(s string) (string, error) {
	// TODO add implementation
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StationServiceImpl) RetrieveStation(s string) (string, error) {
	// TODO add implementation
	return "", ErrEmpty
}

func (StationServiceImpl) UpdateStation(s string) (string, error) {
	// TODO add implementation
	return "", ErrEmpty
}

func (StationServiceImpl) DeleteStation(s string) (string, error) {
	// TODO add implementation
	return "", ErrEmpty
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

////////////////////////////////////////////////////////////////////////////////////
