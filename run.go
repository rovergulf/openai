package openai

type Run struct {
	Id          string `json:"id"`
	AssistantId string `json:"assistant_id"`
	ThreadId    string `json:"thread_id"`
}
