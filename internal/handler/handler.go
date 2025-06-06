package handler

import (
	"net/http"

	"github.com/bismastr/anti-judol-regex/internal/llm"
	"github.com/bismastr/anti-judol-regex/internal/regex"
	"github.com/bismastr/anti-judol-regex/internal/web_analyze"
	"github.com/go-chi/render"
)

type Handler struct {
	regexService      regex.RegexService
	webAnalyzeService web_analyze.WebAnalyzeService
	llmService        llm.LlmService
}

func NewHandler(regexService regex.RegexService, webAnalyzeService web_analyze.WebAnalyzeService, llmService llm.LlmService) *Handler {
	return &Handler{
		regexService:      regexService,
		webAnalyzeService: webAnalyzeService,
		llmService:        llmService,
	}
}

func (h *Handler) GetRegexList(w http.ResponseWriter, r *http.Request) {
	response, err := h.regexService.GetRegexList(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, NewSuccessResponse(200, response))
}

func (h *Handler) InsertRegex(w http.ResponseWriter, r *http.Request) {
	data := &regex.InsertRegexRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	err := h.regexService.InsertRegex(r.Context(), data)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, NewSuccessResponse(200, regex.InsertRegexResponse{
		Word:    data.Word,
		Regex:   data.Regex,
		Message: "Succesfully add regex into db",
	}))
}

func (h *Handler) WebAnalyzeIsJudol(w http.ResponseWriter, r *http.Request) {
	data := &web_analyze.WebAnalyzeRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	response, err := h.webAnalyzeService.WebAnalyzeIsJudol(r.Context(), data)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, NewSuccessResponse(200, response))
}

func (h *Handler) AnalyzeAndConvertToRegex(w http.ResponseWriter, r *http.Request) {
	data := &regex.RegexAnlyzeRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	response, err := h.regexService.RegexAnalyze(r.Context(), data)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, NewSuccessResponse(200, response))
}
