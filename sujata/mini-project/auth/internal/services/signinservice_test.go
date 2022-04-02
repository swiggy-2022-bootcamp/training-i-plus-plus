package services

import (
	"auth/config"
	"auth/internal/dao/mongodao/mocks"
	model "auth/internal/dao/mongodao/models"
	"auth/internal/errors"
	"auth/util"
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
)

var testSync sync.Once
var mockDao *mocks.MongoDAO

func setupTest() *mocks.MongoDAO {
	webserverconfig, err := config.FromEnv()
	if err != nil {
	}

	routerConfig := &util.RouterConfig{
		WebServerConfig: webserverconfig,
	}

	testSync.Do(func() {
		mockDao = &mocks.MongoDAO{}
		InitSigninService(routerConfig, mockDao)
		InitSignupService(routerConfig, mockDao)
	})

	return mockDao
}

func TestSigninValidateRequest(t *testing.T) {
	_ = setupTest()
	service := GetSigninService()

	type functionArgs struct {
		user model.User
	}

	tests := []struct {
		name          string
		args          functionArgs
		expectedError *errors.ServerError
	}{
		{
			name: "ValidateRequestForCorrectData",
			args: functionArgs{
				user: model.User{
					Email:     "sd@gmail.com",
					Firstname: "Sujata",
					Lastname:  "Dwivedi",
					Password:  "abcdefgeh",
					Address:   "Street lane - 123, India",
					Role:      "BUYER",
				},
			},
			expectedError: nil,
		},
		{
			name: "ValidateRequestWithParameterMissingError",
			args: functionArgs{
				user: model.User{},
			},
			expectedError: &errors.ParametersMissingError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := service.ValidateRequest(testCase.args.user)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}

func TestSigninProcessRequest(t *testing.T) {
	mockdao := setupTest()
	service := GetSigninService()
	ctx := context.TODO()
	dummyToken := model.Token{TokenValue: "eyJhbGciOiJSUzI1NiIs"}

	user := model.User{
		Email:     "sd@gmail.com",
		Firstname: "Sujata",
		Lastname:  "Dwivedi",
		Password:  "abcdefgeh",
		Address:   "Street lane - 123, India",
		Role:      "BUYER",
	}

	type functionArgs struct {
		ctx  context.Context
		user model.User
	}

	tests := []struct {
		name          string
		args          functionArgs
		expectedToken model.Token
		expectedError *errors.ServerError
		setupFunc     func()
	}{
		{
			name: "ProcessRequestWithAuthError",
			args: functionArgs{
				ctx: ctx,
				user: model.User{
					Email:    "sd@gmail.com",
					Password: "abcdefgeh",
				},
			},
			expectedToken: dummyToken,
			expectedError: &errors.IncorrectUserPasswordError,
			setupFunc: func() {
				mockdao.On("FindUserByEmail", mock.Anything, mock.Anything).Return(user, nil).Once()
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setupFunc != nil {
				testCase.setupFunc()
			}
			_, err := service.ProcessRequest(testCase.args.ctx, testCase.args.user)

			if err != nil && err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
