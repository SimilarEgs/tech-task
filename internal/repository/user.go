package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strconv"
	"time"

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
func (u *UserStore) CreateUser(user models.CreateUserRequest) (int, error) {

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

	newUser := models.User{
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
func (u *UserStore) UpdateUser(user models.UpdateUserRequest, id string) error {

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
func (u *UserStore) DeleteUser(id string) error {

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
