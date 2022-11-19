package model

import (
	"net/http"
	"time"
)

type User struct {
	CreatedAt   time.Time `json:"created_at"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
}

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *CreateUserRequest) Bind(r *http.Request) error { return nil }
func (c *UpdateUserRequest) Bind(r *http.Request) error { return nil }

type UserService interface {
	SearchUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(user CreateUserRequest) (int, error)
	UpdateUser(user UpdateUserRequest, id string) error
	DeleteUser(id string) error
}

type UserRepository interface {
	SearchUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(user CreateUserRequest) (int, error)
	UpdateUser(user UpdateUserRequest, id string) error
	DeleteUser(id string) error
}
