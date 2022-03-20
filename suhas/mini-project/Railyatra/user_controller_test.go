package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-mongo-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCreateRoute(t *testing.T) {
	rout := SetupRouter()
	arr_body := []models.User{
		{
			Name:  "Elita Moreland",
			Email: "emoreland0@bloglovin.com",
		}, {
			Name:  "Thurston Spurden",
			Email: "tspurden1@netlog.com",
		}, {
			Name:  "Elna Veldman",
			Email: "eveldman2@canalblog.com",
		}, {
			Name:  "Laverne Nutley",
			Email: "lnutley3@rediff.com",
		}, {
			Name:  "Adolphe Pfeuffer",
			Email: "apfeuffer4@com.com",
		}, {
			Name:  "Tobin Denford",
			Email: "tdenford5@mapy.cz",
		}, {
			Name:  "Penni Gibbon",
			Email: "pgibbon6@imdb.com",
		}, {
			Name:  "Martica Rickardsson",
			Email: "mrickardsson7@apple.com",
		}, {
			Name:  "Nataline Stotherfield",
			Email: "nstotherfield8@patch.com",
		}, {
			Name:  "Janek Doumerque",
			Email: "jdoumerque9@omniture.com",
		},
	}
	for _, v := range arr_body {
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(v)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/user", payloadBuf)
		rout.ServeHTTP(w, req)
		fmt.Println(w.Body)
		assert.Equal(t, 201, w.Code)
	}
}
