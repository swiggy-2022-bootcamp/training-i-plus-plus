package handlers

import (
	model "cart/internal/dao/models"
	"cart/internal/errors"
	"cart/internal/services"
	"cart/util"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func AddProductToCartHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var goErr error

		role, email := util.ExtractDetailsFromToken(req)
		if role == "SELLER" {
			log.Error("unauthorized user, user if of type SELLER")
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		// read request body
		b, goErr := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while reading the request body")
			http.Error(w, errors.InternalError.ErrorMessage, errors.InternalError.HttpResponseCode)
			return
		}

		// unmarshal the request to Product model
		var product model.CartProduct
		goErr = json.Unmarshal(b, &product)
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while unmarshalling the request")
			http.Error(w, errors.UnmarshalError.ErrorMessage, errors.UnmarshalError.HttpResponseCode)
			return
		}

		service := services.GetAddProductToCartService()

		// Validate the request
		if err := service.ValidateRequest(product); err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		// Process the request
		err := service.ProcessRequest(ctx, email, product)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
