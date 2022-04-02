package services

import (
	"context"
	model "order/internal/dao/models"
	"order/internal/errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSetOrderStatusValidateRequest(t *testing.T) {
	_ = setupTest()
	service := GetSetOrderStatusService()

	type functionArgs struct {
		ctx       context.Context
		orderInfo model.OrderInfo
	}

	tests := []struct {
		name          string
		args          functionArgs
		expectedError *errors.ServerError
	}{
		{
			name: "ValidateRequestWithWrongOrderStatus",
			args: functionArgs{
				ctx: context.TODO(),
				orderInfo: model.OrderInfo{
					OrderId:     "1234",
					OrderStatus: "NOT_AVAILABLE",
				},
			},
			expectedError: &errors.BadRequest,
		},
		{
			name: "ValidateRequestWithCorrectOrderInfo",
			args: functionArgs{
				ctx: context.TODO(),
				orderInfo: model.OrderInfo{
					OrderId:     "1234",
					OrderStatus: "OUT_FOR_DELIVERY",
				},
			},
			expectedError: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := service.ValidateRequest(testCase.args.ctx, testCase.args.orderInfo)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}

func TestSetOrderStatusProcessRequest(t *testing.T) {
	_ = setupTest()
	service := GetSetOrderStatusService()

	type functionArgs struct {
		ctx       context.Context
		orderInfo model.OrderInfo
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
				orderInfo: model.OrderInfo{
					OrderId:     "1234",
					OrderStatus: "CANCELLED",
				},
			},
			expectedError: nil,
			setupFunc: func() {
				mockDao.On("SetOrderStatus", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "ProcessRequestWithInternalError",
			args: functionArgs{
				ctx: context.TODO(),
				orderInfo: model.OrderInfo{
					OrderId:     "1234",
					OrderStatus: "CANCELLED",
				},
			},
			expectedError: &errors.InternalError,
			setupFunc: func() {
				mockDao.On("SetOrderStatus", mock.Anything, mock.Anything).Return(&errors.InternalError).Once()
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupFunc()
			err := service.ProcessRequest(testCase.args.ctx, testCase.args.orderInfo)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
