package server

import (
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

// NewServer func - creates and configures a new http server
func NewServer(addr string) (*Server, error) {

	srv := http.Server{
		Addr:         addr,
		Handler:      NewAPI(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{&srv}, nil
}
