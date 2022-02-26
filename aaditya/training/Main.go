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

func main(){
	sum:= add(10,15);
	fmt.Println(sum);

	if(inventoryCheck("dolo")){
		fmt.Println("Medicine available");
	}else{
		fmt.Println("Medicine not available")
	}
	var nums = make([]int,4,7);
	var nameAgeMap = map[string]int{
		"James": 50,
		"Ali": 39,
	}
	fmt.Printf("numbers =%v, \nlength = %d, \n capacity = %d",nums,len(nums), cap(nums))
	fmt.Println(nameAgeMap["Ali"])

	


	u1 := User{"john","admin","doctor",39}

	u1.userDetails();
	
}

func (u User) userDetails(){
	fmt.Printf(" name is %s\n",u.username)
	fmt.Printf(" password is %s\n",u.password)
	fmt.Printf(" category is %s\n",u.category)
	fmt.Printf(" age is %d\n",u.age)
}



