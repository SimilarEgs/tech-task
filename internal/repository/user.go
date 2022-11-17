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
