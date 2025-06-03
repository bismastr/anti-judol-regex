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
	Message string `json:"message"`
}

type RegexAnlyzeRequest struct {
	Text []string `json:"text"`
}

func (request *RegexAnlyzeRequest) Bind(r *http.Request) error {
	return nil
}

type InsertRegexRequest struct {
	Word  string `json:"word"`
	Regex string `json:"regex"`
}

type InsertRegexResponse struct {
	Word    string `json:"word"`
	Regex   string `json:"regex"`
	Message string `json:"message"`
}

func (request *InsertRegexRequest) Bind(r *http.Request) error {
	return nil
}
