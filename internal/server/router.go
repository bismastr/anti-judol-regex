package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() {

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/regex/v1", func(w http.ResponseWriter, r *http.Request) {
			// render.JSON(w, r, handler.NewSuccessResponse(http.StatusOK, response))
		})
	})

}
