package main
import "fmt"

func main() {
   var mp map[int]string=map[int]string{0:"uday",1:"kiran",2:"bakka"};
	 
	fmt.Println(mp);
    for key,val := range mp{
		fmt.Println(key,val,mp[key]);
	}
	
}