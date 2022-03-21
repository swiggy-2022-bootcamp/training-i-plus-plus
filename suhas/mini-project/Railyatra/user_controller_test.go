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
	"github.com/tidwall/gjson"
)

//init router
var rout = SetupRouter()

var (
	arr_insertedId []string
	arr_body       = []models.User{
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
)

func TestUserCreateRoute(t *testing.T) {

	for _, v := range arr_body {
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(v)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/user", payloadBuf)
		rout.ServeHTTP(w, req)
		fmt.Println(w.Body)
		assert.Equal(t, 201, w.Code)
		insertedId := gjson.Get(w.Body.String(), "data.data.InsertedID")
		arr_insertedId = append(arr_insertedId, insertedId.String())
		fmt.Print(arr_insertedId)
	}
}

func TestUserGetRoute(t *testing.T) {
	for _, v := range arr_insertedId {
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/user/%s", v)
		req, _ := http.NewRequest(http.MethodPost, url, nil)
		rout.ServeHTTP(w, req)
		fmt.Println(w.Body)
		assert.Equal(t, 201, w.Code)
	}
}
