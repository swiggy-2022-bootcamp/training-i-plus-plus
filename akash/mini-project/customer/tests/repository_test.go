package db

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sample.akash.com/api"
	"testing"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/user/all", api.QueryAll)
	return router
}

func TestCreateEndpoint(t *testing.T) {
	router := Router()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/all", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "your expected output", w.Body.String())
}
