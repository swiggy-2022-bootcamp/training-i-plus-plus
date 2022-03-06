package Controllers;
 
 
func GetSkills()[]string{
	var	skills =[] string{};
	for key,_:= range ExpertMap{
		skills=append(skills,key);
	}
	return skills;
}



func WorkDone(id int,userid int){
	for index,person := range ExpertList{
		if(person.id==id){
			person.isAvailable=true;
			ExpertList[index]=person;
			RemoveRelation(userid,id);
		}
	}
}

func AddRating(rating int, review string,id int){
	for index,value:= range ExpertList{
		if value.id==id{
			value.rating=append(value.rating,RatingStruct{rating,review})
			ExpertList[index]=value;
		}
	}
}

 