package main
import ("fmt"
		"sync"
		"runtime"
)

func afunc(){
	for i:=0;i<100;i++{
		fmt.Println("A function:",i)
	}
	wg.Done()
}
func bfunc(){
	for i:=0;i<100;i++{
		fmt.Println("B function:",i)
	}
	wg.Done()
}
var wg=&sync.WaitGroup{}
func main(){
  runtime.GOMAXPROCS(4)
	wg.Add(2)
	go afunc()
	go bfunc()
	wg.Wait()
}