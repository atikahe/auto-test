package codex

// Reference: https://beta.openai.com/docs/api-reference/completions/create
type CompletionArgs struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
	Stop        string  `json:"stop,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	Stream      bool    `json:"stream,omitempty"`
	N           int     `json:"n,omitempty"`
}

type Choices struct {
	Text         string `json:"text"`
	Index        int64  `json:"index"`
	Longprobs    int64  `json:"longprobs"`
	FinishReason string `json:"finish_reason"`
}

type CompletionResponse struct {
	Object  string    `json:"object"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
}
