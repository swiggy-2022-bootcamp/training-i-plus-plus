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
	"bhatiachahat/order-service/responses"
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

// Place Order godoc
// @Summary To place order in the application
// @Description This request will allow user to place order.
// @Tags Order
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
	// fmt.Println(newusercart)
		//	var usercart = model.ProductUser
			//user.UserCart= newusercart
			update := bson.M{"user_cart": newusercart}

	_, err :=userCollection.UpdateOne(ctx,bson.M{"user_id":userId},bson.M{"$set": update})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
            }
		//	fmt.Println(result)
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

		
	
		_, insertErr := orderCollection.InsertOne(ctx,order)
		// _, insertErr := productCollection.InsertOne(ctx, product)
		if insertErr !=nil {
		//	msg := fmt.Sprintf(product, "Product was not created")
		c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		// defer cancel()
		c.JSON(http.StatusBadRequest,
			responses.OrderResponse{Status:http.StatusBadRequest, Message: "success", Data: map[string]interface{}{"data": order}},
		)
	}else{
		{c.JSON(http.StatusOK,
			responses.OrderResponse{Status: http.StatusOK, Message: "Add products to cart", Data: map[string]interface{}{}},
		)}
	}
	}

}
