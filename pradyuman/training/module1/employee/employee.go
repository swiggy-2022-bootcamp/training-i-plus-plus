package employee

import "fmt"

type Worker struct{
	Name string
	Age  int
}

func (w Worker) HelloWorker(){
	fmt.Printf("hello %s to the company",w.Name)
}