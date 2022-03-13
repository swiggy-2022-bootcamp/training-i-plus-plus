package ordermodel
import(
	"time"
"go.mongodb.org/mongo-driver/bson/primitive"
productmodel "github.com/bhatiachahat/mongoapi/models/productmodel"


)

type Order struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Products		[]productmodel.Product	`json:"products" validate:"required"`
	Ordered_at		time.Time				`json:"ordered_at"`
	Order_id			string				`json:"order_id"`
}