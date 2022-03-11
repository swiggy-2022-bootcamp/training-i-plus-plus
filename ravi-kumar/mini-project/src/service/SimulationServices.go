package service

import (
	"fmt"
	"sync"
	"time"
)

//util
func OrderItem(itemCount *int, wg1, orderingWg *sync.WaitGroup) {
	for i := 0; i < 30; i++ {
		*itemCount--
		if *itemCount <= 0 {
			orderingWg.Add(1)
			orderingWg.Wait()
		}
		fmt.Println("item ordered, number of items left: ", *itemCount)
		//time.Sleep(time.Second * 1)
	}
	wg1.Done()
}

//util
func AddItem(itemCount *int, wg1, orderingWg *sync.WaitGroup) {
	for i := 0; i < 30; i++ {
		*itemCount++
		if *itemCount == 1 {
			orderingWg.Done()
		}
		fmt.Println("item added, number of items left: ", *itemCount)
		time.Sleep(time.Second * 1)
	}
	wg1.Done()
}

//util
func FulfillOrders(ch chan string, itemNo int, fulfillOrdersWg *sync.WaitGroup) {
	defer fulfillOrdersWg.Done()
	ch <- fmt.Sprintf("order id %d fulfilled.", itemNo+1)
}
