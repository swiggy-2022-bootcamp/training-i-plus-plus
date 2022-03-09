package main

import (
	"log"
	"net/http"

	"mini-project/bff/service"
	"mini-project/bff/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var svc service.StationService
	svc = service.StationServiceImpl{}

	addStationHandler := httptransport.NewServer(
		transport.MakeAddStationEndpoint(svc),
		transport.DecodeAddStationRequest,
		transport.EncodeResponse,
	)

	retrieveStationHandler := httptransport.NewServer(
		transport.MakeRetrieveStationEndpoint(svc),
		transport.DecodeRetrieveStationRequest,
		transport.EncodeResponse,
	)

	updateStationHandler := httptransport.NewServer(
		transport.MakeUpdateStationEndpoint(svc),
		transport.DecodeUpdateStationRequest,
		transport.EncodeResponse,
	)

	deleteStationHandler := httptransport.NewServer(
		transport.MakeDeleteStationEndpoint(svc),
		transport.DecodeDeleteStationRequest,
		transport.EncodeResponse,
	)

	http.Handle("/addStation", addStationHandler)
	http.Handle("/updateStation", updateStationHandler)
	http.Handle("/retrieveStation", retrieveStationHandler)
	http.Handle("/deleteStation", deleteStationHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
