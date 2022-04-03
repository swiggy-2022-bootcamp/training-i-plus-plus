package utils

import (
	"log"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) []byte {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Panicf("Error hashing password")
	}
	return hashed
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
