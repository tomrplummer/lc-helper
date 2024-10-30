package lc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tomrplummer/lc-helper/internal/gpt"
)

const (
	GraphqlApiUrl = "https://leetcode.com/graphql"
)

type Scraper struct {
	Url     string
	Content []byte
}

func New(url string) *Scraper {
	return &Scraper{
		Url: url,
	}
}

func Query(slug, lang string) (gpt.QueryResult, error) {
	query := `
		query($titleSlug: String!) {
			question(titleSlug: $titleSlug) {
				questionFrontendId
				titleSlug
				content
				codeSnippets {
					lang
					langSlug
					code 
				}
			}
		}
	`

	variables := map[string]string{
		"titleSlug": slug,
	}

	payload := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return gpt.QueryResult{}, fmt.Errorf("unable to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", GraphqlApiUrl, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return gpt.QueryResult{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return gpt.QueryResult{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result gpt.GraphQLResponse

	err = json.NewDecoder(resp.Body).Decode(&result)

	content := result.Data.Question.Content
	var code string

	for _, val := range result.Data.Question.CodeSnippets {
		if val.LangSlug == lang {
			code = val.Code
		}
	}

	return gpt.QueryResult{
		Code:    code,
		Content: content,
	}, nil
}
