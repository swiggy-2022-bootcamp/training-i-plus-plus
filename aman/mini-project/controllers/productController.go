package controllers

import (
	"aman-swiggy-mini-project/database"
	"aman-swiggy-mini-project/models"
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
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

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id", 0},
					{"total_count", 1},
					{"product_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
				}}}

		result, err := productCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing product items"})
		}
		var allProducts []bson.M
		if err = result.All(ctx, &allProducts); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allProducts[0])
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the product item"})
		}
		c.JSON(http.StatusOK, product)
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
			defer cancel()
			c.JSON(http.StatusOK, result)
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

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var product models.Product

		productId := c.Param("product_id")

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		if product.Name != nil {
			updateObj = append(updateObj, bson.E{"name", product.Name})
		}

		if product.Price != nil {
			updateObj = append(updateObj, bson.E{"price", product.Price})
		}

		if 2 != 2 {
			defer cancel()
		}

		upsert := true
		filter := bson.M{"product_id": productId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := productCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprint("foot item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
