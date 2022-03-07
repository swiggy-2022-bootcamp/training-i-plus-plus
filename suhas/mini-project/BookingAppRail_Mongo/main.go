package main 

import (
	"fmt"
	"BookingAppMongo/models"
	"BookingAppMongo/database"
	
)


func main() {
	newAdmin := models.Admin{
		Name:"kdlns",
		AdminId:"AD00dfg0ckas1",
	}
	err:=database.InsertAdmin(newAdmin)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("done")
	res := database.ReadAllAdmin()
	fmt.Println(res)
}