package domain

type UserService interface {
	CreateUser(string, string, string, string, string, string, Role) (User, error)
	DeleteUserByUsername(string) (bool, error)
}

type service struct {
	userRepository UserRepository
}

func (s service) CreateUser(firstName, lastName, username, phone, email, password string, role Role) (User, error) {
	user := NewUser(firstName, lastName, username, phone, email, password, role)
	persistedUser, err := s.userRepository.Save(*user)
	if err != nil {
		return User{}, err
	}
	user.SetId(persistedUser.Id())
	return *user, nil
}

func (s service) DeleteUserByUsername(username string) (bool, error) {
	_, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return false, err
	}
	return s.userRepository.DeleteUserByUsername(username)
}

func NewUserService(userRepository UserRepository) UserService {
	return &service{
		userRepository: userRepository,
	}
}
