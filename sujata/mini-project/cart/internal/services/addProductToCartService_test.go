package services

import (
	"cart/config"
	model "cart/internal/dao/models"
	"context"

	"cart/internal/dao/mocks"
	"cart/internal/errors"
	"cart/util"
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
		InitAddProductToCartService(routerConfig, mockDao)
		InitDeleteProductFromCartService(routerConfig, mockDao)
		InitGetCartService(routerConfig, mockDao)
	})

	return mockDao
}

func TestAddProductToCartValidateRequest(t *testing.T) {
	_ = setupTest()
	service := GetAddProductToCartService()
	product := model.CartProduct{
		ProductId: "1234",
		Quantity:  2,
	}

	type functionArgs struct {
		product model.CartProduct
	}

	tests := []struct {
		name          string
		args          functionArgs
		expectedError *errors.ServerError
	}{
		{
			name: "ValidateRequestForCorrectData",
			args: functionArgs{
				product: product,
			},
			expectedError: nil,
		},
		{
			name: "ValidateRequestWithParameterMissing",
			args: functionArgs{
				product: model.CartProduct{},
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

func TestAddProductToCartProcessRequest(t *testing.T) {
	mockdao := setupTest()
	service := GetAddProductToCartService()
	ctx := context.TODO()

	product := model.CartProduct{
		ProductId: "1234",
		Quantity:  2,
	}

	type functionArgs struct {
		ctx     context.Context
		email   string
		product model.CartProduct
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
				ctx:     ctx,
				email:   "abc@gmail.com",
				product: product,
			},
			setupFunc: func() {
				mockdao.On("AddProduct", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "ProcessRequestWithError",
			args: functionArgs{
				ctx:     ctx,
				email:   "abc@gmail.com",
				product: product,
			},
			setupFunc: func() {
				mockdao.On("AddProduct", mock.Anything, mock.Anything, mock.Anything).Return(&errors.InternalError).Once()
			},
			expectedError: &errors.InternalError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setupFunc != nil {
				testCase.setupFunc()
			}
			err := service.ProcessRequest(testCase.args.ctx, testCase.args.email, testCase.args.product)

			if err != nil && err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
