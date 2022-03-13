

package products
import (
	"context"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "github.com/bhatiachahat/mongoapi/helper"
	productmodel "github.com/bhatiachahat/mongoapi/models/productmodel"
	database "github.com/bhatiachahat/mongoapi/db"
	
	//"github.com/akhil/golang-jwt-project/models"
	//"github.com/akhil/golang-jwt-project/database"
	//"golang.org/x/crypto/bcrypt"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")
var validate = validator.New()

func AddProduct()gin.HandlerFunc{

	return func(c *gin.Context){
	
		if err := helper.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var product productmodel.Product

		if err := c.BindJSON(&product); err != nil {
			
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(product)
		if validationErr != nil {
			fmt.Println("130")
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return
		}

		count, err := productCollection.CountDocuments(ctx, bson.M{"title":product.Title})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the title"})
		}
		if count >0{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"This product already exists"})
		}

		product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.ID = primitive.NewObjectID()
		product.Product_id = product.ID.Hex()
	

		resultInsertionNumber, insertErr := productCollection.InsertOne(ctx, product)
		if insertErr !=nil {
			msg := fmt.Sprintf("Product was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}

}
func GetAllProducts() gin.HandlerFunc{
	return func(c *gin.Context){
		// if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		// 	return
		// }
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
				{"product_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},}}}
result,err := productCollection.Aggregate(ctx, mongo.Pipeline{
	matchStage, groupStage, projectStage})
defer cancel()
if err!=nil{
	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing products"})
}
var allproducts []bson.M
if err = result.All(ctx, &allproducts); err!=nil{
	log.Fatal(err)
}
c.JSON(http.StatusOK, allproducts[0])}}

func GetProduct() gin.HandlerFunc{
	return func(c *gin.Context){
		productId := c.Param("product_id")

		if err := helper.MatchUserTypeToUid(c, productId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var product productmodel.Product
		err := productCollection.FindOne(ctx, bson.M{"product_id":productId}).Decode(&product)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}
func DeleteProduct()gin.HandlerFunc{

	return func(c *gin.Context){
		productId := c.Param("product_id")
		if err := helper.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
        if err := helper.MatchUserTypeToUid(c, productId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// var product productmodel.Product
		res,err := productCollection.DeleteOne(ctx, bson.M{"product_id":productId})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,res)



		

		// validationErr := validate.Struct(product)
		// if validationErr != nil {
		// 	fmt.Println("130")
		// 	c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
		// 	return
		// }

		

		// password := HashPassword(*user.Password)
		// user.Password = &password

		// count, err = productCollection.CountDocumets(ctx, bson.M{"category":product.Category})
		// defer cancel()
		// if err!= nil {
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the category"})
		// }

		

		// product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// product.ID = primitive.NewObjectID()
		// product.Product_id = product.ID.Hex()
		//token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
	//	user.Token = &token
	//	user.Refresh_token = &refreshToken

		
	}

}

func UpdateProduct()gin.HandlerFunc{

	return func(c *gin.Context){
		productId := c.Param("product_id")
		if err := helper.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
        if err := helper.MatchUserTypeToUid(c, productId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		updated_time, _:=time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)) 
        update := bson.M{"$set":bson.M{"updated_at":updated_time}}
		// var product productmodel.Product
		res,err := productCollection.UpdateOne(ctx, bson.M{"product_id":productId},update)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,res)



		

		// validationErr := validate.Struct(product)
		// if validationErr != nil {
		// 	fmt.Println("130")
		// 	c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
		// 	return
		// }

		

		// password := HashPassword(*user.Password)
		// user.Password = &password

		// count, err = productCollection.CountDocumets(ctx, bson.M{"category":product.Category})
		// defer cancel()
		// if err!= nil {
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the category"})
		// }

		

		// product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// product.ID = primitive.NewObjectID()
		// product.Product_id = product.ID.Hex()
		//token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
	//	user.Token = &token
	//	user.Refresh_token = &refreshToken

		
	}

}



// package products

// import (
//     "net/http"
// 	"encoding/json"
// 	"github.com/gorilla/mux"
// 	"strconv"
// 	"math/rand"
// 	productmodel "github.com/bhatiachahat/mongoapi/models/productmodel"
	
// )
// var products []productmodel.Product
// func init() {
// 	products= append(products, productmodel.Product{ID:"1",Title:"Samsung Galaxy",Category: "Mobile"})
// 	products= append(products, productmodel.Product{ID:"2",Title:"Samsung Note",Category: "Mobile"})
// 	products= append(products, productmodel.Product{ID:"3",Title:"Samsung A50",Category: "Mobile"})
// 	products= append(products, productmodel.Product{ID:"4",Title:"Samsung M52",Category: "Mobile"})
// }


// // List all Products 
// func GetProducts(w http.ResponseWriter,r *http.Request){
//     w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(products)
// }

// // Get single Product 
// func GetProduct(w http.ResponseWriter,r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	params :=mux.Vars(r)
// 	for _,item:=range products{
// 		if item.ID== params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&productmodel.Product{})
// }

// // add product
// func AddProduct(w http.ResponseWriter,r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	var product productmodel.Product
// 	_=json.NewDecoder(r.Body).Decode(&product)
// 	product.ID=strconv.Itoa(rand.Intn(1000000))
// 	products= append(products, product)
// 	json.NewEncoder(w).Encode(product)


// }

// //update product

// func UpdateProduct(w http.ResponseWriter,r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	params := mux.Vars(r)
// 	for index,item := range products {
// 		if item.ID == params["id"]{
// 		products= append(products[:index],products[index+1:]...)
// 		var product productmodel.Product
// 	_=json.NewDecoder(r.Body).Decode(&product)
// 	product.ID=params["id"]
// 	products= append(products, product)
// 	json.NewEncoder(w).Encode(product)
// 	return
// 		}
// 	}
	
// }

// //delete product

// func DeleteProduct(w http.ResponseWriter,r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	params := mux.Vars(r)
// 	for index,item := range products {
// 		if item.ID == params["id"]{
// 		products= append(products[:index],products[index+1:]...)
// 		break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(products)

// }