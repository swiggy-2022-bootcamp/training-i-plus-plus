package controller
import (
"context"
	 "fmt"
//	"log"
	//"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"bhatiachahat/track-stream-service/responses"
	helper "bhatiachahat/track-stream-service/helper"
//	 model "bhatiachahat/track-stream-service/model"
	 database "bhatiachahat/track-stream-service/db"
	// kafkaservice "bhatiachahat/product-service/kafkaservice"
	"go.mongodb.org/mongo-driver/bson"
 //"go.mongodb.org/mongo-driver/bson/primitive"
	 "go.mongodb.org/mongo-driver/mongo"
)
// const (
//     topic         = "Orders"
// )
var trackstreamCollection *mongo.Collection = database.OpenCollection(database.Client, "trackstream")

//var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "cart")
var validate = validator.New()


// GetTrackkStreamData godoc
// @Summary Get the analytics of application(This usecase tracks the count of different modes of payments)
// @Description This request will give analytics data of different types of payment.
// @Tags TrackStream
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	DIGITAL || COD 
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /getTrackingData [GET]
func GetTrackingData()gin.HandlerFunc{

	return func(c *gin.Context){
		if err := helper.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusOK,
				responses.TrackStreamResponse{Status: http.StatusBadRequest,  Message: "Unauthorized to perform this action"},
			)
		return
	}
	
		 var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		 defer cancel()
		 //var trackstreamobj= model.TrackStream
		 count1, err1 := trackstreamCollection.CountDocuments(ctx,  bson.M{"payment_type":"DIGITAL"})
		 count2, err2 := trackstreamCollection.CountDocuments(ctx,  bson.M{"payment_type":"COD"})
		 if err1 != nil || err2 !=nil{
			
			c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
			return
		}
		 fmt.Println(count1)
		 fmt.Println(count2)

	//	if count1>count2 {c.JSON(http.StatusOK,"DIGITAL")}else{c.JSON(http.StatusOK,"COD")}
		c.JSON(http.StatusOK,
			responses.TrackStreamResponse{Status: http.StatusOK, Message: "success", Digital:count1, COD:count2},
		)
		
	}

}
