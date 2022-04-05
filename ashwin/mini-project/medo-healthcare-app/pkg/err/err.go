package err

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

//CheckNilErr ..
func CheckNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//InternalErrHandler ..
func InternalErrHandler(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
	}
}

//HandleCheckForCookiesErr ..
func HandleCheckForCookiesErr(err error, w http.ResponseWriter) {
	if err != nil {
		if err != http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

//HandleJWTSignatureErr ..
func HandleJWTSignatureErr(err error, w http.ResponseWriter) bool {
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("JWT Token Signature Invalid !"))
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JWT Token Parsing Error !"))
		return false
	}
	return true
}
