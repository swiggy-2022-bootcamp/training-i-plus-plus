package model

import(
	"time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type TrackStream struct {
	ID				primitive.ObjectID		    `bson:"_id"`

	PaymentType         string                       `json:"payment_type"`
}

// type Order struct{
// 	ID				primitive.ObjectID		`bson:"_id"`
// 	Order_Cart		    []ProductUser 					`json:"order_cart" bson:"order_cart"`
// //	Category		string					`json:"category" validate:"required"`
// //	Price		    uint			        `json:"price" bson:"price"`
// 	// Seller			string					`json:"seller" validate:"required"`
// 	Ordered_at		time.Time				`json:"ordered_at"`
// 	// Updated_at		time.Time				`json:"updated_at"`
// 	User_id		string				    `json:"user_id"`  
// //	Ratings         uint                    `json:"ratings"      bson:"ratings"`
// //	ImageUrl        string                  `json:"image_url"    bson:"image_url"`
// //	Products        Products[]
//   // Payment_Method   model.Payment                   `json:"payment"    bson:"payment"`

// }

// type Payment struct{
// 	ID            primitive.ObjectID   `json:"_id"                bson:"_id"                 validate:"required"`
// 	Digital       bool                 `json:"digital"           bson:"digital"      validate:"required"`
// 	COD           bool                 `json:"cod"               bson:"cod"            validate:"required"`
// }
// type ProductUser struct{
// 	ID				primitive.ObjectID		`bson:"_id"`
// 	Title		    string					`json:"title" validate:"required"`
// 	Category		string					`json:"category" validate:"required"`
// 	Price		    uint			        `json:"price" validate:"required"`
// 	Seller			string					`json:"seller" validate:"required"`
// 	Created_at		time.Time				`json:"created_at"`
// 	Updated_at		time.Time				`json:"updated_at"`
// 	Product_id		string				    `json:"product_id"`  
// 	Ratings         uint                    `json:"ratings"      bson:"ratings"`
// 	ImageUrl        string                  `json:"image_url"    bson:"image_url"`
// }
type ProductUser struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Title		    string					`json:"title" validate:"required"`
	Category		string					`json:"category" validate:"required"`
	Price		    uint			        `json:"price" validate:"required"`
	Seller			string					`json:"seller" validate:"required"`
	Created_at		time.Time				`json:"created_at"`
	Updated_at		time.Time				`json:"updated_at"`
	Product_id		string				    `json:"product_id"`  
	Ratings         uint                    `json:"ratings"      bson:"ratings"`
	ImageUrl        string                  `json:"image_url"    bson:"image_url"`
}
type Order struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Order_Cart		    []ProductUser 					`json:"order_cart" bson:"order_cart"`
//	Category		string					`json:"category" validate:"required"`
//	Price		    uint			        `json:"price" bson:"price"`
	// Seller			string					`json:"seller" validate:"required"`
	Ordered_at		time.Time				`json:"ordered_at"`
	// Updated_at		time.Time				`json:"updated_at"`
	User_id		string				    `json:"user_id"`  
//	Ratings         uint                    `json:"ratings"      bson:"ratings"`
//	ImageUrl        string                  `json:"image_url"    bson:"image_url"`
//	Products        Products[]
   Payment_Method  string                  `json:"payment"    bson:"payment" eq=COD|eq=DIGITAL`
}