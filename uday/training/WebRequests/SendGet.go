package main

import (
	 
	"fmt"
	"io/ioutil"
	"net/http"
)
const url="https://dog.ceo/api/breeds/image/random";
func main(){
	fmt.Println("LCO web request");

	response, err := http.Get(url);

	if err != nil {
		panic(err)
	}
	// fmt.Println("Response is of type: %T \n", response)
	// response.Body.Close()//  caller's responsibiity to close the connection

	databytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
	response.Body.Close()
}