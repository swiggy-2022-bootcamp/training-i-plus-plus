package domain

type UserService interface {
	CreateUserInMongo(string, string, string, string, string, string, Role) (User, error)
	GetMongoUserByUserId(int) (*User, error)
	DeleteUserByUserId(int) error
	UpdateUser(User, int) (*User, error)
}

type service struct {
	userMongoRepository UserMongoRepository
}

func (s service) CreateUserInMongo(firstName, lastName, username, phone, email, password string, role Role) (User, error) {
	user := NewUser(firstName, lastName, username, phone, email, password, role)
	persistedUser, err := s.userMongoRepository.InsertUser(*user)
	if err != nil {
		return User{}, err
	}
	return persistedUser, nil
}

func (s service) GetMongoUserByUserId(userId int) (*User, error) {
	res, err := s.userMongoRepository.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) DeleteUserByUserId(userId int) error {
	err := s.userMongoRepository.DeleteUserByUserId(userId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateUser(user User, userId int) (*User, error) {
	user.Id = userId
	res, err := s.userMongoRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewUserService(userMongoRepository UserMongoRepository) UserService {
	return &service{
		userMongoRepository: userMongoRepository,
	}
}
