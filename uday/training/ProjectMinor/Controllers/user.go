package Controllers;

import "fmt"
 
import "container/list"


type UserExpert struct{
	userid int
	expertid int
	accepted bool 
	cost int
	skill string
}

var RelationList=[]UserExpert{}
var WaitingList=map[string][]int{}
func AddWaitingList(userid int,skill string){
	WaitingList[skill]=append(WaitingList[skill],userid)
}

func CreateRelation(userid int, expertid int,skill string){
	NewRelation:=UserExpert{userid,expertid,true,-1,skill};
	RelationList=append(RelationList,NewRelation);
}

func RemoveRelation(userid int,expertid int){
	var index int =-1;
	for ind,value:=range RelationList{
		if value.userid==userid && value.expertid==expertid{
			index=ind;
			break;
		}
	}
	if(index==-1){
		fmt.Println("User Not found");
		return;
	}
	tempSkill:=RelationList[index].skill
	if(len(WaitingList[tempSkill])>0){
		
		CreateRelation(WaitingList[tempSkill][0],expertid,tempSkill);
		fmt.Println("For user with user id",WaitingList[tempSkill][0],"Expert with exper id",expertid,"is assigned")
		WaitingList[tempSkill]=append(WaitingList[tempSkill][1:])
	
	}
	RelationList=append(RelationList[:index],RelationList[index+1:]...)

}

 func AddCost(userid int, expertid int, cost int){
	index:=0
	for ind,value:=range RelationList{
		if value.userid==userid && value.expertid==expertid{
			index=ind;
			break;
		}
	}
	RelationList[index].cost=cost;
 }

 func MakePayment(userid int, expertid int){
	index:=0
	for ind,value:=range RelationList{
		if value.userid==userid && value.expertid==expertid{
			if(value.cost==-1){
				fmt.Println("Please ask your Service provider to add cost");
				break;
			}
			fmt.Println("Please make the payment....");
			fmt.Println("Amount Payable is....",value.cost);
			index=ind;
			break;
		}
	}
	RelationList=append(RelationList[:index],RelationList[index+1:]...)	 
 }


func Run(){
	fmt.Println("golang is started",ExpertMap);
}