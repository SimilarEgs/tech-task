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

func NewUserRepository() *UserStore {
	return &UserStore{
		List: UserList{},
	}
}

func (u *UserStore) SearchUsers() []models.User {
	return nil
}

func (u *UserStore) GetUser(id string) (models.User, error) {
	u.List["1"] = models.User{DisplayName: "Test"}
	return models.User{}, nil
}
func (u *UserStore) CreateUser(user models.User) error {
	return nil
}
func (u *UserStore) UpdateUser(user models.User) error {
	return nil
}
func (u *UserStore) DeleteUser(id string) error {
	return nil
}
