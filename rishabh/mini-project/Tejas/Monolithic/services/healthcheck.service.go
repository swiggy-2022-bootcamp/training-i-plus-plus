package services

import (
	"context"
	"tejas/configs"
	"time"
)

type ServiceResponse struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

type HealthCheckResponse struct {
	Status   string            `json:"status"`
	Services []ServiceResponse `json:"services"`
}

func HealthCheck() HealthCheckResponse {
	response := HealthCheckResponse{
		Status: "up",
	}
	return response
}

func DeepHealthCheck() HealthCheckResponse {
	response := HealthCheckResponse{}
	response.Status = "up"
	response.Services = append(response.Services, DBHealthCheck())

	for _, services := range response.Services {
		if services.Status == "down" {
			response.Status = "down"
			break
		}
	}
	return response
}

func DBHealthCheck() ServiceResponse {
	response := ServiceResponse{}
	response.Name = "MongoDB"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := configs.DB.Ping(ctx, nil)

	if err != nil {
		response.Status = "down"
		response.Error = err.Error()
	} else {
		response.Status = "up"
		response.Error = ""
	}

	return response
}
