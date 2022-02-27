package main;
import authen "expert.com/events/Controllers";
import (
	"fmt"
)
var CurrentUser authen.User=authen.User{};
func main(){
	authen.InitDB();
	// go authen.ExpertAccepted(1,2);
	for{
		var value string;
		fmt.Println("Please select one of the following valid options:...")
		fmt.Println("login");
		fmt.Println("signup");
		fmt.Println("createExpert")
		fmt.Scanln(&value);
		if value=="login"{
			break;
		}else if value=="signup"{
			var username,email,password string;
			var phone int;
			fmt.Println("Enter Username")
			fmt.Scanln(&username)
			fmt.Println("Enter email")
			fmt.Scanln(&email)
			fmt.Println("Enter Password")
			fmt.Scanln(&password)
			fmt.Println("Enter Phone numbers")
			fmt.Scanln(&phone)
			authen.SignUp(username,email,password,phone)
		}		else if value=="createExpert"{
			var username,email,skill string;
			fmt.Println("Enter username")
			fmt.Scanln(&username)
			fmt.Println("Enter email")
			fmt.Scanln(&email)
			fmt.Println("Enter skills")
			fmt.Scanln(&skill)
			authen.SignUpExpert(username,skill,email)
		}  
	}
	
	for{
	    var username, password string;
		fmt.Print("Enter Username : ");
		fmt.Scanln(&username);
		fmt.Print("Enter Password : ");
		fmt.Scanln(&password);
		if authen.LogIn(username,password){
			CurrentUser=authen.GetUser(username,password)
			fmt.Println("---------------Login Successful---------------------");
			break;
		} else{
			fmt.Println("---------------Invalid Username or Password------------------");
		}
	}

	
	fmt.Println(authen.GetSkills())
	var cmd string;
	for{
		fmt.Println("--------------------------------------------------------")
		fmt.Scanln(&cmd)
		fmt.Println("Entered Command :",cmd);
		fmt.Println("--------------------------------------------------------");

		if cmd=="end" {
			fmt.Println("--------------Connection Terminated-----------------");
			break;
		} else if  cmd=="skills" {
			fmt.Println(authen.GetSkills())
		} else if cmd=="carpenter"{
			output,code:=authen.BookEmployee("carpenter",CurrentUser.Id)
			if(code==404){
				fmt.Println("Carpenters are busy... Please wait system will allocate automatically...")
			} else{
				fmt.Println("The following person is assigned to you. He will contact you.")
				fmt.Println(output,"output")
			}
		} else if cmd=="painter"{
			output,code:=authen.BookEmployee("painter",CurrentUser.Id)
			if(code==404){
				fmt.Println("Painters are busy... Please wait system will allocate automatically...")
			} else{
				fmt.Println("The following person is assigned to you. He will contact you.")
				fmt.Println(output)
			}
		} else if cmd=="plumber"{
			output,code:=authen.BookEmployee("plumber",CurrentUser.Id)
			if(code==404){
				fmt.Println("Plumbers are busy... Please wait system will allocate automatically...")
			} else{
				fmt.Println("The following person is assigned to you. He will contact you.")
				fmt.Println(output)
			}
		} else if cmd=="completed"{
			var id int;
			fmt.Scanln(&id);
			authen.WorkDone(id,CurrentUser.Id);
		} else if cmd=="rate"{
			var rating,id int;
			var review string;
			fmt.Println("Enter id,skill,rating,review......")
			fmt.Scanln(&id)
			fmt.Scanln(&rating)
			fmt.Scanln(&review)
			authen.AddRating(rating,review,id)
		} else {
			fmt.Println("invalid input")
			break;
		}
		
	}
 

}