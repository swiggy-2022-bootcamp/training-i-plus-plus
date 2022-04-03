package main

import (
	//	"google.golang.org/grpc"
	"fmt"
    "context"
	pb "pb/p2"
	"os"
    "os/signal"
    "strconv"
	//"github.com/gin-gonic/gin"
	grpc "google.golang.org/grpc"
    "net"
)
type server struct{
    pb.UnimplementedServiceServer
}
func(*server) NameService(ctx context.Context,req *pb.NameRequest )(*pb.NameResponse, error){
    fmt.Println("user heartbeat called")
    name:=req.GetName().GetName()
    
    result := "User Name is "+ name +"\n"
    fmt.Println(result)
    nameResponse:=pb.NameResponse{
        Result:result,
    }
    return &(nameResponse),nil
}
func(*server) NumberService(ctx context.Context,req *pb.NumberRequest )(*pb.NumberResponse, error){
    fmt.Println("user heartbeat called")
    name:=req.GetNumber().GetNumber()
    
    result := "User Name is "+ strconv.Itoa(int(name)) +"\n"
    fmt.Println(result)
    nameResponse:=pb.NumberResponse{
        Result:result,
    }
    return &(nameResponse),nil
}


func main() {
    fmt.Println("Server started")

    opts:=[]grpc.ServerOption{}
    s:=grpc.NewServer(opts...)
    pb.RegisterServiceServer(s,&server{})
    Listener,_:=net.Listen("tcp","localhost:8081")
    s.Serve(Listener)
    fmt.Println("welcome")
    ch:=make(chan os.Signal,1)
    signal.Notify(ch,os.Interrupt)
    <-ch
    fmt.Println("closing connection")
}
