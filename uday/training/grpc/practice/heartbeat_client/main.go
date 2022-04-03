package main

import (
	"context"
	"heartbeat/heartbeat_pb"
	"google.golang.org/grpc"
	"fmt"
	//"new.com/events/grpc/heatbeat/heartbeat_pb"
	//	"new.com/events/grpc/heatbeat/heartbeat_pb"
)
func UserHeartBeat(c heartbeat_pb.HeartBeatServiceClient){
	heartbeatRequest:=heartbeat_pb.HeartBeatRequest{
		Heartbeat:&heartbeat_pb.HeartBeat{
			Bpm:75,
			Username:"udaybakka",
		},
	}
	res,err:=c.UserHeartBeat(context.Background(),&heartbeatRequest)
	fmt.Println(err,res)
}
func main(){
	fmt.Println("Client started")
	conn,_:=grpc.Dial("localhost:8081",grpc.WithInsecure())
	defer conn.Close()
	c:=heartbeat_pb.NewHeartBeatServiceClient(conn)
	UserHeartBeat(c)
	UserHeartBeat(c)
	UserHeartBeat(c)
}