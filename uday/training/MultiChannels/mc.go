package main
import "fmt"
import "sync"

func run(job chan int){
	 for i:=range job{
		fmt.Println(i);
	 }
	 
	 wg.Done()
}
func insert(job chan int){
	for i:=1;i<=10;i++{
		job<-i;
	}
	close(job)
}
var wg=sync.WaitGroup{}
func main(){
	wg.Add(1)
	chan1:=make(chan int)
 
	go run(chan1);
	go insert(chan1);
	
	wg.Wait()
}