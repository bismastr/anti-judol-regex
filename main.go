package main

import (
	"github.com/bismastr/anti-judol-regex/internal/handler"
	"github.com/bismastr/anti-judol-regex/internal/regex"
	"github.com/bismastr/anti-judol-regex/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type RegexResponse struct {
	TotalData int      `json:"totalData"`
	RegexList *[]Regex `json:"regexList"`
}

type Regex struct {
	Regex string `json:"regex"`
}

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	regexService := regex.NewRegexService()
	h := handler.NewHandler(regexService)
	s, err := server.NewServer(h)
	if err != nil {
		panic(err)
	}

	s.Start()
}
