package openai

import (
	"context"
	"fmt"
	"net/http"
)

const (
	modelsUrl = "/v1/models"
)

// ChatCompletion ...
func (c *client) ChatCompletion(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error) {
	var result CompletionResponse

	if err := c.makeRequest(
		ctx,
		http.MethodPost,
		chatCompletionUrl,
		nil,
		req,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

// Models ...
func (c *client) Models(ctx context.Context) (*ModelResponse, error) {
	var result ModelResponse

	if err := c.makeRequest(
		ctx,
		http.MethodGet,
		modelsUrl,
		nil,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) ImageGeneration(ctx context.Context, req *ImageGenRequest) (*ImageGenResponse, error) {
	var result ImageGenResponse

	if err := c.makeRequest(
		ctx,
		http.MethodPost,
		imageGenUrl,
		nil,
		req,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) CreateAssistant(ctx context.Context, req *Assistant) (*Assistant, error) {
	var result Assistant

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	if err := c.makeRequest(
		ctx,
		http.MethodPost,
		assistantUrl,
		headers,
		req,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) ListAssistants(ctx context.Context, req *ListRequestParams) (*ListResponse[*Assistant], error) {
	var result *ListResponse[*Assistant]

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	query := req.ToUrlValues()

	if err := c.makeRequest(
		ctx,
		http.MethodGet,
		assistantUrl+"?"+query.Encode(),
		headers,
		req,
		&result,
	); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) RetrieveAssistant(ctx context.Context, id string) (*Assistant, error) {
	var result Assistant

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	requestUrl := fmt.Sprintf(assistantUrlFmt, id)
	if err := c.makeRequest(
		ctx,
		http.MethodGet,
		requestUrl,
		headers,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) ModifyAssistant(ctx context.Context, id string, data *Assistant) (*Assistant, error) {
	var result Assistant

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	requestUrl := fmt.Sprintf(assistantUrlFmt, id)
	if err := c.makeRequest(
		ctx,
		http.MethodPost,
		requestUrl,
		headers,
		data,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) DeleteAssistant(ctx context.Context, id string) (*DeleteResponse, error) {
	requestUrl := fmt.Sprintf(assistantUrlFmt, id)

	return c.deleteEntity(ctx, requestUrl)
}

func (c *client) CreateThread(ctx context.Context) (*Thread, error) {
	var result Thread

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	if err := c.makeRequest(
		ctx,
		http.MethodPost,
		threadUrl,
		headers,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) RetrieveThread(ctx context.Context, id string) (*Thread, error) {
	var result Thread

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	requestUrl := fmt.Sprintf(threadUrlFmt, id)
	if err := c.makeRequest(
		ctx,
		http.MethodGet,
		requestUrl,
		headers,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) ModifyThread(ctx context.Context, id string, data *Thread) (*Thread, error) {
	var result Thread

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	requestUrl := fmt.Sprintf(threadUrlFmt, id)
	if err := c.makeRequest(
		ctx,
		http.MethodPost,
		requestUrl,
		headers,
		data,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) DeleteThread(ctx context.Context, id string) (*DeleteResponse, error) {
	requestUrl := fmt.Sprintf(threadUrlFmt, id)

	return c.deleteEntity(ctx, requestUrl)
}

func (c *client) deleteEntity(ctx context.Context, requestUrl string) (*DeleteResponse, error) {
	var result DeleteResponse

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	if err := c.makeRequest(
		ctx,
		http.MethodDelete,
		requestUrl,
		headers,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) CreateMessage(ctx context.Context, threadId string, req *CompletionMessage) (*Message, error) {
	var result Message

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	requestUrl := fmt.Sprintf(messagesUrlFmt, threadId)
	if err := c.makeRequest(
		ctx,
		http.MethodPost,
		requestUrl,
		headers,
		req,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) ListMessages(ctx context.Context, threadId string, req *MessagesRequest) (*ListResponse[*Message], error) {
	var result ListResponse[*Message]

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	query := req.ToUrlValues()

	requestUrl := fmt.Sprintf(messagesUrlFmt, threadId)
	if err := c.makeRequest(
		ctx,
		http.MethodGet,
		requestUrl+"?"+query.Encode(),
		headers,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) RetrieveMessage(ctx context.Context, threadId, messageId string) (*Message, error) {
	var result Message

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	requestUrl := fmt.Sprintf(messageUrlFmt, threadId, messageId)
	if err := c.makeRequest(
		ctx,
		http.MethodGet,
		requestUrl,
		headers,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) ModifyMessage(ctx context.Context, threadId, messageId string, data *Message) (*Message, error) {
	var result Message

	headers := http.Header{}
	headers.Add("OpenAI-Beta", "assistants=v2")

	requestUrl := fmt.Sprintf(messageUrlFmt, threadId, messageId)
	if err := c.makeRequest(
		ctx,
		http.MethodPost,
		requestUrl,
		headers,
		data,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) DeleteMessage(ctx context.Context, threadId, messageId string) (*DeleteResponse, error) {
	requestUrl := fmt.Sprintf(messageUrlFmt, threadId, messageId)

	return c.deleteEntity(ctx, requestUrl)
}
