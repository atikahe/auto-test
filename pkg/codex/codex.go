// Package codex provides an interface to interact with OpenAI's Codex API
package codex

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type codex struct {
	Client *resty.Client
	Token  string
	Model  string
}

func New(token string, url string, model ...string) *codex {
	return &codex{
		Client: resty.New().SetBaseURL(url),
		Token:  token,
		Model:  "code-davinci-002",
	}
}

func (c *codex) CreateCompletion(args *CompletionArgs) (*CompletionResponse, error) {
	completionURL := fmt.Sprintf("%s/completions", c.Client.BaseURL)
	response, err := c.Client.R().EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetAuthScheme("Bearer").
		SetAuthToken(c.Token).
		SetBody(args).
		Post(completionURL)
	if err != nil {
		return nil, errors.New("unable to send request")
	}

	var body CompletionResponse
	if err = json.Unmarshal(response.Body(), &body); err != nil {
		return nil, errors.New("unable to unmarshal response")
	}

	return &body, nil
}
