package controller

import (
	repository "User-Service/Repository"
	errors "User-Service/errors"
	mockdata "User-Service/model"
	service "User-Service/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
)

var userService service.IUserService

func init() {
	userService = service.InitUserService(&repository.MongoDAO{})
}

func LogInUser(c *gin.Context) {
	service.InfoLogger.Println("User login attempted. Client IP: ", c.ClientIP())
	var logInDTO mockdata.LogInDTO
	json.NewDecoder(c.Request.Body).Decode(&logInDTO)
	jwtToken, error := userService.LogInUser(logInDTO)

	if error != nil {
		userError, ok := error.(*errors.UserError)
		if ok {
			service.ErrorLogger.Println(userError.ErrorMessage+" Client IP: ", c.ClientIP())
			c.JSON(userError.Status, userError.ErrorMessage)
			return
		} else {
			service.ErrorLogger.Println("userError casting error in GetUserById")
			return
		}
	}
	service.InfoLogger.Println("Sucessful login. Client IP: ", c.ClientIP())
	c.JSON(http.StatusOK, jwtToken)
}

func CreateUser(c *gin.Context) {
	service.InfoLogger.Println("User registration attempted. Client IP: ", c.ClientIP())
	var newUser mockdata.User
	json.NewDecoder(c.Request.Body).Decode(&newUser)

	insertedId, jwtToken, err := userService.CreateUser(newUser)
	if err != nil {
		service.ErrorLogger.Println(err.Error()+" Client IP: ", c.ClientIP())
		c.JSON(http.StatusFailedDependency, err)
		return
	}

	//response body
	type ResponseBody struct {
		InsertId string
		JwtToken string
	}

	var responseBody ResponseBody = ResponseBody{insertedId, jwtToken}
	service.InfoLogger.Println("Sucessful registration. Client IP: ", c.ClientIP())
	c.JSON(http.StatusOK, responseBody)
}

func GetAllUsers(c *gin.Context) {
	service.InfoLogger.Println("Get all users attempted. Client IP: ", c.ClientIP())
	//access: only admin
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	if mockdata.Role(acessorUserRole) != mockdata.Admin {
		service.ErrorLogger.Println(errors.AccessDenied().ErrorMessage+" Client IP: ", c.ClientIP())
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}
	users := userService.GetAllUsers()
	service.InfoLogger.Println("Sucessful Get all users. Client IP: ", c.ClientIP())
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	service.InfoLogger.Println("Get User By Id attempted. Client IP: ", c.ClientIP())
	//access: admin or the user itself
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	acessorUserId := c.Param("acessorUserId")
	userId := c.Param("userId")

	if !(mockdata.Role(acessorUserRole) == mockdata.Admin || acessorUserId == userId) {
		service.ErrorLogger.Println(errors.AccessDenied().ErrorMessage+" Client IP: ", c.ClientIP())
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	userRetrieved, error := userService.GetUserById(userId)

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
	service.InfoLogger.Println("Sucessful Get User By Id. Client IP: ", c.ClientIP())
	c.JSON(http.StatusOK, userRetrieved)
}

func UpdateUserById(c *gin.Context) {
	service.InfoLogger.Println("Update User By Id attempted. Client IP: ", c.ClientIP())
	//access: admin or the user itself
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	acessorUserId := c.Param("acessorUserId")
	userId := c.Param("userId")

	if !(mockdata.Role(acessorUserRole) == mockdata.Admin || acessorUserId == userId) {
		service.ErrorLogger.Println(errors.AccessDenied().ErrorMessage+" Client IP: ", c.ClientIP())
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	var updatedUser mockdata.User
	unmarshalErr := json.NewDecoder(c.Request.Body).Decode(&updatedUser)
	if unmarshalErr != nil {
		c.JSON(errors.UnmarshallError().Status, errors.UnmarshallError().ErrorMessage)
		return
	}

	userRetrieved, error := userService.UpdateUserById(userId, updatedUser)

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
	service.InfoLogger.Println("Sucessful Update User By Id. Client IP: ", c.ClientIP())
	c.JSON(http.StatusOK, userRetrieved)

}

func DeleteUserbyId(c *gin.Context) {
	service.InfoLogger.Println("Delete User By Id attempted. Client IP: ", c.ClientIP())

	//access: admin or the user itself
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	acessorUserId := c.Param("acessorUserId")
	userId := c.Param("userId")

	if !(mockdata.Role(acessorUserRole) == mockdata.Admin || acessorUserId == userId) {
		service.ErrorLogger.Println(errors.AccessDenied().ErrorMessage+" Client IP: ", c.ClientIP())
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	successMessage, error := userService.DeleteUserbyId(userId)

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
	service.InfoLogger.Println("Sucessful Delete User By Id attempted. Client IP: ", c.ClientIP())
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
