package service
import (
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"fmt"	
)
var expertid int=1;
type ExpertService interface{
	GetSkills() []string
	WorkDone(int,int)
	// AddRating()
	SignUpExpert(username string,skill string, email string)
	BookEmployee(skill string,userid int) (entity.Expert,int)
	InitDB()
}

type expertService struct{
	ExpertMap map[string][]int
	ExpertList []entity.Expert
}

func ExpertNew() ExpertService{
	return &expertService{map[string][]int{},[]entity.Expert{}}
}

func (service *expertService)SignUpExpert(username string,skill string, email string){
	NewExpert:=entity.Expert{Id:expertid,Username:username,Skill:skill,Email:email,IsAvailable:true,Served:0}
	// NewExpert.rating=[]RatingStruct{}
	service.ExpertList=append(service.ExpertList,NewExpert);
	// _,found:=service.ExpertMap[skill]
	// if found==false{
	// 	service.ExpertMap[skill]=[]int{}
	// }
	service.ExpertMap[skill]=append(service.ExpertMap[skill],expertid);
	expertid=expertid+1
}

func  (service *expertService)GetSkills()[]string{
	var	skills =[] string{};
	for key,_:= range service.ExpertMap{
		skills=append(skills,key);
	}
	return skills;
}



func (service *expertService)WorkDone(id int,userid int){
	for index,person := range service.ExpertList{
		if(person.Id==id){
			person.IsAvailable=true;
			service.ExpertList[index]=person;
			// RemoveRelation(userid,id);
		}
	}
}

func (service *expertService)BookEmployee(skill string,userid int) (entity.Expert,int){
	var availablePerson entity.Expert;
	var g_index int=-1;
	for index,person:= range service.ExpertList{
		if person.IsAvailable==true  && person.Skill==skill{
			if g_index==-1 {
				availablePerson=person;
				g_index=index
			} else if availablePerson.Served>person.Served {
				availablePerson=person;
				g_index=index;
			}
		}
	}

	if(g_index==-1){
		// UserExpertQueue.PushBack(UserExpert{2,-1,false})
		// AddWaitingList(userid,skill);
		return entity.Expert{},404;
	}

	availablePerson.IsAvailable=false;
	availablePerson.Served=availablePerson.Served+1
	service.ExpertList[g_index]=availablePerson;
	// CreateRelation(userid,service.ExpertList[g_index].id,service.ExpertList[g_index].skill);
	return availablePerson,200;
}

 

// func (service *expertService)AddRating(rating int, review string,id int){
// 	for index,value:= range service.ExpertList{
// 		if value.id==id{
// 			value.rating=append(value.rating,RatingStruct{rating,review})
// 			service.ExpertList[index]=value;
// 		}
// 	}
// }



func (service *expertService) InitDB(){
	fmt.Println(service.ExpertList)
	fmt.Println(service.ExpertMap)
	service.SignUpExpert("Mahesh","carpenter","mahesh@gmail.com");
	service.SignUpExpert("NTR","painter","ntr@gmail.com");
	service.SignUpExpert("Surya","carpenter","surya@gmail.com"); 
	service.SignUpExpert("Arjun","carpenter","arjun@gmail.com");
	service.SignUpExpert("Vijay","painter","vijay@gmail.com");
	service.SignUpExpert("Ajeeth","carpenter","ajeeth@gmail.com");
	service.SignUpExpert("Prabhas","plumber","prabhas@gmail.com");
	service.SignUpExpert("Pawan","plumber","pawan@gmail.com");

	fmt.Println("----------------Database Loaded-------------------")
}
