package handler

import (
	"net/http"

	"github.com/bismastr/anti-judol-regex/internal/regex"
	"github.com/go-chi/render"
)

type Handler struct {
	regexService regex.RegexService
}

func NewHandler(regexService regex.RegexService) *Handler {
	return &Handler{
		regexService: regexService,
	}
}

func (h *Handler) GetRegexList(w http.ResponseWriter, r *http.Request) {
	response, err := h.regexService.GetRegexList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	render.JSON(w, r, NewSuccessResponse(200, response))
}
