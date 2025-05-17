package handler

import (
	"net/http"

	"github.com/bismastr/anti-judol-regex/internal/regex"
	"github.com/bismastr/anti-judol-regex/internal/web_analyze"
	"github.com/go-chi/render"
)

type Handler struct {
	regexService      regex.RegexService
	webAnalyzeService web_analyze.WebAnalyzeService
}

func NewHandler(regexService regex.RegexService, webAnalyzeService web_analyze.WebAnalyzeService) *Handler {
	return &Handler{
		regexService:      regexService,
		webAnalyzeService: webAnalyzeService,
	}
}

func (h *Handler) GetRegexList(w http.ResponseWriter, r *http.Request) {
	response, err := h.regexService.GetRegexList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	render.JSON(w, r, NewSuccessResponse(200, response))
}

func (h *Handler) WebAnalyzeIsJudol(w http.ResponseWriter, r *http.Request) {
	data := &web_analyze.WebAnalyzeRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}

	response, err := h.webAnalyzeService.WebAnalyzeIsJudol(r.Context(), data)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}

	render.JSON(w, r, NewSuccessResponse(200, response))
}
