package handlers

import (
	"encoding/json"
	"net/http"
	"order/internal/services"
	"order/util"

	log "github.com/sirupsen/logrus"
)

func GetOrderStatusHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		role, email := util.ExtractDetailsFromToken(req)
		if role == "SELLER" {
			log.Error("unauthorized user, user if of type SELLER")
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		service := services.GetGetOrderService()

		// Process the request
		allOrders, err := service.ProcessRequest(ctx, email)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		respBytes, goErr := json.Marshal(allOrders)
		if goErr != nil {
			log.WithError(goErr).Error("an error occured while marshalling the response")
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(respBytes)
	}
}
