package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewAPI func - Defines the REST apis and returns the new chi router
func NewAPI() *chi.Mux {

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
				r.Get("/", SearchUsers)
				r.Post("/", CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", GetUser)
					r.Patch("/", UpdateUser)
					r.Delete("/", DeleteUser)
				})
			})
		})
	})
	return r
}
