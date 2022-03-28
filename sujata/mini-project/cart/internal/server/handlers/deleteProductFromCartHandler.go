package handlers

import (
	"cart/internal/services"
	"cart/util"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func DeleteProductFromCartHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		role, email := util.ExtractDetailsFromToken(req)
		if role == "SELLER" {
			log.Error("unauthorized user, user if of type SELLER")
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		// extract searched product from query
		params := req.URL.Query()
		productId := params["productId"][0]
		fmt.Println(productId)

		service := services.GetDeleteProductFromCartService()

		// Validate the request
		if err := service.ValidateRequest(productId); err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		// Process the request
		err := service.ProcessRequest(ctx, productId, email)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
