
package products

import (
	"context"
	"fmt"
	"log"
	//"strconv"

	"strconv"
	"net/http"
	"time"

	helper "github.com/bhatiachahat/mongoapi/helper"
	cartmodel "github.com/bhatiachahat/mongoapi/models/cartmodel"
	//productmodel "github.com/bhatiachahat/mongoapi/models/productmodel"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	//	productcontroller "github.com/bhatiachahat/mongoapi/controllers/products"

	database "github.com/bhatiachahat/mongoapi/db"

	
	//"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "cart")
var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

var validate = validator.New()

func AddProductToCart()gin.HandlerFunc{

	return func(c *gin.Context){
	
	
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var cart cartmodel.Cart

		if err := c.BindJSON(&cart); err != nil {
		
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(cart)
		if validationErr != nil {
			
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return
		}

		// count, err := cartCollection.CountDocuments(ctx, bson.M{"product_id":cart.Product_id})
		// defer cancel()
		// if err != nil {
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the title"})
		// }

		// password := HashPassword(*user.Password)
		// user.Password = &password

		// count, err = productCollection.CountDocumets(ctx, bson.M{"category":product.Category})
		// defer cancel()
		// if err!= nil {
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the category"})
		// }

	// 	if count >0{
	// 	//	c.JSON(http.StatusInternalServerError, gin.H{"error":"This product already exists"})
	// 	quantity,_ :=  strconv.Atoi(cart.Quantity)
	//     price , _ := strconv.Atoi(cart.Price)
	// 	// fmt.Printf("%T\n",quantity)
	// 	// fmt.Printf("%T\n",price)
	// 	total:= price*quantity
		
	//    cart.SubTotal= strconv.Itoa(total) 
	// 	}
    //   cart.Quantity= cart.Quantity
	    quantity,_ :=  strconv.Atoi(cart.Quantity)
	    price , _ := strconv.Atoi(cart.Price)
		// fmt.Printf("%T\n",quantity)
		// fmt.Printf("%T\n",price)
		total:= price*quantity
		
	   cart.SubTotal= strconv.Itoa(total) 
		cart.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		//cart.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		cart.ID = primitive.NewObjectID()
		cart.Cart_id = cart.ID.Hex()
		//token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
	//	user.Token = &token
	//	user.Refresh_token = &refreshToken

	_	, insertErr := cartCollection.InsertOne(ctx, cart)
		if insertErr !=nil {
			msg := fmt.Sprintf("Product was not added")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, cart)
	}

}
func GetAllCartItems()gin.HandlerFunc{
	return func(c *gin.Context){
		//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// filter := bson.D{{"pop", bson.D{{"$lte", 500}}}}
		// cursor, err := coll.Find(context.TODO(), filter)
		// if err != nil {
		// 	panic(err)
		// }
		// body := c.Request.Body()
		// value,err:= ioutil.ReadAll(body)
		// userId:= string(value)
		userId := c.Param("user_id")
		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var cart []cartmodel.Cart
		// products,err := cartCollection.Find(ctx, bson.M{"user_id":userId})
		// defer cancel()
		cursor,err := cartCollection.Find(ctx, bson.M{"user_id":userId})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
if err = cursor.All(context.TODO(), &cart); err != nil {
  log.Fatal(err)
}
//fmt.Printf("Found multiple documents: %+v\n", cart)
//var products []productmodel.Product
// for i:=0;i<len(cart);i++ {
// 	product_id:=cart[i].Product_id
// 	var product productmodel.Product
// 		err := productCollection.FindOne(ctx, bson.M{"product_id":product_id}).Decode(&product)
// 		defer cancel()
// 		if err != nil{
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 	//	c.JSON(http.StatusOK, product)
// 	//product:= productcontroller.GetProduct(cart[i].product_id)
// 	products = append(products,product)
// }

		c.JSON(http.StatusOK, cart)
// if err != nil { log.Fatal(err) }
// //defer cur.Close(context.Background())
// for cur.Next(ctx.Background()) {
//   // To decode into a struct, use cursor.Decode()
//   result := struct{
//     Foo string
//     Bar int32
//   }{}
//   res,err := cur.Decode(&result)
//   if err != nil { log.Fatal(err) }}
}
}
func DeleteFromCart()gin.HandlerFunc{

	return func(c *gin.Context){
		cartId := c.Param("cart_id")
		
        if err := helper.MatchUserTypeToUid(c, cartId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// var product productmodel.Product
		res,err := cartCollection.DeleteOne(ctx, bson.M{"cart_id":cartId})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,res)


		
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




