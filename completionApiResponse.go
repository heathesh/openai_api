package main

import "encoding/json"

func UnmarshalCompletionAPIResponse(data []byte) (CompletionAPIResponse, error) {
	var r CompletionAPIResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompletionAPIResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CompletionAPIResponse struct {
	ID      string                        `json:"id"`
	Object  string                        `json:"object"`
	Created int64                         `json:"created"`
	Model   string                        `json:"model"`
	Usage   CompletionAPIResponseUsage    `json:"usage"`
	Choices []CompletionAPIResponseChoice `json:"choices"`
}

type CompletionAPIResponseChoice struct {
	Message       CompletionAPIResponseMessage       `json:"message"`
	FinishDetails CompletionAPIResponseFinishDetails `json:"finish_details"`
	Index         int64                              `json:"index"`
}

type CompletionAPIResponseFinishDetails struct {
	Type string `json:"type"`
	Stop string `json:"stop"`
}

type CompletionAPIResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionAPIResponseUsage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}
