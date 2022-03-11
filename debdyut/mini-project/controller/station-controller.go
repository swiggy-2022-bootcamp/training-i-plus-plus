package controller

import (
	"log"
	"mini-project/bff/model"
	"mini-project/bff/service"

	"github.com/gin-gonic/gin"
)

type StationController interface {
	AddStation(ctx *gin.Context) (model.Station, error)
	RetrieveAllStations() []model.Station
}

type stationController struct {
	svc service.StationService
}

func New(svc service.StationService) StationController {
	return &stationController{
		svc: svc,
	}
}

func (c *stationController) AddStation(ctx *gin.Context) (model.Station, error) {
	var station model.Station
	err := ctx.ShouldBindJSON(&station)
	if err != nil {
		return model.Station{}, err
	}

	station, err = c.svc.AddStation(station)
	if err != nil {
		log.Fatalf("Unable to add station")
	}
	return station, nil
}

func (c *stationController) RetrieveAllStations() []model.Station {
	stations, err := c.svc.RetrieveAllStations()
	if err != nil {
		log.Fatalf("Unable to retrieve all stations")
	}
	return stations
}
