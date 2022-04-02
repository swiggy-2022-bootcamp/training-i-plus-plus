package handlers

import (
	"encoding/json"
	"net/http"
	"order/internal/services"
	"order/util"

	log "github.com/sirupsen/logrus"
)

func CreateOrderHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		role, email := util.ExtractDetailsFromToken(req)
		if role == "SELLER" {
			log.Error("unauthorized user, user if of type SELLER")
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		service := services.GetCreateOrderService()

		reqToken := req.Header.Get("Authorization")
		resp, err := service.ProcessRequest(ctx, email, reqToken)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		respBytes, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respBytes)
	}
}
