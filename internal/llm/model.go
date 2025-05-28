package llm

import "net/http"

type LlmTextAnalyzeToRegexRequest struct {
	Text []string `json:"text"`
}

func (l *LlmTextAnalyzeToRegexRequest) Bind(r *http.Request) error {
	return nil
}

type LlmTextAnalyzeToRegexResponse struct {
	GambeleWord string `json:"gamble_word"`
	Regex       string `json:"regex"`
}

type LlmAnalyzeWebIsJudolRequest struct {
	Domain     string `json:"domain"`
	WebContent string `json:"webContent"`
}

func (l *LlmAnalyzeWebIsJudolRequest) Bind(r *http.Request) error {
	return nil
}

type LlmAnalyzeWebIsJudolResponse struct {
	IsJudol bool `json:"isJudol"`
}
