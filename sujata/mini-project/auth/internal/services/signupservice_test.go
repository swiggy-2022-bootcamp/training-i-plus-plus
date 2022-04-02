package services

import (
	model "auth/internal/dao/mongodao/models"
	"auth/internal/errors"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSignupValidateRequest(t *testing.T) {
	_ = setupTest()
	service := GetSignupService()

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
			name: "ValidateRequestWithWeakPasswordErr",
			args: functionArgs{
				user: model.User{
					Email:     "sd@gmail.com",
					Firstname: "Sujata",
					Lastname:  "Dwivedi",
					Password:  "abc",
					Address:   "Street lane - 123, India",
					Role:      "BUYER",
				},
			},
			expectedError: &errors.WeakPasswordError,
		},
		{
			name: "ValidateRequestWithParameterMissing",
			args: functionArgs{
				user: model.User{},
			},
			expectedError: &errors.ParametersMissingError,
		},
		{
			name: "ValidateRequestWithIncorrectUserRole",
			args: functionArgs{
				user: model.User{
					Email:     "sd@gmail.com",
					Firstname: "Sujata",
					Lastname:  "Dwivedi",
					Password:  "abc",
					Address:   "Street lane - 123, India",
					Role:      "Cobuyer",
				},
			},
			expectedError: &errors.IncorrectUserRoleError,
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

func TestSignupProcessRequest(t *testing.T) {
	mockdao := setupTest()
	service := GetSignupService()
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
			name: "ProcessRequestWithUserAlreadyExistsError",
			args: functionArgs{
				ctx: ctx,
				user: model.User{
					Email:    "sd@gmail.com",
					Password: "abcdefgeh",
				},
			},
			expectedToken: dummyToken,
			expectedError: &errors.UserAlreadyExists,
			setupFunc: func() {
				mockdao.On("FindUserByEmail", mock.Anything, mock.Anything).Return(user, nil).Once()
			},
		},
		{
			name: "ProcessRequestWithNoError",
			args: functionArgs{
				ctx: ctx,
				user: model.User{
					Email:    "de@gmail.com",
					Password: "abcdefgeh",
				},
			},
			expectedToken: dummyToken,
			expectedError: nil,
			setupFunc: func() {
				mockdao.On("FindUserByEmail", mock.Anything, mock.Anything).Return(model.User{}, nil).Once()
				mockdao.On("AddUser", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "ProcessRequestWithAddUserError",
			args: functionArgs{
				ctx: ctx,
				user: model.User{
					Email:    "de@gmail.com",
					Password: "abcdefgeh",
				},
			},
			expectedToken: dummyToken,
			expectedError: &errors.InternalError,
			setupFunc: func() {
				mockdao.On("FindUserByEmail", mock.Anything, mock.Anything).Return(model.User{}, &errors.UserNotFoundError).Once()
				mockdao.On("AddUser", mock.Anything, mock.Anything).Return(&errors.InternalError).Once()
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setupFunc != nil {
				testCase.setupFunc()
			}
			err := service.ProcessRequest(testCase.args.ctx, testCase.args.user)

			if err != nil && err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
