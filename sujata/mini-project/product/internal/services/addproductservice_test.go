package services

import (
	"context"
	"product/config"
	"product/internal/dao/mongodao/mocks"
	model "product/internal/dao/mongodao/models"
	"product/internal/errors"
	"product/util"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
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
		InitAddProductService(routerConfig, mockDao)
		InitGetProductsService(routerConfig, mockDao)
	})

	return mockDao
}

func TestAddProductServiceValidateRequest(t *testing.T) {
	_ = setupTest()
	service := GetAddProductService()

	type functionArgs struct {
		product model.Product
	}

	tests := []struct {
		name          string
		args          functionArgs
		expectedError *errors.ServerError
	}{
		{
			name: "ValidateRequestWithCorrectData",
			args: functionArgs{
				product: model.Product{
					Name:        "Product1",
					Description: "This is sample product 1",
					Price:       30,
					Quantity:    100,
				},
			},
			expectedError: nil,
		},
		{
			name: "ValidateRequestWithIncorrectData",
			args: functionArgs{
				product: model.Product{
					Description: "This is sample product 1",
					Quantity:    100,
				},
			},
			expectedError: &errors.ParametersMissingError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := service.ValidateRequest(testCase.args.product)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}

func TestAddProductServiceProcessRequest(t *testing.T) {
	mockDao := setupTest()
	service := GetAddProductService()

	type functionArgs struct {
		ctx     context.Context
		product model.Product
	}

	tests := []struct {
		name          string
		args          functionArgs
		expectedError *errors.ServerError
		setupFunc     func()
	}{
		{
			name: "ProcessRequestWithNoError",
			args: functionArgs{
				ctx: context.TODO(),
				product: model.Product{
					Name:        "Product1",
					Description: "This is sample product 1",
					Price:       30,
					Quantity:    100,
				},
			},
			expectedError: nil,
			setupFunc: func() {
				mockDao.On("AddProduct", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "ProcessRequestWithError",
			args: functionArgs{
				ctx: context.TODO(),
				product: model.Product{
					Name:        "Product1",
					Description: "This is sample product 1",
					Price:       30,
					Quantity:    100,
				},
			},
			expectedError: &errors.InternalError,
			setupFunc: func() {
				mockDao.On("AddProduct", mock.Anything, mock.Anything).Return(&errors.InternalError).Once()
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupFunc()
			err := service.ProcessRequest(testCase.args.ctx, testCase.args.product)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
