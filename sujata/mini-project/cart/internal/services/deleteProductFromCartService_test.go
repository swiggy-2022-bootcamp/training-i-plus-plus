package services

import (
	"cart/internal/errors"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestDeleteProductFromCartValidateRequest(t *testing.T) {
	_ = setupTest()
	service := GetDeleteProductFromCartService()

	type functionArgs struct {
		productId string
	}

	tests := []struct {
		name          string
		args          functionArgs
		expectedError *errors.ServerError
	}{
		{
			name: "ValidateRequestForCorrectData",
			args: functionArgs{
				productId: "1234",
			},
			expectedError: nil,
		},
		{
			name: "ValidateRequestWithMalformedIdError",
			args: functionArgs{
				productId: "\"123",
			},
			expectedError: &errors.MalformedQueryParamError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := service.ValidateRequest(testCase.args.productId)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}

func TestDeleteProductFromCartProcessRequest(t *testing.T) {
	mockdao := setupTest()
	service := GetDeleteProductFromCartService()
	ctx := context.TODO()

	type functionArgs struct {
		ctx       context.Context
		email     string
		productId string
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
				ctx:       ctx,
				email:     "abc@gmail.com",
				productId: "1234",
			},
			setupFunc: func() {
				mockdao.On("DeleteProduct", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "ProcessRequestWithError",
			args: functionArgs{
				ctx:       ctx,
				email:     "abc@gmail.com",
				productId: "1234",
			},
			setupFunc: func() {
				mockdao.On("DeleteProduct", mock.Anything, mock.Anything, mock.Anything).Return(&errors.InternalError).Once()
			},
			expectedError: &errors.InternalError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setupFunc != nil {
				testCase.setupFunc()
			}
			err := service.ProcessRequest(testCase.args.ctx, testCase.args.productId, testCase.args.email)

			if err != nil && err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
