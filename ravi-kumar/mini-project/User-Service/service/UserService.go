package service

import (
	repository "User-Service/Repository"
	errors "User-Service/errors"
	"User-Service/middleware"
	mockdata "User-Service/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	LogInUser(logInDTO mockdata.LogInDTO) (jwtToken string, err error)
	CreateUser(newUser mockdata.User) (insertedId string, jwtToken string, err error)
	GetAllUsers() (allUsers []mockdata.User)
	GetUserById(userId string) (userRetrieved *mockdata.User, err error)
	UpdateUserById(userId string, updatedUser mockdata.User) (userRetrieved *mockdata.User, err error)
	DeleteUserbyId(userId string) (successMessage *string, err error)
}

type UserService struct {
	mongoDAO repository.IMongoDAO
}

func InitUserService(initMongoDAO repository.IMongoDAO) IUserService {
	userService := new(UserService)
	userService.mongoDAO = initMongoDAO
	return userService
}

func (userService *UserService) LogInUser(logInDTO mockdata.LogInDTO) (jwtToken string, err error) {
	user, errr := userService.mongoDAO.MongoUserLogin(logInDTO)
	if errr != nil {
		return "", errr
	}

	jwtToken, _ = middleware.GenerateJWT(user.Id.Hex(), user.Role)
	return
}

func (userService *UserService) CreateUser(newUser mockdata.User) (insertedId string, jwtToken string, err error) {
	userPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	newUser.Password = string(userPassword)
	insertedId = userService.mongoDAO.MongoCreateUser(newUser)

	if insertedId == "" {
		return "", "", errors.UserNameAlreadyTaken()
	}

	jwtToken, err = middleware.GenerateJWT(insertedId, newUser.Role)
	fmt.Println(jwtToken)
	if err != nil {
		fmt.Println(err.Error())
		return "", "", err
	}

	return insertedId, jwtToken, nil
}

func (userService *UserService) GetAllUsers() (allUsers []mockdata.User) {
	allUsers = userService.mongoDAO.MongoGetAllUsers()
	return
}

func (userService *UserService) GetUserById(userId string) (userRetrieved *mockdata.User, err error) {
	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, errors.MalformedIdError()
	}

	return userService.mongoDAO.MongoGetUserById(objectId)
}

func (userService *UserService) UpdateUserById(userId string, updatedUser mockdata.User) (userRetrieved *mockdata.User, err error) {
	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	return userService.mongoDAO.MongoUpdateUserById(objectId, updatedUser)
}

func (userService *UserService) DeleteUserbyId(userId string) (successMessage *string, err error) {
	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	return userService.mongoDAO.MongoDeleteUserById(objectId)
}
