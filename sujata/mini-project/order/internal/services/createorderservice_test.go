package services

import (
	"context"
	"order/config"
	"order/internal/dao/mocks"
	"order/internal/errors"
	"order/util"
	"sync"
	"testing"
)

var testSync sync.Once
var mockDao *mocks.MongoDAO

func setupTest() *mocks.MongoDAO {
	webserverconfig, _ := config.FromEnv()

	routerConfig := &util.RouterConfig{
		WebServerConfig: webserverconfig,
	}

	testSync.Do(func() {
		mockDao = &mocks.MongoDAO{}
		InitCreateOrderService(routerConfig, mockDao)
		InitGetOrderService(routerConfig, mockDao)
		InitSetOrderStatusService(routerConfig, mockDao)
	})

	return mockDao
}

func TestGetCreateOrderProcessRequest(t *testing.T) {
	_ = setupTest()
	service := GetCreateOrderService()

	type functionArgs struct {
		ctx   context.Context
		email string
		token string
	}

	tests := []struct {
		name          string
		args          functionArgs
		expectedError *errors.ServerError
	}{
		{
			name: "ProcessRequestWithProductHttpCallFailure",
			args: functionArgs{
				ctx:   context.TODO(),
				email: "sd@gmail.com",
				// dummy value for token
				token: "eyz12324bsmdfbkdknkaslj",
			},
			expectedError: &errors.InternalError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := service.ProcessRequest(testCase.args.ctx, testCase.args.email, testCase.args.token)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
