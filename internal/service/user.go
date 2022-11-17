package service

import "github.com/SimilarEgs/tech-task/internal/models"

type UserService interface {
	searchUsers() ([]models.User, error)
	getUser(id string) (models.User, error)
	createUser(user models.User) error
	updateUser(user models.User) error
	deleteUser(id string) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) searchUsers() ([]models.User, error) {
	return nil, nil
}
func (u *userService) getUser(id string) (models.User, error) {
	return models.User{}, nil
}
func (u *userService) createUser(user models.User) error {
	return nil
}
func (u *userService) updateUser(user models.User) error {
	return nil
}
func (u *userService) deleteUser(id string) error {
	return nil
}
