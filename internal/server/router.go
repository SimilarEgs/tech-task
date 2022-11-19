package server

import (
	"net/http"
	"time"

	"github.com/SimilarEgs/tech-task/internal/handler"
	"github.com/SimilarEgs/tech-task/internal/repository"
	"github.com/SimilarEgs/tech-task/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewAPI() *chi.Mux {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/user", func(r chi.Router) {
				r.Get("/", userHandler.SearchUsers)
				r.Post("/", userHandler.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", userHandler.GetUser)
					r.Patch("/", userHandler.UpdateUser)
					r.Delete("/", userHandler.DeleteUser)
				})
			})
		})
	})

	return r
}
