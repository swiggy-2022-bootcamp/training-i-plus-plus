package main
import "fmt"
import "time"
import "sync"
func worker(id int,job chan int,result chan int){
	for i:=range job{
		fmt.Println("worker ",id,"job",i,"--Started")
		
		fmt.Println("worker",id,"job",i,"--Finished")
		result<-i;
		time.Sleep(time.Millisecond*1000);
	}
	wg.Done()
}
var wg=sync.WaitGroup{}
func main(){
	jobs:=make(chan int);
	result:=make(chan int);
	wg.Add(5)
	for i:=0;i<5;i++{
		
	go worker(i,jobs,result);
	}
	go func(){
		for i:=0;i<10;i++{
		jobs<-i;
	}
	// close(jobs)
		}()
	
  go func(){
	  	for j:=0;j<10;j++{
		fmt.Println("Result",<-result)
	}
	// close(result)
		}()
	
	
	// close(result)
	wg.Wait()
	// close(jobs)

}