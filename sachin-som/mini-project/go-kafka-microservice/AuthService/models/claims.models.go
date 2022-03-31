package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
