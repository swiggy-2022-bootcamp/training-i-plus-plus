package main

import (
	"net/http"
	"os"

	"mini-project/bff/config"
	"mini-project/bff/service"
	"mini-project/bff/transport"

	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "station_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "station_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "station_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var svc service.StationService
	svc = service.StationServiceImpl{}
	svc = config.LoggingMiddleware{logger, svc}
	svc = config.InstrumentingMiddleware{requestCount, requestLatency, countResult, svc}

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
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
