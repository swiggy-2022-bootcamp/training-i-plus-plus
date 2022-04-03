package model

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User Model ..
type User struct {
	UserID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username"`
	Password     string             `json:"password" validate:"required"`
	Email        string             `json:"email" validate:"email,required"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refreshtoken"`
	Phone        string             `json:"string"`
	VerifiedUser bool               `json:"verified"`
	PremiumUser  bool               `json:"premiumuser"`
	UserType     string             `json:"usertype" validate:"required"`
}

//CoreUserData ...
type CoreUserData struct {
	Username     string             `json:"username" validate:"required" bson:"username"`
	Password     string             `json:"password" validate:"required" bson:"password"`
	UserType     string             `json:"userType" validate:"required"`
	VerifiedUser bool               `json:"verifiedUser"`
	PremiumUser  bool               `json:"premiumUser"`
	Email        string             `json:"email" validate:"email,required" bson:"email"`
	Name         string             `json:"name" bson:"name"`
	Phone        string             `json:"phone"`
	Location     string             `json:"location"`
	Age          string             `json:"age"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refreshToken"`
	UserID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

//Credentials - Used while logging in
type Credentials struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

//Claims ..
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
