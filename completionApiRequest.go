package main

import "encoding/json"

func UnmarshalCompletionAPIRequest(data []byte) (CompletionAPIRequest, error) {
	var r CompletionAPIRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompletionAPIRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CompletionAPIRequest struct {
	Model     string                        `json:"model"`
	Messages  []CompletionAPIRequestMessage `json:"messages"`
	MaxTokens int64                         `json:"max_tokens"`
}

type CompletionAPIRequestMessage struct {
	Role    string                        `json:"role"`
	Content []CompletionAPIRequestContent `json:"content"`
}

type CompletionAPIRequestContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
