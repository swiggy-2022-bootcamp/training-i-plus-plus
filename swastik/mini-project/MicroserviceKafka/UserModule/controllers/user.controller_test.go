package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	mocks "github.com/swastiksahoo153/train-reservation-system/mocks"
	"github.com/swastiksahoo153/train-reservation-system/models"
	// "github.com/swastiksahoo153/train-reservation-system/services"
	"github.com/golang/mock/gomock"
	// "github.com/stretchr/testify/require"
)

var (
	firstname string = "swastik"
	lastname  string = "sahoo"
	email     string = "swastik@gmail.com"
	phone     string = "12345677890"
	password  string = "test123"
)

func TestRegisterUser(t *testing.T) {

	testUser := &models.User{
		Name: 		"swastik sahoo",
		Username:    "swastiksah",  
		Password: 	"test123",
		Age:    22,
	}

	userJson, _ := json.Marshal(testUser)

	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/userRegistration", bytes.NewBuffer(userJson))

	userMockCtrl := gomock.NewController(t)
	defer userMockCtrl.Finish()

	userMockService := mocks.NewMockUserService(userMockCtrl)
	userMockService.EXPECT().CreateUser(testUser).Return(nil)
	userMockController := New(userMockService)
	userMockController.RegisterUser(context)
	if response.Code != http.StatusCreated {
		t.Error("Did not receive response status: 201")
	}
}
