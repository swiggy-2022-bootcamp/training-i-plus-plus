package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	model "product/internal/dao/mongodao/models"
	"product/internal/errors"
	"product/internal/services"
	"product/util"

	log "github.com/sirupsen/logrus"
)

const (
	BUYER string = "BUYER"
)

func GetProductsHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Info("Inside get product handler")
		ctx := req.Context()

		// read request body
		b, goErr := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if goErr != nil {
			log.WithError(goErr).Error("an error occurrec while reading the request body")
			http.Error(w, goErr.Error(), http.StatusInternalServerError)
			return
		}

		// unmarshal the request to Product model
		var productIds model.ProductIds
		goErr = json.Unmarshal(b, &productIds)
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while unmarshalling the request")
			http.Error(w, errors.UnmarshalError.ErrorMessage, errors.UnmarshalError.HttpResponseCode)
			return
		}

		service := services.GetGetProductsService()

		// Process the request
		products, err := service.ProcessRequest(ctx, productIds)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		resp := model.Products{Products: products}
		respBytes, goErr := json.Marshal(resp)
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while marshalling the response")
			http.Error(w, errors.MarshalError.ErrorMessage, errors.MarshalError.HttpResponseCode)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, goErr = w.Write(respBytes)
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while writing the response")

		}
	}
}
