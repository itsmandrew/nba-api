package server

import (
	"fmt"
	internal "nba-api/internal/database"
	h "nba-api/internal/handlers"
	"net/http"
	"time"
)

type Server struct {
	store  *internal.Store
	server *http.Server
}

func New(store *internal.Store) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.HelloHandler)

	return &Server{
		store: store,
		server: &http.Server{
			Addr:           ":8080",
			Handler:        mux,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			IdleTimeout:    15 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}

}

func (s *Server) Start() error {
	fmt.Println("Server starting on :8080...")
	return s.server.ListenAndServe()
}
