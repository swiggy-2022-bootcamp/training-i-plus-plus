package services

import (
	"context"
	model "product/internal/dao/mongodao/models"
	"product/internal/errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGetProductServiceProcessRequest(t *testing.T) {
	mockDao := setupTest()
	service := GetGetProductsService()

	type functionArgs struct {
		ctx        context.Context
		productIds model.ProductIds
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
				productIds: model.ProductIds{
					Ids: []string{"123"},
				},
			},
			expectedError: nil,
			setupFunc: func() {
				mockDao.On("GetProduct", mock.Anything, mock.Anything).Return(model.Product{}, nil).Once()
			},
		},
		{
			name: "ProcessRequestWithNoError",
			args: functionArgs{
				ctx: context.TODO(),
				productIds: model.ProductIds{
					Ids: []string{"123"},
				},
			},
			expectedError: &errors.InternalError,
			setupFunc: func() {
				mockDao.On("GetProduct", mock.Anything, mock.Anything).Return(model.Product{}, &errors.InternalError).Once()
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupFunc()
			_, err := service.ProcessRequest(testCase.args.ctx, testCase.args.productIds)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
