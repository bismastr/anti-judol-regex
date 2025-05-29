package main

import (
	"context"

	"github.com/bismastr/anti-judol-regex/internal/db"
	"github.com/bismastr/anti-judol-regex/internal/handler"
	"github.com/bismastr/anti-judol-regex/internal/llm"
	"github.com/bismastr/anti-judol-regex/internal/regex"
	"github.com/bismastr/anti-judol-regex/internal/repository"
	"github.com/bismastr/anti-judol-regex/internal/server"
	"github.com/bismastr/anti-judol-regex/internal/web_analyze"
)

type RegexResponse struct {
	TotalData int      `json:"totalData"`
	RegexList *[]Regex `json:"regexList"`
}

type Regex struct {
	Regex string `json:"regex"`
}

func main() {
	ctx := context.Background()

	db, err := db.NewDb()
	if err != nil {
		panic(err)
	}

	llmService, err := llm.NewLlmService(ctx)
	if err != nil {
		panic(err)
	}

	repository := repository.New(db.Conn)
	regexService := regex.NewRegexService(llmService, repository)
	webAnalyzeService := web_analyze.NewWebAnalyzeService(llmService)
	h := handler.NewHandler(regexService, webAnalyzeService, llmService)

	s, err := server.NewServer(h)
	if err != nil {
		panic(err)
	}

	s.Start()
}
