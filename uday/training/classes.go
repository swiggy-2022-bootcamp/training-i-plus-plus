package main;
import "fmt";
type points struct{
	a,b int
}
func main(){
 
	point3:=points{b:4};
 
	point3.a=10;

	a,b := point3.call();
	
	fmt.Println(a,b);
}

func (obj points) call() (m int,n int){
	fmt.Println(obj.a,obj.b);
	m=obj.a;
	n=obj.b;
	return;
}