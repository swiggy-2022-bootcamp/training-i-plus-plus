package handlers

import (
	"cart/internal/errors"
	"cart/internal/services"
	"cart/util"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetCartHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		role, userId := util.ExtractDetailsFromToken(req)
		if role == "SELLER" {
			log.Error("unauthorized user, user if of type SELLER")
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		service := services.GetGetCartService()

		// Process the request
		cart, err := service.ProcessRequest(ctx, userId)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		respBytes, goErr := json.Marshal(cart)
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while marhsalling the cart response")
			http.Error(w, errors.MarshalError.ErrorMessage, errors.MarshalError.HttpResponseCode)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(respBytes)
	}
}
