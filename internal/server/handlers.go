package server

import (
	"net/http"

	"github.com/SimilarEgs/tech-task/internal/service"
)

type Handler struct {
	userService service.UserService
}

func (h *Handler) SearchUsers(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request)     {}
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request)  {}
