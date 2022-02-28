package auth

import (
	"context"
	"fmt"
	jwtmanager "swiggy/train_reservation/helpers/lib"
	db "swiggy/train_reservation/helpers/utils"
	authpb "swiggy/train_reservation/services/auth/authpb"
	"swiggy/train_reservation/services/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	JWTManager *jwtmanager.JWTManager
}

func (server *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	user := &user.User{}
	if err := db.DataStore.Collection("user").FindOne(ctx, bson.M{}).Decode(&user); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.JWTManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &authpb.LoginResponse{AccessToken: token}
	return res, nil
}

func (server *AuthServer) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	user, err := user.NewUser(req.GetUsername(), req.GetPassword(), req.GetRole())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res, err := db.DataStore.Collection("user").InsertOne(ctx, user)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("can not convert to oid %v", err),
		)
	}
	return &authpb.SignupResponse{
		Username: req.Username,
		Id:       oid.Hex(),
	}, nil
}

func (server *AuthServer) CheckAuth(ctx context.Context, req *authpb.AuthRequest) (*authpb.AuthResponse, error) {
	return &authpb.AuthResponse{}, nil
}
