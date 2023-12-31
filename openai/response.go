package openai

type ResponseBody struct {
	Id      string   `json:"id,omitempty"`
	Object  string   `json:"object,omitempty"`
	Created int64    `json:"created,omitempty"`
	Model   string   `json:"model,omitempty"`
	Usage   Usage    `json:"usage,omitempty"`
	Choices []Choice `json:"choices,omitempty"`
}

type Usage struct {
	PromptTokens     int32 `json:"prompt_tokens,omitempty"`
	CompletionTokens int32 `json:"completion_tokens,omitempty"`
	TotalTokens      int32 `json:"total_tokens,omitempty"`
}

type Choice struct {
	Message      Message     `json:"message,omitempty"`
	Logprobs     string      `json:"logprobs,omitempty"`
	FinishReason interface{} `json:"finish_reason,omitempty"`
	Index        int32       `json:"index,omitempty"`
}
