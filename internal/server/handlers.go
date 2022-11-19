package server

import (
	"log"
	"net/http"

	"github.com/SimilarEgs/tech-task/internal/models"
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
		if err != httperrors.NotFound {
			log.Println(err)
		}
		render.JSON(w, r, httperrors.ErrorResponse(err))
		return
	}

	render.JSON(w, r, data)

}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	request := models.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		log.Printf("[Error] %s\n", err.Error())
		_ = render.Render(w, r, httperrors.NewBadRequestError(err))
		return
	}

	id, err := h.userService.CreateUser(request)
	if err != nil {
		log.Println(err)
		render.JSON(w, r, httperrors.ErrorResponse(err))
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})

}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	request := models.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		log.Printf("[Error] %s\n", err.Error())
		_ = render.Render(w, r, httperrors.NewBadRequestError(err))
		return
	}

	id := chi.URLParam(r, "id")

	err := h.userService.UpdateUser(request, id)
	if err != nil {
		if err != httperrors.NotFound {
			log.Println(err)
		}
		render.JSON(w, r, httperrors.ErrorResponse(err))
	}

	render.Status(r, http.StatusNoContent)
}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.userService.DeleteUser(id)

	if err != nil {
		if err != httperrors.NotFound {
			log.Println(err)
		}
		render.JSON(w, r, httperrors.ErrorResponse(err))
		return
	}

	render.Status(r, http.StatusNoContent)

}
