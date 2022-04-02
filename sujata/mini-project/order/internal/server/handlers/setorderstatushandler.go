package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	model "order/internal/dao/models"
	"order/internal/services"
	"order/util"

	log "github.com/sirupsen/logrus"
)

func SetOrderStatusHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		// read request body
		b, goErr := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while reading the request body")
			http.Error(w, goErr.Error(), 500)
			return
		}

		// unmarshal the request to Product model
		var orderInfo model.OrderInfo
		goErr = json.Unmarshal(b, &orderInfo)
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while unmarshalling the request")
			//http.Error(w, errors.UnmarshalError.ErrorMessage, errors.UnmarshalError.HttpResponseCode)
			return
		}

		role, _ := util.ExtractDetailsFromToken(req)
		if role == "BUYER" && orderInfo.OrderStatus != model.CANCELLED {
			log.Error("unauthorized user, user of type BUYER can not change the order status to ", orderInfo.OrderStatus)
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		service := services.GetSetOrderStatusService()

		err := service.ValidateRequest(ctx, orderInfo)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		err = service.ProcessRequest(ctx, orderInfo)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
