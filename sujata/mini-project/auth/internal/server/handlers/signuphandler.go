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

// Signup godoc
// @Summary Create a new user in the database
// @Description It checks if the user email exists in database or not, if it exists then it doesn't create new user. Otherwise it creates new user in the database along with his/her details.
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 201 {object} model.User
// @Router /auth/v1/signup [post]
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

		w.WriteHeader(http.StatusCreated)
	}
}
