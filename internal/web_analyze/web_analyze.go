package web_analyze

import (
	"context"

	"github.com/bismastr/anti-judol-regex/internal/llm"
)

type WebAnalyzeService interface {
	WebAnalyzeIsJudol(ctx context.Context, request *WebAnalyzeRequest) (*WebAnalyzeResponse, error)
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
