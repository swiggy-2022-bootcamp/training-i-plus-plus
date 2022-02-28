package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	jwtmanager "swiggy/train_reservation/helpers/lib"
	db "swiggy/train_reservation/helpers/utils"
	"swiggy/train_reservation/services/auth"
	authpb "swiggy/train_reservation/services/auth/authpb"

	grpc "google.golang.org/grpc"
)

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	client, ctx, cancel := db.ConnectDB()
	defer cancel()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
		fmt.Println("MongoDB Connection Closed")
	}()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	JWTManager := jwtmanager.NewJWTManager("sbjabhbk", 40)

	authpb.RegisterAuthServiceServer(s, &auth.AuthServer{
		authpb.UnimplementedAuthServiceServer{},
		JWTManager,
	})

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//Block Untill Signal is Received
	<-ch

	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("End of program")
}
