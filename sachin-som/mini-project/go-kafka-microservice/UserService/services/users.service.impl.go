package services

import (
	"context"
	"errors"

	pb "github.com/go-kafka-microservice/AuthProto"
	"github.com/go-kafka-microservice/UserService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type UserServiceImpl struct {
	Ctx               context.Context
	UserCollection    *mongo.Collection
	AuthServiceClient pb.AuthServicesClient
	GrpcConn          *grpc.ClientConn
}

func NewUserServiceImpl(userCollection *mongo.Collection, authServiceClient pb.AuthServicesClient, grpcConn *grpc.ClientConn, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		Ctx:               ctx,
		UserCollection:    userCollection,
		AuthServiceClient: authServiceClient,
		GrpcConn:          grpcConn,
	}
}

func (us *UserServiceImpl) CreateUser(user *models.User) (string, error) {
	if user == nil || user.Email == "" || user.Fullname == "" || user.Phone == "" || user.Password == "" {
		return "", errors.New("Provide valid user details.")
	}
	user.ID = primitive.NewObjectID() // Generate Unique IDs
	// Hash the original password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)
	if _, err := us.UserCollection.InsertOne(us.Ctx, user); err != nil {
		return "", err
	}
	return user.ID.Hex(), nil
}

func (us *UserServiceImpl) GetUser(userId primitive.ObjectID) (*models.User, error) {
	filter := bson.D{bson.E{Key: "_id", Value: userId}}
	var user models.User
	if err := us.UserCollection.FindOne(us.Ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserServiceImpl) UpdateUser(userId primitive.ObjectID, updatedUser *models.User) error {
	filterQ := bson.D{bson.E{Key: "_id", Value: userId}}
	updateQ := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "fullname", Value: updatedUser.Fullname},
			bson.E{Key: "email", Value: updatedUser.Email},
			bson.E{Key: "password", Value: updatedUser.Password},
			bson.E{Key: "phone", Value: updatedUser.Phone},
		},
		},
	}
	res, err := us.UserCollection.UpdateOne(us.Ctx, filterQ, updateQ)
	if err != nil {
		return err
	}
	if res.MatchedCount != 1 {
		return errors.New("No User Found.")
	}
	if res.ModifiedCount != 1 {
		return errors.New("User didn't update.")
	}
	return nil
}

func (us *UserServiceImpl) DeleteUser(userId primitive.ObjectID) error {
	filterQ := bson.D{bson.E{Key: "_id", Value: userId}}
	res, err := us.UserCollection.DeleteOne(us.Ctx, filterQ)
	if err != nil {
		return err
	}
	if res.DeletedCount != 1 {
		return errors.New("User not deleted.")
	}
	return nil
}
func (us *UserServiceImpl) Login(credentials *models.Credentials) (string, error) {
	var opt []grpc.CallOption
	tokenRes, err := us.AuthServiceClient.Authenticate(us.Ctx, &pb.Credentials{
		Email:    credentials.Email,
		Password: credentials.Password,
	}, opt...)
	if err != nil {
		return "", err
	}
	return tokenRes.Token, nil
}
