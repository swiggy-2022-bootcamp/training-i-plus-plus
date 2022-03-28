package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"swiggy/gin/services/auth"
	"testing"

	"swiggy/gin/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/tidwall/gjson"
)

var (
	token string
	rout  *gin.Engine = router.ApplyRoutes()
)

func TestLogin(t *testing.T) {
	username := "ishan"
	password := "ishan123"

	body := auth.LoginBody{
		Username: username, Password: password,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/login", payloadBuf)
	rout.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	token = gjson.Get(w.Body.String(), "token").String()
}

func TestFetchReservations(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/reservation", nil)
	req.Header = http.Header{
		"Authorization": []string{"Bearer " + token},
	}

	rout.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
