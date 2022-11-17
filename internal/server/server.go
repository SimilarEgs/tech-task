package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	*http.Server
}

// NewServer func - creates and configures a new http server
func NewServer(addr string, r *chi.Mux) (*Server, error) {

	srv := http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{&srv}, nil
}

// Start func - runs http server via ListenAndServe with graceful shutdown
func (srv *Server) Start() {

	log.Println("[Info] starting server")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[Error] occurred while starting the server: %s", err.Error())
		}
	}()
	log.Println("[Info] server has been successfully started")
	srv.gracefulShutdown()
}

// gracefulShutdown func - implementation of gracful shutdown
func (srv *Server) gracefulShutdown() {

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("[Info] server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[Error] occured during server shutdown: %s", err.Error())
	}

	log.Println("[Info] server exited")
}
