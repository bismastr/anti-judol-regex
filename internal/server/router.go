package server

import (
	"github.com/bismastr/anti-judol-regex/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(h *handler.Handler) (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/regex", h.GetRegexList)
			r.Post("/analyze", h.WebAnalyzeIsJudol)
		})
	})

	return r, nil
}
