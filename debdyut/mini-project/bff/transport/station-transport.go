package transport

import (
	"context"
	"encoding/json"
	"mini-project/bff/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

////////////////////////////////////////////////////////////////////////////////////

// For each method, we define request and response structs
type addStationRequest struct {
	S string `json:"s"`
}

type addStationResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type retrieveStationRequest struct {
	S string `json:"s"`
}

type retrieveStationResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type updateStationRequest struct {
	S string `json:"s"`
}

type updateStationResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type deleteStationRequest struct {
	S string `json:"s"`
}

type deleteStationResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

////////////////////////////////////////////////////////////////////////////////////

func MakeAddStationEndpoint(svc service.StationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addStationRequest)
		v, err := svc.AddStation(req.S)
		if err != nil {
			return addStationResponse{v, err.Error()}, nil
		}
		return addStationResponse{v, ""}, nil
	}
}

func MakeUpdateStationEndpoint(svc service.StationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(updateStationRequest)
		v, err := svc.UpdateStation(req.S)
		if err != nil {
			return updateStationResponse{v, err.Error()}, nil
		}
		return updateStationResponse{v, ""}, nil
	}
}

func MakeRetrieveStationEndpoint(svc service.StationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(retrieveStationRequest)
		v, err := svc.RetrieveStation(req.S)
		if err != nil {
			return retrieveStationResponse{v, err.Error()}, nil
		}
		return retrieveStationResponse{v, ""}, nil
	}
}

func MakeDeleteStationEndpoint(svc service.StationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteStationRequest)
		v, err := svc.DeleteStation(req.S)
		if err != nil {
			return deleteStationResponse{v, err.Error()}, nil
		}
		return deleteStationResponse{v, ""}, nil
	}
}

////////////////////////////////////////////////////////////////////////////////////

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeAddStationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addStationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeUpdateStationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request updateStationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRetrieveStationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request retrieveStationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeDeleteStationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteStationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
