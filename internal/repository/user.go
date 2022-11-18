package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/SimilarEgs/tech-task/internal/models"
)

const (
	store = "users.json"
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

func (u *UserStore) SearchUsers() ([]models.User, error) {

	f, err := ioutil.ReadFile(store)
	if err != nil {
		return nil, fmt.Errorf("error occurred during file read: %s", err.Error())
	}

	s := UserStore{}

	err = json.Unmarshal(f, &s)
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing: %s", err.Error())
	}

	res := make([]models.User, 0, 10)

	for _, user := range s.List {
		res = append(res, user)
	}

	return res, nil
}

func (u *UserStore) GetUser(id string) (models.User, error) {
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
