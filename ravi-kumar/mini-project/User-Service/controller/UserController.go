package controller

import (
	errors "User-Service/errors"
	mockdata "User-Service/model"
	service "User-Service/service"
	"encoding/json"
	"fmt"
	"net/http"

	//"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LogInUser(c *gin.Context) {
	var logInDTO mockdata.LogInDTO
	json.NewDecoder(c.Request.Body).Decode(&logInDTO)
	jwtToken, error := service.LogInUser(logInDTO)

	if error != nil {
		userError, ok := error.(*errors.UserError)
		if ok {
			c.JSON(userError.Status, userError.ErrorMessage)
			return
		} else {
			fmt.Println("userError casting error in GetUserById")
			return
		}
	}
	c.JSON(http.StatusOK, jwtToken)
}

func CreateUser(c *gin.Context) {
	result, jwtToken, err := service.CreateUser(&c.Request.Body)
	if err != nil {
		c.JSON(http.StatusFailedDependency, err)
		return
	}

	//response body
	type ResponseBody struct {
		InsertId string
		JwtToken string
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	var responseBody ResponseBody = ResponseBody{insertedId, jwtToken}

	c.JSON(http.StatusOK, responseBody)
}

func GetAllUsers(c *gin.Context) {
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	userId := c.Param("userId")
	userRetrieved, error := service.GetUserById(userId)

	if error != nil {
		userError, ok := error.(*errors.UserError)
		if ok {
			c.JSON(userError.Status, userError.ErrorMessage)
			return
		} else {
			fmt.Println("userError casting error in GetUserById")
			return
		}
	}
	c.JSON(http.StatusOK, userRetrieved)
}

func UpdateUserById(c *gin.Context) {
	userId := c.Param("userId")
	userRetrieved, error := service.UpdateUserById(userId, &c.Request.Body)

	if error != nil {
		userError, ok := error.(*errors.UserError)
		if ok {
			c.JSON(userError.Status, userError.ErrorMessage)
			return
		} else {
			fmt.Println("userError casting error in UpdateUserById")
			return
		}
	}
	c.JSON(http.StatusOK, userRetrieved)

}

func DeleteUserbyId(c *gin.Context) {
	userId := c.Param("userId")
	successMessage, error := service.DeleteUserbyId(userId)

	if error != nil {
		userError, ok := error.(*errors.UserError)
		if ok {
			c.JSON(userError.Status, userError.ErrorMessage)
			return
		} else {
			fmt.Println("userError casting error in DeleteUserbyId")
			return
		}
	}
	c.JSON(http.StatusOK, successMessage)
}

// func Authorize(userName string, password string) bool {
// 	authorized := mockdata.Authenticate(userName, password)
// 	if !authorized {
// 		//use this deferred function when panicked on incorrect credentials
// 		defer func() {
// 			v := recover()
// 			fmt.Println("\nPanic recovered: ", v)
// 		}()
// 		panic("Incorrect Credentials")
// 	}
// 	return authorized
// }

// func SimulateOrders() {
// 	wg1.Add(2)
// 	itemCount := 20
// 	go service.OrderItem(&itemCount, &wg1, &orderingWg)
// 	go service.AddItem(&itemCount, &wg1, &orderingWg)
// 	wg1.Wait()
// }

// func SimulateFulfillmentViaChannels() {
// 	//channel with 30 buffer size
// 	ch := make(chan string, 30)
// 	for i := 0; i < 30; i++ {
// 		fulfillOrdersWg.Add(1)
// 		go service.FulfillOrders(ch, i+1, &fulfillOrdersWg)
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
// 		for _, orderItem := range order {
// 			PrintProduct(&orderItem)
// 		}
// 	}
//}
