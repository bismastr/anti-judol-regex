package regex

import "net/http"

type RegexResponse struct {
	TotalData int      `json:"totalData"`
	RegexList []*Regex `json:"regexList"`
}

type Regex struct {
	Word  string `json:"word"`
	Regex string `json:"regex"`
}

type RegexAnalyzeResponse struct {
	TotalJudolText int    `json:"total_judol_text"`
	Message        string `json:"message"`
}

type RegexAnlyzeRequest struct {
	Text []string `json:"text"`
}

func (request *RegexAnlyzeRequest) Bind(r *http.Request) error {
	return nil
}
