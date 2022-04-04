package domain_test

import (
	"testing"
	"user/domain"
	"user/mocks"
	"user/utils"
	"user/utils/errs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserService(t *testing.T) {

	testDefaultUserSvc := &domain.DefaultUserService{
		UserDB: nil,
	}

	newDefaultUserSvc := domain.NewUserService(nil)
	assert.EqualValues(t, testDefaultUserSvc, newDefaultUserSvc)
}

func TestRegisterForNewUser(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"
	password := "aBc&@=+dfe35"
	name := "Ayan Dutta"
	address := "1, Pqr St., Howrah"
	zipcode := 951478
	mobileNo := "9874563210"
	role := "buyer"
	hashedPassword, _ := utils.HashPassword(password)
	user := domain.NewUser(email, password, name, address, int32(zipcode), mobileNo, role)

	mockUserRepo.On("FetchUserByEmail", email).Return(nil, &errs.AppError{})
	mockUserRepo.On("Save", mock.Anything).Return(user, nil)

	usr, err := userSvc.Register(*user)
	usr.Password = hashedPassword

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 1)
	mockUserRepo.AssertNumberOfCalls(t, "Save", 1)

	assert.EqualValues(t, user, usr)
	assert.Nil(t, err)
}

func TestRegisterForExistingUser(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"
	password := "aBc&@=+dfe35"
	name := "Ayan Dutta"
	address := "1, Pqr St., Howrah"
	zipcode := 951478
	mobileNo := "9874563210"
	role := "buyer"
	user := domain.NewUser(email, password, name, address, int32(zipcode), mobileNo, role)

	mockUserRepo.On("FetchUserByEmail", email).Return(user, nil)
	mockUserRepo.On("Save", mock.Anything).Return(nil, nil)

	usr, err := userSvc.Register(*user)

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 1)
	mockUserRepo.AssertNumberOfCalls(t, "Save", 0)

	assert.Nil(t, usr)
	assert.EqualError(t, err.Error(), "user already exists")
}

func TestLoginForGoodCredentials(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"
	password := "aBc&@=+dfe35"
	name := "Ayan Dutta"
	address := "1, Pqr St., Howrah"
	zipcode := 951478
	mobileNo := "9874563210"
	role := "buyer"
	hashedPassword, _ := utils.HashPassword(password)
	user := domain.NewUser(email, hashedPassword, name, address, int32(zipcode), mobileNo, role)

	mockUserRepo.On("FetchUserByEmail", email).Return(user, nil)

	_, err := userSvc.Login(email, password)

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 2)

	assert.Nil(t, err)

}

func TestVerifyCredentialsForBadCredentials(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"
	password := "aBc&@=+dfe35"
	name := "Ayan Dutta"
	address := "1, Pqr St., Howrah"
	zipcode := 951478
	mobileNo := "9874563210"
	role := "buyer"
	hashedPassword, _ := utils.HashPassword(password)
	user := domain.NewUser(email, hashedPassword, name, address, int32(zipcode), mobileNo, role)

	mockUserRepo.On("FetchUserByEmail", email).Return(user, nil)

	_, err := userSvc.Login(email, password+"12345")

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 1)

	assert.EqualError(t, err.Error(), "invalid credentials")

}

func TestVerifyValidToken(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"
	password := "aBc&@=+dfe35"
	name := "Ayan Dutta"
	address := "1, Pqr St., Howrah"
	zipcode := 951478
	mobileNo := "9874563210"
	role := "buyer"
	hashedPassword, _ := utils.HashPassword(password)
	user := domain.NewUser(email, hashedPassword, name, address, int32(zipcode), mobileNo, role)

	mockUserRepo.On("FetchUserByEmail", email).Return(user, nil)

	token, _ := utils.GenerateJWT(email, role)

	isVerified, err := userSvc.VerifyToken(email, role, token)

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 1)

	assert.True(t, isVerified)
	assert.Nil(t, err)
}

func TestVerifyInvalidToken(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"
	password := "aBc&@=+dfe35"
	name := "Ayan Dutta"
	address := "1, Pqr St., Howrah"
	zipcode := 951478
	mobileNo := "9874563210"
	role := "buyer"
	hashedPassword, _ := utils.HashPassword(password)
	user := domain.NewUser(email, hashedPassword, name, address, int32(zipcode), mobileNo, role)

	mockUserRepo.On("FetchUserByEmail", email).Return(user, nil)

	invalidToken, _ := utils.GenerateJWT(role, email)

	isVerified, err := userSvc.VerifyToken(email, role, invalidToken)

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 1)

	assert.False(t, isVerified)
	assert.Nil(t, err)
}

func TestUpdateForValidUser(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"
	password := "aBc&@=+dfe35"
	name := "Ayan Dutta"
	address := "1, Pqr St., Howrah"
	zipcode := 951478
	mobileNo := "9874563210"
	role := "buyer"
	user := domain.NewUser(email, password, name, address, int32(zipcode), mobileNo, role)

	mockUserRepo.On("FetchUserByEmail", email).Return(user, nil)
	mockUserRepo.On("UpdateUser", mock.Anything).Return(user, nil)

	usr, err := userSvc.Update(*user)

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 1)
	mockUserRepo.AssertNumberOfCalls(t, "UpdateUser", 1)

	assert.EqualValues(t, user, usr)
	assert.Nil(t, err)
}

func TestUpdateForInvalidUser(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"
	password := "aBc&@=+dfe35"
	name := "Ayan Dutta"
	address := "1, Pqr St., Howrah"
	zipcode := 951478
	mobileNo := "9874563210"
	role := "buyer"
	user := domain.NewUser(email, password, name, address, int32(zipcode), mobileNo, role)

	mockUserRepo.On("FetchUserByEmail", email).Return(nil, &errs.AppError{})
	mockUserRepo.On("UpdateUser", mock.Anything).Return(user, nil)

	usr, err := userSvc.Update(*user)

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 1)
	mockUserRepo.AssertNumberOfCalls(t, "UpdateUser", 0)

	assert.Nil(t, usr)
	assert.NotNil(t, err.Error())
}

func TestDeleteForValidUser(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"
	password := "aBc&@=+dfe35"
	name := "Ayan Dutta"
	address := "1, Pqr St., Howrah"
	zipcode := 951478
	mobileNo := "9874563210"
	role := "buyer"
	user := domain.NewUser(email, password, name, address, int32(zipcode), mobileNo, role)

	mockUserRepo.On("FetchUserByEmail", email).Return(user, nil)
	mockUserRepo.On("DeleteUserByEmail", mock.Anything).Return(nil)

	usr, err := userSvc.DeleteByEmail(email)

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 1)
	mockUserRepo.AssertNumberOfCalls(t, "DeleteUserByEmail", 1)

	assert.EqualValues(t, user, usr)
	assert.Nil(t, err)

}

func TestDeleteForInvalidUser(t *testing.T) {

	mockUserRepo := mocks.UserRepositoryDB{}
	userSvc := domain.NewUserService(&mockUserRepo)

	email := "ayan59dutta@gmail.com"

	mockUserRepo.On("FetchUserByEmail", email).Return(nil, &errs.AppError{})
	mockUserRepo.On("DeleteUserByEmail", mock.Anything).Return(nil)

	usr, err := userSvc.DeleteByEmail(email)

	mockUserRepo.AssertNumberOfCalls(t, "FetchUserByEmail", 1)
	mockUserRepo.AssertNumberOfCalls(t, "DeleteUserByEmail", 0)

	assert.Nil(t, usr)
	assert.NotNil(t, err.Error())
}
