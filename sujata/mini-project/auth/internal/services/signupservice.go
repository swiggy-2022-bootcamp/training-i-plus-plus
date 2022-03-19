package services

import (
	"auth/internal/dao/mongodao"
	model "auth/internal/dao/mongodao/models"
	"auth/internal/errors"
	"auth/util"
	"context"
	"net/mail"
	"sync"
	"unicode"

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
}

func InitSignupService(config *util.RouterConfig) SignupService {
	signupServiceOnce.Do(func() {
		signupServiceStruct = &signupService{
			config: config,
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
	if user.Firstname == "" || user.Lastname == "" || user.Email == "" || user.Address == "" || user.Password == "" {
		return &errors.ParametersMissingError
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		log.WithError(err).Error("an error occurred while checking the format of user email")
		return &errors.InvalidEmailFormatError
	}

	// if !verifyPassword(user.Password) {
	// 	log.Error("user password does not fulfill the password criteria")
	// 	return &errors.WeakPasswordError
	// }

	return nil
}

func (service *signupService) ProcessRequest(ctx context.Context, user model.User) *errors.ServerError {
	dao := mongodao.GetMongoDAO()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.WithError(err).Error("an error occurred while hashing the password")
		return &errors.InternalError
	}

	user.Password = string(hashedPassword)
	dbErr := dao.AddUser(ctx, user)
	if err != nil {
		log.WithError(err).Error("an error occurred while inserting user in database")
		return dbErr
	}

	return nil
}

func verifyPassword(password string) bool {
	var hasNumber bool
	var hasUpperCaseLetter bool
	var hasSpecialCharacter bool
	letters := 0

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperCaseLetter = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecialCharacter = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			return false
		}
	}

	return letters >= 7 && hasNumber && hasUpperCaseLetter && hasSpecialCharacter
}
