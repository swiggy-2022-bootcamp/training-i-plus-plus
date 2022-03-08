package Controllers;

func BookEmployee(skill string,userid int) (Expert,int){
	var availablePerson Expert;
	var g_index int=-1;
	for index,person:= range ExpertList{
		if person.isAvailable==true  && person.skill==skill{
			if g_index==-1 {
				availablePerson=person;
				g_index=index
			} else if availablePerson.served>person.served {
				availablePerson=person;
				g_index=index;
			}
		}
	}
	if(g_index==-1){
		// UserExpertQueue.PushBack(UserExpert{2,-1,false})
		AddWaitingList(userid,skill);
		return Expert{},404;
	}
	availablePerson.isAvailable=false;
	availablePerson.served=availablePerson.served+1
	ExpertList[g_index]=availablePerson;
	CreateRelation(userid,ExpertList[g_index].id,ExpertList[g_index].skill);
	return availablePerson,200;
}


