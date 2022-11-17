package service

type ServerService struct {
	userService UserService
}

func NewServerService(userService UserService) *ServerService {
	return &ServerService{userService: userService}
}
