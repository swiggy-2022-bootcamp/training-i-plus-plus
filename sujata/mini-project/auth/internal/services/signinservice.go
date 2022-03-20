package services

import (
	"auth/internal/dao/mongodao"
	model "auth/internal/dao/mongodao/models"
	"auth/internal/errors"
	"auth/util"
	"context"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type SigninService interface {
	ValidateRequest(user model.User) *errors.ServerError
	ProcessRequest(ctx context.Context, user model.User) (model.Token, *errors.ServerError)
}

var signinServiceStruct SigninService
var signinServiceOnce sync.Once

type signinService struct {
	config *util.RouterConfig
}

func InitSigninService(config *util.RouterConfig) SignupService {
	signinServiceOnce.Do(func() {
		signinServiceStruct = &signinService{
			config: config,
		}
	})

	return signupServiceStruct
}

func GetSigninService() SigninService {
	if signinServiceStruct == nil {
		panic("Signin service not initialised")
	}

	return signinServiceStruct
}

func (s *signinService) ValidateRequest(user model.User) *errors.ServerError {
	if user.Email == "" || user.Password == "" {
		log.Error("Either user email or password missing from sign in request")
		return &errors.ParametersMissingError
	}

	return nil
}

func (s *signinService) ProcessRequest(ctx context.Context, user model.User) (model.Token, *errors.ServerError) {
	dao := mongodao.GetMongoDAO()

	// check email if email does not exists return unauthorised
	userRecord, err := dao.FindUserByEmail(ctx, user.Email)
	if err != nil {
		log.WithField("Error:", err).Error("user not found with email ", user.Email)
		return model.Token{}, err
	}

	// password not matched return unauthorised
	byteHash := []byte(userRecord.Password)
	passwordCheckErr := bcrypt.CompareHashAndPassword(byteHash, []byte(user.Password))
	if passwordCheckErr != nil {
		log.WithError(passwordCheckErr).Error("invalid password for user with email ", user.Email)
		return model.Token{}, &errors.IncorrectUserPasswordError
	}

	// create the signer for rsa 256
	signer := jwt.New(jwt.GetSigningMethod("RS256"))
	// create a map to store our claims
	claims := signer.Claims.(jwt.MapClaims)

	// set our claims
	claims["iss"] = "admin"
	claims["CustomUserInfo"] = struct {
		Email string
		Role  model.Role
	}{userRecord.Email, userRecord.Role}

	// set the expire time
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	// sigining the token with RSA private key so that
	// when we receive the token we can authenticate it with our RSA
	// public key.
	tokenString, jwtErr := signer.SignedString(util.SignKey)
	if jwtErr != nil {
		log.WithError(jwtErr).Error("Error while signing the token")
		return model.Token{}, &errors.InternalError
	}

	tokenResponse := model.Token{TokenValue: tokenString}
	return tokenResponse, nil
}
