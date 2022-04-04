package productcontroller
import (
	//"context"
	// "fmt"
	//"log"
	//"strconv"
	"net/http"
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

// DoPayment godoc
// @Summary To make payment for the order in the application
// @Description This request will allow customer to make payment.
// @Tags Payment
// @Accept json
// @Produce json
// @Success	201  {string} 	Payment Successful
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /payment/:orderId [POST]
func DoPayment()gin.HandlerFunc{

	return func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message":"Payment Successfull"})
	}

}
