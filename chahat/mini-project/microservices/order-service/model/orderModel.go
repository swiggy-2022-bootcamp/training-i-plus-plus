package model

import(
	"time"
    "go.mongodb.org/mongo-driver/bson/primitive"
	//model "bhatiachahat/payment-service/model"
)
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
type User struct{
	ID       primitive.ObjectID `json:"_id"           bson:"_id"`
	Firstname *string             `json:"first_name"      bson:"first_name"      validate:"required"`
	Lastname *string             `json:"last_name"      bson:"last_name"      validate:"required"`
	Email    *string             `json:"email"         bson:"email"         validate:" required"`
	Phone    *string             `json:"phone"         bson:"phone"         validate:"required"`
	Password *string             `json:"password"      bson:"password"      validate:"required"`
	Token    *string              `json:"token"        bson:"token"`
	Refresh_token *string           `json:"refresh_token"  bson:"refresh_token"`
	Created_at		time.Time				`json:"created_at"`
	Updated_at		time.Time				`json:"updated_at"`
	User_id       string             `json:"user_id"`
	Usertype *string				`json:"usertype"      bson:"usertype"      validate:"required,eq=ADMIN|eq=USER"`
    UserCart  []ProductUser      `json:"user_cart"    bson:"user_cart"`
}