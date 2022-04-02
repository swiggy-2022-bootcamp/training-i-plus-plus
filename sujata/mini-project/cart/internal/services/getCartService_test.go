package services

import (
	model "cart/internal/dao/models"
	"cart/internal/errors"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGetCartProcessRequest(t *testing.T) {
	mockdao := setupTest()
	service := GetGetCartService()
	ctx := context.TODO()

	type functionArgs struct {
		ctx   context.Context
		email string
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
				ctx:   ctx,
				email: "abc@gmail.com",
			},
			setupFunc: func() {
				mockdao.On("GetCart", mock.Anything, mock.Anything).Return(model.Cart{}, nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "ProcessRequestWithError",
			args: functionArgs{
				ctx:   ctx,
				email: "abc@gmail.com",
			},
			setupFunc: func() {
				mockdao.On("GetCart", mock.Anything, mock.Anything).Return(model.Cart{}, &errors.InternalError).Once()
			},
			expectedError: &errors.InternalError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setupFunc != nil {
				testCase.setupFunc()
			}
			_, err := service.ProcessRequest(testCase.args.ctx, testCase.args.email)

			if err != nil && err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
