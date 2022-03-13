package service;
import (
	"fmt"
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
)
type UserExpertService interface{
	AddWaitingList(userid int,skill string)
	CreateRelation(userid int, expertid int,skill string)
	RemoveRelation(userid int,expertid int) bool
	AddCost(userid int, expertid int, cost int)
	MakePayment(userid int, expertid int)
}

type userexpertService struct{
	 RelationList []entity.UserExpert
	 WaitingList map[string][]int
}
 
func NewUserExpertService() UserExpertService{
	 
	return &userexpertService{[]entity.UserExpert{},map[string][]int{}}
	
}

func (s *userexpertService)AddWaitingList(userid int,skill string){
	s.WaitingList[skill]=append(s.WaitingList[skill],userid)
}

func (s *userexpertService)CreateRelation(userid int, expertid int,skill string){
	NewRelation:=entity.UserExpert{userid,expertid,true,-1,skill};
	s.RelationList=append(s.RelationList,NewRelation);
}

func (s *userexpertService)RemoveRelation(userid int,expertid int) bool{
	var index int =-1;
	for ind,value:=range s.RelationList{
		if value.Userid==userid && value.Expertid==expertid{
			index=ind;
			break;
		}
	}
	fmt.Println("entered remove relation",index)
	if(index==-1){
		fmt.Println("User Not found");
		return false;
	}
	var bl=false;
	tempSkill:=s.RelationList[index].Skill
	if(len(s.WaitingList[tempSkill])>0){
		
		s.CreateRelation(s.WaitingList[tempSkill][0],expertid,tempSkill);
		fmt.Println("For user with user id",s.WaitingList[tempSkill][0],"Expert with exper id",expertid,"is assigned")
		s.WaitingList[tempSkill]=append(s.WaitingList[tempSkill][1:])
	    bl=true;
	}
	s.RelationList=append(s.RelationList[:index],s.RelationList[index+1:]...)
	return bl;
}

 func (s *userexpertService)AddCost(userid int, expertid int, cost int){
	index:=0
	for ind,value:=range s.RelationList{
		if value.Userid==userid && value.Expertid==expertid{
			index=ind;
			break;
		}
	}
	s.RelationList[index].Cost=cost;
 }

 func (s *userexpertService)MakePayment(userid int, expertid int){
	index:=0
	for ind,value:=range s.RelationList{
		if value.Userid==userid && value.Expertid==expertid{
			if(value.Cost==-1){
				fmt.Println("Please ask your Service provider to add cost");
				break;
			}
			fmt.Println("Please make the payment....");
			fmt.Println("Amount Payable is....",value.Cost);
			index=ind;
			break;
		}
	}
	s.RelationList=append(s.RelationList[:index],s.RelationList[index+1:]...)	 
 }


// func Run(){
// 	fmt.Println("golang is started");
// }
