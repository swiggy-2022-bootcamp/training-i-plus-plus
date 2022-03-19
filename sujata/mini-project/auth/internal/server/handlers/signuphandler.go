package handlers

import (
	model "auth/internal/dao/mongodao/models"
	"auth/internal/errors"
	"auth/internal/services"
	"auth/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SignupHandler handles the http request, performs marshalling/unmarshalling and writes the
// the http response.
func SignupHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		// read request body
		b, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// unmarshal the request to User model
		var user model.User
		err = json.Unmarshal(b, &user)
		if err != nil {
			http.Error(w, errors.UnmarshalError.ErrorMessage, errors.UnmarshalError.HttpResponseCode)
			return
		}

		service := services.GetSignupService()

		// Validate the request
		if err := service.ValidateRequest(user); err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		// Process the request
		if err := service.ProcessRequest(ctx, user); err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		// marshall the response

		// Return the response along with header
		w.WriteHeader(http.StatusCreated)
	}
}
