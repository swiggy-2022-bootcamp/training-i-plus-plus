package service;
import (
	log	"github.com/Udaysonu/SwiggyGoLangProject/config"
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/Udaysonu/SwiggyGoLangProject/database"
	"context"
	"time"
 	"go.mongodb.org/mongo-driver/bson/primitive"
)
 
var ueCollection *mongo.Collection = database.GetCollection(database.DB, "user_expert")
var waitlistCollection *mongo.Collection = database.GetCollection(database.DB, "waiting_list")



func  AddWaitingList(userid primitive.ObjectID,skill string){
	newRelation:=entity.UserExpert{userid,userid,true,-1,skill,time.Now().Format("01-02-2006 15:04:05"),"","Pending"};
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	 waitlistCollection.InsertOne(ctx, newRelation)
	// s.WaitingList[skill]=append(s.WaitingList[skill],userid)
}

func CreateRelation(userid primitive.ObjectID, expertid primitive.ObjectID,skill string){
	newRelation:=entity.UserExpert{userid,expertid,false,-1,skill,time.Now().Format("01-02-2006 15:04:05"),"","Pending"};
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_,err:= ueCollection.InsertOne(ctx, newRelation)
	 if err!=nil{
		 log.Error.Println("CreateRelation Error: ",err)
	 }
}

func RemoveRelation(userid primitive.ObjectID,expertid primitive.ObjectID) (int,bool){

	var result entity.UserExpert
	var boolean=false
	filter := bson.M{"userid": userid,"expertid":expertid}
	 
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := ueCollection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		log.Error.Println("RemoveRelation Error: Record does not exists")
	} else if err != nil {
		log.Error.Println("RemoveRelation Error: ",err)
	} 

	var waiting_result entity.UserExpert
	 
	filter2 := bson.M{"skill": result.Skill}
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
 
	err2 := waitlistCollection.FindOne(ctx2, filter2).Decode(&waiting_result)
	if err2==mongo.ErrNoDocuments{
		// Do something when no record was found
		log.Error.Println("RemoveRelation Error: Record does not exists")
		expertCollection.UpdateOne(ctx,bson.M{"_id":expertid},bson.D{{"$set", bson.D{{"isavailable", true}}}})
	} else if err2 != nil {
		log.Error.Println("CreateRelation Error: ",err)
		}  else {
  		waiting_result.Expertid=result.Expertid
		ctx4, cancel4 := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel4()
		boolean=true
		ueCollection.InsertOne(ctx4,waiting_result)
	}
	
	ctx5, cancel5 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel5()
	_,err11:=ueCollection.DeleteOne(ctx5,bson.M{"userid":userid,"expertid":expertid})
	if err11!=nil{
		log.Error.Println("CreateRelation Error: ",err11)
	}
	ctx6, cancel6 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel6()
	_,err22:=waitlistCollection.DeleteOne(ctx6,bson.M{"userid":result.Userid,"expertid":result.Userid})
	if err22!=nil{
		log.Error.Println("CreateRelation Error: ",err22)
	}
	return 200,boolean 
}

 func  AddCost(userid primitive.ObjectID, expertid primitive.ObjectID, cost int){
	var result entity.UserExpert
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
 
	ueCollection.FindOne(ctx, bson.M{"userid":userid,"expertid":expertid}).Decode(&result)
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	result.Cost=cost
	ueCollection.UpdateOne(ctx2,bson.M{"userid":userid,"expertid":expertid},result)

 }

 func  MakePayment(userid primitive.ObjectID, expertid primitive.ObjectID){
	 
 }

 