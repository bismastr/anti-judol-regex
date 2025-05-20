package llm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bismastr/anti-judol-regex/internal/config"
	"google.golang.org/genai"
)

type LlmService interface {
	LlmWebAnalyzeIsJudol(ctx context.Context, request *LlmAnalyzeWebIsJudolRequest) (*LlmAnalyzeWebIsJudolResponse, error)
}

type LlmAnalyzeWebIsJudolRequest struct {
	Domain     string `json:"domain"`
	WebContent string `json:"webContent"`
}

type LlmAnalyzeWebIsJudolResponse struct {
	IsJudol bool `json:"isJudol"`
}

type LlmServiceImpl struct {
	*genai.Client
}

func NewLlmService(ctx context.Context) (*LlmServiceImpl, error) {
	config := &genai.ClientConfig{
		APIKey:  config.Envs.GeminiApiKey,
		Backend: genai.BackendGeminiAPI,
	}
	client, err := genai.NewClient(ctx, config)
	if err != nil {
		return nil, err
	}

	return &LlmServiceImpl{client}, nil
}

func (llm *LlmServiceImpl) LlmWebAnalyzeIsJudol(ctx context.Context, request *LlmAnalyzeWebIsJudolRequest) (*LlmAnalyzeWebIsJudolResponse, error) {
	config := genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				&genai.Part{
					Text: "You are an analyst of a website to find out whether the website is gambling online (Judol) or not, and your job is to analyze starting from the entire structure that will be given to you such as Domain, Header, Body, or so on. You also have to be smarter and not be fooled by the tricks implemented on the website where the tricks are used to fool the system so that the website is not detected as Judol. Provide your results in raw JSON format, for example:{ “isJudol” : True or False} Dont use markdown ```json",
				},
			},
		},
	}

	reqString, err := requestToContent(request)
	if err != nil {
		return nil, err
	}

	fmt.Println("request", reqString)

	parts := []*genai.Part{
		{Text: reqString},
	}

	result, err := llm.Models.GenerateContent(ctx, "gemini-1.5-flash", []*genai.Content{{Parts: parts}}, &config)
	if err != nil {
		return nil, err
	}

	fmt.Println("response ", result.Text())

	var response *LlmAnalyzeWebIsJudolResponse
	json.Unmarshal([]byte(result.Text()), &response)

	return response, nil
}

func requestToContent(request *LlmAnalyzeWebIsJudolRequest) (string, error) {
	requestString := fmt.Sprintf("domain: %s webContent: %s", request.Domain, request.WebContent)
	return requestString, nil
}
