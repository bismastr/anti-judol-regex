package web_analyze

import (
	"errors"
	"net/http"
)

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
