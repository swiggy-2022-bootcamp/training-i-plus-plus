package main

import (
	"fmt"
	"programs/db"
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
		
		cfg := db.CreateLocalClient()
		//fmt.Println(cfg)
		//db.CreateTable(cfg);
		db.TableExists(cfg,"my-table");
		db.InsertItem(cfg,"101")
		//db.InsertItem(cfg,"102")
		db.InsertItem(cfg,"103")
		db.GetAllItems(cfg)
		db.GetItem(cfg,"101")
		db.DeleteItem(cfg,"101")
		db.GetAllItems(cfg)
	}

func (u User) userDetails(){
	fmt.Printf(" name is %s\n",u.username)
	fmt.Printf(" password is %s\n",u.password)
	fmt.Printf(" category is %s\n",u.category)
	fmt.Printf(" age is %d\n",u.age)
}



