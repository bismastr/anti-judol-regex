package regex

import (
	"context"

	"github.com/bismastr/anti-judol-regex/internal/llm"
)

type RegexService interface {
	GetRegexList() (*RegexResponse, error)
	RegexAnalyze(ctx context.Context, request *RegexAnlyzeRequest) (*RegexAnalyzeResponse, error)
}

type RegexServiceImpl struct {
	LlmService llm.LlmService
}

func NewRegexService(llmService llm.LlmService) *RegexServiceImpl {
	return &RegexServiceImpl{
		LlmService: llmService,
	}
}

func (s *RegexServiceImpl) RegexAnalyze(ctx context.Context, request *RegexAnlyzeRequest) (*RegexAnalyzeResponse, error) {
	llmResponse, err := s.LlmService.LlmTextAnalyzeToRegex(ctx, &llm.LlmTextAnalyzeToRegexRequest{Text: request.Text})
	if err != nil {
		return nil, err
	}

	response := RegexAnalyzeResponse{
		TotalJudolText: len(llmResponse),
	}

	//TODO save the judol regex into databse
	if len(llmResponse) < 1 {
		response.Message = "Your reported text is not containing judol"
	} else {
		response.Message = "Thank you for reporting, our LLM analyzed that your reported text is containing judol."
	}

	return &response, nil
}

func (s *RegexServiceImpl) GetRegexList() (*RegexResponse, error) {
	response := &RegexResponse{
		TotalData: 10,
		RegexList: []*Regex{
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

	return response, nil
}
