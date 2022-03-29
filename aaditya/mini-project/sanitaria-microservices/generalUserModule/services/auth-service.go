package services

import (
	"errors"
	"log"
	"sanitaria-microservices/generalUserModule/configs"
	"time"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET_KEY string

func init(){
	SECRET_KEY = configs.EnvJWTSecretKey()
}

type SignedDetails struct{
	Name 		string	`json:"name"`
	EmailId 	string  `json:"emailId"`
	jwt.StandardClaims
}

func CreateToken(emailId, name string) (singedToken string, err error) {
	claims:=&SignedDetails{
		Name: name,
		EmailId: emailId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token ,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	
	if err != nil {
		log.Panic(err)
		return 
	}

	return token, err
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

func HashPassword(password string) string{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string)(bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err!= nil {
		msg = "Wrong password"
		check=false
	}
	return check, msg
}