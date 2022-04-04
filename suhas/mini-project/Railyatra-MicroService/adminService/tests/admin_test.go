package main

import (
	"adminService/config"
	"adminService/models"
	"adminService/routes"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

//init router
func SetupAdminRouter() *gin.Engine {

	config.ConnectDB()

	router := gin.Default()
	routes.AdminRoute(router)
	router.Run("localhost:6001")
	return router
}

var rout = SetupAdminRouter()

var (
	arr_insertedId []string
	arr_body       = []models.Admin{
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

func TestAdminCreateRoute(t *testing.T) {

	for _, v := range arr_body {
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(v)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/admin", payloadBuf)
		rout.ServeHTTP(w, req)
		fmt.Println(w.Body)
		assert.Equal(t, 201, w.Code)
		insertedId := gjson.Get(w.Body.String(), "data.data.InsertedID")
		arr_insertedId = append(arr_insertedId, insertedId.String())
		fmt.Print(arr_insertedId)
	}
}

func TestAdminGetRoute(t *testing.T) {
	for _, v := range arr_insertedId {
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/admin/%s", v)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		rout.ServeHTTP(w, req)
		fmt.Println(w.Body)
		assert.Equal(t, 201, w.Code)
	}
}

func TestAdminUpdateRoute(t *testing.T) {
	arr_body_update := []models.Admin{
		{
			Name:  "Elita Moreland",
			Email: "emoreland0@bloglovin.com",
		}, {
			Name:  "Thurston Spurden",
			Email: "tspurden1@netlog.com",
		}, {
			Name:  "Elna Veldman",
			Email: "eveldman2@canalblog.com",
		},
	}
	for i, v := range arr_insertedId {
		if i < 3 {
			obj := arr_body_update[i]
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(obj)
			w := httptest.NewRecorder()
			url := fmt.Sprintf("/admin/%s", v)
			req, _ := http.NewRequest(http.MethodPut, url, payloadBuf)
			rout.ServeHTTP(w, req)
			fmt.Println(w.Body)
			assert.Equal(t, 201, w.Code)
		}
	}
}

func TestAdminDeleteRoute(t *testing.T) {
	for _, v := range arr_insertedId {
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/admin/%s", v)
		req, _ := http.NewRequest(http.MethodDelete, url, nil)
		rout.ServeHTTP(w, req)
		fmt.Println(w.Body)
		assert.Equal(t, 201, w.Code)
	}
}
