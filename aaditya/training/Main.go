package main

import (
	"fmt"
)
var inventory [5]string = [5]string{"paracetamol", "dolo", "aspirin","covaxin","covidshield"}

func add(a int, b int) int{
	return a + b;
}


func inventoryCheck(item string) bool{
	for i := 0; i < len(inventory); i++ {
		if(inventory[i]==item){
			return true;
		}

	}
	return false;
}

type User struct{
	username string
	password string
	category string
	age int
}
		
	func main() {
		//Assigning values to the fields in the person struct:
		s := make([]int,4)
		s[0]=10
		s[1]=20
		s[2]=30
		s[3]=40
		s =append(s[:2],s[3:]...)
		fmt.Println(s)
	}

func (u User) userDetails(){
	fmt.Printf(" name is %s\n",u.username)
	fmt.Printf(" password is %s\n",u.password)
	fmt.Printf(" category is %s\n",u.category)
	fmt.Printf(" age is %d\n",u.age)
}



