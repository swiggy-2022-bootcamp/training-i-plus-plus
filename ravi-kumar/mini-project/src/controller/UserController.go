package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	mockdata "src/model"
	"sync"
	"time"

	"src/config"

	//"go.mongodb.org/mongo-driver/bson"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var wg1 sync.WaitGroup
var orderingWg sync.WaitGroup
var fulfillOrdersWg sync.WaitGroup

var client *mongo.Client
var mongoURL string = config.MONGO_URL

func init() {
	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var newUser mockdata.User
	json.NewDecoder(req.Body).Decode(&newUser)
	collection := client.Database("swiggy_mini").Collection("users")

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, newUser)

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var allUsers []mockdata.User
	collection := client.Database("swiggy_mini").Collection("users")

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, _ := collection.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var user mockdata.User
		cursor.Decode(&user)
		allUsers = append(allUsers, user)
	}

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(allUsers)
}

func GetUserById(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//storing route variables for a request
	variables := mux.Vars(req)
	var userId string = variables["userId"]
	var userRetrieved mockdata.User

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("swiggy_mini").Collection("users")

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Malformed user id")
		return
	}

	result := collection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("user with given id not found")
		return
	}

	result.Decode(&userRetrieved)

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(userRetrieved)
}

func UpdateUserById(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//storing route variables for a request
	variables := mux.Vars(req)
	userId := variables["userId"]

	reqBody, _ := ioutil.ReadAll(req.Body)

	updatedUser := &mockdata.User{}
	unmarshalErr := json.Unmarshal(reqBody, updatedUser)
	if unmarshalErr != nil {
		log.Print("Couldn't unmarshall user body in request")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("swiggy_mini").Collection("users")

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Malformed user id")
		return
	}

	result, error := collection.UpdateByID(ctx, objectId, bson.M{"$set": updatedUser})
	if error != nil {
		fmt.Println(error)
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode("Internal server error")
		return
	}
	if result.MatchedCount == 0 {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("user with given id not found")
		return
	}

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("user updated")
}

func DeleteUserbyId(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	variables := mux.Vars(req)
	userId := variables["userId"]

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("swiggy_mini").Collection("users")

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Malformed user id")
		return
	}

	result, error := collection.DeleteOne(ctx, bson.M{"_id": objectId})

	if error != nil {
		fmt.Println(error)
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode("Internal server error")
		return
	}

	if result.DeletedCount == 0 {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("user with given id not found")
		return
	}

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("user deleted")
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
