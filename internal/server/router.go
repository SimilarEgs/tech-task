package server

import (
	"net/http"
	"time"

	"github.com/SimilarEgs/tech-task/internal/repository"
	"github.com/SimilarEgs/tech-task/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewAPI func - Defines the REST apis and returns the new chi router
func NewAPI() *chi.Mux {

	// init repository
	userRepo := repository.NewUserRepository()

	// init services
	userService := service.NewUserService(*userRepo)

	// init handlers
	h := Handler{
		userService: userService,
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", h.SearchUsers)
				r.Post("/", h.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", h.GetUser)
					r.Patch("/", h.UpdateUser)
					r.Delete("/", h.DeleteUser)
				})
			})
		})
	})

	return r
}
