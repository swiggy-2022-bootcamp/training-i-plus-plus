package main

import (
	"fmt"
	"encoding/json"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int 	`json:"price"`
	Platform string	`json:"platform"`
	Password string	`json:"-"`
	Tags     []string`json:"tags,omitempty"`
}

func main(){
	EncodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{"ReactJS Bootcamp",299,"LearnCodeOnlin.in","abc123",[]string{"web-dev","js"}},
		{"Angular Bootcamp",337,"LearnCodeOnlin.in","xyz456",[]string{"full-stack","google"}},
	}
	finalJson,_:=json.MarshalIndent(lcoCourses, "","\t")
	fmt.Printf("%s",finalJson)
}


func DecodeJson(){
	jsonDataFromWeb:=[]byte{
		`{
			"coursename":"ReactJs Bootcamp"
		 }`
	}

	var lcoCourse course;
	checkValid:=json.Valid(jsonDataFromWeb)
	if checkValid{
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	fmt.Printf("%#v\n",lcoCourse) 
	
}