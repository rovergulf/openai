package openai

import (
	"fmt"
	"net/url"
	"strconv"
)

// Reference: https://platform.openai.com/docs/models
const (
	GPT35Turbo = "gpt-3.5-turbo"
	GPT4o      = "gpt-4o"
	GPT4Turbo  = "gpt-4-turbo"
	O1         = "o1"
	O1Mini     = "o1-mini"
	O3Mini     = "o3-mini"
)

const (
	SystemRole    = "system"
	UserRole      = "user"
	AssistantRole = "assistant"
	FunctionRole  = "function"
	ToolRole      = "tool"
)

const (
	defaultTemperature = 1
)

const (
	DefaultSystemTask = "You're a helpful assistant."
)

type ErrorResponse struct {
	Data Err `json:"error"`
}

type Err struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   any    `json:"param"`
	Code    string `json:"code"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("[%s] %s", e.Data.Code, e.Data.Message)
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

type Tool struct {
	Type string `json:"type"`
}

type ListRequestParams struct {
	Limit  int    `json:"limit"` // defaults to 20
	Order  string `json:"order"`
	After  string `json:"after"`
	Before string `json:"before"`
}

func (req *ListRequestParams) toUrlValues() url.Values {
	query := url.Values{}
	if req.Limit > 0 {
		query.Set("limit", strconv.Itoa(req.Limit))
	}

	if req.Order != "" {
		query.Set("order", req.Order)
	}

	if req.Before != "" {
		query.Set("before", req.Before)
	}

	if req.After != "" {
		query.Set("after", req.After)
	}

	return query
}

type ListResponse[T any] struct {
	Object  string `json:"object"`
	Data    []T    `json:"data"`
	FirstId string `json:"first_id"`
	LastId  string `json:"last_id"`
	HasMore bool   `json:"has_more"`
}

type DeleteResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}
