package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/UserService/mocks"
	"github.com/golang/mock/gomock"
)

var (
	firstname string = "sachin"
	lastname  string = "som"
	email     string = "sachin@gmail.com"
	phone     string = "12345677890"
	password  string = "test123"
)

func TestCreateUser(t *testing.T) {
	data := url.Values{}
	data.Set("_id", "1")
	data.Set("first_name", firstname)
	data.Set("last_name", lastname)
	data.Set("email", email)
	data.Set("phone", phone)
	data.Set("password", password)

	// Create test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(`{"full_name": "sachin som", "phone": "1234", "email": "sachinsom@gmail.com", "password": "12344577"}`))

	// Create mockcontroller for user
	userMockCtrl := gomock.NewController(t)
	defer userMockCtrl.Finish()

	// Create mock user service
	userMockService := mocks.NewMockUserService(userMockCtrl)

	// Craete userController instance
	userMockController := NewUserControllers(userMockService)

	// Call CreateUser using test context
	userMockController.CreateUser(context)
	fmt.Println(response.Body)

	if response.Code != http.StatusBadGateway {
		t.Error("empty body should give error code for bad gateway.")
	}
}
