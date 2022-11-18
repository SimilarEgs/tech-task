package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/SimilarEgs/tech-task/internal/models"
	httperrors "github.com/SimilarEgs/tech-task/pkg/httpErrors"
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
		return nil, fmt.Errorf("[Error] occurred during file read: %s", err.Error())
	}

	err = json.Unmarshal(f, &u)
	if err != nil {
		return nil, fmt.Errorf("[Error] occurred while parsing: %s", err.Error())
	}

	res := make([]models.User, 0, 10)

	for _, user := range u.List {
		res = append(res, user)
	}

	return res, nil
}

func (u *UserStore) GetUser(id string) (models.User, error) {

	f, err := ioutil.ReadFile(store)
	if err != nil {
		return models.User{}, fmt.Errorf("[Error] occurred during file read: %s", err.Error())
	}

	err = json.Unmarshal(f, &u)
	if err != nil {
		return models.User{}, fmt.Errorf("[Error] occurred while parsing: %s", err.Error())
	}

	if _, ok := u.List[id]; !ok {
		return models.User{}, httperrors.NotFound
	}

	return u.List[id], nil

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
