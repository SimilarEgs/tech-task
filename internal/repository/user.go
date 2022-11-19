package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/SimilarEgs/tech-task/internal/model"
	httperrors "github.com/SimilarEgs/tech-task/pkg/httpErrors"
)

const (
	store = "users.json"
)

type userStore struct {
	Increment int      `json:"increment"`
	List      UserList `json:"list"`
}

type UserList map[string]model.User

func NewUserRepository() *userStore {
	return &userStore{
		List: UserList{},
	}
}

func (u *userStore) SearchUsers() ([]model.User, error) {
	f, err := ioutil.ReadFile(store)
	if err != nil {
		return nil, fmt.Errorf("[Error] occurred during file read: %s", err.Error())
	}

	err = json.Unmarshal(f, &u)
	if err != nil {
		return nil, fmt.Errorf("[Error] occurred while parsing: %s", err.Error())
	}

	res := make([]model.User, 0, 10)

	for _, user := range u.List {
		res = append(res, user)
	}

	return res, nil
}

func (u *userStore) GetUser(id string) (model.User, error) {
	f, err := ioutil.ReadFile(store)
	if err != nil {
		return model.User{}, fmt.Errorf("[Error] occurred during file read: %s", err.Error())
	}

	err = json.Unmarshal(f, &u)
	if err != nil {
		return model.User{}, fmt.Errorf("[Error] occurred while parsing: %s", err.Error())
	}

	if _, ok := u.List[id]; !ok {
		return model.User{}, httperrors.NotFound
	}

	return u.List[id], nil

}
func (u *userStore) CreateUser(user model.CreateUserRequest) (int, error) {
	if user.DisplayName == "" || user.Email == "" {
		return 0, errors.New("[Error] empty fields")
	}

	f, err := ioutil.ReadFile(store)
	if err != nil {
		return 0, fmt.Errorf("[Error] occurred during file read: %s", err.Error())
	}

	err = json.Unmarshal(f, &u)
	if err != nil {
		return 0, fmt.Errorf("[Error] occurred while parsing: %s", err.Error())
	}

	u.Increment++

	newUser := model.User{
		CreatedAt:   time.Now(),
		DisplayName: user.DisplayName,
		Email:       user.Email,
	}

	id := strconv.Itoa(u.Increment)
	u.List[id] = newUser

	b, err := json.Marshal(&u)
	if err != nil {
		return 0, fmt.Errorf("[Error] occurred while marshalling: %s", err.Error())
	}

	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		return 0, fmt.Errorf("[Error] occurred while writing file: %s", err.Error())
	}

	return u.Increment, nil
}

func (u *userStore) UpdateUser(user model.UpdateUserRequest, id string) error {
	f, err := ioutil.ReadFile(store)
	if err != nil {
		return fmt.Errorf("[Error] occurred during file read: %s", err.Error())
	}

	err = json.Unmarshal(f, &u)
	if err != nil {
		return fmt.Errorf("[Error] occurred while parsing: %s", err.Error())
	}

	if _, ok := u.List[id]; !ok {
		return httperrors.NotFound
	}

	updateUser := u.List[id]

	updateUser.DisplayName = user.DisplayName
	updateUser.Email = user.Email

	u.List[id] = updateUser

	b, err := json.Marshal(&u)
	if err != nil {
		return fmt.Errorf("[Error] occurred while marshalling: %s", err.Error())
	}

	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("[Error] occurred while writing file: %s", err.Error())
	}

	return nil
}

func (u *userStore) DeleteUser(id string) error {
	f, err := ioutil.ReadFile(store)
	if err != nil {
		return fmt.Errorf("[Error] occurred during file read: %s", err.Error())
	}

	err = json.Unmarshal(f, &u)
	if err != nil {
		return fmt.Errorf("[Error] occurred while parsing: %s", err.Error())
	}

	if _, ok := u.List[id]; !ok {
		return httperrors.NotFound
	}

	delete(u.List, id)

	b, err := json.Marshal(&u)
	if err != nil {
		return fmt.Errorf("[Error] occurred while marshalling: %s", err.Error())
	}

	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("[Error] occurred while writing file: %s", err.Error())
	}

	return nil
}
