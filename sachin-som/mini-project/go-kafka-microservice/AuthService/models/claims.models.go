package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Email string `json:"email" bson:"email"`
	jwt.StandardClaims
}
