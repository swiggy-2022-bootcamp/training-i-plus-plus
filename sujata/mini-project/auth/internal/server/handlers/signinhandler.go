package handlers

import (
	"auth/internal/services"
	"auth/util"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// SignupHandler handles the http request, performs marshalling/unmarshalling and writes the
// the http response.
func SigninHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Context())
		log.Info(r.Body)

		var service services.SignupService
		// Check for Content-Type or Accept header

		// Content-Type or Accept header not set, default to JSON

		// unmarshal the request to User model

		// Validate the request
		if err := service.ValidateRequest(); err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		// Process the request
		if err := service.ProcessRequest(); err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		// marshall the response

		// Return the response along with header
		w.WriteHeader(http.StatusOK)
	}
}
