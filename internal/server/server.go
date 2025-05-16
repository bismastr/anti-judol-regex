package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bismastr/anti-judol-regex/internal/handler"
)

type Server struct {
	*http.Server
}

func NewServer(h *handler.Handler) (*Server, error) {
	r, err := NewRouter(h)
	if err != nil {
		return nil, err
	}

	srv := http.Server{
		Addr:    ":4527",
		Handler: r,
	}

	return &Server{&srv}, nil
}

func (s *Server) Start() {
	log.Println("starting server...")
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", s.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)

	if err := s.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}
