package service

import (
	"context"
	// "os/exec"
	"time"
	log "github.com/Udaysonu/SwiggyGoLangProject/config"
	"github.com/Udaysonu/SwiggyGoLangProject/database"
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

 	// "github.com/Udaysonu/SwiggyGoLangProject/config"
)

var expertCollection *mongo.Collection = database.GetCollection(database.DB, "experts")
var completedCollection *mongo.Collection = database.GetCollection(database.DB, "completed")

func GetAllExperts()(int,[]entity.Expert){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var experts []entity.Expert
	defer cancel()
	results, _ := expertCollection.Find(ctx, bson.M{})
	for results.Next(ctx) {
		var singleUser entity.Expert
		results.Decode(&singleUser) 
		experts = append(experts, singleUser)
	}
	return 200,experts
}
func Delete(id primitive.ObjectID)int{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
 	_,err:= expertCollection.DeleteOne(ctx, bson.M{"_id":id})
	if err!=nil{
		return 500
	} 
	return 200
}

func  SignUpExpert(username string,skill string, email string,location int)(int,entity.Expert){
	newExpert:=entity.Expert{Id:primitive.NewObjectID(),Username:username,Skill:skill,Email:email,IsAvailable:true,Served:0,Rating:0.0,Location:location,Reviews:[]entity.RatingStruct{}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
 	_,err:= expertCollection.InsertOne(ctx, newExpert)
	 if err!=nil{
		 return 400,newExpert
	 }
	return 200,newExpert
}

func  GetSkills()(int,[]string){
	var	skills =[] string{"carpenter","plumber","painter","beautician","labour","barber"};
	return 200,skills;
}

func  WorkDone(userid primitive.ObjectID,id primitive.ObjectID)(int,bool){
	return RemoveRelation(userid,id);
}

func GetWaitingRequest(expertid primitive.ObjectID)(int,entity.UserExpert){
	var result entity.UserExpert
	
	filter := bson.M{"expertid": expertid}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ueCollection.FindOne(ctx, filter).Decode(&result)
	return 200, result
}

func AcceptWaitingRequest(expertid primitive.ObjectID)(int,entity.UserExpert){
	var result entity.UserExpert
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ueCollection.UpdateOne(ctx, bson.M{"expertid":expertid}, bson.D{{"$set", bson.D{{"accepted",true},{"status","Accepted"},{"acceptedat", time.Now().Format("01-02-2006 15:04:05")}}}}	)
	ueCollection.FindOne(ctx,bson.M{"expertid":expertid}).Decode(&result)
	return 200,result
}

func RejectWaitingResult(expertid primitive.ObjectID)(int,bool){
	var result entity.UserExpert
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ueCollection.FindOne(ctx,bson.M{"expertid":expertid}).Decode(&result)
	RemoveRelation(result.Userid,result.Expertid)
	return 200,true
}

func CompletedRequest(expertid primitive.ObjectID,cost int)(int,entity.UserExpert){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result entity.UserExpert
	ueCollection.FindOne(ctx,bson.M{"expertid":expertid}).Decode(&result)
 	result.Cost=cost
	result.Status="Completed"
	completedCollection.InsertOne(ctx,result)
	RemoveRelation(result.Userid,result.Expertid)
	return 200,result
}


func  BookEmployee(skill string,userid primitive.ObjectID) (int,entity.Expert){
	var availablePerson entity.Expert; 
	filter := bson.M{"isavailable": true,"skill":skill}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var experts []entity.Expert
	opts := options.Find().SetSort(bson.D{{"rating", 1}})
	results, err := expertCollection.Find(ctx, filter, opts)
	for results.Next(ctx) {
		var singleUser entity.Expert
		results.Decode(&singleUser) 
		experts = append(experts, singleUser)
	}

	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		AddWaitingList(userid,skill);
		return 404,entity.Expert{};
		log.Error.Println("BookEmployee Error: Record does not exists")
	} else if err != nil {
		log.Error.Println("BookEmployee Error: ",err)
	} else if len(experts)>0{
		availablePerson=experts[0]
 		availablePerson.Served=availablePerson.Served+1
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	 
		expertCollection.UpdateOne(ctx, bson.M{"_id":availablePerson.Id}, bson.D{{"$set", bson.D{{"isavailable",false},{"served",availablePerson.Served}}}}	)
		CreateRelation(userid,availablePerson.Id,skill)
	} else if len(experts)==0{
		AddWaitingList(userid,skill);
	}

	return 200,availablePerson;
}

func  AddRating(rating int, review string,id primitive.ObjectID)int{
	var result entity.Expert
	 
	filter := bson.D{{"_id", id}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	expertCollection.FindOne(ctx, filter).Decode(&result)
	
	result.Rating=((result.Rating+float64(rating))/(float64)(len(result.Reviews)+1))
	result.Reviews=append(result.Reviews,entity.RatingStruct{rating,review})
	update := bson.D{{"$set", bson.D{{"rating", result.Rating},{"reviews",result.Reviews}}}}

	expertCollection.UpdateOne(  ctx,filter,   update)

 	return 200;
}

func   GetExperts(skill string)(int,[]entity.Expert){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var experts []entity.Expert
	defer cancel()
	results, _ := expertCollection.Find(ctx, bson.M{"skill":skill})
	for results.Next(ctx) {
		var singleUser entity.Expert
		results.Decode(&singleUser) 
		experts = append(experts, singleUser)
	}
	return 200,experts
}

func  GetExpertByID(id primitive.ObjectID) (int,entity.Expert){
	var result entity.Expert
	 
	filter := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
 
	err := expertCollection.FindOne(ctx, filter).Decode(&result)
	
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		log.Error.Println("GetExpertByID Error: No documents")
	} else if err != nil {
		log.Error.Println("GetExpertByID Error: ",err)
	} 
	return 200,result 
} 


func   FilterExpert(skill string, rating int )(int,[]entity.Expert){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var experts []entity.Expert
	defer cancel()
	results, _ := expertCollection.Find(ctx, bson.M{})
	for results.Next(ctx) {
		var singleUser entity.Expert
		results.Decode(&singleUser) 
		if singleUser.Rating>=float64(rating){
			experts = append(experts, singleUser)
		}
	}
	return 200, experts
}

 