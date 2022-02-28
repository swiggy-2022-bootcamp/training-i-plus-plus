package app

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var secret string = "Alohomora"

type Service interface {
	AuthenticateUser(string, string) (string, error)
	VerifyToken(string) (int, error)

	GetAllUsers() []User
	GetUserByUsername(string) User
	AddUser() (bool, error)
	DeleteUserByUsername(string)
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Username  string
	Password  string
	Phone     string
	Email     string
}

type service struct {
	db []User
}

func (s service) GetAllUsers() []User {
	return s.db
}

func (s service) GetUserByUsername(username string) (*User, error) {
	for _, val := range s.GetAllUsers() {
		if username == val.Username {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("no users found")
}

func (s service) AuthenticateUser(username string, password string) (string, error) {
	user, err := s.GetUserByUsername(username)
	if err != nil {
		panic("something went wrong")
	}

	if user.Password == password {
		return generateJWT(user.Id)
	}
	return "", fmt.Errorf("invalid credentials")
}

func (s service) ParseAuthToken(tokenString string) (float64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(float64), nil
	}
	return -1, err
}

func NewService() *service {
	return &service{
		db: []User{
			{
				FirstName: "Murtaza",
				LastName:  "Sadriwala",
				Phone:     "9090887112",
				Email:     "murtaza896@gmail.com",
				Password:  "Pass!23",
				Username:  "murtaza896",
			},
		},
	}
}

func generateJWT(id int) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	}).SignedString([]byte(secret + "sadaldjakl"))
}
