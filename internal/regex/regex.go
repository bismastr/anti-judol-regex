package regex

import (
	"context"
	"errors"

	"github.com/bismastr/anti-judol-regex/internal/llm"
	"github.com/bismastr/anti-judol-regex/internal/repository"
)

type RegexService interface {
	GetRegexList(ctx context.Context) (*RegexResponse, error)
	InsertRegex(ctx context.Context, regex *InsertRegexRequest) error
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

func (s *RegexServiceImpl) InsertRegex(ctx context.Context, request *InsertRegexRequest) error {
	exist, err := s.Repository.WordExists(ctx, request.Word)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("Word is duplicated")
	}

	err = s.Repository.InsertRegex(ctx, &repository.Regex{
		Word:  request.Word,
		Regex: request.Regex,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *RegexServiceImpl) RegexAnalyze(ctx context.Context, request *RegexAnlyzeRequest) (*RegexAnalyzeResponse, error) {
	response := RegexAnalyzeResponse{
		Message: "Thank you for reporting, our AI will analyze your reported text.",
	}

	go s.processLLMResponse(context.Background(), request.Text)

	return &response, nil
}

func (s *RegexServiceImpl) processLLMResponse(ctx context.Context, text []string) error {
	llmResponse, err := s.LlmService.LlmTextAnalyzeToRegex(ctx, &llm.LlmTextAnalyzeToRegexRequest{Text: text})
	if err != nil {
		return err
	}

	for _, regex := range llmResponse {
		exist, err := s.Repository.WordExists(ctx, regex.GambeleWord)
		if err != nil {
			return err
		}

		if !exist {
			err = s.Repository.InsertRegex(ctx, &repository.Regex{
				Regex: regex.Regex,
				Word:  regex.GambeleWord,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
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
