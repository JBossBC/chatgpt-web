package dao

import "time"

type ChatResponse struct {
	Id      string        `json:"id"`
	Object  string        `json:"object"`
	Created time.Duration `json:"created"`
	Choices []Choice      `json:"choices"`
	Usage   Usage         `json:"usage"`
}

type Choice struct {
	Index         int64   `json:"index"`
	Message       Message `json:"message"`
	Finish_reason string  `json:"finish_reason"`
}

type Usage struct {
	Prompt_tokens     int64 `json:"prompt_tokens"`
	Completion_tokens int64 `json:"completion_tokens"`
	Total_tokens      int64 `json:"total_tokens"`
}
