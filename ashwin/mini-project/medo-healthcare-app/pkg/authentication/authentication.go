package authentication

import (
	"encoding/json"
	"fmt"
	"medo-healthcare-app/cmd/model"
	"medo-healthcare-app/pkg/database"
	"medo-healthcare-app/pkg/err"
	"medo-healthcare-app/pkg/logger"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

//ValidateLogin ...
func ValidateLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var credentials model.Credentials
	cancel := json.NewDecoder(r.Body).Decode(&credentials)
	logger.Error("Error : ", cancel)
	currentUser := database.FindOne(credentials.Email)
	expectedPassword := currentUser.Password
	if expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect Password"))
		return
	}
	tokenExpirationTime := time.Now().Add(time.Minute * 1)
	claims := &model.Claims{
		Username: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, cancel := token.SignedString(jwtKey)
	err.InternalErrHandler(cancel, w)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: tokenExpirationTime,
	})
	w.Write([]byte("Login Successful ! 😍"))
}

//AuthenticateLogin ..
func AuthenticateLogin(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Content-Type", "application/json")
	cookie, cancel := r.Cookie("token")
	if cancel != nil {
		if cancel == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	tokenStr := cookie.Value
	claims := &model.Claims{}
	tkn, cancel := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	err.HandleJWTSignatureErr(cancel, w)
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("JWT Token Invalid !"))
		return false
	}
	return true
}

//GetUsernameFromToken ...
func GetUsernameFromToken(w http.ResponseWriter, r *http.Request) string {
	w.Header().Set("Content-Type", "application/json")
	cookie, cancel := r.Cookie("token")
	result := ""
	if cancel != nil {
		if cancel == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			result = ""
		}
		w.WriteHeader(http.StatusBadRequest)
		result = ""
	}
	tokenStr := cookie.Value
	claims := &model.Claims{}
	tkn, cancel := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	err.HandleJWTSignatureErr(cancel, w)
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("JWT Token Invalid !"))
		result = ""
	}
	result = (fmt.Sprint(claims.Username))
	return result
}

// //IsLoginSuccessful ...
// func IsLoginSuccessful(w http.ResponseWriter, r *http.Request) bool {
// 	if AuthenticateLogin(w, r) {
// 		claims := &model.Claims{}
// 		fmt.Println("Login Successful !")
// 		w.Write([]byte("Login Successful ! 😍"))
// 		w.Write([]byte(fmt.Sprintf("Hello, %s \n", claims.Username)))
// 		return true
// 	}
// 	return false
// }
