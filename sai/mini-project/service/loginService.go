package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "sai",
		authorizedPassword: "kumar",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return username == service.authorizedUsername && password == service.authorizedPassword
}
