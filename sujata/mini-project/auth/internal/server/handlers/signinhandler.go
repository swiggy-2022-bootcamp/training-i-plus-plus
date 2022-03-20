package handlers

import (
	model "auth/internal/dao/mongodao/models"
	"auth/internal/services"
	"auth/util"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// SignupHandler handles the http request, performs marshalling/unmarshalling and writes the
// the http response.
func SigninHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		// read request body
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// unmarshal the request to User model
		var user model.User
		json.Unmarshal(b, &user)

		service := services.GetSigninService()

		// Validate the request
		if err := service.ValidateRequest(user); err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		// Process the request
		token, serviceErr := service.ProcessRequest(ctx, user)
		if serviceErr != nil {
			http.Error(w, serviceErr.ErrorMessage, serviceErr.HttpResponseCode)
			return
		}

		// marshall the response
		responseBytes, err := json.Marshal(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(responseBytes)
		if err != nil {
			log.WithError(err).Error("error while writing the response")
			return
		}
	}
}
