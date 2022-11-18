package server

import (
	"log"
	"net/http"

	"github.com/SimilarEgs/tech-task/internal/service"
	httperrors "github.com/SimilarEgs/tech-task/pkg/httpErrors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Handler struct {
	userService service.UserService
}

func (h *Handler) SearchUsers(w http.ResponseWriter, r *http.Request) {

	data, err := h.userService.SearchUsers()

	if err != nil {
		log.Println(err)
		render.JSON(w, r, httperrors.ErrorResponse(err))
		return
	}

	render.JSON(w, r, data)

}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	data, err := h.userService.GetUser(id)

	if err != nil {
		log.Println(err)
		render.JSON(w, r, httperrors.ErrorResponse(err))
		return
	}

	render.JSON(w, r, data)

}
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {}
