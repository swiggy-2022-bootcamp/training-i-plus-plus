package Controllers;

var expertid int=1;
var userid int=1;
type User struct{
	Id int
	username string
	password string
	email string
	phone int
}

type Expert struct{
	id              int
	username     	string
	skill   	    string
	email 			string
	isAvailable     bool
	served			int
	rating          []RatingStruct
}

var userList=[]User{{Id:1,username:"uday",password:"kiran",email:"udaysonubakka123@gmail.com",phone:9440}}
var ExpertMap = map[string][]int{};
var ExpertList=[]Expert{}

func SignUpExpert(username string,skill string, email string){
	NewExpert:=Expert{id:expertid,username:username,skill:skill,email:email,isAvailable:true,served:0}
	NewExpert.rating=[]RatingStruct{}
	ExpertList=append(ExpertList,NewExpert);
	ExpertMap[skill]=append(ExpertMap[skill],expertid);
	expertid=expertid+1
}

func LogIn(username string, password string) bool{
	for _,user:=range userList{
		if user.username==username && user.password==password{
			 
			return true;
		}
	}
	return false;
}

func GetUser(username string, password string) User{
	for _,user:=range userList{
		if user.username==username && user.password==password{
			 
			return user;
		}
	}
	return User{};
}

func SignUp(username string,email string,password string,phone int)[]User{
	userList=append(userList,User{Id:userid,username:username,password:password,email:email,phone:phone});
	userid=userid+1;
	return userList;
}