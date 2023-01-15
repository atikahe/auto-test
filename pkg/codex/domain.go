package codex

// Reference: https://beta.openai.com/docs/api-reference/completions/create
type CompletionArgs struct {
	Model        string  `json:"model"`
	Prompt       string  `json:"prompt"`
	MaxTokens    int     `json:"max_tokens"`
	Temp         int     `json:"temperature"`
	StopSequence string  `json:"stop,omitempty"`
	TopP         float64 `json:"top_p,omitempty"`
	Stream       bool    `json:"stream,omitempty"`
}

type Choices struct {
	Text         string `json:"text"`
	Index        int64  `json:"index"`
	Longprobs    int64  `json:"longprobs"`
	FinishReason string `json:"finish_reason"`
}

type CompletionResponse struct {
	Choices []Choices `json:"choices"`
}
