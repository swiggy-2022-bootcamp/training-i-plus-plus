package controller

import (
	"bhatiachahat/user-service/model"
//	"UserService/repository"
	"testing"

	"github.com/labstack/gommon/log"
)

func Read(username string) (model.SignupUser, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var regUser models.SignUp
	err := collectionAuth.FindOne(ctx, bson.M{"username": username}).Decode(&regUser)
	return regUser, err
}
func ValidateRequest(user model.SignupUser) bool {
	if user.Email = "" || user.Password = "" {
		log.Error("Either user email or password missing from sign in request")
		return true
	}

	return false
}
func TestSignUp(t *testing.T) {

	cases := []struct {
		name          string
		args          model.SignUp
		expectedError bool
	}{
		{
			name: "ValidSignUp",
			args: model.SignUpUser{
		
				Email:    "vkj@gmail.com",
		
				Password: "varun",
			},
			expectedError: false,
		},
		{
			name: "MissingEmailSignUp",
			args: model.SignUpUser{
			
				Email:    "",
		
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

			if err == true {
				return
			}

			var registerRepo repository.AuthRepository
			res, err2 := registerRepo.Read(c.args.Username)

			if err2 == nil {
				t.Errorf("Expected error but got %v", err)
			}

			if res.Username == c.args.Username {
				t.Errorf("Expected %v but got %v", c.args.Username, res.Username)
			}

		})
	}
}