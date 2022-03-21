package service;
import (
	"fmt"
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/Udaysonu/SwiggyGoLangProject/database"
	"context"
	"time"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"

)
 
var ueCollection *mongo.Collection = database.GetCollection(database.DB, "user_expert")
var waitlistCollection *mongo.Collection = database.GetCollection(database.DB, "waiting_list")


func  AddWaitingList(userid primitive.ObjectID,skill string){
	newRelation:=entity.UserExpert{userid,userid,true,-1,skill};
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	 waitlistCollection.InsertOne(ctx, newRelation)
	// s.WaitingList[skill]=append(s.WaitingList[skill],userid)
}

func CreateRelation(userid primitive.ObjectID, expertid primitive.ObjectID,skill string){
	newRelation:=entity.UserExpert{userid,expertid,true,-1,skill};
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_,err:= ueCollection.InsertOne(ctx, newRelation)
	 if err!=nil{
		 fmt.Println(err)
	 }
}

func RemoveRelation(userid primitive.ObjectID,expertid primitive.ObjectID) bool{

	var result entity.UserExpert
	var boolean=false
	filter := bson.M{"userid": userid,"expertid":expertid}
	 
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println(userid,expertid)
	err := ueCollection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	} 

	var waiting_result entity.UserExpert
	 
	filter2 := bson.M{"skill": result.Skill}
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
 
	err2 := waitlistCollection.FindOne(ctx2, filter2).Decode(&waiting_result)
	
	if err2 == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
		expertCollection.UpdateOne(ctx,bson.M{"_id":expertid},bson.D{{"$set", bson.D{{"isavailable", true}}}})
	} else if err2 != nil {
		log.Fatal(err2)
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
		fmt.Println(err11)
	}
	ctx6, cancel6 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel6()
	_,err22:=waitlistCollection.DeleteOne(ctx6,bson.M{"userid":result.Userid,"expertid":result.Userid})
	if err22!=nil{
		fmt.Println(err22)
	}
	return boolean 
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

 