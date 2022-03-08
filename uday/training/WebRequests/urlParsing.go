package main;
import (
	"fmt"
	"net/http"
	"io/url"
)
var myurl="www.google.com"
func main(){
	fmt.Println("Welcome to handlign URLs in golang")
	fmt.Println(myurl)
	result,_:=url.Parse(myurl)
	if result!=nil{
		fmt.Println(result)
	}
}