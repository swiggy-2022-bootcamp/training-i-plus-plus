package user

type UserDbService interface {
	AddUser(User) (User, error)
	DeleteUserByEmail(string) (User, error)
	FetchUsers() (map[string]User, error)
	FetchUserByEmail(string) (User, error)
}

type UserDB struct {
	userMap map[string]User
}

func NewUserDbService() *UserDbService {
	return &UserDB{
		map[string]User{"abc": User{
			email:    "abc",
			password: "!@#",
			name:     "ABC",
			address:  "98, aBcD",
			pincode:  999999,
			mobileNo: 9874563210,
		},
		},
	}
}

func (udb *UserDB) AddUser(user User) (User, error) {

	udb.userMap[user.email] = user

	return user, nil
}

func (udb *UserDB) DeleteUserByEmail(email string) (User, error) {

	user := udb.userMap[email]
	delete(udb.userMap, email)

	return user, nil
}

func (udb *UserDB) FetchUsers() (map[string]User, error) {

	return udb.userMap, nil
}

func (udb *UserDB) FetchUserByEmail(email string) (User, error) {

	user := udb.userMap[email]

	return user, nil
}
