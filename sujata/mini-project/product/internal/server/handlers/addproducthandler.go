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

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const (
	SELLER string = "SELLER"
)

func AddProductHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		token, _ := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, keyLookupFunc)

		// check if the user adding the product is of type SELLER
		claims := token.Claims.(jwt.MapClaims)
		userInfo := claims["CustomUserInfo"].(map[string]interface{})

		log.Info(userInfo)
		if userInfo["Role"] != SELLER {
			log.WithField("User Role: ", userInfo["Role"]).Error("user forbidden as user is not of type SELLER")
			http.Error(w, "User not allowed", http.StatusForbidden)
			return
		}

		ctx := req.Context()

		// read request body
		b, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.WithError(err).Error("an error occurrec while reading the request body")
			http.Error(w, err.Error(), 500)
			return
		}

		// unmarshal the request to Product model
		var product model.Product
		err = json.Unmarshal(b, &product)
		if err != nil {
			log.WithError(err).Error("an error occurred while unmarshalling the request")
			http.Error(w, errors.UnmarshalError.ErrorMessage, errors.UnmarshalError.HttpResponseCode)
			return
		}

		service := services.GetAddProductService()

		// Validate the request
		if err := service.ValidateRequest(product); err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		// Process the request
		if err := service.ProcessRequest(ctx, product); err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// keyLookupFunc returns the public key for JWT authentication
func keyLookupFunc(*jwt.Token) (interface{}, error) {
	return util.VerifyKey, nil
}
