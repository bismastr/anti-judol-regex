package regex

import (
	"context"

	"github.com/bismastr/anti-judol-regex/internal/llm"
	"github.com/bismastr/anti-judol-regex/internal/repository"
)

type RegexService interface {
	GetRegexList(ctx context.Context) (*RegexResponse, error)
	RegexAnalyze(ctx context.Context, request *RegexAnlyzeRequest) (*RegexAnalyzeResponse, error)
}

type RegexServiceImpl struct {
	LlmService llm.LlmService
	Repository *repository.Queries
}

func NewRegexService(llmService llm.LlmService, repository *repository.Queries) *RegexServiceImpl {
	return &RegexServiceImpl{
		LlmService: llmService,
		Repository: repository,
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

func (s *RegexServiceImpl) GetRegexList(ctx context.Context) (*RegexResponse, error) {
	regex, err := s.Repository.GetRegexList(ctx)
	if err != nil {
		return nil, err
	}

	var totalRegex []*Regex
	for _, r := range regex {
		tempRegex := &Regex{
			Word:  r.Word,
			Regex: r.Regex,
		}

		totalRegex = append(totalRegex, tempRegex)
	}

	response := &RegexResponse{
		TotalData: len(regex),
		RegexList: totalRegex,
	}

	return response, nil
}
