package productcontroller
import (
	//"context"
	// "fmt"
	//"log"
	//"strconv"
///	"net/http"
//	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	//helper "github.com/bhatiachahat/mongoapi/helper"
//	model "bhatiachahat/payment-service/model"
//	database "bhatiachahat/product-service/db"
//	kafkaservice "bhatiachahat/product-service/kafkaservice"
//	"go.mongodb.org/mongo-driver/bson"
//	 "go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
)
const (
    topic         = "Cart"
)
//var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "cart")
var validate = validator.New()

// AddProduct godoc
// @Summary To add a new product in the application
// @Description This request will adds a new product.
// @Tags Product
// @Schemes
// @Accept json
// @Produce json
// @Success	201  {string} 	Payment Successful
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /payment/:orderId [POST]
func DoPayment()gin.HandlerFunc{

	return func(c *gin.Context){
	// 	inventoryId, _ := primitive.ObjectIDFromHex(gctx.Param("inventoryId"))
	// var prouduct models.Product
	// if err := gctx.ShouldBindJSON(&prouduct); err != nil {
	// 	gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// if err := ic.InventoryService.AddProduct(inventoryId, &prouduct); err != nil {
	// 	gctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	// 	return
	// }
	// gctx.JSON(http.StatusCreated, gin.H{"message": "Product Added to Inventory."})
	
	    // cartId := primitive.ObjectIDFromHex(c.Param("cart_id"))
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
		// filterCart := bson.D{bson.E{Key: "_id", Value: cartId}}
	    // pushProuductID := bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "products", Value: product}}}}
		// result, err := cartCollection.UpdateOne(ctx, filterCart, pushProuductID)
		// defer cancel()
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":err})
		// }
		// if result.MatchedCount != 1 {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"Wrong Cart Id"})
		// }
		// if result.ModifiedCount != 1 {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"Product not added in cart Id"})
		// }

		// count, err := productCollection.CountDocuments(ctx, bson.M{"title":product.Title})
		// defer cancel()
		// if err != nil {
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the title"})
		// }
		// if count >0{
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"This product already exists"})
		// 	return
		// }
	

		// 	p, err_ :=  kafkaservice.CreateProducer()
		// 	if err_ != nil{
		// 		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error in kafka"})
		// 		return
		// 	}
		// v,err:= kafkaservice.ProduceProduct(p,topic,product)

		// // product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// // product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// // product.ID = primitive.NewObjectID()
		// // product.Product_id = product.ID.Hex()
	

		// // _, insertErr := productCollection.InsertOne(ctx, product)
		// if err !=nil {
		// //	msg := fmt.Sprintf(product, "Product was not created")
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":err})
		// 	return
		// }
		// // defer cancel()
		// if v {c.JSON(http.StatusOK, product)}
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
// // @Success	200  {object} 	model.Product
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