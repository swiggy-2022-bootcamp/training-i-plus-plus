package main

import (
	"context"
	"fmt"
	heartbeat_pb "heartbeat/heartbeat_pb"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
	heartbeat_pb.UnimplementedHeartBeatServiceServer
}

type heart_item struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Bpm      int32              `bson:"bpm"`
	Username string             `bson:"username"`
}

func pushUserToDb(ctx context.Context, item heart_item) primitive.ObjectID {
	res, err := collection.InsertOne(ctx, item)
	handleError(err)

	return res.InsertedID.(primitive.ObjectID)
}

func (*server) NormalAbnormalHeartBeat(stream heartbeat_pb.HeartBeatService_NormalAbnormalHeartBeatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		bpm := req.GetHeartbeat().Bpm
		var result string
		if bpm >= 60 && bpm <= 100 {
			result = fmt.Sprintf("BPM of %v is Normal", bpm)
		} else {
			result = fmt.Sprintf("BPM of %v is Abnormal", bpm)
		}
		stream.Send(&heartbeat_pb.NormalAbnormalHeatBeatResponse{
			Result: result,
		})
	}
}

func (*server) HeartBeatHistory(req *heartbeat_pb.HeartBeatHistoryRequest, stream heartbeat_pb.HeartBeatService_HeartBeatHistoryServer) error {
	fmt.Println("HeartBeatHistory() called")
	username := req.GetUsername()

	filter := bson.M{
		"username": username,
	}
	var result_data []heart_item
	cursor, err := collection.Find(context.TODO(), filter)
	handleError(err)

	cursor.All(context.Background(), &result_data)

	for _, v := range result_data {
		res := &heartbeat_pb.HeartBeatHistoryResponse{
			Heartbeat: &heartbeat_pb.HeartBeat{
				Bpm:      v.Bpm,
				Username: v.Username,
			},
		}
		stream.Send(res)
	}

	return nil
}

func (*server) LiveHeartBeat(stream heartbeat_pb.HeartBeatService_LiveHeartBeatServer) error {
	result := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&heartbeat_pb.LiveHeartBeatResponse{
				Result: result,
			})
		}
		if err != nil {
			return err
		}

		bpm := req.GetHeartbeat().GetBpm()
		docid := pushUserToDb(context.TODO(), heart_item{
			Bpm:      req.GetHeartbeat().GetBpm(),
			Username: req.GetHeartbeat().GetUsername(),
		})
		result += fmt.Sprintf("User HeartBeat = %v, docid = %v\n", bpm, docid)
	}
}

func (*server) UserHeartBeat(ctx context.Context, req *heartbeat_pb.HeartBeatRequest) (*heartbeat_pb.HeartBeatResponse, error) {
	fmt.Println("HeartBeat() called")
	heartbeat := req.GetHeartbeat().GetBpm()
	username := req.GetHeartbeat().GetUsername()

	newHeartItem := heart_item{
		Bpm:      int32(heartbeat),
		Username: username,
	}

	docid := pushUserToDb(ctx, newHeartItem)

	result := "User HeartBeat is " + strconv.Itoa(int(heartbeat)) + ", docid = " + docid.String() + "\n"

	response := heartbeat_pb.HeartBeatResponse{
		Result: result,
	}

	return &response, nil
}

func main() {
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	handleError(err)
	heartbeat_pb.RegisterHeartBeatServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting Server")
		if err := s.Serve(lis); err != nil {
			handleError(err)
		}
	}()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError(err)
	fmt.Println("MongoDB connected")

	err = client.Connect(context.TODO())
	handleError(err)

	collection = client.Database("hb").Collection("heartbeat")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Closing MongoDB connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		handleError(err)
	}

	s.Stop()
}
