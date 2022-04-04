
package controllers

import (
    "context"
    "inventoryService/configs"
    "inventoryService/models"
    "inventoryService/responses"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
var orderCollection *mongo.Collection = configs.GetCollection(configs.DB, "orders")
var validate = validator.New()

func AddProduct() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var product models.Product
        defer cancel()

        if err := c.BindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            configs.ErrorLogger.Println(err)
            return
        }

        if validationErr := validate.Struct(&product); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            configs.ErrorLogger.Println(validationErr)
            return
        }
		
        product.ProductId = primitive.NewObjectID()

        result, err := productCollection.InsertOne(ctx, product)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            configs.ErrorLogger.Println(err)
            return
        }

        c.JSON(http.StatusCreated, responses.ProductResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
        configs.InfoLogger.Println(result)
    }
}

func GetAProduct() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        productId := c.Param("productId")
        var product models.Product
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(productId)

        err := productCollection.FindOne(ctx, bson.M{"productid": objId}).Decode(&product)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            configs.ErrorLogger.Println(err)
            return
        }

        c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": product}})
    }
}

func EditAProduct() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        productId := c.Param("productId")
        var product models.Product
        defer cancel()
        objId, _ := primitive.ObjectIDFromHex(productId)

        //validate the request body
        if err := c.BindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            configs.ErrorLogger.Println(err)
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&product); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            configs.ErrorLogger.Println(validationErr)
            return
        }

        update := bson.M{"name": product.Name, "price": product.Price, "quantity": product.Quantity,"sellerid":product.SellerId}
        result, err := productCollection.UpdateOne(ctx, bson.M{"productid": objId}, bson.M{"$set": update})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            configs.ErrorLogger.Println(err)
            return
        }

        //get updated product details
        var updatedProduct models.Product
        if result.MatchedCount == 1 {
            err := productCollection.FindOne(ctx, bson.M{"productid": objId}).Decode(&updatedProduct)
            if err != nil {
                c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                configs.ErrorLogger.Println(err)
                return
            }
        }

        c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedProduct}})
    }
}

func DeleteAProduct() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        productId := c.Param("productId")
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(productId)

        result, err := productCollection.DeleteOne(ctx, bson.M{"productid": objId})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        if result.DeletedCount < 1 {
            c.JSON(http.StatusNotFound,
                responses.ProductResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Product with specified ID not found!"}},
            )
            return
        }

        c.JSON(http.StatusOK,
            responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Product successfully deleted!"}},
        )
    }
}

func GetAllProducts() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var products []models.Product
        defer cancel()

        results, err := productCollection.Find(ctx, bson.M{})

        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleProduct models.Product
            if err = results.Decode(&singleProduct); err != nil {
                c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }
          
            products = append(products, singleProduct)
        }

        c.JSON(http.StatusOK,
            responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": products}},
        )
    }
}

