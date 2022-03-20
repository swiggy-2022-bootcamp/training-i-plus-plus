package controller

import (
	// "github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type UserExpertController interface{
// 	AddWaitingList(userid int,skill string)
// 	CreateRelation(userid int, expertid int,skill string)
// 	RemoveRelation(userid int,expertid int) bool
// 	AddCost(userid int, expertid int, cost int)
// 	MakePayment(userid int, expertid int)
// }

// type userexpertController struct{
// 	service service.UserExpertService
// }

// func NewUserExpert(service service.UserExpertService) service.UserExpertService{
// 	return &userexpertController{service:service}
// }

func  AddWaitingList(userid primitive.ObjectID,skill string){
	service.AddWaitingList(userid,skill)
}

func  CreateRelation(userid primitive.ObjectID, expertid primitive.ObjectID,skill string){
	 service.CreateRelation(userid , expertid ,skill )
}

func  RemoveRelation(userid primitive.ObjectID,expertid primitive.ObjectID) bool{	 
	return service.RemoveRelation(userid,expertid)
}

 func  AddCost(userid primitive.ObjectID, expertid primitive.ObjectID, cost int){
	service.AddCost(userid,expertid,cost)
 }

 func MakePayment(userid primitive.ObjectID, expertid primitive.ObjectID){
	 service.MakePayment(expertid,userid)
 }
