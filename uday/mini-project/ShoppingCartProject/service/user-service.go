package service
import (
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
)

type UserService interface{
	SignUpUser(username string, password string, email string, phone int)
	IsUserPresent(username string, password string)bool
	GetUser(username string, password string)entity.User
	SignIn(username string, password string) entity.User
}

type userService struct{
	UserList []entity.User
}

func UserNew() UserService{
	return	&userService{[]entity.User{}}
}

func (s *userService)	SignIn(username string, password string) entity.User{
	var user entity.User
	for _,value:=range s.UserList{
		if value.Username==username && value.Password==password{
			user=value;
			break;
		}
	}
	return user
}
var userid=0
func (s *userService)SignUpUser(username string, password string, email string, phone int){
	NewUser:=entity.User{Id:userid,Username:username,Password:password,Email:email,Phone:phone}
 	s.UserList=append(s.UserList,NewUser)
 	expertid=expertid+1
}

func (s *userService)IsUserPresent(username string, password string)bool{
	var isPresent bool=false;
	for _,value:=range s.UserList{
		if value.Password==password && value.Username==username{
			isPresent=true
			break
		}
	}
	return isPresent
}

func (s *userService)GetUser(username string, password string)entity.User{
	var user entity.User;
	for _,value:=range s.UserList{
		if value.Password==password && value.Username==username{
			user=value
			break
		}
	}
	return user
}
