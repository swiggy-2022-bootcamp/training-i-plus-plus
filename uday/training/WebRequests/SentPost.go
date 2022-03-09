package main;
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)
func main(){
   fmt.Println("Welcome to web vervideo - LCO")  
   PerformPostFORMRequest()
}

func PerformPostJSONRequest(){
	const myurl="https://www.google.com/post"
	//fake json payload

	requestBody:=strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"learnCodeOnline.in"
		}
	`)
	response,err:=http.Post(myurl,"application/json",requestBody)
		if err!=nil{
			panic(err)
		}
		defer response.Body.Close()
		content,_:=ioutil.ReadAll(response.Body)
		fmt.Println(string(content))
}


func PerformPostFORMRequest(){
	const myurl="https://google.com/postform"
	//formdata
	data:=url.Values{}
	data.Add("firstname","Hitesh")
	data.Add("lastname","Choudary")
	data.Add("email","hitesh@go.dev")
	response,_:=http.PostForm(myurl,data)
	defer response.Body.Close()
	content,_:=ioutil.ReadAll(response.Body)
    fmt.Println(string(content))
}