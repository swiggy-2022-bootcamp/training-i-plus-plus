package services

import (
	"auth/internal/dao/mongodao"
	model "auth/internal/dao/mongodao/models"
	"auth/internal/errors"
	"auth/util"
	"context"
	"net/mail"
	"sync"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type SignupService interface {
	ValidateRequest(user model.User) *errors.ServerError
	ProcessRequest(ctx context.Context, user model.User) *errors.ServerError
}

var signupServiceStruct SignupService
var signupServiceOnce sync.Once

type signupService struct {
	config *util.RouterConfig
	dao    mongodao.MongoDAO
}

func InitSignupService(config *util.RouterConfig, dao mongodao.MongoDAO) SignupService {
	signupServiceOnce.Do(func() {
		signupServiceStruct = &signupService{
			config: config,
			dao:    dao,
		}
	})

	return signupServiceStruct
}

func GetSignupService() SignupService {
	if signupServiceStruct == nil {
		panic("Signup service not initialised")
	}

	return signupServiceStruct
}

func (s *signupService) ValidateRequest(user model.User) *errors.ServerError {
	if user.Firstname == "" || user.Lastname == "" || user.Email == "" || user.Address == "" || user.Password == "" || user.Role == "" {
		return &errors.ParametersMissingError
	}

	if !(user.Role == "SELLER" || user.Role == "BUYER") {
		log.Error("user role have invalid value: ", user.Role, " Required value can be SELLER or BUYER")
		return &errors.IncorrectUserRoleError
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		log.WithError(err).Error("an error occurred while checking the format of user email")
		return &errors.InvalidEmailFormatError
	}

	if len(user.Password) < 7 {
		log.Error("user password should be greater than or equal to 7 characters")
		return &errors.WeakPasswordError
	}

	return nil
}

func (service *signupService) ProcessRequest(ctx context.Context, user model.User) *errors.ServerError {
	userRecord, err := service.dao.FindUserByEmail(ctx, user.Email)
	if err != nil && err != &errors.UserNotFoundError {
		log.WithField("Error: ", err).Error("an error occurred while trying to search the existence of the user")
		return err
	}

	if userRecord.Email == user.Email && err != &errors.UserNotFoundError {
		log.Error("User already exists")
		return &errors.UserAlreadyExists
	}

	hashedPassword, goErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if goErr != nil {
		log.WithError(goErr).Error("an error occurred while hashing the password")
		return &errors.InternalError
	}

	user.Password = string(hashedPassword)
	err = service.dao.AddUser(ctx, user)
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while inserting user in database")
		return err
	}

	return nil
}
