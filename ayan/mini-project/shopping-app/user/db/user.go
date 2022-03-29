package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Name      string             `bson:"name"`
	Address   string             `bson:"address"`
	Zipcode   int32              `bson:"zipcode"`
	MobileNo  string             `bson:"mobile_no"`
	Role      string             `bson:"role"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewUser(email string, password string, name string, address string, zipcode int32, mobileNo string, role string) *User {
	return &User{
		Email:    email,
		Password: password,
		Name:     name,
		Address:  address,
		Zipcode:  zipcode,
		MobileNo: mobileNo,
		Role:     role,
	}
}

func (u *User) SetId(id primitive.ObjectID) {
	u.Id = id
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}
