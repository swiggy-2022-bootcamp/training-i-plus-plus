package controller

import (
	"fmt"
	"math/rand"
	"src/mockdata"
	"src/service"
	"sync"
)

var wg1 sync.WaitGroup
var orderingWg sync.WaitGroup
var fulfillOrdersWg sync.WaitGroup

func GetAllUsers() []mockdata.User {
	return mockdata.GetAllUsers()
}

func GetUserById(userId int) (user mockdata.User) {
	for _, userDetail := range mockdata.GetAllUsers() {
		if userDetail.UserId == userId {
			user = userDetail
			return
		}
	}
	return
}

func Authorize(userName string, password string) bool {
	authorized := mockdata.Authenticate(userName, password)
	if !authorized {
		//use this deferred function when panicked on incorrect credentials
		defer func() {
			v := recover()
			fmt.Println("\nPanic recovered: ", v)
		}()
		panic("Incorrect Credentials")
	}
	return authorized
}

func SimulateOrders() {
	wg1.Add(2)
	itemCount := 20
	go service.OrderItem(&itemCount, &wg1, &orderingWg)
	go service.AddItem(&itemCount, &wg1, &orderingWg)
	wg1.Wait()
}

func SimulateFulfillmentViaChannels() {
	//channel with 30 buffer size
	ch := make(chan string, 30)
	for i := 0; i < 30; i++ {
		fulfillOrdersWg.Add(1)
		go service.FulfillOrders(ch, i+1, &fulfillOrdersWg)
	}
	//wait till all channels are filled and only then, close
	fulfillOrdersWg.Wait()
	close(ch)

	//print all strings in channel
	fmt.Println("Order fulfillment via channels")
	for str := range ch {
		fmt.Println(str)
	}
}

func populateAndPrintOrders(catalog *[]mockdata.Product) {
	//Maps
	orders := make(map[string][]mockdata.Product)
	//populate orders for each user with 3 random products
	for _, user := range mockdata.GetAllUsers() {
		for i := 0; i < 3; i++ {
			orders[user.UserName] = append(orders[user.UserName], (*catalog)[rand.Intn(len(*catalog))])
		}
	}

	//print orders
	for userName, order := range orders {
		fmt.Print("\nOrders of user with username: ", userName)
		for _, orderItem := range order {
			PrintProduct(&orderItem)
		}
	}
}
