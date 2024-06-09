package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

const (
	apiUrl = "https://api.openai.com"
)

type Client interface {
	Models(ctx context.Context) (*ModelResponse, error)

	ChatCompletion(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error)

	ImageGeneration(ctx context.Context, req *ImageGenRequest) (*ImageGenResponse, error)

	ListAssistants(ctx context.Context, req *ListRequestParams) (*ListResponse[*Assistant], error)
	CreateAssistant(ctx context.Context, req *Assistant) (*Assistant, error)
	RetrieveAssistant(ctx context.Context, id string) (*Assistant, error)
	ModifyAssistant(ctx context.Context, id string, data *Assistant) (*Assistant, error)
	DeleteAssistant(ctx context.Context, id string) (*DeleteResponse, error)

	CreateThread(ctx context.Context) (*Thread, error)
	RetrieveThread(ctx context.Context, id string) (*Thread, error)
	ModifyThread(ctx context.Context, id string, thread *Thread) (*Thread, error)
	DeleteThread(ctx context.Context, id string) (*DeleteResponse, error)

	ListMessages(ctx context.Context, threadId string, req *MessagesRequest) (*ListResponse[*Message], error)
	CreateMessage(ctx context.Context, threadId string, req *CompletionMessage) (*Message, error)
	RetrieveMessage(ctx context.Context, threadId, messageId string) (*Message, error)
	ModifyMessage(ctx context.Context, threadId, messageId string, data *Message) (*Message, error)
	DeleteMessage(ctx context.Context, threadId, messageId string) (*DeleteResponse, error)
}

type client struct {
	apiKey string
	client *http.Client
}

func NewClient(apikey string) Client {
	return &client{
		apiKey: apikey,
		client: &http.Client{Timeout: 25 * time.Second},
	}
}

func (c *client) makeRequest(ctx context.Context, method, url string, headers http.Header, body any, result any) error {
	if result != nil && reflect.ValueOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("result must be a pointer")
	}
	var payload *bytes.Buffer
	if body != nil {
		payload = new(bytes.Buffer)
		encoder := json.NewEncoder(payload)
		if err := encoder.Encode(body); err != nil {
			return err
		}
	}

	requestUrl := fmt.Sprintf("%s%s", apiUrl, url)
	req, err := http.NewRequestWithContext(ctx, method, requestUrl, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	for k, v := range headers {
		req.Header.Add(k, v[0])
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var openAIErr ErrorResponse
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&openAIErr); err != nil {
			return fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}

		return &openAIErr
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(result); err != nil {
		return err
	}

	return nil
}
