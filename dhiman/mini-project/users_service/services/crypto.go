package services

import (
	"github.com/dhi13man/healthcare-app/users_service/configs"
	"golang.org/x/crypto/bcrypt"
)

// Hash Password using bcrypt
func HashPassword(password string) (string, error) {
	// Salt and hash the password using the bcrypt algorithm
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), configs.BCryptHashRounds())
	if err != nil {
		return password, err
	} else {
		return string(hashedPassword), nil
	}
}

// Compare Password with hashed password and returns true if they match
func ComparePassword(hashedPassword string, password string) bool {
	// Compare the password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
