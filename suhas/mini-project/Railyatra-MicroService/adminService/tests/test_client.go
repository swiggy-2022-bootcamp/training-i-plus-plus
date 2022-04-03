package main

import (
	"context"
	"fmt"

	//"authService/routes"

	pb "adminService/protobuf"

	"google.golang.org/grpc"
)

func main() {
	// Set up connection with the grpc server

	conn, err := grpc.Dial("localhost:6011", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	c := pb.NewAvailTicketServiceClient(conn)

	// Lets invoke the remote function from client on the server
	// resp, err := c.AvailTicketService(
	// 	context.Background(),
	// 	&pb.ChargeRequest{
	// 		Amount:       500,
	// 		Receiptemail: "syt@gmail.com",
	// 		Ticketid:     "8098a12f8901ad",
	// 	},
	// )

	fmt.Print("sim")
	resp, err := c.GetTicketConfirmation(
		context.Background(),
		&pb.AvailTicketRequest{
			TrainId:      "6235859657c5241404ee2ac8",
			NumOfTickets: 2,
		},
	)
	if err != nil {
		fmt.Printf("Error while making tx, %v\n", err)
	} else {
		fmt.Println(resp.Message)
		fmt.Println(resp.Station1, resp.Station2)
	}
}
