package services

import (
	"context"
	model "order/internal/dao/models"
	"order/internal/errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGetAllOrdersProcessRequest(t *testing.T) {
	_ = setupTest()
	service := GetGetOrderService()

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
				ctx:   context.TODO(),
				email: "sd@gmail.com",
			},
			expectedError: nil,
			setupFunc: func() {
				mockDao.On("GetOrders", mock.Anything, mock.Anything).Return(model.AllOrders{}, nil).Once()
			},
		},
		{
			name: "ProcessRequestWithError",
			args: functionArgs{
				ctx:   context.TODO(),
				email: "sd@gmail.com",
			},
			expectedError: &errors.InternalError,
			setupFunc: func() {
				mockDao.On("GetOrders", mock.Anything, mock.Anything).Return(model.AllOrders{}, &errors.InternalError).Once()
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupFunc()
			_, err := service.ProcessRequest(testCase.args.ctx, testCase.args.email)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
