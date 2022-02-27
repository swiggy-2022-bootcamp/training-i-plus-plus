package main
import "fmt"


type calculate interface{
	rate_of_interest() int
	get_interest_rate() int
}

type principal_amt struct{
	amount int
}

func (pa principal_amt) rate_of_interest() int{
	return pa.amount*100;
}

func main(){
	var res calculate;
	res=principal_amt{560}
	fmt.Println(res.rate_of_interest());
}