package main

import (
	"context"
	"fmt"

	//"authService/routes"

	pb "paymentService/protobuf"

	"google.golang.org/grpc"
)

func main() {
	// Set up connection with the grpc server

	conn, err := grpc.Dial("localhost:6010", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	c := pb.NewChargeServiceClient(conn)

	// Lets invoke the remote function from client on the server
	resp, err := c.Charge(
		context.Background(),
		&pb.ChargeRequest{
			Amount:       500,
			Receiptemail: "syt@gmail.com",
			Ticketid:     "8098a12f8901ad",
		},
	)
	if err != nil {
		fmt.Printf("Error while making tx, %v\n", err)
	} else {
		fmt.Println(resp.Confirmation)
		fmt.Println(resp.Message)
	}
}
