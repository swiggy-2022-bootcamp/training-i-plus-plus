package main

import (
	"context"
	pb "pb/p2"
	"google.golang.org/grpc"
	"fmt"
	//"new.com/events/grpc/heatbeat/heartbeat_pb"
	//	"new.com/events/grpc/heatbeat/heartbeat_pb"
)
func NameService(c pb.ServiceClient){
	nameRequest:=pb.NameRequest{
		Name:&pb.Name{
			Name:"Udaykiran",
		},
	}
	res,err:=c.NameService(context.Background(),&nameRequest)
	fmt.Println(err,res)
}
func NumberService(c pb.ServiceClient){
	nameRequest:=pb.NumberRequest{
		Number:&pb.Number{
			Number:9765,
		},
	}
	res,err:=c.NumberService(context.Background(),&nameRequest)
	fmt.Println(err,res)
}
func main(){
	fmt.Println("Client started")
	conn,_:=grpc.Dial("localhost:8081",grpc.WithInsecure())
	defer conn.Close()
	c:=pb.NewServiceClient(conn)
	NumberService(c)
	NameService(c)
}