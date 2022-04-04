package controller
import ( 
		"context"
		"fmt"
		grpc "google.golang.org/grpc"
        "expertserver/database"
        "expertserver/model"
        "expertserver/config"
        "go.mongodb.org/mongo-driver/bson"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/bson/primitive"
        "time"
        "net"
        "os"
        "os/signal"
     //   "github.com/gin-gonic/gin"

)

func(*server) ExpertService(ctx context.Context,req *pb.ExpertRequest )(*pb.ExpertResponse, error){
    fmt.Println("user heartbeat called")
    id_s:=req.GetExpertid().GetId()
    var result entity.Expert
    id,_:=primitive.ObjectIDFromHex(id_s)
	filter := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
 
	expertCollection.FindOne(ctx, filter).Decode(&result)
	reactStruct:=[]*pb.Ratingstruct{};
    for _,val:=range result.Reviews{
        reactStruct=append(reactStruct,&pb.Ratingstruct{Rating:int32(val.Rating),Review:val.Review,})
    }
    fmt.Println(result)
    result1:=pb.Expert{
        Id:result.Id.Hex(),
        Username:result.Username,
        Skill:result.Skill,
        Email:result.Email,
        Isavailable:result.IsAvailable,
        Served:int32(result.Served),
        Rating:float32(result.Rating),
        Location:int64(result.Location),
        Ratingstruct:reactStruct,
    }
    nameResponse:=pb.ExpertResponse{
        Expert:&result1,
    }
    return &(nameResponse),nil
}
 