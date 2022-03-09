package main;
import "fmt";
 
func main(){
	 var arr[] int=[]int{1,2,3};
	 arr=append(arr,4);
	 arr=append(arr,5);
	 for _,val := range arr{
		 fmt.Println(val);
	 }
}