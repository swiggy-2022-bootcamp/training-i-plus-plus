package main

import (
	"context"
	"fmt"
	"log"
	authpb "swiggy/train_reservation/services/auth/authpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("coul not connet %f", err)
	}
	defer cc.Close()
	c := authpb.NewAuthServiceClient(cc)
	login(c)
	// signUp(c)
}

func login(c authpb.AuthServiceClient) {
	req := &authpb.LoginRequest{
		Username: "Ishan",
		Password: "123",
	}
	res, err := c.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greet RPC %v", err)
	}
	log.Printf("Response from Auth: %v", res.GetAccessToken())
}

func signUp(c authpb.AuthServiceClient) {
	req := &authpb.SignupRequest{
		Username: "Ishan",
		Password: "123",
	}
	res, err := c.Signup(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greet RPC %v", err)
	}
	log.Printf("Response from Auth: %v", res.GetId())
}
