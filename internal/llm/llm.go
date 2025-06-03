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
	LlmTextAnalyzeToRegex(ctx context.Context, request *LlmTextAnalyzeToRegexRequest) ([]*LlmTextAnalyzeToRegexResponse, error)
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
	systemPrompt := "You are tasked to analyze wether a site is a gambling site or not. you must return with json response, with key: isJudol and value: true/false. Make it a raw json string. Dont use markdown ```json"
	config := genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(systemPrompt, genai.RoleModel),
	}

	reqString := fmt.Sprintf("domain: %s webContent: %s", request.Domain, request.WebContent)

	parts := []*genai.Part{
		{Text: reqString},
	}

	result, err := llm.Models.GenerateContent(ctx, "gemini-2.0-flash", []*genai.Content{{Parts: parts}}, &config)
	if err != nil {
		return nil, err
	}

	var response *LlmAnalyzeWebIsJudolResponse
	json.Unmarshal([]byte(result.Text()), &response)

	return response, nil
}

func (llm *LlmServiceImpl) LlmTextAnalyzeToRegex(ctx context.Context, request *LlmTextAnalyzeToRegexRequest) ([]*LlmTextAnalyzeToRegexResponse, error) {
	systemPrompt := "Your tasks are:\n\n" +
		"1. Analyze the following sentences to identify all words or phrases related to **online gambling**. Examples: gambling site names, game codes, or popular terms commonly used in the context of online gambling.\n\n" +
		"2. For each word identified as a gambling element, convert it to a regular expression (regex) **that is resistant to obfuscation** and **various writing variations**, including:\n\n" +
		"   - **Use of upper and lower case letters** (regex should be case-insensitive)\n" +
		"   - Use of symbols, spaces, or underscores between letters**\n" +
		"   - Repetition of letters**, for example: `gambling` can be written as `jjuuddii`\n" +
		"   - **Common character substitution**, for example:\n" +
		"     - `a` → `[a4@]+`\n" +
		"     - `i` → `[i1!|]+`\n" +
		"     - `l` → `[l1|]+`\n" +
		"     - `o` → `[o0]+`\n" +
		"     - `e` → `[e3]+`\n" +
		"     - `s` → `[s5$z]+`\n" +
		"     - `t` → `[t7+]+`\n" +
		"     - `b` → `[b8]+`\n" +
		"     - etc. (use common substitutions often used in spam/gambling contexts)\n\n" +
		"   - Between letters, add `[\\W_]*` to handle spaces, punctuation, or symbols.\n\n" +
		"3. Display the final result in JSON format as follows:\n\n" +
		"[\n" +
		"  {\n" +
		"    \"gamble_word\": \"take4d\",\n" +
		"    \"regex\": \"[a4@]+[\\\\W_]*[m]+[\\\\W_]*[b8]+[\\\\W_]*[i1!|]+[\\\\W_]*[l1|]+[\\\\W_]*[4a@]+[\\\\W_]*[d]+\"\n" +
		"  },\n" +
		"  {\n" +
		"    \"gamble_word\": \"poipet308\",\n" +
		"    \"regex\": \"[p]+[\\\\W_]*[o0]+[\\\\W_]*[i1!|]+[\\\\W_]*[p]+[\\\\W_]*[e3]+[\\\\W_]*[t7+]+[\\\\W_]*[308o]+\"\n" +
		"  }\n" +
		"]"

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(systemPrompt, genai.RoleModel),
		ResponseMIMEType:  "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"gamble_word": {Type: genai.TypeString},
					"regex":       {Type: genai.TypeString},
				},
				PropertyOrdering: []string{"gamble_word", "regex"},
			},
		},
	}

	reqString := fmt.Sprintf("inputPrompt: %s", request.Text)

	parts := []*genai.Part{
		{Text: reqString},
	}

	result, err := llm.Models.GenerateContent(ctx, "gemini-2.0-flash", []*genai.Content{{Parts: parts}}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	var response []*LlmTextAnalyzeToRegexResponse
	json.Unmarshal([]byte(result.Text()), &response)

	return response, nil
}
