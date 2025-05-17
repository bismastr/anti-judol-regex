package web_analyze

import (
	"context"
	"errors"
	"net/http"

	"github.com/bismastr/anti-judol-regex/internal/llm"
)

type WebAnalyzeService interface {
	WebAnalyzeIsJudol(ctx context.Context, request *WebAnalyzeRequest) (*WebAnalyzeResponse, error)
}

type WebAnalyzeResponse struct {
	IsJudol     bool   `json:"isJudol"`
	IsNewDomain bool   `json:"isNewDomain"`
	Message     string `json:"message"`
}

type WebAnalyzeRequest struct {
	Domain string `json:"domain"`
	Header string `json:"header"`
}

func (req *WebAnalyzeRequest) Bind(r *http.Request) error {
	if req.Header == "" {
		return errors.New("missing required field header")
	}

	return nil
}

type WebAnalyzeImpl struct {
	LlmService llm.LlmService
}

func NewWebAnalyzeService(llmService llm.LlmService) *WebAnalyzeImpl {
	return &WebAnalyzeImpl{
		LlmService: llmService,
	}
}

func (w *WebAnalyzeImpl) WebAnalyzeIsJudol(ctx context.Context, request *WebAnalyzeRequest) (*WebAnalyzeResponse, error) {
	llmResponse, err := w.LlmService.LlmWebAnalyzeIsJudol(ctx, &llm.LlmAnalyzeWebIsJudolRequest{
		Domain:     request.Domain,
		WebContent: request.Header,
	})
	if err != nil {
		return nil, err
	}

	return &WebAnalyzeResponse{
		IsJudol:     llmResponse.IsJudol,
		IsNewDomain: true,
	}, nil
}
