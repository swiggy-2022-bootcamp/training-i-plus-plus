package productcontroller
import (
	"context"
	"fmt"
	"log"
	//"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	//helper "github.com/bhatiachahat/mongoapi/helper"
	model "github.com/bhatiachahat/sample-gin-project-with-mongodb/model"
	database "github.com/bhatiachahat/sample-gin-project-with-mongodb/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")
var validate = validator.New()
func AddProduct()gin.HandlerFunc{

	return func(c *gin.Context){
	
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var product model.Product

		if err := c.BindJSON(&product); err != nil {
			
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(product)
		if validationErr != nil {
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
			return
		}

		product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.ID = primitive.NewObjectID()
		product.Product_id = product.ID.Hex()
	

		_, insertErr := productCollection.InsertOne(ctx, product)
		if insertErr !=nil {
			msg := fmt.Sprintf("Product was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, product)
	}

}

func GetProduct() gin.HandlerFunc{
	return func(c *gin.Context){
		productId := c.Param("product_id")

		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var product model.Product
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
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		
		res,err := productCollection.DeleteOne(ctx, bson.M{"product_id":productId})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,res)

	}

}


func UpdateProduct() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        productId := c.Param("product_id")
        var product model.Product
        defer cancel()
	

       // validate the request body
        if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

      

        update := bson.M{"title": product.Title, "category": product.Category, "Price": product.Price,"seller":product.Seller}
		
	
        result, err := productCollection.UpdateOne(ctx,bson.M{"product_id":productId}, bson.M{"$set": update})
	
        if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
        }

        //get updated user details
        var updatedproduct model.Product
        if result.MatchedCount == 1 {
            err := productCollection.FindOne(ctx, bson.M{"product_id":productId}).Decode(&updatedproduct)
            if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
            }
        }

		c.JSON(http.StatusOK, updatedproduct)    }}


func GetAllProducts() gin.HandlerFunc {
			return func(c *gin.Context) {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				var products []model.Product
				defer cancel()
		
				results, err := productCollection.Find(ctx, bson.M{})
		
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				//	c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
				}
		
				//reading from the db in an optimal way
				defer results.Close(ctx)
				for results.Next(ctx) {
					var singleProduct model.Product
					if err = results.Decode(&singleProduct); err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						//c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					}
				  
					products = append(products, singleProduct)
				}
				c.JSON(http.StatusOK, products)
				// c.JSON(http.StatusOK,
				// 	//responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": products}},
				// )
			}}