package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName       string             `json:"firstname" bson:"firstname"`
	LastName        string             `json:"lastname" bson:"lastname"`
	Email           string             `json:"email" bson:"email"`
	Password        string             `json:"password" bson:"password"`
	Phone           string             `json:"phone" bson:"phone"`
	Token           string             `json:"token" bson:"token"`
	Refresh_token   string             `json:"refresh_token" bson:"refresh_token"`
	Created_At      time.Time          `json:"created_at" bson:"created_at"`
	Updated_At      time.Time          `json:"updated_at" bson:"updated_at"`
	User_Id         string             `json:"user_id" bson:"user_id"`
	UserCart        []ProductUser      `json:"user_cart" bson:"user_cart"`
	Address_Details []Address          `json:"address_details" bson:"address_details"`
	Order_Status    []Order            `json:"order_status" bson:"order_status"`
}
