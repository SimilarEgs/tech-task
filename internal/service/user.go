package service

import (
	"github.com/SimilarEgs/tech-task/internal/model"
)

type userService struct {
	userRepository model.UserRepository
}

func NewUserService(userRepo model.UserRepository) model.UserRepository {
	return &userService{userRepository: userRepo}
}

func (u *userService) SearchUsers() ([]model.User, error) {
	return u.userRepository.SearchUsers()
}

func (u *userService) GetUser(id string) (model.User, error) {
	return u.userRepository.GetUser(id)
}

func (u *userService) CreateUser(user model.CreateUserRequest) (int, error) {
	return u.userRepository.CreateUser(user)
}

func (u *userService) UpdateUser(user model.UpdateUserRequest, id string) error {
	return u.userRepository.UpdateUser(user, id)
}

func (u *userService) DeleteUser(id string) error {
	return u.userRepository.DeleteUser(id)
}
