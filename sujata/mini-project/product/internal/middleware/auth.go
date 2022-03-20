package middleware

import (
	"encoding/json"
	"net/http"
	model "product/internal/dao/mongodao/models"
	"product/internal/errors"
	"product/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	log "github.com/sirupsen/logrus"
)

// authHandler validate the JWT token. Header should have the token in the following schema
// Authorization: Bearer <token>
func AuthHandler(nextHandler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, keyLookupFunc)

		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError: // something was wrong during the validation
				vErr := err.(*jwt.ValidationError)

				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					log.WithError(err).Error("token expired, get a new one")
					jsonResponse(w, "You are not authenticated")
					return
				default:
					log.WithError(err).Error("an error while parsing JWT token")
					http.Error(w, errors.InternalError.ErrorMessage, errors.InternalError.HttpResponseCode)
					return
				}

			default: // something else went wrong
				log.WithError(err).Error("an error while parsing JWT token")
				http.Error(w, errors.InternalError.ErrorMessage, errors.InternalError.HttpResponseCode)
				return
			}
		}

		if token.Valid {
			nextHandler.ServeHTTP(w, r)
		} else {
			jsonResponse(w, "You are not authenticated")
		}
	})
}

// keyLookupFunc returns the public key for JWT authentication
func keyLookupFunc(*jwt.Token) (interface{}, error) {
	return util.VerifyKey, nil
}

// jsonResponse does json marshalling and in case of error returns the internal server error
// while marshalling. Otherwise, returns the auth response string along with 401 HTTP Status code.
func jsonResponse(w http.ResponseWriter, responseString string) {
	response := model.AuthResponse{Text: responseString}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.WithError(err).Error("an error while marshalling json response")
		http.Error(w, errors.InternalError.ErrorMessage, errors.InternalError.HttpResponseCode)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
