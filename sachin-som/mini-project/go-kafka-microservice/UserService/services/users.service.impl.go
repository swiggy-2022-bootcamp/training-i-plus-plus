package services

import (
	"context"

	pb "github.com/go-kafka-microservice/AuthProto"
	"github.com/go-kafka-microservice/UserService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (us *UserServiceImpl) CreateUser(user *models.User) error {
	user.ID = primitive.NewObjectID()
	if _, err := us.UserCollection.InsertOne(us.Ctx, user); err != nil {
		return err
	}
	return nil
}

func (us *UserServiceImpl) GetUser(userId primitive.ObjectID) (*models.User, error) {
	filter := bson.D{bson.E{Key: "_id", Value: userId}}
	var user models.User
	if err := us.UserCollection.FindOne(us.Ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserServiceImpl) Login(credentials *models.Credentials) (string, error) {

	var opt []grpc.CallOption
	tokenRes, err := us.AuthServiceClient.Authenticate(us.Ctx, &pb.Credentials{
		Email:    credentials.Email,
		Password: credentials.Password,
	}, opt...)
	// out := new(pb.Response)
	// err := us.GrpcConn.Invoke(us.Ctx, "github.com/go-kafka-microservice/AuthServices/controllers.AuthServices/Authorize", pb.Credentials{
	// 	Email:    credentials.Email,
	// 	Password: credentials.Password,
	// }, &out, opt...)
	if err != nil {
		return "", err
	}
	return tokenRes.Token, nil
}
