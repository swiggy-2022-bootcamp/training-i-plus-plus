package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/UserService/mocks"
	"github.com/go-kafka-microservice/UserService/models"
	"github.com/go-kafka-microservice/UserService/responses"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	firstname string = "sachin"
	lastname  string = "som"
	email     string = "sachin@gmail.com"
	phone     string = "12345677890"
	password  string = "test123"
)

func TestSuccesfullCreateUser(t *testing.T) {

	// Mock User
	_user := &models.User{
		Fullname: "sachin som",
		Email:    "sachinsom@gmail.com",
		Password: "test123",
		Phone:    "1234",
	}
	userStr, _ := json.Marshal(_user)

	// Create test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(userStr))

	// Create mockcontroller for user
	userMockCtrl := gomock.NewController(t)
	defer userMockCtrl.Finish()

	// Create mock user service
	userMockService := mocks.NewMockUserService(userMockCtrl)

	_userId := primitive.NewObjectID().Hex()
	userMockService.EXPECT().CreateUser(_user).Return(_userId, nil)

	// Craete userController instance
	userMockController := NewUserControllers(userMockService)

	// Call CreateUser using test context
	userMockController.CreateUser(context)

	if response.Code != http.StatusCreated {
		t.Error("Succesfull user creation should give 201 response.")
	}
}

func TestBadGatewayCreateUser(t *testing.T) {

	// Mock User
	_user := &models.User{
		Fullname: "sachin som",
		Password: "test123",
		Phone:    "1234",
	}
	userStr, _ := json.Marshal(_user)

	// Create test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(userStr))

	// Create mockcontroller for user
	userMockCtrl := gomock.NewController(t)
	defer userMockCtrl.Finish()

	// Create mock user service
	userMockService := mocks.NewMockUserService(userMockCtrl)

	// _userId := primitive.NewObjectID().Hex()
	userMockService.EXPECT().CreateUser(_user).Return("", errors.New("Provide valid user details."))

	// Craete userController instance
	userMockController := NewUserControllers(userMockService)

	// Call CreateUser using test context
	userMockController.CreateUser(context)

	if response.Code != http.StatusBadGateway {
		t.Error("Invalid user details should give 502 response.")
	}
}

func TestBadRequestCreateUser(t *testing.T) {

	// Create test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/create", nil)

	// Create mockcontroller for user
	userMockCtrl := gomock.NewController(t)
	defer userMockCtrl.Finish()

	// Create mock user service
	userMockService := mocks.NewMockUserService(userMockCtrl)

	// Craete userController instance
	userMockController := NewUserControllers(userMockService)

	// Call CreateUser using test context
	userMockController.CreateUser(context)
	if response.Code != http.StatusBadRequest {
		t.Error("Invalid user details should give 400 response.")
	}
}

func TestSuccesfullLogin(t *testing.T) {

	// Mock User
	_cred := &models.Credentials{
		Email:    "sachinsom@gmail.com",
		Password: "test123",
	}
	credStr, _ := json.Marshal(_cred)

	// Create test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(credStr))

	// Create mockcontroller for user
	userMockCtrl := gomock.NewController(t)
	defer userMockCtrl.Finish()

	// Create mock user service
	userMockService := mocks.NewMockUserService(userMockCtrl)

	// _userId := primitive.NewObjectID().Hex()
	userMockService.EXPECT().Login(_cred).Return("dummytoken", nil)

	// Craete userController instance
	userMockController := NewUserControllers(userMockService)

	// Call CreateUser using test context
	userMockController.Login(context)
	var _token responses.TokenResponse

	err := json.NewDecoder(response.Body).Decode(&_token)
	require.Nil(t, err)

	if _token.Token == "" {
		t.Error("Succesfull login should return JWT token in response body.")
	}
}

func TestBadRequestlLogin(t *testing.T) {

	// Create test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/login", nil)

	// Create mockcontroller for user
	userMockCtrl := gomock.NewController(t)
	defer userMockCtrl.Finish()

	// Create mock user service
	userMockService := mocks.NewMockUserService(userMockCtrl)

	// Craete userController instance
	userMockController := NewUserControllers(userMockService)

	// Call CreateUser using test context
	userMockController.Login(context)

	if response.Code != http.StatusBadRequest {
		t.Error("Empty credentials should give 400 response code.")
	}
}

func TestWrongCredLogin(t *testing.T) {

	// Mock User
	_cred := &models.Credentials{
		Email:    "sachinsom@gmail.com",
		Password: "wrongpassword",
	}
	credStr, _ := json.Marshal(_cred)

	// Create test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(credStr))

	// Create mockcontroller for user
	userMockCtrl := gomock.NewController(t)
	defer userMockCtrl.Finish()

	// Create mock user service
	userMockService := mocks.NewMockUserService(userMockCtrl)

	// _userId := primitive.NewObjectID().Hex()
	userMockService.EXPECT().Login(_cred).Return("", errors.New("Password or email or matching"))

	// Craete userController instance
	userMockController := NewUserControllers(userMockService)

	// Call CreateUser using test context
	userMockController.Login(context)
	var _token responses.TokenResponse

	err := json.NewDecoder(response.Body).Decode(&_token)
	require.Nil(t, err)

	if _token.Token != "" {
		t.Error("Wrong credentials should not return JWT token in response body.")
	}
}
