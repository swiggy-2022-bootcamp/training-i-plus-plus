// package productmodel

// type Product struct {
// 	ID string `json:"id"`
//     Title string `json:"title"`
// 	Category string `json:"category"`
// }

package productmodel

import(
	"time"
"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Title		*string					    `json:"title" validate:"required"`
	Category		*string					`json:"category" validate:"required"`
	Price		*string				        `json:"price" validate:"required"`
	Seller			*string					`json:"seller" validate:"required"`
	Created_at		time.Time				`json:"created_at"`
	Updated_at		time.Time				`json:"updated_at"`
	Product_id			string				`json:"product_id"`
}