package Controllers;
import "fmt"
type points struct{
	a,b int
}

type RatingStruct struct{
	rating   int
	review   string
}

func InitDB(){
	SignUpExpert("Mahesh","carpenter","mahesh@gmail.com");
	SignUpExpert("NTR","painter","ntr@gmail.com");
	SignUpExpert("Surya","carpenter","surya@gmail.com"); 
	SignUpExpert("Arjun","carpenter","arjun@gmail.com");
	SignUpExpert("Vijay","painter","vijay@gmail.com");
	SignUpExpert("Ajeeth","carpenter","ajeeth@gmail.com");
	SignUpExpert("Prabhas","plumber","prabhas@gmail.com");
	SignUpExpert("Pawan","plumber","pawan@gmail.com");

	fmt.Println("----------------Database Loaded-------------------")
}
