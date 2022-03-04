package domain

type UserService interface {
	CreateUser() (User, error)
	DeleteUserByUsername(string) (bool, error)
}

type service struct {
	userRepository UserRepository
}

func (s service) CreateUser(firstName, lastName, username, phone, email, password string, role Role) (User, error) {
	user := NewUser(firstName, lastName, username, phone, email, password, role)
	return s.userRepository.Save(*user)
}

func (s service) DeleteUserByUsername(username string) (bool, error) {
	_, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return false, err
	}
	return s.userRepository.DeleteUserByUsername(username)
}

type UserRepository interface {
	GetAllUsers() ([]*User, error)
	FindByUsername(string) (User, error)
	FindByEmail(string) (User, error)
	Save(User) (User, error)
	DeleteUserByUsername(string) (bool, error)
}
