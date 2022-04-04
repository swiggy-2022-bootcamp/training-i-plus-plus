package controllers

import (
	"UserService/models"
	"testing"

	"github.com/labstack/gommon/log"
)

func ValidateRequest(user models.SignUp) bool {
	if user.Email == "" || user.Password == "" || user.Username == "" {
		log.Error("Either user email or password missing from sign in request")
		return true
	}

	return false
}

func TestSignIn(t *testing.T) {

	cases := []struct {
		name          string
		args          models.SignUp
		expectedError bool
	}{
		{
			name: "ValidSignUp",
			args: models.SignUp{
				Username: "varun",
				Email:    "vkj@gmail.com",
				TypeOf:   "user",
				Password: "varun",
			},
			expectedError: false,
		},
		{
			name: "MissingUsernameSignUp",
			args: models.SignUp{
				Username: "",
				Email:    "am@gmail.com",
				TypeOf:   "student",
				Password: "aditi",
			},
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := ValidateRequest(c.args)
			if c.expectedError && err == false {
				t.Error("Expected an error but didn't get one")
			}
			if !c.expectedError && err != false {
				t.Errorf("Expected no error but got %v", err)
			}
		})
	}
}
