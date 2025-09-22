package server

import (
	"context"
	"log"
	internal "nba-api/internal/database"
	h "nba-api/internal/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

func (s *Server) Start() {

	// Run server in a goroutine so we can listen for shutdown
	go func() {
		log.Println("Server starting on :8080")
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Channel to listen for OS interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Context w/ timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server shutdown complete")
}
