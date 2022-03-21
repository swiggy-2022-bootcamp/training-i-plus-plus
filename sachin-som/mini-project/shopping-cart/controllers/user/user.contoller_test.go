package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sachinsom93/shopping-cart/mocks"
	"github.com/sachinsom93/shopping-cart/models"
)

var (
	firstname string = "sachin"
	lastname  string = "som"
	email     string = "sachin@gmail.com"
	phone     string = "12345677890"
	password  string = "test123"
)

func TestCreateUser(t *testing.T) {
	var user models.User = models.User{
		FirstName: &firstname,
		LastName:  &lastname,
		Email:     &email,
		Phone:     &phone,
		Password:  &password,
	}
	var userBuf bytes.Buffer
	err := json.NewEncoder(&userBuf).Encode(user)

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	// Create test context
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/v1/users/create", &userBuf)

	// Create mockcontroller for user
	userMockCtrl := gomock.NewController(t)
	defer userMockCtrl.Finish()

	// Create mock user service
	userMockService := mocks.NewMockUserServices(userMockCtrl)

	// Craete userController instance
	userMockController := NewUserController(userMockService)

	// Call CreateUser using test context
	// engine.POST("/v1/users/create", userMockController.CreateUser)
	// engine.ServeHTTP(response, context.Request)
	userMockController.CreateUser(context)

	if response.Code != http.StatusCreated {
		t.Error("reposne code should be 201")
	}
}
