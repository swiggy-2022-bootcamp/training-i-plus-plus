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
//	helper "bhatiachahat/product-service/helper"
	"bhatiachahat/product-service/responses"
	model "bhatiachahat/product-service/model"
	database "bhatiachahat/product-service/db"
	kafkaservice "bhatiachahat/product-service/kafkaservice"
	"go.mongodb.org/mongo-driver/bson"
	 "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
const (
    topic         = "Products"
)
var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")
var validate = validator.New()

// AddProduct godoc
// @Summary To add a new product in the application
// @Description This request will adds a new product.
// @Tags Product
// @Schemes
// @Accept json
// @Produce json
// @Success	201  {object} 	model.Product
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /products [POST]
func AddProduct()gin.HandlerFunc{

	return func(c *gin.Context){
		// if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusOK,
		// 		responses.ProductResponse{Status: http.StatusBadRequest,  Message: "Unauthorized to perform this action"},
		// 	)
		// return
//	}
		
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
fmt.Println(product)
			p, err_ :=  kafkaservice.CreateProducer()
			if err_ != nil{
				c.JSON(http.StatusInternalServerError, gin.H{"error":"Error in kafka"})
				return
			}
		v,err:= kafkaservice.ProduceProduct(p,topic,product)

		
		if err !=nil {
		//	msg := fmt.Sprintf(product, "Product was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":err})
			return
		}
		// defer cancel()
		if v {	c.JSON(http.StatusCreated, responses.ProductResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": product}})}
	}

}

// GetProductByID godoc
// @Summary Get Product by ID.
// @Description View of a particular product.
// @Tags Product
// @Schemes
// @Param id path string true "Product id"
// @Accept json
// @Produce json
// @Success	200  {object} 	model.Product
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /products/{id} [GET]
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
		c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": product}})
	}
}

// DeleteProductByID godoc
// @Summary Delete product by ID.
// @Description Delete product.
// @Tags Product
// @Schemes
// @Param id path string true "Product id"
// @Accept json
// @Produce json
// @Success	200  {string} 	Product successfully deleted!
// @Failure	404  {number} 	http.http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /products/{id} [DELETE]
func DeleteProduct()gin.HandlerFunc{

	return func(c *gin.Context){
	// 	if err := helper.CheckUserType(c, "ADMIN"); err != nil {
	// 		c.JSON(http.StatusOK,
	// 			responses.ProductResponse{Status: http.StatusBadRequest,  Message: "Unauthorized to perform this action"},
	// 		)
	// 	return
	// }
		productId := c.Param("product_id")
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		
		_,err := productCollection.DeleteOne(ctx, bson.M{"product_id":productId})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,
			responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Product successfully deleted!"}},
		)

	}

}

// UpdateProductByID godoc
// @Summary Update product by ID.
// @Description Update details of a product.
// @Tags Product
// @Schemes
// @Param id path string true "product id"
// @Accept json
// @Produce json
// @Success	200  {object} 	model.Product
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /products/{id} [PUT]
func UpdateProduct() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusOK,
		// 		responses.ProductResponse{Status: http.StatusBadRequest,  Message: "Unauthorized to perform this action"},
		// 	)
		// 		// c.JSON(http.StatusOK,
		// 		// 	responses.ProductResponse{http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Unauthorized to perform this action", Data: map[string]interface{}{"data":err.Error()}}},
		// 		// )
		// 	return
		// }
        productId := c.Param("product_id")
        var product model.Product
        defer cancel()
	    // validate the request body
        if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

      
// fmt.Println(product.ImageUrl);
        update := bson.M{"title": product.Title, "category": product.Category, "Price": product.Price,"Seller":product.Seller,"Ratings":product.Ratings,"ImageUrl":product.ImageUrl}
		
	
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

		c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedproduct}})  }}

// GetAllProducts godoc
// @Summary Get all products.
// @Description Get all products.
// @Tags Product
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	model.Product
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /products [GET]
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
				c.JSON(http.StatusOK,
					responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": products}},
				)
				// c.JSON(http.StatusOK,
				// 	//responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": products}},
				// )
			}}