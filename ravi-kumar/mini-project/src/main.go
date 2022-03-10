package main

import (
	"fmt"
)

//runtime configs
// const enableAuthorization = false
// const enableGoRoutineOrdering = false
// const enableFulfillmentViaChannel = false
// const enablePrintOrders = false

//synchronization utils
// var orderingWg sync.WaitGroup
// var wg1 sync.WaitGroup
// var fulfillOrdersWg sync.WaitGroup

func main() {
	fmt.Println("Hello world, Welcome to Online Cart Management System")
	// fmt.Println("Please enter login credentials")

	// if enableAuthorization && !authorize() {
	// 	//use this deferred function when panicked on incorrect credentials
	// 	defer func() {
	// 		v := recover()
	// 		fmt.Println("\nPanic recovered: ", v)
	// 	}()
	// 	panic("Incorrect Credentials")
	// }

	// //package handling
	// var catalog = mockdata.GetProductCatalog()
	// //pass by reference
	// fmt.Println("Catalog: ")
	// printCatalog(&catalog)

	// //go routines
	// if enableGoRoutineOrdering {
	// 	wg1.Add(2)
	// 	itemCount := 20
	// 	go orderItem(&itemCount)
	// 	go addItem(&itemCount)
	// 	wg1.Wait()
	// }

	// //channels
	// if enableFulfillmentViaChannel {
	// 	//channel with 30 buffer size
	// 	ch := make(chan string, 30)
	// 	for i := 0; i < 30; i++ {
	// 		fulfillOrdersWg.Add(1)
	// 		go fulfillOrders(ch, i+1)
	// 	}
	// 	//wait till all channels are filled and only then, close
	// 	fulfillOrdersWg.Wait()
	// 	close(ch)

	// 	//print all strings in channel
	// 	fmt.Println("Order fulfillment via channels")
	// 	for str := range ch {
	// 		fmt.Println(str)
	// 	}
	// }

	// 	//Maps
	// 	if enablePrintOrders {
	// 		populateAndPrintOrders(&catalog)
	// 	}
}

// func populateAndPrintOrders(catalog *[]mockdata.Product) {
// 	//Maps
// 	orders := make(map[string][]mockdata.Product)
// 	//populate orders for each user with 3 random products
// 	for _, user := range mockdata.GetAllUsers() {
// 		for i := 0; i < 3; i++ {
// 			orders[user.UserName] = append(orders[user.UserName], (*catalog)[rand.Intn(len(*catalog))])
// 		}
// 	}

// 	//print orders
// 	for userName, order := range orders {
// 		fmt.Print("\nOrders of user with username: ", userName)
// 		printCatalog(&order)
// 	}
// }

// func fulfillOrders(ch chan string, itemNo int) {
// 	defer fulfillOrdersWg.Done()
// 	ch <- fmt.Sprintf("order id %d fulfilled.", itemNo+1)
// }

// func addItem(itemCount *int) {
// 	for i := 0; i < 30; i++ {
// 		*itemCount++
// 		if *itemCount == 1 {
// 			orderingWg.Done()
// 		}
// 		fmt.Println("item added, number of items left: ", *itemCount)
// 		time.Sleep(time.Second * 1)
// 	}
// 	wg1.Done()
// }

// func orderItem(itemCount *int) {
// 	for i := 0; i < 30; i++ {
// 		*itemCount--
// 		if *itemCount <= 0 {
// 			orderingWg.Add(1)
// 			orderingWg.Wait()
// 		}
// 		fmt.Println("item ordered, number of items left: ", *itemCount)
// 		//time.Sleep(time.Second * 1)
// 	}
// 	wg1.Done()
// }

// func authorize() bool {
// 	var userName, password string
// 	fmt.Print("user name: ")
// 	fmt.Scan(&userName)
// 	fmt.Print("password: ")
// 	fmt.Scan(&password)
// 	return mockdata.Authenticate(userName, password)
// }

// func printCatalog(catalog *[]mockdata.Product) {
// 	for _, product := range *catalog {
// 		fmt.Println("\nname: ", product.Name,
// 			"\nprice: ", product.Price,
// 			"\ndescription: ", product.Description,
// 			"\nseller: ", product.Seller,
// 			"\nrating: ", product.Rating,
// 			"\nreview: ", strings.Join(product.Review, ", "))
// 	}
// }
