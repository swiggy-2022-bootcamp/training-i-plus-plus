package main;
import "fmt";
type points struct{
	a,b int
}
func main(){
	point1:=points{2,3};
	point2:=points{a:2};
	point3:=points{b:4};
	fmt.Println(point1);
	fmt.Println(point2);
	fmt.Println(point3);

	fmt.Println(point3.a);
	point3.a=10;
}