package controller
import (
	"context"
	 "fmt"
	"log"
	"math/rand"
	//"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	//helper "github.com/bhatiachahat/mongoapi/helper"
	model "bhatiachahat/order-service/model"
	database "bhatiachahat/order-service/db"
	kafkaservice "bhatiachahat/order-service/kafkaservice"
	"go.mongodb.org/mongo-driver/bson"
	 "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
const (
    topic         = "Orders"
)
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

var validate = validator.New()

// AddProduct godoc
// @Summary To add a new product in the application
// @Description This request will adds a new product.
// @Tags Product
// @Schemes
// @Accept json
// @Produce json
// @Success	201  {object} 	model.Order
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /orders/place_order/:user_id [POST]
func PlaceOrder()gin.HandlerFunc{

	return func(c *gin.Context){
	
		userId := c.Param("user_id")
     //  fmt.Println(userId)
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var order model.Order
		var user model.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		if err!=nil{
			fmt.Println(err)
		}
				defer cancel()
					usercart := []model.ProductUser{}
					usercart=user.UserCart
					if(len(usercart)>=1){

				
				order.Order_Cart=usercart
			//	fmt.Println(user)
	//	fmt.Println(user.UserCart)
	newusercart := []model.ProductUser{}
	fmt.Println(newusercart)
		//	var usercart = model.ProductUser
			//user.UserCart= newusercart
			update := bson.M{"user_cart": newusercart}

			result, err :=userCollection.UpdateOne(ctx,bson.M{"user_id":userId},bson.M{"$set": update})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
            }
			fmt.Println(result)
			order.Payment_Method="DIGITAL"
			num:=rand.Intn(200)
			if num%2==0{ order.Payment_Method="COD"}
			
		order.Ordered_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.User_id	=userId
		order.ID = primitive.NewObjectID()
		count, err := orderCollection.CountDocuments(ctx, bson.M{"Id":order.ID})
		defer cancel()
		if err!= nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while placing the order"})
		}
		//var order ordermodel.Order
// cart := model.{}
// result.Decode(user)
		// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// var product model.Product

		// if err := c.BindJSON(&product); err != nil {
			
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// validationErr := validate.Struct(product)
		// if validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
		// 	return
		// }

		// count, err := productCollection.CountDocuments(ctx, bson.M{"title":product.Title})
		// defer cancel()
		// if err != nil {
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the title"})
		// }
		if count >0{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Order has been already placed"})
			return
		}
	
		
			p, err_ :=  kafkaservice.CreateProducer()
			if err_ != nil{
				c.JSON(http.StatusInternalServerError, gin.H{"error":"Error in kafka"})
				return
			}
            kafkaservice.ProduceOrder(p,topic,order)

		// product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// product.ID = primitive.NewObjectID()
		// product.Product_id = product.ID.Hex()
	
		_, insertErr := orderCollection.InsertOne(ctx,order)
		// _, insertErr := productCollection.InsertOne(ctx, product)
		if insertErr !=nil {
		//	msg := fmt.Sprintf(product, "Product was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":err})
			return
		}
		// defer cancel()
		 {c.JSON(http.StatusOK, order)}
	}else{
		{c.JSON(http.StatusInternalServerError, gin.H{"error":"Add products to cart"})}
	}
	}

}

// // GetProductByID godoc
// // @Summary Get Product by ID.
// // @Description View of a particular product.
// // @Tags Product
// // @Schemes
// // @Param id path string true "Product id"
// // @Accept json
// // @Produce json
// // @Success	200  {object} 	model.Order
// // @Failure	500  {number} 	http.StatusInternalServerError
// // @Security Bearer Token
// // @Router /products/{id} [GET]
// func GetProduct() gin.HandlerFunc{
// 	return func(c *gin.Context){
// 		productId := c.Param("product_id")

		
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 		var product model.Product
// 		err := productCollection.FindOne(ctx, bson.M{"product_id":productId}).Decode(&product)
// 		defer cancel()
// 		if err != nil{
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, product)
// 	}
// }

// // DeleteProductByID godoc
// // @Summary Delete product by ID.
// // @Description Delete product.
// // @Tags Product
// // @Schemes
// // @Param id path string true "Product id"
// // @Accept json
// // @Produce json
// // @Success	200  {string} 	Product successfully deleted!
// // @Failure	404  {number} 	http.http.StatusNotFound
// // @Failure	500  {number} 	http.StatusInternalServerError
// // @Security Bearer Token
// // @Router /products/{id} [DELETE]
// func DeleteProduct()gin.HandlerFunc{

// 	return func(c *gin.Context){
// 		productId := c.Param("product_id")
		
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		
// 		res,err := productCollection.DeleteOne(ctx, bson.M{"product_id":productId})
// 		defer cancel()
// 		if err != nil{
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK,res)

// 	}

// }

// // UpdateProductByID godoc
// // @Summary Update product by ID.
// // @Description Update details of a product.
// // @Tags Product
// // @Schemes
// // @Param id path string true "product id"
// // @Accept json
// // @Produce json
// // @Success	200  {object} 	model.Product
// // @Failure	400  {number} 	http.http.StatusBadRequest
// // @Failure	500  {number} 	http.StatusInternalServerError
// // @Security Bearer Token
// // @Router /products/{id} [PUT]
// func UpdateProduct() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//         productId := c.Param("product_id")
//         var product model.Product
//         defer cancel()
// 	    // validate the request body
//         if err := c.BindJSON(&product); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//             return
//         }

      
// // fmt.Println(product.ImageUrl);
//         update := bson.M{"title": product.Title, "category": product.Category, "Price": product.Price,"Seller":product.Seller,"Ratings":product.Ratings,"ImageUrl":product.ImageUrl}
		
	
//         result, err := productCollection.UpdateOne(ctx,bson.M{"product_id":productId}, bson.M{"$set": update})
	
//         if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
//         }

//         //get updated user details
//         var updatedproduct model.Product
//         if result.MatchedCount == 1 {
//             err := productCollection.FindOne(ctx, bson.M{"product_id":productId}).Decode(&updatedproduct)
//             if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 				return
//             }
//         }

// 		c.JSON(http.StatusOK, updatedproduct)    }}

// // GetAllProducts godoc
// // @Summary Get all products.
// // @Description Get all products.
// // @Tags Product
// // @Schemes
// // @Accept json
// // @Produce json
// // @Success	200  {array} 	model.Product
// // @Failure	500  {number} 	http.StatusInternalServerError
// // @Security Bearer Token
// // @Router /products [GET]
// func GetAllProducts() gin.HandlerFunc {
// 			return func(c *gin.Context) {
// 				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 				var products []model.Product
// 				defer cancel()
		
// 				results, err := productCollection.Find(ctx, bson.M{})
		
// 				if err != nil {
// 					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 				//	c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 					return
// 				}
		
// 				//reading from the db in an optimal way
// 				defer results.Close(ctx)
// 				for results.Next(ctx) {
// 					var singleProduct model.Product
// 					if err = results.Decode(&singleProduct); err != nil {
// 						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 						//c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 					}
				  
// 					products = append(products, singleProduct)
// 				}
// 				c.JSON(http.StatusOK, products)
// 				// c.JSON(http.StatusOK,
// 				// 	//responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": products}},
// 				// )
// 			}}