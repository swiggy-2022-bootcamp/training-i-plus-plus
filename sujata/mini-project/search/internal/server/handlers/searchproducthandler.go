package handlers

import (
	"net/http"
	"search/internal/errors"
	"search/internal/services"
	"search/util"

	log "github.com/sirupsen/logrus"
)

func SearchProductHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var goErr error
		ctx := req.Context()

		log.Info("Inside search product handler")
		// extract searched product from query
		params := req.URL.Query()
		searchedProduct := params["product"][0]

		log.Info(searchedProduct)
		service := services.GetSearchProductService()

		// Process the request
		resp, err := service.ProcessRequest(ctx, searchedProduct)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, goErr = w.Write(resp)
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while writing the response")
			http.Error(w, errors.InternalError.ErrorMessage, errors.InternalError.HttpResponseCode)
			return
		}
	}
}
