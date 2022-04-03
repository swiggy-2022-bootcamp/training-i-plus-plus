package services

import (
	"errors"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET_KEY string

func init(){
	SECRET_KEY = "secret5"
}

type SignedDetails struct{
	Name 		string	`json:"name"`
	EmailId 	string  `json:"emailId"`
	jwt.StandardClaims
}


func ValidateToken(tokenReceived string) (bool,error) {
	token,err := jwt.ParseWithClaims(
		tokenReceived,
		&SignedDetails{},
		func(token *jwt.Token)(interface{}, error){
			return []byte(SECRET_KEY), nil
		},
	)

	if err!=nil{
		return false,err
	}

	claims, ok:= token.Claims.(*SignedDetails)
	if !ok{
		err_ := errors.New("Token is invalid")
		return false,err_
	}

	if claims.ExpiresAt < time.Now().Local().Unix(){
		err_ := errors.New("Your session has expired. Please re-login")
		return false, err_
	}

	return true, nil
}
