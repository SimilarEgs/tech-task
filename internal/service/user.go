package service

import (
	"github.com/SimilarEgs/tech-task/internal/models"
	"github.com/SimilarEgs/tech-task/internal/repository"
)

type UserService interface {
	SearchUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(user models.CreateUserRequest) (int, error)
	UpdateUser(user models.UpdateUserRequest, id string) error
	DeleteUser(id string) error
}

type userService struct {
	userRepository repository.UserStore
}

func NewUserService(userRepo repository.UserStore) UserService {
	return &userService{userRepository: userRepo}
}

func (u *userService) SearchUsers() ([]models.User, error) {
	return u.userRepository.SearchUsers()
}

func (u *userService) GetUser(id string) (models.User, error) {
	return u.userRepository.GetUser(id)
}

func (u *userService) CreateUser(user models.CreateUserRequest) (int, error) {
	return u.userRepository.CreateUser(user)
}

func (u *userService) UpdateUser(user models.UpdateUserRequest, id string) error {
	return u.userRepository.UpdateUser(user, id)
}

func (u *userService) DeleteUser(id string) error {
	return u.userRepository.DeleteUser(id)
}
