package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UserCredential struct {
	Username string
	Password string
}

var userCredentials = []UserCredential{}

func setUp() {

	fileContent, err := ioutil.ReadFile("sample.json")
	handleErr(err)

	if err := json.Unmarshal(fileContent, &userCredentials); err != nil {
		handleErr(err)
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic("Something went wrong")
	}
}

func (u UserCredential) authenticateUser() bool {
	for _, userCredential := range userCredentials {
		if userCredential.Username == u.Username && userCredential.Password == u.Password {
			return true
		}
	}
	return false
}

func (u UserCredential) addUserCredentialToFile() {
	userCredentials = append(userCredentials, u)
	result, err := json.Marshal(userCredentials)
	handleErr(err)
	os.WriteFile("sample.json", result, 0666)
}

func main() {
	setUp()
	myUser := UserCredential{"murtaza", "123Pass"}
	fmt.Println(myUser.authenticateUser())
	myUser.addUserCredentialToFile()

	myOtherUser := UserCredential{"john@gmail.com", "pass123"}
	fmt.Println(myOtherUser.authenticateUser())
}
