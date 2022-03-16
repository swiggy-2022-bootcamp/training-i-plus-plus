package controllers

import (
    "context"
    "gin-mongo-api/config"
    "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "net/http"
    "time"
  
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var adminCollection *mongo.Collection = config.GetCollection(config.DB, "admins")
var avalidate = validator.New()

func CreateAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var admin models.Admin
        defer cancel()
      
        //validate the request body
        if err := c.BindJSON(&admin); err != nil {
            c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
      
        //use the validator library to validate required fields
        if validationErr := avalidate.Struct(&admin); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }
      
        newAdmin := models.Admin{
            Id:       primitive.NewObjectID(),
            Name:     admin.Name,
            Location: admin.Location,
            Title:    admin.Title,
        }
      
        result, err := adminCollection.InsertOne(ctx, newAdmin)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
      
        c.JSON(http.StatusCreated, responses.AdminResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}

func GetAAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        adminId := c.Param("adminId")
        var admin models.Admin
        defer cancel()
      
        objId, _ := primitive.ObjectIDFromHex(adminId)
      
        err := adminCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&admin)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
      
        c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admin}})
    }
}

func EditAAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        adminId := c.Param("adminId")
        var admin models.Admin
        defer cancel()
      
        objId, _ := primitive.ObjectIDFromHex(adminId)
      
        //validate the request body
        if err := c.BindJSON(&admin); err != nil {
            c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
      
        //use the validator library to validate required fields
        if validationErr := avalidate.Struct(&admin); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }
      
        update := bson.M{"name": admin.Name, "location": admin.Location, "title": admin.Title}
        result, err := adminCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
      
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
      
        //get updated admin details
        var updatedAdmin models.Admin
        if result.MatchedCount == 1 {
            err := adminCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedAdmin)
            if err != nil {
                c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                return
            }
        }
      
        c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedAdmin}})
    }
}

func DeleteAAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        adminId := c.Param("adminId")
        defer cancel()
      
        objId, _ := primitive.ObjectIDFromHex(adminId)
      
        result, err := adminCollection.DeleteOne(ctx, bson.M{"id": objId})
      
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
      
        if result.DeletedCount < 1 {
            c.JSON(http.StatusNotFound,
                responses.AdminResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Admin with specified ID not found!"}},
            )
            return
        }
      
        c.JSON(http.StatusOK,
            responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Admin successfully deleted!"}},
        )
    }
}

func GetAllAdmins() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var admins []models.Admin
        defer cancel()
      
        results, err := adminCollection.Find(ctx, bson.M{})
      
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
      
        //reading from the db in an optimal way
        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleAdmin models.Admin
            if err = results.Decode(&singleAdmin); err != nil {
                c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }
          
            admins = append(admins, singleAdmin)
        }
      
        c.JSON(http.StatusOK,
            responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admins}},
        )
    }
}