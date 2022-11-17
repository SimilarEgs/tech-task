package repository

import (
	"github.com/SimilarEgs/tech-task/internal/models"
)

// UserStroe struct - responsible for the User data storage
type UserStore struct {
	Increment int      `json:"increment"`
	List      UserList `json:"list"`
}

type UserList map[string]models.User

func NewUserStore() *UserStore {
	return &UserStore{
		List: UserList{},
	}
}

func (u *UserStore) searchUsers() []models.User {
	return nil
}

func (u *UserStore) getUser(id string) (models.User, error) {
	return models.User{}, nil
}
func (u *UserStore) createUser(user models.User) error {
	return nil
}
func (u *UserStore) updateUser(user models.User) error {
	return nil
}
func (u *UserStore) deleteUser(id string) error {
	return nil
}
