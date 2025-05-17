package llm_test

import (
	"context"
	"testing"

	"github.com/bismastr/anti-judol-regex/internal/llm"
	"github.com/stretchr/testify/assert"
)

type LlmService struct {
	llm.LlmService
}

func TestLlmIsJudol(t *testing.T) {
	ctx := context.Background()
	llmClient, err := llm.NewLlmService(ctx)
	if err != nil {
		t.Fatalf("cannot create llm client %v", err)
	}

	judolHeader := `<head><meta charset="UTF-8"><meta content="width=device-width,initial-scale=1.0,minimum-scale=1.0" name="viewport"><meta content="telephone=no" name="format-detection"><meta content="address=no" name="format-detection"><meta content="dark light" name="color-scheme"><meta content="origin" name="referrer"><meta content="notranslate" name="google"><meta content="Al+K52juR0aEQRQLoQCQAaRGSl3jdaBxoCZb3LNsVUIRMea6Bwf8rvpnyTLry3bPmIyZuar7DGcJTZYXgeJ24Q8AAAB2eyJvcmlnaW4iOiJodHRwczovL2dvb2dsZS5jb206NDQzIiwiZmVhdHVyZSI6IkF0dHJpYnV0aW9uUmVwb3J0aW5nQ3Jvc3NBcHBXZWIiLCJleHBpcnkiOjE3MTQ1MjE1OTksImlzU3ViZG9tYWluIjp0cnVlfQ==" http-equiv="origin-trial"><link href="/images/branding/product/1x/gsa_android_144dp.png" rel="icon"><meta content="/images/branding/googleg/1x/googleg_standard_color_128dp.png" itemprop="image"><title>PAWPAW4D: LINK LOGIN ALTERNATIF TERBAIK No.1 Di INDONESIA</title>`
	llmService := LlmService{
		llmClient,
	}

	response, err := llmService.LlmWebAnalyzeIsJudol(ctx, &llm.LlmAnalyzeWebIsJudolRequest{
		Domain:     "wwww.p4w4d.com",
		WebContent: judolHeader,
	})
	if err != nil {
		t.Fatalf("Cannot anlyze is judol %v", err)
	}

	assert.True(t, response.IsJudol)

	nonJudolHeader := `<head><meta charset="UTF-8"><meta content="width=device-width,initial-scale=1.0,minimum-scale=1.0" name="viewport"><meta content="telephone=no" name="format-detection"><meta content="address=no" name="format-detection"><meta content="dark light" name="color-scheme"><meta content="origin" name="referrer"><meta content="notranslate" name="google"><meta content="Al+K52juR0aEQRQLoQCQAaRGSl3jdaBxoCZb3LNsVUIRMea6Bwf8rvpnyTLry3bPmIyZuar7DGcJTZYXgeJ24Q8AAAB2eyJvcmlnaW4iOiJodHRwczovL2dvb2dsZS5jb206NDQzIiwiZmVhdHVyZSI6IkF0dHJpYnV0aW9uUmVwb3J0aW5nQ3Jvc3NBcHBXZWIiLCJleHBpcnkiOjE3MTQ1MjE1OTksImlzU3ViZG9tYWluIjp0cnVlfQ==" http-equiv="origin-trial"><link href="/images/branding/product/1x/gsa_android_144dp.png" rel="icon"><meta content="/images/branding/googleg/1x/googleg_standard_color_128dp.png" itemprop="image"><title>Google</title>`
	response, err = llmService.LlmWebAnalyzeIsJudol(ctx, &llm.LlmAnalyzeWebIsJudolRequest{
		Domain:     "www.google.com",
		WebContent: nonJudolHeader,
	})
	if err != nil {
		t.Fatalf("Cannot anlyze is judol %v", err)
	}

	assert.False(t, response.IsJudol)
}
