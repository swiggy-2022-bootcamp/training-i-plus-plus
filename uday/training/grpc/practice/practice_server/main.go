package main

import (
	//	"google.golang.org/grpc"
	"fmt"
    "context"
	heart_pb "heartbeat/heartbeat_pb"
	"os"
    "os/signal"
    "strconv"
	//"github.com/gin-gonic/gin"
	grpc "google.golang.org/grpc"
    "net"
)
type server struct{
    heart_pb.UnimplementedHeartBeatServiceServer
}
func(*server) UserHeartBeat(ctx context.Context,req *heart_pb.HeartBeatRequest )(*heart_pb.HeartBeatResponse, error){
    fmt.Println("user heartbeat called")
    bpm:=req.GetHeartbeat().GetBpm()
    username:=req.GetHeartbeat().GetUsername()
    
    result := "User HeartBeat is " + strconv.Itoa(int(bpm)) + ", docid = " + username + "\n"
    fmt.Println(result)
    heartBeatResponse:=heart_pb.HeartBeatResponse{
        Result:result,
    }
    return &(heartBeatResponse),nil
}

func main() {
    fmt.Println("Server started")

    opts:=[]grpc.ServerOption{}
    s:=grpc.NewServer(opts...)
    heart_pb.RegisterHeartBeatServiceServer(s,&server{})
    Listener,_:=net.Listen("tcp","localhost:8081")
    s.Serve(Listener)
    fmt.Println("welcome")
    ch:=make(chan os.Signal,1)
    signal.Notify(ch,os.Interrupt)
    <-ch
    fmt.Println("closing connection")
}
