package app

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var secret string = "Alohomora"

type Service interface {
	AuthenticateUser(string, string) (string, error)
	ParseAuthToken(string) (int, role, error)

	GetAllUsers() []User
	GetUserByUsername(string) User
	GetUserByEmail(string) User
	AddUser(User) (bool, error)
	DeleteUserByUsername(string) (bool, error)

	SetEmail(string)
	SetUsername(string)
	SetPassword(string)
	SetPhone(string)
	SetFirstName(string)
	SetLastName(string)
}

type role int

const (
	Admin    role = iota + 1 // EnumIndex = 1
	Seller                   // EnumIndex = 2
	Customer                 // EnumIndex = 3
)

func (r role) String() string {
	return [...]string{"admin", "seller", "customer"}[r-1]
}

func (r role) EnumIndex() int {
	return int(r)
}

func GetEnumByIndex(idx int) role {
	switch idx {
	case 1:
		return Admin
	case 2:
		return Seller
	case 3:
		return Customer
	default:
		panic("invalid enum index")
	}
}

type User struct {
	id        int
	firstName string
	lastName  string
	username  string
	password  string
	phone     string
	email     string
	role      role
}

//---- getter and setter --------
func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u User) GetUsername() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u User) GetPhone() string {
	return u.phone
}

func (u *User) SetPhone(phone string) {
	u.phone = phone
}

//-------------------------------------------

type service struct {
	db []User
}

func (s service) GetAllUsers() []User {
	return s.db
}

func (s service) GetUserByUsername(username string) (*User, error) {
	for _, val := range s.GetAllUsers() {
		if username == val.username {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("no users found")
}

func (s service) GetUserByEmail(email string) (*User, error) {
	for _, val := range s.GetAllUsers() {
		if email == val.email {
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

	if user.password == password {
		return generateJWT(user.id, user.role)
	}
	return "", fmt.Errorf("invalid credentials")
}

func (s service) ParseAuthToken(tokenString string) (int, role, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["id"].(float64)), GetEnumByIndex(int(claims["role"].(float64))), nil
	}
	return -1, -1, err
}

func (s *service) AddUser(u User) (bool, error) {
	s.db = append(s.db, u)
	return true, nil
}

func (s *service) DeleteUserByUsername(username string) (bool, error) {
	var pos int = -1
	users := s.GetAllUsers()
	for i, v := range users {
		if v.username == username {
			pos = i
			break
		}
	}
	s.db = append(users[0:pos], users[pos+1:]...)
	return true, nil
}

func NewService() *service {
	return &service{
		db: []User{
			{
				firstName: "Murtaza",
				lastName:  "Sadriwala",
				phone:     "9090887112",
				email:     "murtaza896@gmail.com",
				password:  "Pass!23",
				username:  "murtaza896",
				role:      Admin,
			},
		},
	}
}

func generateJWT(id int, role role) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role.EnumIndex(),
	}).SignedString([]byte(secret))
}
