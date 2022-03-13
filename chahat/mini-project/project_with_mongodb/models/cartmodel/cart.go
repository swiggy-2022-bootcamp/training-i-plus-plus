package cartmodel
import(
	"time"
"go.mongodb.org/mongo-driver/bson/primitive"
//productmodel "github.com/bhatiachahat/mongoapi/models/productmodel"


)

type Cart struct{
	ID				primitive.ObjectID	`bson:"_id"`
	Product_id 	string					`json:"product_id" validate:"required"`
	Title		*string					`json:"title" validate:"required"`
	Category		*string				`json:"category" validate:"required"`
	Price		   string				`json:"price" validate:"required"`
	Seller			*string				`json:"seller" validate:"required"`
	SubTotal       string                `json:"subtotal"`
	User_id	        string				`json:"user_id" validate:"required"`
	Created_at		time.Time			 `json:"created_at"`
	Quantity          string             `json:"quantity" validate:"required,min=1,max=10"`
	Cart_id			string				 `json:"cart_id"`
}