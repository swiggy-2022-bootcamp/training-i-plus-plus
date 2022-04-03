package main

import (
	"context"
	"fmt"

	//"authService/routes"

	pb "authService/protobuf"

	"google.golang.org/grpc"
)

func main() {
	// Set up connection with the grpc server

	conn, err := grpc.Dial("localhost:6010", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	c := pb.NewAuthenticationServiceClient(conn)

	// Lets invoke the remote function from client on the server
	resp, err := c.Authenticate(
		context.Background(),
		&pb.AuthenticateRequest{
			Group: "ADMIN",
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDg4OTMyMzMsImdyb3VwIjoiQURNSU4ifQ.lwxqjAF0Z00qVkyUTAppjJSUyqG8fkog4p8z45iayuA",
		},
	)
	if err != nil {
		fmt.Printf("Error while making tx, %v\n", err)
	} else {
		fmt.Println(resp.Confirmation)
		fmt.Println(resp.Message)
	}
}
