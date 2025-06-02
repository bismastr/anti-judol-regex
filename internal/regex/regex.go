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
	llmResponse, err := s.LlmService.LlmTextAnalyzeToRegex(ctx, &llm.LlmTextAnalyzeToRegexRequest{Text: request.Text})
	if err != nil {
		return nil, err
	}

	response := RegexAnalyzeResponse{
		TotalJudolText: len(llmResponse),
	}

	if len(llmResponse) == 0 {
		response.Message = "Your reported text is not containing judol"
		return &response, nil
	}

	response.Message = "Thank you for reporting, our LLM analyzed that your reported text is containing judol."

	duplicatedCount := 0
	for _, regex := range llmResponse {
		exist, err := s.Repository.WordExists(ctx, regex.GambeleWord)
		if err != nil {
			return nil, err
		}

		if exist {
			duplicatedCount++
		} else {
			err = s.Repository.InsertRegex(ctx, &repository.Regex{
				Regex: regex.Regex,
				Word:  regex.GambeleWord,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	response.DuplicatedWord = duplicatedCount

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
