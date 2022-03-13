
package users

import(
"context"
"fmt"
"log"
"strconv"
"net/http"
"time"
"github.com/gin-gonic/gin"
"github.com/go-playground/validator/v10"
helper "github.com/bhatiachahat/mongoapi/helper"
usermodel "github.com/bhatiachahat/mongoapi/models/usermodel"
database "github.com/bhatiachahat/mongoapi/db"


"golang.org/x/crypto/bcrypt"

"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func HashPassword(password string) string{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string)(bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err!= nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check=false
	}
	return check, msg
}

func Signup()gin.HandlerFunc{

	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user usermodel.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the email"})
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})
		defer cancel()
		if err!= nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the phone number"})
		}

		if count >0{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this email or phone number already exists"})
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr !=nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}

}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user usermodel.User
		var foundUser usermodel.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return 
		}

		err := userCollection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true{
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundUser.Email == nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"user not found"})
		}
		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}

func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		if err := helper.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage <1{
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 !=nil || page<1{
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{
			{"_id", bson.D{{"_id", "null"}}}, 
			{"total_count", bson.D{{"$sum", 1}}}, 
			{"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},}}}
result,err := userCollection.Aggregate(ctx, mongo.Pipeline{
	matchStage, groupStage, projectStage})
defer cancel()
if err!=nil{
	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing user items"})
}
var allusers []bson.M
if err = result.All(ctx, &allusers); err!=nil{
	log.Fatal(err)
}
c.JSON(http.StatusOK, allusers[0])}}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user usermodel.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}



// // package users

// // import (
// //     "net/http"
// // 	"encoding/json"
// // 	"github.com/gorilla/mux"
// // 	"strconv"
// // 	"math/rand"
// // 	usermodel "github.com/bhatiachahat/mongoapi/models/usermodel"
	
// // )
// // var users []usermodel.User
// // func init() {
// // 	users= append(users, usermodel.User{ID:"1",Email:"john@gmail.com",FirstName: "John", LastName: "Doe"})
// // 	users= append(users, usermodel.User{ID:"2",Email:"larry@gmail.com",FirstName: "Larry", LastName: "Wheels"})
// // 	users= append(users, usermodel.User{ID:"3",Email:"jack@gmail.com",FirstName: "Jack", LastName: "Parr"})
// //     users= append(users, usermodel.User{ID:"4",Email:"tim@gmail.com",FirstName: "Tim", LastName: "Burner"})

// // }


// // // // List all Products 
// // // func GetProducts(w http.ResponseWriter,r *http.Request){
// // //     w.Header().Set("Content-Type","application/json")
// // // 	json.NewEncoder(w).Encode(products)
// // // }

// // // Get user details
// // func GetUser(w http.ResponseWriter,r *http.Request){
// // 	w.Header().Set("Content-Type","application/json")
// // 	params :=mux.Vars(r)
// // 	for _,user:=range users{
// // 		if user.Email== params["email"] {
// // 			json.NewEncoder(w).Encode(user)
// // 			return
// // 		}
// // 	}
// // 	json.NewEncoder(w).Encode(&usermodel.User{})
// // }

// // // Register User
// // func RegisterUser(w http.ResponseWriter,r *http.Request){
// // 	w.Header().Set("Content-Type","application/json")
// // 	var user usermodel.User
// // 	_=json.NewDecoder(r.Body).Decode(&user)
// // 	user.ID=strconv.Itoa(rand.Intn(1000000))
// // 	users= append(users, user)
// // 	json.NewEncoder(w).Encode(user)


// // }

// // // Login User
// // func LoginUser(w http.ResponseWriter,r *http.Request){
// // 	w.Header().Set("Content-Type","application/json")
// // 	var user usermodel.User
// // 	_=json.NewDecoder(r.Body).Decode(&user)
// // 	user.ID=strconv.Itoa(rand.Intn(1000000))
// // 	users= append(users, user)
// // 	json.NewEncoder(w).Encode(user)


// // }

// // // //update product

// // // func UpdateProduct(w http.ResponseWriter,r *http.Request){
// // // 	w.Header().Set("Content-Type","application/json")
// // // 	params := mux.Vars(r)
// // // 	for index,item := range products {
// // // 		if item.ID == params["id"]{
// // // 		products= append(products[:index],products[index+1:]...)
// // // 		var product productmodel.Product
// // // 	_=json.NewDecoder(r.Body).Decode(&product)
// // // 	product.ID=params["id"]
// // // 	products= append(products, product)
// // // 	json.NewEncoder(w).Encode(product)
// // // 	return
// // // 		}
// // // 	}
	
// // // }

// // //delete user

// // func DeleteUser(w http.ResponseWriter,r *http.Request){
// // 	w.Header().Set("Content-Type","application/json")
// // 	params := mux.Vars(r)
// // 	for index,item := range users {
// // 		if item.Email == params["email"]{
// // 		users= append(users[:index],users[index+1:]...)
// // 		break
// // 		}
// // 	}
// // 	json.NewEncoder(w).Encode(users)

// // }


