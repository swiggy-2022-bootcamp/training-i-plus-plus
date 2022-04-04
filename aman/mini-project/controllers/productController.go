package controllers

import (
	"aman-swiggy-mini-project/database"
	"aman-swiggy-mini-project/logger"
	"aman-swiggy-mini-project/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "product")
var validate = validator.New()

// CreateResource godoc
// @Summary Get All Products
// @Description Gets all the available products on the portal
// @Tags Portal
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /products [get]
func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		findOptions := options.Find()
		var results []models.Product
		cur, err := productCollection.Find(ctx, bson.M{}, findOptions)
		if err != nil {
			logger.Log.Println(err)
		}
		for cur.Next(ctx) {
			var elem models.Product
			err := cur.Decode(&elem)
			if err != nil {
				logger.Log.Println(err)
			}
			results = append(results, elem)
		}

		var aList []interface{}
		for _, elem := range results {
			aList = append(aList, gin.H{"Name": elem.Name, "Price": elem.Price, "Product ID": elem.Product_id})
		}

		if err := cur.Err(); err != nil {
			logger.Log.Println(err)
		}
		cur.Close(ctx)
		logger.Log.Println("All Products Requested")
		defer cancel()
		c.JSON(http.StatusOK, aList)
	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		productId := c.Param("product_id")
		var product models.Product

		err := productCollection.FindOne(ctx, bson.M{"product_id": productId}).Decode(&product)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while fetching the product item"})
		}
		logger.Log.Println("Product Requested")
		c.JSON(http.StatusOK, gin.H{"Name": product.Name, "Price": product.Price, "Stocks Remaining": product.Stock_units})
	}
}

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientType := c.Request.Header.Get("type")

		if clientType == "Buyer" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Users are not allowed to add products to the portal. You need to register as a Seller"})
			return
		} else if clientType == "Seller" {

			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			var product models.Product

			if err := c.BindJSON(&product); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			product.Seller_id = c.Request.Header.Get("id")

			validationErr := validate.Struct(product)
			if validationErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
				return
			}
			defer cancel()
			product.ID = primitive.NewObjectID()
			product.Product_id = product.ID.Hex()
			var num = toFixed(*product.Price, 2)
			product.Price = &num

			result, insertErr := productCollection.InsertOne(ctx, product)
			if insertErr != nil {
				msg := fmt.Sprintf("Product item was not created")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			logger.Log.Println("Product created")
			defer cancel()
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please register yourself as a seller to add your items to the portal."})
			return
		}
	}
}

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientType := c.Request.Header.Get("type")

		if clientType == "Buyer" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Users are not allowed to update products on the portal. You need to register as a Seller"})
			return
		} else if clientType == "Seller" {
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			sellerId := c.Request.Header.Get("id")
			productId := c.Param("product_id")
			var bMap map[string]interface{}

			body := c.Request.Body
			b, err := io.ReadAll(body)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			json.Unmarshal([]byte(b), &bMap)
			updatedQuantity := bMap["stock_units"]
			result, insertErr := productCollection.UpdateOne(ctx, bson.M{"seller_id": sellerId, "product_id": productId}, bson.D{{"$set",
				bson.D{
					{"stock_units", updatedQuantity},
				},
			}})

			if insertErr != nil {
				msg := fmt.Sprintf("order item was not updated")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			logger.Log.Println("Product updated")
			defer cancel()
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized Request. Please login."})
			return
		}
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
