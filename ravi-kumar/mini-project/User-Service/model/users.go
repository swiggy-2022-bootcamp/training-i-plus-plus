package mockdata

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Fullname string             `json: "fullName" bson: "fullName"`
	UserName string             `json: "userName" bson: "userName"`
	Password string             `json: "password" bson: "password"`
	Address  string             `json: "address" bson: "address"`
}

type LogInDTO struct {
	UserName string `json: "userName" bson: "userName"`
	Password string `json: "password" bson: "password"`
}

// var AllUsers = []User{
// 	{
// 		//UserId:   1,
// 		Fullname: "Ravi Kumar",
// 		UserName: "ravi",
// 		Password: "Password",
// 		Address:  "Bangalore",
// 	},
// 	{
// 		//UserId:   2,
// 		Fullname: "Ajay",
// 		UserName: "ajay99",
// 		Password: "Password",
// 		Address:  "Delhi",
// 	},
// }

// func GetAllUsers() []User {
// 	return AllUsers
// }

// func Authenticate(UserName string, Password string) bool {
// 	users := GetAllUsers()
// 	for i := 0; i < len(users); i++ {
// 		if users[i].UserName == UserName && users[i].Password == Password {
// 			return true
// 		}
// 	}
// 	return false
// }
