package services

import (
	"context"
	"time"

	"github.com/go-kafka-microservice/AuthService/models"
	"github.com/go-kafka-microservice/AuthService/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthServicesImpl struct {
	JWTUtils       utils.JWTUtils
	UserCollection *mongo.Collection
	Ctx            context.Context
}

func NewAuthServiceImpl(jwtUtils utils.JWTUtils, userCollection *mongo.Collection, ctx context.Context) *AuthServicesImpl {
	return &AuthServicesImpl{
		JWTUtils:       jwtUtils,
		UserCollection: userCollection,
		Ctx:            ctx,
	}
}

func (as *AuthServicesImpl) Authenticate(credentials *models.Credentials) (string, error) {

	// Get the stored pwd in userdb
	// TODO: Can communicate through gRPC to UserService
	var user models.User
	filter := bson.D{bson.E{Key: "email", Value: credentials.Email}}
	if err := as.UserCollection.FindOne(as.Ctx, filter).Decode(&user); err != nil {
		return "", nil
	}
	// Check for password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		// err means password not matching
		return "", err
	}
	token, err := as.JWTUtils.GenerateToken(credentials, time.Now().Add(5*time.Hour))

	return token, err
}

// Authorize takes a token as string
// validates the string
// returns nil if case of valid credentials
// else returns error with proper reason
func (as *AuthServicesImpl) Authorize(tokenStr string) (string, error) {
	refreshedTkn, err := as.JWTUtils.ValidateToken(tokenStr, time.Now().Add(5*time.Minute))
	if err != nil {
		return "", err
	}
	return refreshedTkn, nil
}
