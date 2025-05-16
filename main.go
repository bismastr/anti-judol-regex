package main

import (
	"net/http"

	"github.com/bismastr/anti-judol-regex/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
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

	response := &RegexResponse{
		TotalData: 10,
		RegexList: &[]Regex{
			{
				Regex: "m+\\s*[a4]+\\s*x+\\s*w+\\s*i+\\s*n+",
			},
			{
				Regex: "j+\\s*[a4]+\\s*c+\\s*k+\\s*p+\\s*[o0]+\\s*t+",
			},
			{
				Regex: "p+\\s*e+\\s*t+\\s*i+\\s*r+",
			},
			{
				Regex: "z+\\s*e+\\s*u+\\s*s+",
			},
			{
				Regex: "k+\\s*[a4]+\\s*k+\\s*e+\\s*k+",
			},
			{
				Regex: "g+\\s*[a4]+\\s*c+\\s*[o0]+\\s*r+",
			},
			{
				Regex: "g+\\s*u+\\s*a+\\s*c+\\s*[o0]+\\s*r+",
			},
			{
				Regex: "t+\\s*[e3]+\\s*r+\\s*p+\\s*[e3]+\\s*r+\\s*c+\\s*[a4]+\\s*y+\\s*[a4]+",
			},
			{
				Regex: "c+\\s*u+\\s*[a4]+\\s*n+",
			},
		},
	}

	r.Get("/regex/v1", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, handler.NewSuccessResponse(http.StatusOK, response))
	})

	http.ListenAndServe(":4527", r)
}
