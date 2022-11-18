package service

import (
	"github.com/SimilarEgs/tech-task/internal/models"
	"github.com/SimilarEgs/tech-task/internal/repository"
)

type UserService interface {
	SearchUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
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
	test, _ := u.userRepository.GetUser(id)
	return test, nil
}
func (u *userService) CreateUser(user models.User) error {
	return nil
}
func (u *userService) UpdateUser(user models.User) error {
	return nil
}
func (u *userService) DeleteUser(id string) error {
	return nil
}
