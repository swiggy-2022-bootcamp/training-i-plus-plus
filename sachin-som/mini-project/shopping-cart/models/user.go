package models

import "time"

type User struct {
	ID           int           `json:"_id"           bson:"_id"`
	FirstName    *string       `json:"first_name"    bson:"first_name"    validate:"required, min=2, max=50"`
	LastName     *string       `json:"last_name"     bson:"last_name"     validate:"required, min=2, max=50"`
	Email        *string       `json:"email"         bson:"email"         validate:"email required"`
	Phone        *string       `json:"phone"         bson:"phone"         validate:"required"`
	Password     *string       `json:"password"      bson:"password"      validate:"required"`
	Token        *string       `json:"token"         bson:"token"`
	RefreshToken *string       `json:"refresh_token" bson:"refresh_token"`
	CreatedAt    time.Time     `json:"created_at"    bson:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"    bson:"updated_at"`
	UserID       *string       `json:"user_id"       bson:"user_id"`
	Addresses    []Address     `json:"addresses"     bson:"addresses"`
	UserCart     []ProductUser `json:"user_cart"     bson:"user_cart"`
	UserOrders   []Order       `json:"orders"        bson:"orders"`
}
