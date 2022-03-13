package controller;
import (
	// "github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
)
type UserExpertController interface{
	AddWaitingList(userid int,skill string)
	CreateRelation(userid int, expertid int,skill string)
	RemoveRelation(userid int,expertid int) bool
	AddCost(userid int, expertid int, cost int)
	MakePayment(userid int, expertid int)
}

type userexpertController struct{
	service service.UserExpertService
}

func NewUserExpert(service service.UserExpertService) service.UserExpertService{
	return &userexpertController{service:service}
}

func (c *userexpertController)AddWaitingList(userid int,skill string){
	c.service.AddWaitingList(userid,skill)
}

func (c *userexpertController)CreateRelation(userid int, expertid int,skill string){
	 c.service.CreateRelation(userid , expertid ,skill )
}

func (c *userexpertController)RemoveRelation(userid int,expertid int) bool{	 
	return c.service.RemoveRelation(userid,expertid)
}

 func (c *userexpertController)AddCost(userid int, expertid int, cost int){
	c.service.AddCost(userid,expertid,cost)
 }

 func (c *userexpertController)MakePayment(userid int, expertid int){
	 c.service.MakePayment(expertid,userid)
 }
